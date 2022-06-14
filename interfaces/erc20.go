package iface

import (
	"math/big"

	"memoc/contracts/erc20"

	"github.com/ethereum/go-ethereum/common"
)

// ERC20Info contains information about memo primary token
type ERC20Info interface {
	DeployERC20(string, string, uint16, [5]common.Address) (common.Address, *erc20.ERC20, error)
	GetName() (string, error)
	GetSymbol() (string, error)
	GetDecimals() (uint8, error)
	GetTotalSupply() (*big.Int, error)
	GetMaxSupply() (*big.Int, error)
	GetVersion() (uint16, error)
	MultiSigAddrs() ([5]common.Address, error)
	BalanceOf(common.Address) (*big.Int, error)
	Allowance(common.Address, common.Address) (*big.Int, error)

	Transfer(common.Address, *big.Int) error
	Approve(common.Address, *big.Int) error
	TransferFrom(common.Address, common.Address, *big.Int) error
	IncreaseAllowance(common.Address, *big.Int) error
	DecreaseAllowance(common.Address, *big.Int) error
	MintToken(common.Address, *big.Int, [5][]byte) error
	Burn(*big.Int) error

	// AccessControl
	SetUpRole(uint8, common.Address) error
	RevokeRole(uint8, common.Address) error
	RenounceRole(uint8) error
	Pause() error
	Unpause() error

	GetPaused() (bool, error)
	HasRole(uint8, common.Address) (bool, error)
}
