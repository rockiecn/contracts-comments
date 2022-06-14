package callconts

import (
	"log"
	"math/big"
	"memoc/contracts/erc20"
	iface "memoc/interfaces"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// NewERC20 new a instance of ContractModule
func NewERC20(erc20Addr, addr common.Address, hexSk string, txopts *TxOpts, endPoint string, status chan error) iface.ERC20Info {
	e := &ContractModule{
		addr:            addr,
		hexSk:           hexSk,
		txopts:          txopts,
		contractAddress: erc20Addr,
		endPoint:        endPoint,
		Status:          status, // 用于接收：后台goroutine检查交易是否执行成功， nil代表成功
	}

	return e
}

// DeployERC20 deploy an ERC20 contract, called by admin, specify name and symbol.
func (e *ContractModule) DeployERC20(name, symbol string, version uint16, addrs [5]common.Address) (common.Address, *erc20.ERC20, error) {
	var erc20Addr common.Address
	var erc20Ins *erc20.ERC20

	log.Println("begin deploy ERC20 contract...")
	client := getClient(e.endPoint)
	defer client.Close()

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(e.endPoint, e.hexSk, nil, e.txopts)
	if errMA != nil {
		return erc20Addr, nil, errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	erc20Addr, tx, erc20Ins, err := erc20.DeployERC20(auth, client, name, symbol, version, addrs)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("DeployERC20 Err:", err)
		return erc20Addr, nil, err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(e.endPoint, tx, e.Status, "DeployERC20")

	log.Println("ERC20 address is ", erc20Addr.Hex())
	return erc20Addr, erc20Ins, nil
}

// newERC20 new an instance of ERC20 contract, 'erc20Addr' indicates ERC20 contract address
func newERC20(erc20Addr common.Address, client *ethclient.Client) (*erc20.ERC20, error) {
	erc20Ins, err := erc20.NewERC20(erc20Addr, client)
	if err != nil {
		return nil, err
	}
	return erc20Ins, nil
}

// Transfer the account represented by the e.hexsk transfers the specified amount to recipient.
func (e *ContractModule) Transfer(recipient common.Address, value *big.Int) error {
	client := getClient(e.endPoint)
	defer client.Close()
	erc20Ins, err := newERC20(e.contractAddress, client)
	if err != nil {
		return err
	}

	paused, err := e.GetPaused()
	if err != nil {
		return err
	}
	if paused {
		return errPaused
	}

	if recipient.Hex() == InvalidAddr {
		log.Println("you try to transfer money to an invalidAddr")
		return ErrInValAddr
	}
	if e.addr.Hex() == InvalidAddr {
		log.Println("sender in transfer is an invalidAddr")
		return ErrInValAddr
	}

	// need to determine whether the account balance is enough to transfer.
	bal, err := e.BalanceOf(e.addr)
	if err != nil {
		return err
	}
	log.Println("Your balance is ", bal)
	if bal.Cmp(value) < 0 {
		log.Println("Balance is not enough, please reconfirm the transfer amount.")
		return ErrBalNotE
	}

	log.Println("begin Transfer to", recipient.Hex(), " with value", value, " in ERC20 contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(e.endPoint, e.hexSk, nil, e.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := erc20Ins.Transfer(auth, recipient, value)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("Transfer Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(e.endPoint, tx, e.Status, "Transfer")

	return nil
}

// Approve The account represented by the e.hexsk authorizes the balance of the specified amount to addr.
func (e *ContractModule) Approve(addr common.Address, value *big.Int) error {
	client := getClient(e.endPoint)
	defer client.Close()
	erc20Ins, err := newERC20(e.contractAddress, client)
	if err != nil {
		return err
	}

	if addr.Hex() == InvalidAddr {
		log.Println("you try to approve an invalidAddr")
		return ErrInValAddr
	}
	if e.addr.Hex() == InvalidAddr {
		log.Println("from in approve is an invalidAddr")
		return ErrInValAddr
	}

	// need to determine whether the account balance is enough to approve.
	bal, err := e.BalanceOf(e.addr)
	if err != nil {
		return err
	}
	log.Println("Your balance is ", bal)
	if bal.Cmp(value) < 0 {
		log.Println("Balance is not enough, please reconfirm the approve amount.")
		return ErrBalNotE
	}

	log.Println("begin Approve", addr.Hex(), " with value", value, " in ERC20 contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(e.endPoint, e.hexSk, nil, e.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := erc20Ins.Approve(auth, addr, value)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("Approve Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(e.endPoint, tx, e.Status, "Approve")

	return nil
}

// TransferFrom The account represented by the e.hexsk transfer value of the specified amount from sender to recipient.
func (e *ContractModule) TransferFrom(sender, recipient common.Address, value *big.Int) error {
	client := getClient(e.endPoint)
	defer client.Close()
	erc20Ins, err := newERC20(e.contractAddress, client)
	if err != nil {
		return err
	}

	if sender.Hex() == InvalidAddr {
		log.Println("sender is an invalidAddr in transferFrom")
		return ErrInValAddr
	}
	if recipient.Hex() == InvalidAddr {
		log.Println("recipient is an invalidAddr in transferFrom")
		return ErrInValAddr
	}

	// need to determine whether the account allowance is enough to TransferFrom.
	allo, err := e.Allowance(sender, e.addr)
	if err != nil {
		return err
	}
	log.Println(sender.Hex(), "=>", e.addr.Hex(), "allowance is:", allo)
	if allo.Cmp(value) < 0 {
		log.Println("Allowance is not enough, please reconfirm the TransferFrom amount.")
		return ErrAlloNotE
	}

	// need to determine whether the sender balance is enough to transfer.
	bal, err := e.BalanceOf(sender)
	if err != nil {
		return err
	}
	log.Println("Sender ", sender.Hex(), " balance is", bal)
	if bal.Cmp(value) < 0 {
		log.Println("Sender balance is not enough.")
		return ErrBalNotE
	}

	log.Println("begin TransferFrom from", sender.Hex(), "to", recipient.Hex(), " with value", value, " in ERC20 contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(e.endPoint, e.hexSk, nil, e.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := erc20Ins.TransferFrom(auth, sender, recipient, value)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("TransferFrom Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(e.endPoint, tx, e.Status, "TransferFrom")

	return nil
}

// IncreaseAllowance The account represented by the e.hexsk increase the allowance for recipient.
func (e *ContractModule) IncreaseAllowance(recipient common.Address, value *big.Int) error {
	client := getClient(e.endPoint)
	defer client.Close()
	erc20Ins, err := newERC20(e.contractAddress, client)
	if err != nil {
		return err
	}

	if recipient.Hex() == InvalidAddr {
		log.Println("increaseAllowance: recipient is an invalidAddr")
		return ErrInValAddr
	}
	if e.addr.Hex() == InvalidAddr {
		log.Println("increaseAllowance: sender is an invalidAddr")
		return ErrInValAddr
	}

	// the given allowance shouldn't exceeds the balance
	bal, err := e.BalanceOf(e.addr)
	if err != nil {
		return err
	}
	allo, err := e.Allowance(e.addr, recipient)
	if err != nil {
		return err
	}
	sum := big.NewInt(0)
	sum.Add(value, allo)
	if bal.Cmp(sum) < 0 {
		log.Println(e.addr.Hex(), " balance is", bal, ", allowance sum is", sum)
		return errAllowanceExc
	}

	log.Println("begin IncreaseAllowance to", recipient.Hex(), " with value", value, " in ERC20 contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(e.endPoint, e.hexSk, nil, e.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := erc20Ins.IncreaseAllowance(auth, recipient, value)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("IncreaseAllowance Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(e.endPoint, tx, e.Status, "IncreaseAllowance")

	return nil
}

// DecreaseAllowance The account represented by the e.hexsk decrease the allowance for recipient.
func (e *ContractModule) DecreaseAllowance(recipient common.Address, value *big.Int) error {
	client := getClient(e.endPoint)
	defer client.Close()
	erc20Ins, err := newERC20(e.contractAddress, client)
	if err != nil {
		return err
	}

	if recipient.Hex() == InvalidAddr {
		log.Println("decreaseAllowance: recipient is an invalidAddr")
		return ErrInValAddr
	}
	if e.addr.Hex() == InvalidAddr {
		log.Println("increaseAllowance: sender is an invalidAddr")
		return ErrInValAddr
	}

	allo, err := e.Allowance(e.addr, recipient)
	if err != nil {
		return err
	}
	if allo.Cmp(value) < 0 {
		log.Println("the allowance is ", allo, ", but you want to decrease allowance ", value)
		return ErrAlloNotE
	}

	log.Println("begin DecreaseAllowance to", recipient.Hex(), " with value", value, " in ERC20 contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(e.endPoint, e.hexSk, nil, e.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := erc20Ins.DecreaseAllowance(auth, recipient, value)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("DecreaseAllowance Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(e.endPoint, tx, e.Status, "DecreaseAllowance")

	return nil
}

// MintToken The account represented by the e.hexsk mint token to target. Called by who has MINTER_ROLE.
func (e *ContractModule) MintToken(target common.Address, mintValue *big.Int, signs [5][]byte) error {
	client := getClient(e.endPoint)
	defer client.Close()
	erc20Ins, err := newERC20(e.contractAddress, client)
	if err != nil {
		return err
	}

	hasMinterRole, err := e.HasRole(uint8(1), e.addr)
	if err != nil {
		return err
	}
	if !hasMinterRole {
		return ErrNoMintRight
	}

	if target.Hex() == InvalidAddr {
		log.Println("mintToken: target is invalid")
		return ErrInValAddr
	}

	log.Println("begin MintToken to", target.Hex(), " with value", mintValue, " in ERC20 contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(e.endPoint, e.hexSk, nil, e.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := erc20Ins.MintToken(auth, target, mintValue, signs)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("MintToken Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(e.endPoint, tx, e.Status, "MintToken")

	return nil
}

// Burn The account represented by the e.hexsk burn it's balance. Called by who has DEFAULT_ADMIN_ROLE.
func (e *ContractModule) Burn(burnValue *big.Int) error {
	client := getClient(e.endPoint)
	defer client.Close()
	erc20Ins, err := newERC20(e.contractAddress, client)
	if err != nil {
		return err
	}

	hasAdminRole, err := e.HasRole(uint8(0), e.addr)
	if err != nil {
		return err
	}
	if !hasAdminRole {
		return ErrNoAdminRight
	}

	bal, err := e.BalanceOf(e.addr)
	if err != nil {
		return err
	}
	if bal.Cmp(burnValue) < 0 {
		log.Println("burn: your balance is ", bal, ", but you want to burn ", burnValue)
		return ErrBalNotE
	}

	log.Println("begin Burn with value", burnValue, " in ERC20 contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(e.endPoint, e.hexSk, nil, e.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := erc20Ins.Burn(auth, burnValue)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("Burn Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(e.endPoint, tx, e.Status, "Burn")

	return nil
}

// GetName get the name of erc20 token.
func (e *ContractModule) GetName() (string, error) {
	var name string

	client := getClient(e.endPoint)
	defer client.Close()
	erc20Ins, err := newERC20(e.contractAddress, client)
	if err != nil {
		return name, err
	}

	retryCount := 0
	for {
		retryCount++
		name, err = erc20Ins.GetName(&bind.CallOpts{
			From: e.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return name, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return name, nil
	}
}

// GetSymbol get the symbol of erc20 token.
func (e *ContractModule) GetSymbol() (string, error) {
	var name string

	client := getClient(e.endPoint)
	defer client.Close()
	erc20Ins, err := newERC20(e.contractAddress, client)
	if err != nil {
		return name, err
	}

	retryCount := 0
	for {
		retryCount++
		name, err = erc20Ins.GetSymbol(&bind.CallOpts{
			From: e.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return name, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return name, nil
	}
}

// GetDecimals get the decimals of erc20 token. For example, it's 18 in eth.
func (e *ContractModule) GetDecimals() (uint8, error) {
	var decimals uint8

	client := getClient(e.endPoint)
	defer client.Close()
	erc20Ins, err := newERC20(e.contractAddress, client)
	if err != nil {
		return decimals, err
	}

	retryCount := 0
	for {
		retryCount++
		decimals, err = erc20Ins.GetDecimals(&bind.CallOpts{
			From: e.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return decimals, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return decimals, nil
	}
}

// GetTotalSupply get the total supply of erc20 token.
func (e *ContractModule) GetTotalSupply() (*big.Int, error) {
	var totalSupply *big.Int

	client := getClient(e.endPoint)
	defer client.Close()
	erc20Ins, err := newERC20(e.contractAddress, client)
	if err != nil {
		return totalSupply, err
	}

	retryCount := 0
	for {
		retryCount++
		totalSupply, err = erc20Ins.GetTotalSupply(&bind.CallOpts{
			From: e.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return totalSupply, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return totalSupply, nil
	}
}

// GetMaxSupply get the maxSupply of erc20 token.
func (e *ContractModule) GetMaxSupply() (*big.Int, error) {
	ms := big.NewInt(0)

	client := getClient(e.endPoint)
	defer client.Close()
	erc20Ins, err := newERC20(e.contractAddress, client)
	if err != nil {
		return ms, err
	}

	retryCount := 0
	for {
		retryCount++
		ms, err = erc20Ins.MaxSupply(&bind.CallOpts{
			From: e.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return ms, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return ms, nil
	}
}

// GetVersion get the version of erc20 token.
func (e *ContractModule) GetVersion() (uint16, error) {
	var v uint16

	client := getClient(e.endPoint)
	defer client.Close()
	erc20Ins, err := newERC20(e.contractAddress, client)
	if err != nil {
		return v, err
	}

	retryCount := 0
	for {
		retryCount++
		v, err = erc20Ins.Version(&bind.CallOpts{
			From: e.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return v, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return v, nil
	}
}

// BalanceOf get the balance of addr in erc20 token.
func (e *ContractModule) BalanceOf(addr common.Address) (*big.Int, error) {
	var balance *big.Int

	client := getClient(e.endPoint)
	defer client.Close()
	erc20Ins, err := newERC20(e.contractAddress, client)
	if err != nil {
		return balance, err
	}

	retryCount := 0
	for {
		retryCount++
		balance, err = erc20Ins.BalanceOf(&bind.CallOpts{
			From: e.addr,
		}, addr)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return balance, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return balance, nil
	}
}

// Allowance get the allowance of sender to addr.
func (e *ContractModule) Allowance(sender, addr common.Address) (*big.Int, error) {
	var allowance *big.Int

	client := getClient(e.endPoint)
	defer client.Close()
	erc20Ins, err := newERC20(e.contractAddress, client)
	if err != nil {
		return allowance, err
	}

	retryCount := 0
	for {
		retryCount++
		allowance, err = erc20Ins.Allowance(&bind.CallOpts{
			From: e.addr,
		}, sender, addr)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return allowance, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return allowance, nil
	}
}

// MultiSigAddrs get the addresses that used to multi-signature.
func (e *ContractModule) MultiSigAddrs() ([5]common.Address, error) {
	addrs := [5]common.Address{}

	client := getClient(e.endPoint)
	defer client.Close()
	erc20Ins, err := newERC20(e.contractAddress, client)
	if err != nil {
		return addrs, err
	}

	for i := int64(0); i < 5; i++ {
		retryCount := 0
		for {
			retryCount++
			addrs[i], err = erc20Ins.Addrs(&bind.CallOpts{
				From: e.addr,
			}, big.NewInt(i))
			if err != nil {
				if retryCount > sendTransactionRetryCount {
					return addrs, err
				}
				time.Sleep(retryGetInfoSleepTime)
				continue
			}
			break
		}
	}
	return addrs, nil
}
