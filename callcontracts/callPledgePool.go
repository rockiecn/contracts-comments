// only prepare to complete calling basic functions of contract now, after that, will judge the input parameters of functions
// And also need to add Getter functions.
// Also need to add 'Approve' function in some functions related to transfer operations.

package callconts

import (
	"log"
	"math/big"
	"memoc/contracts/pledgepool"
	iface "memoc/interfaces"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// NewPledgePool new an instance of ContractModule. 'pledgePoolAddr' indicates PledgePool contract address.
func NewPledgePool(pledgePoolAddr, addr common.Address, hexSk string, txopts *TxOpts, endPoint string, status chan error) iface.PledgePoolInfo {
	p := &ContractModule{
		addr:            addr,
		hexSk:           hexSk,
		txopts:          txopts,
		contractAddress: pledgePoolAddr,
		endPoint:        endPoint,
		Status:          status, // 用于接收：后台goroutine检查交易是否执行成功， nil代表成功
	}

	return p
}

// newPledgePool new an instance of PledgePool contract, 'pledgepAddr' indicates PledgePool contract address.
func newPledgePool(pledgepAddr common.Address, client *ethclient.Client) (*pledgepool.PledgePool, error) {
	ppIns, err := pledgepool.NewPledgePool(pledgepAddr, client)
	if err != nil {
		return nil, err
	}
	return ppIns, nil
}

// DeployPledgePool deploy a PledgePool contract, called by admin.
// primeToken、rToken contract address、role contract address.
func (p *ContractModule) DeployPledgePool(primeToken common.Address, rToken common.Address, role common.Address) (common.Address, *pledgepool.PledgePool, error) {
	var pledgepAddr common.Address
	var pledgepIns *pledgepool.PledgePool

	log.Println("begin deploy PledgePool contract...")
	client := getClient(p.endPoint)
	defer client.Close()

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(p.endPoint, p.hexSk, nil, p.txopts)
	if errMA != nil {
		return pledgepAddr, pledgepIns, errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	pledgepAddr, tx, pledgepIns, err := pledgepool.DeployPledgePool(auth, client, primeToken, rToken, role)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("DeployPledgePool Err:", err)
		return pledgepAddr, pledgepIns, err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(p.endPoint, tx, p.Status, "DeployPledgePool")

	log.Println("PledgePool address is ", pledgepAddr.Hex())
	return pledgepAddr, pledgepIns, nil
}

// Pledge money.
// Called by the account itself or by another account on its behalf.
// 调用前需要index指示的账户approve本合约（也就是pledgePool合约）账户指定的金额（也就是value）,如果是账户本身调用，则会由代码自动approve
func (p *ContractModule) Pledge(erc20Addr, roleAddr common.Address, rindex uint64, value *big.Int, sign []byte) error {
	client := getClient(p.endPoint)
	defer client.Close()
	pledgepIns, err := newPledgePool(p.contractAddress, client)
	if err != nil {
		return err
	}

	r := NewR(roleAddr, p.addr, p.hexSk, p.txopts, p.endPoint, p.Status)
	addr, err := r.GetAddr(rindex)
	if err != nil {
		return err
	}

	// check the rindex is not being banned
	_, isBanned, _, _, _, _, err := r.GetRoleInfo(addr)
	if err != nil {
		return err
	}
	if isBanned {
		log.Println("pledge: account with rindex is banned")
		return ErrIsBanned
	}

	// check whether the allowance[addr][pledgePoolAddr] is not less than value, if not, will approve automatically by code.
	e := NewERC20(erc20Addr, p.addr, p.hexSk, p.txopts, p.endPoint, p.Status)
	allo, err := e.Allowance(addr, p.contractAddress)
	if err != nil {
		return err
	}
	if allo.Cmp(value) < 0 {
		tmp := big.NewInt(0)
		tmp.Sub(value, allo)
		log.Println("The allowance of ", addr.Hex(), " to ", p.contractAddress.Hex(), " is not enough, also need to add allowance", tmp)
		// if called by the account itself， then call IncreaseAllowance directly.
		if sign == nil && p.addr.Hex() == addr.Hex() {
			err = e.IncreaseAllowance(p.contractAddress, tmp)
			if err != nil {
				return err
			}
			if err = <-p.Status; err != nil {
				return err
			}
		} else { // otherwise, quit Pledge
			log.Println("pledge: call pledge by other account, but the allowance is not enough")
			return ErrAlloNotE
		}
	}

	// check balance
	bal, err := e.BalanceOf(addr)
	if err != nil {
		return err
	}
	if bal.Cmp(value) < 0 {
		log.Println("pledge: addr balance is ", bal, " is not enough to pledge ", value)
		return ErrBalNotE
	}

	log.Println("begin Pledge in PledgePool contract with value", value, " and rindex", rindex, " ...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(p.endPoint, p.hexSk, nil, p.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := pledgepIns.Pledge(auth, rindex, value, sign)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("Pledge Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(p.endPoint, tx, p.Status, "Pledge")

	return nil
}

// Withdraw Called by the account itself or by another account on its behalf.
// withdraw its balance from PledgePool.
func (p *ContractModule) Withdraw(roleAddr, rTokenAddr common.Address, rindex uint64, tindex uint32, value *big.Int, sign []byte) error {
	client := getClient(p.endPoint)
	defer client.Close()
	pledgepIns, err := newPledgePool(p.contractAddress, client)
	if err != nil {
		return err
	}

	// check if rindex is banned
	r := NewR(roleAddr, p.addr, p.hexSk, p.txopts, p.endPoint, p.Status)
	addr, err := r.GetAddr(rindex)
	if err != nil {
		return err
	}
	_, isBanned, _, _, _, _, err := r.GetRoleInfo(addr)
	if err != nil {
		return err
	}
	if isBanned {
		log.Println("withdraw: account with rindex ", rindex, " is banned")
		return ErrIsBanned
	}
	// check if tindex is valid
	rt := NewRT(rTokenAddr, p.addr, p.hexSk, p.txopts, p.endPoint)
	isValid, err := rt.IsValid(tindex)
	if err != nil {
		return err
	}
	if !isValid {
		return ErrTIndex
	}

	log.Println("begin Withdraw in PledgePool contract with value", value, " and rindex", rindex, " and tindex", tindex, " ...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(p.endPoint, p.hexSk, nil, p.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := pledgepIns.Withdraw(auth, rindex, tindex, value, sign)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("Withdraw Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(p.endPoint, tx, p.Status, "Withdraw")

	return nil
}

// GetPledge Get all pledge amount in specified token.
func (p *ContractModule) GetPledge(tindex uint32) (*big.Int, error) {
	var amount *big.Int

	client := getClient(p.endPoint)
	defer client.Close()
	pledgepIns, err := newPledgePool(p.contractAddress, client)
	if err != nil {
		return amount, err
	}

	retryCount := 0
	for {
		retryCount++
		amount, err = pledgepIns.GetPledge(&bind.CallOpts{
			From: p.addr,
		}, tindex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return amount, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return amount, nil
	}
}

// GetBalanceInPPool Get balance of the account related rindex in specified token.
func (p *ContractModule) GetBalanceInPPool(rindex uint64, tindex uint32) (*big.Int, error) {
	var amount *big.Int

	client := getClient(p.endPoint)
	defer client.Close()
	pledgepIns, err := newPledgePool(p.contractAddress, client)
	if err != nil {
		return amount, err
	}

	retryCount := 0
	for {
		retryCount++
		amount, err = pledgepIns.GetBalance(&bind.CallOpts{
			From: p.addr,
		}, rindex, tindex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return amount, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return amount, nil
	}
}

// TotalPledge Get all pledge amount in specified token.
func (p *ContractModule) TotalPledge() (*big.Int, error) {
	var amount *big.Int

	client := getClient(p.endPoint)
	defer client.Close()
	pledgepIns, err := newPledgePool(p.contractAddress, client)
	if err != nil {
		return amount, err
	}

	retryCount := 0
	for {
		retryCount++
		amount, err = pledgepIns.TotalPledge(&bind.CallOpts{
			From: p.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return amount, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return amount, nil
	}
}
