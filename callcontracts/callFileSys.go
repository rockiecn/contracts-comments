package callconts

import (
	"log"
	"math/big"
	filesys "memoc/contracts/filesystem"
	iface "memoc/interfaces"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// NewFileSys new a instance of ContractModule, fsAddr:FileSys contract address
func NewFileSys(fsAddr, addr common.Address, hexSk string, txopts *TxOpts, endPoint string, status chan error) iface.FileSysInfo {
	fs := &ContractModule{
		addr:            addr,
		hexSk:           hexSk,
		txopts:          txopts,
		contractAddress: fsAddr,
		endPoint:        endPoint,
		Status:          status, // 用于接收：后台goroutine检查交易是否执行成功， nil代表成功
	}

	return fs
}

// newFileSys new an instance of FileSys contract, 'fsAddr' indicates FileSys contract address.
func newFileSys(fsAddr common.Address, client *ethclient.Client) (*filesys.FileSys, error) {
	fsIns, err := filesys.NewFileSys(fsAddr, client)
	if err != nil {
		return nil, err
	}
	return fsIns, nil
}

// DeployFileSys deploy a FileSys contract, called by admin.
// Called after the admin calls the CreateGroup function in the Role Contract.
// 'r' indicates role-contract address, 'rfs' indicates RoleFS-contract address.
func (fs *ContractModule) DeployFileSys(founder, gIndex uint64, r, rfs common.Address, keepers []uint64) (common.Address, *filesys.FileSys, error) {
	var fsAddr common.Address
	var fsIns *filesys.FileSys

	log.Println("begin deploy FileSys contract...")
	client := getClient(fs.endPoint)
	defer client.Close()

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(fs.endPoint, fs.hexSk, nil, fs.txopts)
	if errMA != nil {
		return fsAddr, nil, errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	fsAddr, tx, fsIns, err := filesys.DeployFileSys(auth, client, founder, gIndex, r, rfs, keepers)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("DeployFileSys Err:", err)
		return fsAddr, nil, err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	df := make(chan error)
	go checkTx(fs.endPoint, tx, df, "DeployFileSys")

	// NOTE： 此处需等待checkTx执行完毕,从而后续执行SetGF函数
	err = <-df
	if err != nil {
		return fsAddr, fsIns, err
	}

	log.Println("FileSys address is ", fsAddr.Hex())
	return fsAddr, fsIns, nil
}

// GetFsInfo Get information of filesystem. return repairFs info when uIndex is 0
func (fs *ContractModule) GetFsInfo(uIndex uint64) (bool, uint32, error) {
	var isActive bool
	var tokenIndex uint32

	client := getClient(fs.endPoint)
	defer client.Close()
	fsIns, err := newFileSys(fs.contractAddress, client)
	if err != nil {
		return isActive, tokenIndex, err
	}

	retryCount := 0
	for {
		retryCount++
		// multiple return values are returned as a structure because named return values are used in the contract function
		isActive, err = fsIns.GetFsInfo(&bind.CallOpts{
			From: fs.addr,
		}, uIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return isActive, tokenIndex, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return isActive, tokenIndex, nil
	}
}

// GetFsProviderSum get providers sum in fs. return repairFs info when uIndex is 0
func (fs *ContractModule) GetFsProviderSum(uIndex uint64) (uint64, error) {
	var pSum uint64

	client := getClient(fs.endPoint)
	defer client.Close()
	fsIns, err := newFileSys(fs.contractAddress, client)
	if err != nil {
		return pSum, err
	}

	retryCount := 0
	for {
		retryCount++
		pSum, err = fsIns.GetFsPSum(&bind.CallOpts{
			From: fs.addr,
		}, uIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return pSum, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return pSum, nil
	}
}

// GetFsProvider get pIndex in fs by array index. return repairFs info when uIndex is 0
func (fs *ContractModule) GetFsProvider(uIndex uint64, index uint64) (uint64, error) {
	var pIndex uint64

	client := getClient(fs.endPoint)
	defer client.Close()
	fsIns, err := newFileSys(fs.contractAddress, client)
	if err != nil {
		return pIndex, err
	}

	retryCount := 0
	for {
		retryCount++
		pIndex, err = fsIns.GetFsPro(&bind.CallOpts{
			From: fs.addr,
		}, uIndex, index)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return pIndex, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return pIndex, nil
	}
}

// GetFsInfoAggOrder Get information of aggOrder. return repairFs info when uIndex is 0
func (fs *ContractModule) GetFsInfoAggOrder(uIndex uint64, pIndex uint64) (uint64, uint64, error) {
	var nonce uint64
	var subNonce uint64

	client := getClient(fs.endPoint)
	defer client.Close()
	fsIns, err := newFileSys(fs.contractAddress, client)
	if err != nil {
		return nonce, subNonce, err
	}

	retryCount := 0
	for {
		retryCount++
		nonce, subNonce, err = fsIns.GetFsInfoAggOrder(&bind.CallOpts{
			From: fs.addr,
		}, uIndex, pIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return nonce, subNonce, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return nonce, subNonce, nil
	}
}

// GetStoreInfo Get information of storage order. return repairFs info when uIndex is 0
func (fs *ContractModule) GetStoreInfo(uIndex uint64, pIndex uint64, tIndex uint32) (uint64, uint64, *big.Int, error) {
	var _time, size uint64
	var price *big.Int

	client := getClient(fs.endPoint)
	defer client.Close()
	fsIns, err := newFileSys(fs.contractAddress, client)
	if err != nil {
		return _time, size, price, err
	}

	retryCount := 0
	for {
		retryCount++
		_time, size, price, err = fsIns.GetStoreInfo(
			&bind.CallOpts{From: fs.addr},
			uIndex,
			pIndex,
			tIndex,
		)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return _time, size, price, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return _time, size, price, nil
	}
}

// GetChannelInfo Get information of channel. return repairFs info when uIndex is 0
func (fs *ContractModule) GetChannelInfo(uIndex uint64, pIndex uint64, tIndex uint32) (*big.Int, uint64, uint64, error) {
	var nonce, expire uint64
	var amount *big.Int

	client := getClient(fs.endPoint)
	defer client.Close()
	fsIns, err := newFileSys(fs.contractAddress, client)
	if err != nil {
		return amount, nonce, expire, err
	}

	retryCount := 0
	for {
		retryCount++
		amount, nonce, expire, err = fsIns.GetChannelInfo(
			&bind.CallOpts{From: fs.addr},
			uIndex,
			pIndex,
			tIndex,
		)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return amount, nonce, expire, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return amount, nonce, expire, nil
	}
}

// GetSettleInfo Get information of settlement.
func (fs *ContractModule) GetSettleInfo(pIndex uint64, tIndex uint32) (uint64, uint64, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var _time, size uint64
	var price, maxPay, hasPaid, canPay, lost, lostPaid, managePay, endPaid, linearPaid *big.Int

	client := getClient(fs.endPoint)
	defer client.Close()
	fsIns, err := newFileSys(fs.contractAddress, client)
	if err != nil {
		return _time, size, price, maxPay, hasPaid, canPay, lost, lostPaid, managePay, endPaid, linearPaid, err
	}

	retryCount := 0
	for {
		retryCount++
		_time, size, price, maxPay, hasPaid, canPay, lost, lostPaid, managePay, endPaid, linearPaid, err = fsIns.GetSettleInfo(&bind.CallOpts{
			From: fs.addr,
		}, pIndex, tIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return _time, size, price, maxPay, hasPaid, canPay, lost, lostPaid, managePay, endPaid, linearPaid, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return _time, size, price, maxPay, hasPaid, canPay, lost, lostPaid, managePay, endPaid, linearPaid, nil
	}
}

// GetBalance Get income balance of rindex.
// rIndex:0 -> foundation
func (fs *ContractModule) GetBalance(rIndex uint64, tIndex uint32) (*big.Int, *big.Int, error) {
	var avail, tmp *big.Int

	client := getClient(fs.endPoint)
	defer client.Close()
	fsIns, err := newFileSys(fs.contractAddress, client)
	if err != nil {
		return avail, tmp, err
	}

	retryCount := 0
	for {
		retryCount++
		avail, tmp, err = fsIns.GetBalance(&bind.CallOpts{
			From: fs.addr,
		}, rIndex, tIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return avail, tmp, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return avail, tmp, nil
	}
}
