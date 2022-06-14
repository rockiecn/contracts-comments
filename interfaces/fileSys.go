package iface

import (
	"math/big"
	filesys "memoc/contracts/filesystem"

	"github.com/ethereum/go-ethereum/common"
)

// FileSysInfo contains operations related to payment about storage order
type FileSysInfo interface {
	DeployFileSys(founder uint64, gIndex uint64, r common.Address, rfs common.Address, keepers []uint64) (common.Address, *filesys.FileSys, error)

	GetFsInfo(uIndex uint64) (isActive bool, tokenIndex uint32, err error)
	GetFsProviderSum(uIndex uint64) (pSum uint64, err error)
	GetFsProvider(uIndex uint64, index uint64) (pIndex uint64, err error)
	GetFsInfoAggOrder(uIndex uint64, provider uint64) (uint64, uint64, error)
	GetStoreInfo(uIndex uint64, pIndex uint64, tIndex uint32) (uint64, uint64, *big.Int, error)
	GetChannelInfo(uIndex uint64, pIndex uint64, tIndex uint32) (*big.Int, uint64, uint64, error)
	GetSettleInfo(pIndex uint64, tIndex uint32) (uint64, uint64, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error)
	GetBalance(rIndex uint64, tIndex uint32) (*big.Int, *big.Int, error)
}
