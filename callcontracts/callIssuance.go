package callconts

import (
	"log"
	"math/big"
	"memoc/contracts/issu"
	iface "memoc/interfaces"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// NewIssu new a instance of ContractModule. issuAddr: Issuance contract address
func NewIssu(issuAddr, addr common.Address, hexSk string, txopts *TxOpts, endPoint string, status chan error) iface.IssuanceInfo {
	issuance := &ContractModule{
		addr:            addr,
		hexSk:           hexSk,
		txopts:          txopts,
		contractAddress: issuAddr,
		endPoint:        endPoint,
		Status:          status, // 用于接收：后台goroutine检查交易是否执行成功， nil代表成功
	}

	return issuance
}

// DeployIssuance deploy an Issuance contract, called by admin
func (issuance *ContractModule) DeployIssuance(rolefsAddr common.Address) (common.Address, *issu.Issuance, error) {
	var issuAddr common.Address
	var issuIns *issu.Issuance

	log.Println("begin deploy Issuance contract...")
	client := getClient(issuance.endPoint)
	defer client.Close()

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(issuance.endPoint, issuance.hexSk, nil, issuance.txopts)
	if errMA != nil {
		return issuAddr, nil, errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	issuAddr, tx, issuIns, err := issu.DeployIssuance(auth, client, rolefsAddr)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("DeployIssuance Err:", err)
		return issuAddr, nil, err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(issuance.endPoint, tx, issuance.Status, "DeployIssuance")

	log.Println("Issuance address is ", issuAddr.Hex())
	return issuAddr, issuIns, nil
}

func newIssuance(issuAddr common.Address, client *ethclient.Client) (*issu.Issuance, error) {
	issuIns, err := issu.NewIssuance(issuAddr, client)
	if err != nil {
		return nil, err
	}
	return issuIns, nil
}

// MintLevel get mintLevel in Issuance contract
func (issuance *ContractModule) MintLevel() (*big.Int, error) {
	m := big.NewInt(0)

	client := getClient(issuance.endPoint)
	defer client.Close()
	issuIns, err := newIssuance(issuance.contractAddress, client)
	if err != nil {
		return m, err
	}

	retryCount := 0
	for {
		retryCount++
		m, err = issuIns.MintLevel(&bind.CallOpts{
			From: issuance.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return m, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return m, nil
	}
}

// LastMint get lastMint in Issuance contract
func (issuance *ContractModule) LastMint() (*big.Int, error) {
	m := big.NewInt(0)

	client := getClient(issuance.endPoint)
	defer client.Close()
	issuIns, err := newIssuance(issuance.contractAddress, client)
	if err != nil {
		return m, err
	}

	retryCount := 0
	for {
		retryCount++
		m, err = issuIns.LastMint(&bind.CallOpts{
			From: issuance.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return m, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return m, nil
	}
}

// Price get price in Issuance contract
func (issuance *ContractModule) Price() (*big.Int, error) {
	m := big.NewInt(0)

	client := getClient(issuance.endPoint)
	defer client.Close()
	issuIns, err := newIssuance(issuance.contractAddress, client)
	if err != nil {
		return m, err
	}

	retryCount := 0
	for {
		retryCount++
		m, err = issuIns.Price(&bind.CallOpts{
			From: issuance.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return m, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return m, nil
	}
}

// Size get size in Issuance contract
func (issuance *ContractModule) Size() (*big.Int, error) {
	m := big.NewInt(0)

	client := getClient(issuance.endPoint)
	defer client.Close()
	issuIns, err := newIssuance(issuance.contractAddress, client)
	if err != nil {
		return m, err
	}

	retryCount := 0
	for {
		retryCount++
		m, err = issuIns.Size(&bind.CallOpts{
			From: issuance.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return m, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return m, nil
	}
}

// SpaceTime get spaceTime in Issuance contract
func (issuance *ContractModule) SpaceTime() (*big.Int, error) {
	m := big.NewInt(0)

	client := getClient(issuance.endPoint)
	defer client.Close()
	issuIns, err := newIssuance(issuance.contractAddress, client)
	if err != nil {
		return m, err
	}

	retryCount := 0
	for {
		retryCount++
		m, err = issuIns.SpaceTime(&bind.CallOpts{
			From: issuance.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return m, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return m, nil
	}
}

// TotalPay get totalPay in Issuance contract
func (issuance *ContractModule) TotalPay() (*big.Int, error) {
	m := big.NewInt(0)

	client := getClient(issuance.endPoint)
	defer client.Close()
	issuIns, err := newIssuance(issuance.contractAddress, client)
	if err != nil {
		return m, err
	}

	retryCount := 0
	for {
		retryCount++
		m, err = issuIns.TotalPay(&bind.CallOpts{
			From: issuance.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return m, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return m, nil
	}
}

// TotalPaid get totalPaid in Issuance contract
func (issuance *ContractModule) TotalPaid() (*big.Int, error) {
	m := big.NewInt(0)

	client := getClient(issuance.endPoint)
	defer client.Close()
	issuIns, err := newIssuance(issuance.contractAddress, client)
	if err != nil {
		return m, err
	}

	retryCount := 0
	for {
		retryCount++
		m, err = issuIns.TotalPaid(&bind.CallOpts{
			From: issuance.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return m, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return m, nil
	}
}

// PeriodTarget get periodTarget in Issuance contract
func (issuance *ContractModule) PeriodTarget() (*big.Int, error) {
	m := big.NewInt(0)

	client := getClient(issuance.endPoint)
	defer client.Close()
	issuIns, err := newIssuance(issuance.contractAddress, client)
	if err != nil {
		return m, err
	}

	retryCount := 0
	for {
		retryCount++
		m, err = issuIns.PeriodTarget(&bind.CallOpts{
			From: issuance.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return m, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}
		return m, nil
	}
}

// PeriodTotalReward get periodTotalReward in Issuance contract
func (issuance *ContractModule) PeriodTotalReward() (*big.Int, error) {
	m := big.NewInt(0)

	client := getClient(issuance.endPoint)
	defer client.Close()
	issuIns, err := newIssuance(issuance.contractAddress, client)
	if err != nil {
		return m, err
	}

	retryCount := 0
	for {
		retryCount++
		m, err = issuIns.PeriodTotalReward(&bind.CallOpts{
			From: issuance.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return m, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}
		return m, nil
	}
}

// IssuRatio get issuRatio in Issuance contract
func (issuance *ContractModule) IssuRatio() (uint16, error) {
	ir := uint16(0)

	client := getClient(issuance.endPoint)
	defer client.Close()
	issuIns, err := newIssuance(issuance.contractAddress, client)
	if err != nil {
		return ir, err
	}

	retryCount := 0
	for {
		retryCount++
		ir, err = issuIns.IssuRatio(&bind.CallOpts{
			From: issuance.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return ir, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}
		return ir, nil
	}
}

// MinRatio get minRatio in Issuance contract
func (issuance *ContractModule) MinRatio() (uint16, error) {
	mr := uint16(0)

	client := getClient(issuance.endPoint)
	defer client.Close()
	issuIns, err := newIssuance(issuance.contractAddress, client)
	if err != nil {
		return mr, err
	}

	retryCount := 0
	for {
		retryCount++
		mr, err = issuIns.MinRatio(&bind.CallOpts{
			From: issuance.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return mr, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}
		return mr, nil
	}
}
