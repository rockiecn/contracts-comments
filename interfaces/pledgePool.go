package iface

import (
	"math/big"
	"memoc/contracts/pledgepool"

	"github.com/ethereum/go-ethereum/common"
)

// PledgePoolInfo contains operations related to pledge on chain
type PledgePoolInfo interface {
	// primeToken、rToken、role
	DeployPledgePool(common.Address, common.Address, common.Address) (common.Address, *pledgepool.PledgePool, error)
	Pledge(common.Address, common.Address, uint64, *big.Int, []byte) error
	Withdraw(common.Address, common.Address, uint64, uint32, *big.Int, []byte) error
	GetPledge(uint32) (*big.Int, error)
	GetBalanceInPPool(uint64, uint32) (*big.Int, error)
	TotalPledge() (*big.Int, error)
}
