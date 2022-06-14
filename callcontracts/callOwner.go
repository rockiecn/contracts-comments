package callconts

import (
	"log"
	iface "memoc/interfaces"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// NewOwn new a instance of ContractModule. roleAddr: Role contract address
func NewOwn(roleAddr, addr common.Address, hexSk string, txopts *TxOpts, endPoint string, status chan error) iface.OwnerInfo {
	own := &ContractModule{
		addr:            addr,
		hexSk:           hexSk,
		txopts:          txopts,
		contractAddress: roleAddr,
		endPoint:        endPoint,
		Status:          status, // 用于接收：后台goroutine检查交易是否执行成功， nil代表成功
	}

	return own
}

// AlterOwner called by admin, to alter Role-contract's owner
func (own *ContractModule) AlterOwner(newOwnerAddr common.Address) error {
	client := getClient(own.endPoint)
	defer client.Close()
	roleIns, err := newRole(own.contractAddress, client)
	if err != nil {
		return err
	}

	if newOwnerAddr.Hex() == InvalidAddr {
		return ErrInValAddr
	}

	owner, err := own.GetOwner()
	if err != nil {
		return err
	}
	if owner.Hex() != own.addr.Hex() {
		log.Println("own.addr:", own.addr.Hex())
		return errNotOwner
	}

	log.Println("begin AlterOwner in Role contract...")

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(own.endPoint, own.hexSk, nil, own.txopts)
	if errMA != nil {
		return errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	tx, err := roleIns.AlterOwner(auth, newOwnerAddr)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("AlterOwner Err:", err)
		return err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(own.endPoint, tx, own.Status, "AlterOwner")

	return nil
}

// GetOwner get the owner-address of Role contract
// 'own.contractAddress' indicates the Role contract address
func (own *ContractModule) GetOwner() (common.Address, error) {
	var ownAddr common.Address

	client := getClient(own.endPoint)
	defer client.Close()
	roleIns, err := newRole(own.contractAddress, client)
	if err != nil {
		return ownAddr, err
	}

	log.Println("begin GetOwner in Role contract...")
	retryCount := 0

	for {
		retryCount++
		ownAddr, err = roleIns.GetOwner(&bind.CallOpts{
			From: own.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return ownAddr, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return ownAddr, nil
	}
}
