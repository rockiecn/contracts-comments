package iface

import (
	"math/big"
	"memoc/contracts/issu"
	"memoc/contracts/role"
	"memoc/contracts/rolefs"

	"github.com/ethereum/go-ethereum/common"
)

// OwnerInfo contains information about the owner of the Role contract
type OwnerInfo interface {
	// called by owner
	AlterOwner(common.Address) error

	// called by anyone
	GetOwner() (common.Address, error)
}

// RoleInfo contains operations related to memo roles
type RoleInfo interface {
	// deploy Role contract. specify foundation、primaryToken、pledgeK、pledgeP
	DeployRole(common.Address, common.Address, *big.Int, *big.Int, uint16) (common.Address, *role.Role, error)

	// callled by owner, set pledgePool-address、issuance-address and rolefs-address
	SetPI(common.Address, common.Address, common.Address) error

	// called by anyone to get index
	Register(common.Address, []byte) error

	// called by anyone to register Keeper
	RegisterKeeper(common.Address, uint64, []byte, []byte) error

	// called by anyone to register Provider
	RegisterProvider(common.Address, uint64, []byte) error

	// called by anyone to register User
	RegisterUser(common.Address, uint64, uint64, []byte, []byte) error

	// called by owner to add token
	RegisterToken(common.Address) error

	// called by owner
	CreateGroup(common.Address, uint64, []uint64, uint16) (uint64, error)

	// called by owner to set fsAddress for group after CreateGroup and deployFileSys
	SetGF(common.Address, uint64) error

	// called by owner
	AddKeeperToGroup(uint64, uint64) error

	// called by provider or called by others
	AddProviderToGroup(uint64, uint64, []byte) error

	// called by owner to set the amount that the keeper and provider needs to pledge
	SetPledgeMoney(*big.Int, *big.Int) error

	// called by user or called by others
	Recharge(common.Address, uint64, uint32, *big.Int, []byte) error

	// called by memo-role or called by others
	WithdrawFromFs(common.Address, uint64, uint32, *big.Int, []byte) error

	PledgePool() (common.Address, error)
	Foundation() (common.Address, error)
	PledgeK() (*big.Int, error)
	PledgeP() (*big.Int, error)
	RToken() (common.Address, error)
	Issuance() (common.Address, error)
	Rolefs() (common.Address, error)

	// get the number of registered addresses
	GetAddrsNum() (uint64, error)

	// get address by role index
	GetAddr(uint64) (common.Address, error)

	// get the account role index by address
	GetRoleIndex(common.Address) (uint64, error)

	// get account information by address
	GetRoleInfo(common.Address) (bool, bool, uint8, uint64, uint64, []byte, error)

	// get the number of groups
	GetGroupsNum() (uint64, error)

	// get group information by gIndex
	GetGroupInfo(uint64) (bool, bool, bool, uint16, *big.Int, *big.Int, common.Address, error)

	// get address and gIndex by array index value
	GetAddrGindex(uint64) (common.Address, uint64, error)

	// get the number of keepers in the group
	GetGKNum(uint64) (uint64, error)

	// get the number of providers in the group
	GetGUPNum(uint64) (uint64, uint64, error)

	// get keeper role index by group array index and keeper array index
	GetGroupK(uint64, uint64) (uint64, error)

	// get provider role index by group array index and provider array index
	GetGroupP(uint64, uint64) (uint64, error)

	// get user role index by group array index and provider array index
	GetGroupU(uint64, uint64) (uint64, error)

	GetRVersion() (uint16, error)
}

// RoleFSInfo contains operations related to memo roles and filesystem-payment
type RoleFSInfo interface {
	DeployRoleFS() (common.Address, *rolefs.RoleFS, error)
	// called by owner, which is the deployer
	SetAddr(common.Address, common.Address, common.Address) error
	AddOrder(common.Address, common.Address, uint64, uint64, uint64, uint64, uint64, uint64, uint32, *big.Int, []byte, []byte) error
	SubOrder(common.Address, common.Address, uint64, uint64, uint64, uint64, uint64, uint64, uint32, *big.Int, []byte, []byte) error
	AddRepair(common.Address, common.Address, uint64, uint64, uint64, uint64, uint64, uint64, uint32, *big.Int, []byte, [][]byte) error
	SubRepair(common.Address, common.Address, uint64, uint64, uint64, uint64, uint64, uint64, uint32, *big.Int, []byte, [][]byte) error
	ProWithdraw(common.Address, common.Address, uint64, uint32, *big.Int, *big.Int, []uint64, [][]byte) error
}

// RTokenInfo contains operations related to tokens that memo supported
type RTokenInfo interface {
	IsValid(uint32) (bool, error)
	GetTA(uint32) (common.Address, error)
	GetTI(common.Address) (uint32, bool, error)
	GetTNum() (uint32, error)
}

// IssuanceInfo contains deploy Isuance-contract function
type IssuanceInfo interface {
	DeployIssuance(rfs common.Address) (common.Address, *issu.Issuance, error)
	MintLevel() (*big.Int, error)
	LastMint() (*big.Int, error)
	Price() (*big.Int, error)
	Size() (*big.Int, error)
	SpaceTime() (*big.Int, error)
	TotalPay() (*big.Int, error)
	TotalPaid() (*big.Int, error)
	PeriodTarget() (*big.Int, error)
	PeriodTotalReward() (*big.Int, error)
	IssuRatio() (uint16, error)
	MinRatio() (uint16, error)
}
