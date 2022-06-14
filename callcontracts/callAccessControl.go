package callconts

import (
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// SetUpRole Called by who has DEFAULT_ADMIN_ROLE. Set role to addr.
// role: DEFAULT_ADMIN_ROLE(0)、MINTER_ROLE(1)、PAUSER_ROLE(2)
func (ac *ContractModule) SetUpRole(role uint8, addr common.Address) error {
	client := getClient(ac.endPoint)
	defer client.Close()
	acIns, err := newERC20(ac.contractAddress, client)
	if err != nil {
		return err
	}

	hasAdmin, err := ac.HasRole(uint8(0), ac.addr)
	if err != nil {
		return err
	}
	if !hasAdmin {
		return ErrNoAdminRight
	}

	if role > 2 {
		return errAccessControlRole
	}

	if addr.Hex() == InvalidAddr {
		return ErrInValAddr
	}

	log.Println("begin SetUpRole to", addr.Hex(), " with role", role, " in AccessControl contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(ac.endPoint, ac.hexSk, nil, ac.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := acIns.SetUpRole(auth, role, addr)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("SetUpRole Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(ac.endPoint, tx, ac.Status, "SetUpRole in AccessControl")

	return nil
}

// RevokeRole Called by who has DEFAULT_ADMIN_ROLE. Revoke other account's role.
func (ac *ContractModule) RevokeRole(role uint8, addr common.Address) error {
	client := getClient(ac.endPoint)
	defer client.Close()
	acIns, err := newERC20(ac.contractAddress, client)
	if err != nil {
		return err
	}

	hasAdmin, err := ac.HasRole(uint8(0), ac.addr)
	if err != nil {
		return err
	}
	if !hasAdmin {
		return ErrNoAdminRight
	}

	if role > 2 {
		return errAccessControlRole
	}

	if addr.Hex() == InvalidAddr {
		return ErrInValAddr
	}

	log.Println("begin RevokeRole to", addr.Hex(), " with role", role, " in AccessControl contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(ac.endPoint, ac.hexSk, nil, ac.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := acIns.RevokeRole(auth, role, addr)

	if err != nil {
		log.Println("RevokeRole Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")

	go checkTx(ac.endPoint, tx, ac.Status, "RevokeRole in AccessControl")

	return nil
}

// RenounceRole Account renounce its role .
func (ac *ContractModule) RenounceRole(role uint8) error {
	client := getClient(ac.endPoint)
	defer client.Close()
	acIns, err := newERC20(ac.contractAddress, client)
	if err != nil {
		return err
	}

	log.Println("begin RenounceRole", role, " in AccessControl contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(ac.endPoint, ac.hexSk, nil, ac.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := acIns.RenounceRole(auth, role)

	if err != nil {
		log.Println("RenounceRole Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")

	go checkTx(ac.endPoint, tx, ac.Status, "RenounceRole in AccessControl")

	return nil
}

// Pause Set to true to prohibit transfer operation in erc20. Called by who has PAUSER_ROLE.
func (ac *ContractModule) Pause() error {
	client := getClient(ac.endPoint)
	defer client.Close()
	acIns, err := newERC20(ac.contractAddress, client)
	if err != nil {
		return err
	}

	hasPause, err := ac.HasRole(uint8(2), ac.addr)
	if err != nil {
		return err
	}
	if !hasPause {
		return ErrNoPauseRight
	}

	log.Println("begin Pause in AccessControl contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(ac.endPoint, ac.hexSk, nil, ac.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := acIns.Pause(auth)

	if err != nil {
		log.Println("Pause Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")

	go checkTx(ac.endPoint, tx, ac.Status, "Pause in AccessControl")

	return nil
}

// Unpause Set to false to allow transfer operation in erc20. Called by who has PAUSER_ROLE.
func (ac *ContractModule) Unpause() error {
	client := getClient(ac.endPoint)
	defer client.Close()
	acIns, err := newERC20(ac.contractAddress, client)
	if err != nil {
		return err
	}

	hasPause, err := ac.HasRole(uint8(2), ac.addr)
	if err != nil {
		return err
	}
	if !hasPause {
		return ErrNoPauseRight
	}

	log.Println("begin Unpause in AccessControl contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(ac.endPoint, ac.hexSk, nil, ac.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := acIns.Unpause(auth)

	if err != nil {
		log.Println("Unpause Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")

	go checkTx(ac.endPoint, tx, ac.Status, "Unpause in AccessControl")

	return nil
}

// HasRole Check whether the account has the right.
func (ac *ContractModule) HasRole(role uint8, addr common.Address) (bool, error) {
	var has bool

	client := getClient(ac.endPoint)
	defer client.Close()
	acIns, err := newERC20(ac.contractAddress, client)
	if err != nil {
		return has, err
	}

	retryCount := 0
	for {
		retryCount++
		has, err = acIns.HasRole(&bind.CallOpts{
			From: ac.addr,
		}, role, addr)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return has, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return has, nil
	}
}

// GetPaused Check whether the account has the right.
func (ac *ContractModule) GetPaused() (bool, error) {
	var isPaused bool

	client := getClient(ac.endPoint)
	defer client.Close()
	acIns, err := newERC20(ac.contractAddress, client)
	if err != nil {
		return isPaused, err
	}

	retryCount := 0
	for {
		retryCount++
		isPaused, err = acIns.GetPaused(&bind.CallOpts{
			From: ac.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return isPaused, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return isPaused, nil
	}
}
