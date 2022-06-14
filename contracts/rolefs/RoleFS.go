// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rolefs

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AOParams is an auto generated low-level Go binding around an user-defined struct.
type AOParams struct {
	UIndex uint64
	PIndex uint64
	Start  uint64
	End    uint64
	Size   uint64
	Nonce  uint64
	TIndex uint32
	SPrice *big.Int
	Usign  []byte
	Psign  []byte
}

// IssuParams is an auto generated low-level Go binding around an user-defined struct.
type IssuParams struct {
	Start  uint64
	End    uint64
	Size   uint64
	SPrice *big.Int
}

// PWParams is an auto generated low-level Go binding around an user-defined struct.
type PWParams struct {
	PIndex   uint64
	TIndex   uint32
	PAddr    common.Address
	TAddr    common.Address
	Pay      *big.Int
	Lost     *big.Int
	KIndexes []uint64
	Ksigns   [][]byte
}

// SGParams is an auto generated low-level Go binding around an user-defined struct.
type SGParams struct {
	Index  uint64
	IsAdd  bool
	Size   *big.Int
	SPrice *big.Int
}

// SOParams is an auto generated low-level Go binding around an user-defined struct.
type SOParams struct {
	KIndex uint64
	UIndex uint64
	PIndex uint64
	Start  uint64
	End    uint64
	Size   uint64
	Nonce  uint64
	TIndex uint32
	SPrice *big.Int
	Usign  []byte
	Psign  []byte
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"mintToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"42966c68": "burn(uint256)",
		"a457c2d7": "decreaseAllowance(address,uint256)",
		"f0141d84": "getDecimals()",
		"17d7de7c": "getName()",
		"15070401": "getSymbol()",
		"c4e41b22": "getTotalSupply()",
		"39509351": "increaseAllowance(address,uint256)",
		"6bd1ce89": "mintToken(address,uint256,bytes[5])",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// GetDecimals is a free data retrieval call binding the contract method 0xf0141d84.
//
// Solidity: function getDecimals() view returns(uint8)
func (_IERC20 *IERC20Caller) GetDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "getDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0xf0141d84.
//
// Solidity: function getDecimals() view returns(uint8)
func (_IERC20 *IERC20Session) GetDecimals() (uint8, error) {
	return _IERC20.Contract.GetDecimals(&_IERC20.CallOpts)
}

// GetDecimals is a free data retrieval call binding the contract method 0xf0141d84.
//
// Solidity: function getDecimals() view returns(uint8)
func (_IERC20 *IERC20CallerSession) GetDecimals() (uint8, error) {
	return _IERC20.Contract.GetDecimals(&_IERC20.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_IERC20 *IERC20Caller) GetName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "getName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_IERC20 *IERC20Session) GetName() (string, error) {
	return _IERC20.Contract.GetName(&_IERC20.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_IERC20 *IERC20CallerSession) GetName() (string, error) {
	return _IERC20.Contract.GetName(&_IERC20.CallOpts)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_IERC20 *IERC20Caller) GetSymbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "getSymbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_IERC20 *IERC20Session) GetSymbol() (string, error) {
	return _IERC20.Contract.GetSymbol(&_IERC20.CallOpts)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_IERC20 *IERC20CallerSession) GetSymbol() (string, error) {
	return _IERC20.Contract.GetSymbol(&_IERC20.CallOpts)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) GetTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "getTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) GetTotalSupply() (*big.Int, error) {
	return _IERC20.Contract.GetTotalSupply(&_IERC20.CallOpts)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) GetTotalSupply() (*big.Int, error) {
	return _IERC20.Contract.GetTotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns(bool)
func (_IERC20 *IERC20Transactor) Burn(opts *bind.TransactOpts, burnAmount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "burn", burnAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns(bool)
func (_IERC20 *IERC20Session) Burn(burnAmount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Burn(&_IERC20.TransactOpts, burnAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Burn(burnAmount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Burn(&_IERC20.TransactOpts, burnAmount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_IERC20 *IERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_IERC20 *IERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.DecreaseAllowance(&_IERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_IERC20 *IERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.DecreaseAllowance(&_IERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_IERC20 *IERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_IERC20 *IERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.IncreaseAllowance(&_IERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_IERC20 *IERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.IncreaseAllowance(&_IERC20.TransactOpts, spender, addedValue)
}

// MintToken is a paid mutator transaction binding the contract method 0x6bd1ce89.
//
// Solidity: function mintToken(address target, uint256 mintedAmount, bytes[5] signs) returns(bool)
func (_IERC20 *IERC20Transactor) MintToken(opts *bind.TransactOpts, target common.Address, mintedAmount *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "mintToken", target, mintedAmount, signs)
}

// MintToken is a paid mutator transaction binding the contract method 0x6bd1ce89.
//
// Solidity: function mintToken(address target, uint256 mintedAmount, bytes[5] signs) returns(bool)
func (_IERC20 *IERC20Session) MintToken(target common.Address, mintedAmount *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _IERC20.Contract.MintToken(&_IERC20.TransactOpts, target, mintedAmount, signs)
}

// MintToken is a paid mutator transaction binding the contract method 0x6bd1ce89.
//
// Solidity: function mintToken(address target, uint256 mintedAmount, bytes[5] signs) returns(bool)
func (_IERC20 *IERC20TransactorSession) MintToken(target common.Address, mintedAmount *big.Int, signs [5][]byte) (*types.Transaction, error) {
	return _IERC20.Contract.MintToken(&_IERC20.TransactOpts, target, mintedAmount, signs)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileSysMetaData contains all meta data concerning the IFileSys contract.
var IFileSysMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"kIndex\",\"type\":\"uint64\"}],\"name\":\"addKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"usign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"psign\",\"type\":\"bytes\"}],\"internalType\":\"structAOParams\",\"name\":\"ps\",\"type\":\"tuple\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"kIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newPro\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"addRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"}],\"name\":\"createFs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"user\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"provider\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"token\",\"type\":\"uint32\"}],\"name\":\"getChannelInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"user\",\"type\":\"uint64\"}],\"name\":\"getFsInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"user\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"provider\",\"type\":\"uint64\"}],\"name\":\"getFsInfoAggOrder\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"user\",\"type\":\"uint64\"}],\"name\":\"getFsPSum\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"user\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getFsPro\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"getSettleInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"user\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"provider\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"token\",\"type\":\"uint32\"}],\"name\":\"getStoreInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"pAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"},{\"internalType\":\"uint64[]\",\"name\":\"kIndexes\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"internalType\":\"structPWParams\",\"name\":\"ps\",\"type\":\"tuple\"}],\"name\":\"proWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"uAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"kIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"usign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"psign\",\"type\":\"bytes\"}],\"internalType\":\"structSOParams\",\"name\":\"ps\",\"type\":\"tuple\"}],\"name\":\"subOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"kIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newPro\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"subRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"roleType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"50cbb46f": "addKeeper(uint64)",
		"0e09662b": "addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes))",
		"0f60c7b3": "addRepair(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256)",
		"e49c0724": "createFs(uint64)",
		"761966b7": "getBalance(uint64,uint32)",
		"cc122253": "getChannelInfo(uint64,uint64,uint32)",
		"324ab551": "getFsInfo(uint64)",
		"03eb18aa": "getFsInfoAggOrder(uint64,uint64)",
		"650c6a89": "getFsPSum(uint64)",
		"25d931e8": "getFsPro(uint64,uint64)",
		"7b31a24d": "getSettleInfo(uint64,uint32)",
		"3f5f363e": "getStoreInfo(uint64,uint64,uint32)",
		"edf26f5c": "proWithdraw((uint64,uint32,address,address,uint256,uint256,uint64[],bytes[]))",
		"e04f98ed": "recharge(uint64,uint32,address,address,uint256)",
		"abfc55f1": "subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes))",
		"75818519": "subRepair(uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256)",
		"bab8d63a": "withdraw(uint64,uint32,uint8,address,address,uint256)",
	},
}

// IFileSysABI is the input ABI used to generate the binding from.
// Deprecated: Use IFileSysMetaData.ABI instead.
var IFileSysABI = IFileSysMetaData.ABI

// Deprecated: Use IFileSysMetaData.Sigs instead.
// IFileSysFuncSigs maps the 4-byte function signature to its string representation.
var IFileSysFuncSigs = IFileSysMetaData.Sigs

// IFileSys is an auto generated Go binding around an Ethereum contract.
type IFileSys struct {
	IFileSysCaller     // Read-only binding to the contract
	IFileSysTransactor // Write-only binding to the contract
	IFileSysFilterer   // Log filterer for contract events
}

// IFileSysCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFileSysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFileSysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFileSysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFileSysSession struct {
	Contract     *IFileSys         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFileSysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFileSysCallerSession struct {
	Contract *IFileSysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IFileSysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFileSysTransactorSession struct {
	Contract     *IFileSysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IFileSysRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFileSysRaw struct {
	Contract *IFileSys // Generic contract binding to access the raw methods on
}

// IFileSysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFileSysCallerRaw struct {
	Contract *IFileSysCaller // Generic read-only contract binding to access the raw methods on
}

// IFileSysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFileSysTransactorRaw struct {
	Contract *IFileSysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFileSys creates a new instance of IFileSys, bound to a specific deployed contract.
func NewIFileSys(address common.Address, backend bind.ContractBackend) (*IFileSys, error) {
	contract, err := bindIFileSys(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFileSys{IFileSysCaller: IFileSysCaller{contract: contract}, IFileSysTransactor: IFileSysTransactor{contract: contract}, IFileSysFilterer: IFileSysFilterer{contract: contract}}, nil
}

// NewIFileSysCaller creates a new read-only instance of IFileSys, bound to a specific deployed contract.
func NewIFileSysCaller(address common.Address, caller bind.ContractCaller) (*IFileSysCaller, error) {
	contract, err := bindIFileSys(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFileSysCaller{contract: contract}, nil
}

// NewIFileSysTransactor creates a new write-only instance of IFileSys, bound to a specific deployed contract.
func NewIFileSysTransactor(address common.Address, transactor bind.ContractTransactor) (*IFileSysTransactor, error) {
	contract, err := bindIFileSys(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFileSysTransactor{contract: contract}, nil
}

// NewIFileSysFilterer creates a new log filterer instance of IFileSys, bound to a specific deployed contract.
func NewIFileSysFilterer(address common.Address, filterer bind.ContractFilterer) (*IFileSysFilterer, error) {
	contract, err := bindIFileSys(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFileSysFilterer{contract: contract}, nil
}

// bindIFileSys binds a generic wrapper to an already deployed contract.
func bindIFileSys(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFileSysABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileSys *IFileSysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileSys.Contract.IFileSysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileSys *IFileSysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileSys.Contract.IFileSysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileSys *IFileSysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileSys.Contract.IFileSysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileSys *IFileSysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileSys.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileSys *IFileSysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileSys.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileSys *IFileSysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileSys.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x761966b7.
//
// Solidity: function getBalance(uint64 index, uint32 tIndex) view returns(uint256, uint256)
func (_IFileSys *IFileSysCaller) GetBalance(opts *bind.CallOpts, index uint64, tIndex uint32) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getBalance", index, tIndex)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetBalance is a free data retrieval call binding the contract method 0x761966b7.
//
// Solidity: function getBalance(uint64 index, uint32 tIndex) view returns(uint256, uint256)
func (_IFileSys *IFileSysSession) GetBalance(index uint64, tIndex uint32) (*big.Int, *big.Int, error) {
	return _IFileSys.Contract.GetBalance(&_IFileSys.CallOpts, index, tIndex)
}

// GetBalance is a free data retrieval call binding the contract method 0x761966b7.
//
// Solidity: function getBalance(uint64 index, uint32 tIndex) view returns(uint256, uint256)
func (_IFileSys *IFileSysCallerSession) GetBalance(index uint64, tIndex uint32) (*big.Int, *big.Int, error) {
	return _IFileSys.Contract.GetBalance(&_IFileSys.CallOpts, index, tIndex)
}

// GetChannelInfo is a free data retrieval call binding the contract method 0xcc122253.
//
// Solidity: function getChannelInfo(uint64 user, uint64 provider, uint32 token) view returns(uint256, uint64, uint64)
func (_IFileSys *IFileSysCaller) GetChannelInfo(opts *bind.CallOpts, user uint64, provider uint64, token uint32) (*big.Int, uint64, uint64, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getChannelInfo", user, provider, token)

	if err != nil {
		return *new(*big.Int), *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return out0, out1, out2, err

}

// GetChannelInfo is a free data retrieval call binding the contract method 0xcc122253.
//
// Solidity: function getChannelInfo(uint64 user, uint64 provider, uint32 token) view returns(uint256, uint64, uint64)
func (_IFileSys *IFileSysSession) GetChannelInfo(user uint64, provider uint64, token uint32) (*big.Int, uint64, uint64, error) {
	return _IFileSys.Contract.GetChannelInfo(&_IFileSys.CallOpts, user, provider, token)
}

// GetChannelInfo is a free data retrieval call binding the contract method 0xcc122253.
//
// Solidity: function getChannelInfo(uint64 user, uint64 provider, uint32 token) view returns(uint256, uint64, uint64)
func (_IFileSys *IFileSysCallerSession) GetChannelInfo(user uint64, provider uint64, token uint32) (*big.Int, uint64, uint64, error) {
	return _IFileSys.Contract.GetChannelInfo(&_IFileSys.CallOpts, user, provider, token)
}

// GetFsInfo is a free data retrieval call binding the contract method 0x324ab551.
//
// Solidity: function getFsInfo(uint64 user) view returns(bool)
func (_IFileSys *IFileSysCaller) GetFsInfo(opts *bind.CallOpts, user uint64) (bool, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getFsInfo", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFsInfo is a free data retrieval call binding the contract method 0x324ab551.
//
// Solidity: function getFsInfo(uint64 user) view returns(bool)
func (_IFileSys *IFileSysSession) GetFsInfo(user uint64) (bool, error) {
	return _IFileSys.Contract.GetFsInfo(&_IFileSys.CallOpts, user)
}

// GetFsInfo is a free data retrieval call binding the contract method 0x324ab551.
//
// Solidity: function getFsInfo(uint64 user) view returns(bool)
func (_IFileSys *IFileSysCallerSession) GetFsInfo(user uint64) (bool, error) {
	return _IFileSys.Contract.GetFsInfo(&_IFileSys.CallOpts, user)
}

// GetFsInfoAggOrder is a free data retrieval call binding the contract method 0x03eb18aa.
//
// Solidity: function getFsInfoAggOrder(uint64 user, uint64 provider) view returns(uint64, uint64)
func (_IFileSys *IFileSysCaller) GetFsInfoAggOrder(opts *bind.CallOpts, user uint64, provider uint64) (uint64, uint64, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getFsInfoAggOrder", user, provider)

	if err != nil {
		return *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

// GetFsInfoAggOrder is a free data retrieval call binding the contract method 0x03eb18aa.
//
// Solidity: function getFsInfoAggOrder(uint64 user, uint64 provider) view returns(uint64, uint64)
func (_IFileSys *IFileSysSession) GetFsInfoAggOrder(user uint64, provider uint64) (uint64, uint64, error) {
	return _IFileSys.Contract.GetFsInfoAggOrder(&_IFileSys.CallOpts, user, provider)
}

// GetFsInfoAggOrder is a free data retrieval call binding the contract method 0x03eb18aa.
//
// Solidity: function getFsInfoAggOrder(uint64 user, uint64 provider) view returns(uint64, uint64)
func (_IFileSys *IFileSysCallerSession) GetFsInfoAggOrder(user uint64, provider uint64) (uint64, uint64, error) {
	return _IFileSys.Contract.GetFsInfoAggOrder(&_IFileSys.CallOpts, user, provider)
}

// GetFsPSum is a free data retrieval call binding the contract method 0x650c6a89.
//
// Solidity: function getFsPSum(uint64 user) view returns(uint64)
func (_IFileSys *IFileSysCaller) GetFsPSum(opts *bind.CallOpts, user uint64) (uint64, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getFsPSum", user)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetFsPSum is a free data retrieval call binding the contract method 0x650c6a89.
//
// Solidity: function getFsPSum(uint64 user) view returns(uint64)
func (_IFileSys *IFileSysSession) GetFsPSum(user uint64) (uint64, error) {
	return _IFileSys.Contract.GetFsPSum(&_IFileSys.CallOpts, user)
}

// GetFsPSum is a free data retrieval call binding the contract method 0x650c6a89.
//
// Solidity: function getFsPSum(uint64 user) view returns(uint64)
func (_IFileSys *IFileSysCallerSession) GetFsPSum(user uint64) (uint64, error) {
	return _IFileSys.Contract.GetFsPSum(&_IFileSys.CallOpts, user)
}

// GetFsPro is a free data retrieval call binding the contract method 0x25d931e8.
//
// Solidity: function getFsPro(uint64 user, uint64 i) view returns(uint64)
func (_IFileSys *IFileSysCaller) GetFsPro(opts *bind.CallOpts, user uint64, i uint64) (uint64, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getFsPro", user, i)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetFsPro is a free data retrieval call binding the contract method 0x25d931e8.
//
// Solidity: function getFsPro(uint64 user, uint64 i) view returns(uint64)
func (_IFileSys *IFileSysSession) GetFsPro(user uint64, i uint64) (uint64, error) {
	return _IFileSys.Contract.GetFsPro(&_IFileSys.CallOpts, user, i)
}

// GetFsPro is a free data retrieval call binding the contract method 0x25d931e8.
//
// Solidity: function getFsPro(uint64 user, uint64 i) view returns(uint64)
func (_IFileSys *IFileSysCallerSession) GetFsPro(user uint64, i uint64) (uint64, error) {
	return _IFileSys.Contract.GetFsPro(&_IFileSys.CallOpts, user, i)
}

// GetSettleInfo is a free data retrieval call binding the contract method 0x7b31a24d.
//
// Solidity: function getSettleInfo(uint64 index, uint32 tIndex) view returns(uint64, uint64, uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_IFileSys *IFileSysCaller) GetSettleInfo(opts *bind.CallOpts, index uint64, tIndex uint32) (uint64, uint64, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getSettleInfo", index, tIndex)

	if err != nil {
		return *new(uint64), *new(uint64), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	out8 := *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	out9 := *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	out10 := *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, out9, out10, err

}

// GetSettleInfo is a free data retrieval call binding the contract method 0x7b31a24d.
//
// Solidity: function getSettleInfo(uint64 index, uint32 tIndex) view returns(uint64, uint64, uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_IFileSys *IFileSysSession) GetSettleInfo(index uint64, tIndex uint32) (uint64, uint64, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IFileSys.Contract.GetSettleInfo(&_IFileSys.CallOpts, index, tIndex)
}

// GetSettleInfo is a free data retrieval call binding the contract method 0x7b31a24d.
//
// Solidity: function getSettleInfo(uint64 index, uint32 tIndex) view returns(uint64, uint64, uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_IFileSys *IFileSysCallerSession) GetSettleInfo(index uint64, tIndex uint32) (uint64, uint64, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IFileSys.Contract.GetSettleInfo(&_IFileSys.CallOpts, index, tIndex)
}

// GetStoreInfo is a free data retrieval call binding the contract method 0x3f5f363e.
//
// Solidity: function getStoreInfo(uint64 user, uint64 provider, uint32 token) view returns(uint64, uint64, uint256)
func (_IFileSys *IFileSysCaller) GetStoreInfo(opts *bind.CallOpts, user uint64, provider uint64, token uint32) (uint64, uint64, *big.Int, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getStoreInfo", user, provider, token)

	if err != nil {
		return *new(uint64), *new(uint64), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetStoreInfo is a free data retrieval call binding the contract method 0x3f5f363e.
//
// Solidity: function getStoreInfo(uint64 user, uint64 provider, uint32 token) view returns(uint64, uint64, uint256)
func (_IFileSys *IFileSysSession) GetStoreInfo(user uint64, provider uint64, token uint32) (uint64, uint64, *big.Int, error) {
	return _IFileSys.Contract.GetStoreInfo(&_IFileSys.CallOpts, user, provider, token)
}

// GetStoreInfo is a free data retrieval call binding the contract method 0x3f5f363e.
//
// Solidity: function getStoreInfo(uint64 user, uint64 provider, uint32 token) view returns(uint64, uint64, uint256)
func (_IFileSys *IFileSysCallerSession) GetStoreInfo(user uint64, provider uint64, token uint32) (uint64, uint64, *big.Int, error) {
	return _IFileSys.Contract.GetStoreInfo(&_IFileSys.CallOpts, user, provider, token)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 kIndex) returns()
func (_IFileSys *IFileSysTransactor) AddKeeper(opts *bind.TransactOpts, kIndex uint64) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "addKeeper", kIndex)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 kIndex) returns()
func (_IFileSys *IFileSysSession) AddKeeper(kIndex uint64) (*types.Transaction, error) {
	return _IFileSys.Contract.AddKeeper(&_IFileSys.TransactOpts, kIndex)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 kIndex) returns()
func (_IFileSys *IFileSysTransactorSession) AddKeeper(kIndex uint64) (*types.Transaction, error) {
	return _IFileSys.Contract.AddKeeper(&_IFileSys.TransactOpts, kIndex)
}

// AddOrder is a paid mutator transaction binding the contract method 0x0e09662b.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes) ps) returns()
func (_IFileSys *IFileSysTransactor) AddOrder(opts *bind.TransactOpts, ps AOParams) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "addOrder", ps)
}

// AddOrder is a paid mutator transaction binding the contract method 0x0e09662b.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes) ps) returns()
func (_IFileSys *IFileSysSession) AddOrder(ps AOParams) (*types.Transaction, error) {
	return _IFileSys.Contract.AddOrder(&_IFileSys.TransactOpts, ps)
}

// AddOrder is a paid mutator transaction binding the contract method 0x0e09662b.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes) ps) returns()
func (_IFileSys *IFileSysTransactorSession) AddOrder(ps AOParams) (*types.Transaction, error) {
	return _IFileSys.Contract.AddOrder(&_IFileSys.TransactOpts, ps)
}

// AddRepair is a paid mutator transaction binding the contract method 0x0f60c7b3.
//
// Solidity: function addRepair(uint64 kIndex, uint64 pIndex, uint64 newPro, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) returns()
func (_IFileSys *IFileSysTransactor) AddRepair(opts *bind.TransactOpts, kIndex uint64, pIndex uint64, newPro uint64, start uint64, end uint64, size uint64, nonce uint64, tokenIndex uint32, sprice *big.Int) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "addRepair", kIndex, pIndex, newPro, start, end, size, nonce, tokenIndex, sprice)
}

// AddRepair is a paid mutator transaction binding the contract method 0x0f60c7b3.
//
// Solidity: function addRepair(uint64 kIndex, uint64 pIndex, uint64 newPro, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) returns()
func (_IFileSys *IFileSysSession) AddRepair(kIndex uint64, pIndex uint64, newPro uint64, start uint64, end uint64, size uint64, nonce uint64, tokenIndex uint32, sprice *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.AddRepair(&_IFileSys.TransactOpts, kIndex, pIndex, newPro, start, end, size, nonce, tokenIndex, sprice)
}

// AddRepair is a paid mutator transaction binding the contract method 0x0f60c7b3.
//
// Solidity: function addRepair(uint64 kIndex, uint64 pIndex, uint64 newPro, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) returns()
func (_IFileSys *IFileSysTransactorSession) AddRepair(kIndex uint64, pIndex uint64, newPro uint64, start uint64, end uint64, size uint64, nonce uint64, tokenIndex uint32, sprice *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.AddRepair(&_IFileSys.TransactOpts, kIndex, pIndex, newPro, start, end, size, nonce, tokenIndex, sprice)
}

// CreateFs is a paid mutator transaction binding the contract method 0xe49c0724.
//
// Solidity: function createFs(uint64 uIndex) returns()
func (_IFileSys *IFileSysTransactor) CreateFs(opts *bind.TransactOpts, uIndex uint64) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "createFs", uIndex)
}

// CreateFs is a paid mutator transaction binding the contract method 0xe49c0724.
//
// Solidity: function createFs(uint64 uIndex) returns()
func (_IFileSys *IFileSysSession) CreateFs(uIndex uint64) (*types.Transaction, error) {
	return _IFileSys.Contract.CreateFs(&_IFileSys.TransactOpts, uIndex)
}

// CreateFs is a paid mutator transaction binding the contract method 0xe49c0724.
//
// Solidity: function createFs(uint64 uIndex) returns()
func (_IFileSys *IFileSysTransactorSession) CreateFs(uIndex uint64) (*types.Transaction, error) {
	return _IFileSys.Contract.CreateFs(&_IFileSys.TransactOpts, uIndex)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xedf26f5c.
//
// Solidity: function proWithdraw((uint64,uint32,address,address,uint256,uint256,uint64[],bytes[]) ps) returns()
func (_IFileSys *IFileSysTransactor) ProWithdraw(opts *bind.TransactOpts, ps PWParams) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "proWithdraw", ps)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xedf26f5c.
//
// Solidity: function proWithdraw((uint64,uint32,address,address,uint256,uint256,uint64[],bytes[]) ps) returns()
func (_IFileSys *IFileSysSession) ProWithdraw(ps PWParams) (*types.Transaction, error) {
	return _IFileSys.Contract.ProWithdraw(&_IFileSys.TransactOpts, ps)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xedf26f5c.
//
// Solidity: function proWithdraw((uint64,uint32,address,address,uint256,uint256,uint64[],bytes[]) ps) returns()
func (_IFileSys *IFileSysTransactorSession) ProWithdraw(ps PWParams) (*types.Transaction, error) {
	return _IFileSys.Contract.ProWithdraw(&_IFileSys.TransactOpts, ps)
}

// Recharge is a paid mutator transaction binding the contract method 0xe04f98ed.
//
// Solidity: function recharge(uint64 uIndex, uint32 tokenIndex, address uAddr, address tAddr, uint256 money) returns()
func (_IFileSys *IFileSysTransactor) Recharge(opts *bind.TransactOpts, uIndex uint64, tokenIndex uint32, uAddr common.Address, tAddr common.Address, money *big.Int) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "recharge", uIndex, tokenIndex, uAddr, tAddr, money)
}

// Recharge is a paid mutator transaction binding the contract method 0xe04f98ed.
//
// Solidity: function recharge(uint64 uIndex, uint32 tokenIndex, address uAddr, address tAddr, uint256 money) returns()
func (_IFileSys *IFileSysSession) Recharge(uIndex uint64, tokenIndex uint32, uAddr common.Address, tAddr common.Address, money *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.Recharge(&_IFileSys.TransactOpts, uIndex, tokenIndex, uAddr, tAddr, money)
}

// Recharge is a paid mutator transaction binding the contract method 0xe04f98ed.
//
// Solidity: function recharge(uint64 uIndex, uint32 tokenIndex, address uAddr, address tAddr, uint256 money) returns()
func (_IFileSys *IFileSysTransactorSession) Recharge(uIndex uint64, tokenIndex uint32, uAddr common.Address, tAddr common.Address, money *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.Recharge(&_IFileSys.TransactOpts, uIndex, tokenIndex, uAddr, tAddr, money)
}

// SubOrder is a paid mutator transaction binding the contract method 0xabfc55f1.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes) ps) returns()
func (_IFileSys *IFileSysTransactor) SubOrder(opts *bind.TransactOpts, ps SOParams) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "subOrder", ps)
}

// SubOrder is a paid mutator transaction binding the contract method 0xabfc55f1.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes) ps) returns()
func (_IFileSys *IFileSysSession) SubOrder(ps SOParams) (*types.Transaction, error) {
	return _IFileSys.Contract.SubOrder(&_IFileSys.TransactOpts, ps)
}

// SubOrder is a paid mutator transaction binding the contract method 0xabfc55f1.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes) ps) returns()
func (_IFileSys *IFileSysTransactorSession) SubOrder(ps SOParams) (*types.Transaction, error) {
	return _IFileSys.Contract.SubOrder(&_IFileSys.TransactOpts, ps)
}

// SubRepair is a paid mutator transaction binding the contract method 0x75818519.
//
// Solidity: function subRepair(uint64 kIndex, uint64 newPro, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) returns()
func (_IFileSys *IFileSysTransactor) SubRepair(opts *bind.TransactOpts, kIndex uint64, newPro uint64, start uint64, end uint64, size uint64, nonce uint64, tokenIndex uint32, sprice *big.Int) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "subRepair", kIndex, newPro, start, end, size, nonce, tokenIndex, sprice)
}

// SubRepair is a paid mutator transaction binding the contract method 0x75818519.
//
// Solidity: function subRepair(uint64 kIndex, uint64 newPro, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) returns()
func (_IFileSys *IFileSysSession) SubRepair(kIndex uint64, newPro uint64, start uint64, end uint64, size uint64, nonce uint64, tokenIndex uint32, sprice *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.SubRepair(&_IFileSys.TransactOpts, kIndex, newPro, start, end, size, nonce, tokenIndex, sprice)
}

// SubRepair is a paid mutator transaction binding the contract method 0x75818519.
//
// Solidity: function subRepair(uint64 kIndex, uint64 newPro, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) returns()
func (_IFileSys *IFileSysTransactorSession) SubRepair(kIndex uint64, newPro uint64, start uint64, end uint64, size uint64, nonce uint64, tokenIndex uint32, sprice *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.SubRepair(&_IFileSys.TransactOpts, kIndex, newPro, start, end, size, nonce, tokenIndex, sprice)
}

// Withdraw is a paid mutator transaction binding the contract method 0xbab8d63a.
//
// Solidity: function withdraw(uint64 index, uint32 tokenIndex, uint8 roleType, address tAddr, address addr, uint256 amount) returns()
func (_IFileSys *IFileSysTransactor) Withdraw(opts *bind.TransactOpts, index uint64, tokenIndex uint32, roleType uint8, tAddr common.Address, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "withdraw", index, tokenIndex, roleType, tAddr, addr, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xbab8d63a.
//
// Solidity: function withdraw(uint64 index, uint32 tokenIndex, uint8 roleType, address tAddr, address addr, uint256 amount) returns()
func (_IFileSys *IFileSysSession) Withdraw(index uint64, tokenIndex uint32, roleType uint8, tAddr common.Address, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.Withdraw(&_IFileSys.TransactOpts, index, tokenIndex, roleType, tAddr, addr, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xbab8d63a.
//
// Solidity: function withdraw(uint64 index, uint32 tokenIndex, uint8 roleType, address tAddr, address addr, uint256 amount) returns()
func (_IFileSys *IFileSysTransactorSession) Withdraw(index uint64, tokenIndex uint32, roleType uint8, tAddr common.Address, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.Withdraw(&_IFileSys.TransactOpts, index, tokenIndex, roleType, tAddr, addr, amount)
}

// IIssuanceMetaData contains all meta data concerning the IIssuance contract.
var IIssuanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIssuParams\",\"name\":\"ps\",\"type\":\"tuple\"}],\"name\":\"issu\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_add\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sub\",\"type\":\"uint256\"}],\"name\":\"setTP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c862ea50": "issu((uint64,uint64,uint64,uint256))",
		"11e65fc0": "setTP(uint256,uint256)",
	},
}

// IIssuanceABI is the input ABI used to generate the binding from.
// Deprecated: Use IIssuanceMetaData.ABI instead.
var IIssuanceABI = IIssuanceMetaData.ABI

// Deprecated: Use IIssuanceMetaData.Sigs instead.
// IIssuanceFuncSigs maps the 4-byte function signature to its string representation.
var IIssuanceFuncSigs = IIssuanceMetaData.Sigs

// IIssuance is an auto generated Go binding around an Ethereum contract.
type IIssuance struct {
	IIssuanceCaller     // Read-only binding to the contract
	IIssuanceTransactor // Write-only binding to the contract
	IIssuanceFilterer   // Log filterer for contract events
}

// IIssuanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IIssuanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IIssuanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IIssuanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IIssuanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IIssuanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IIssuanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IIssuanceSession struct {
	Contract     *IIssuance        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IIssuanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IIssuanceCallerSession struct {
	Contract *IIssuanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IIssuanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IIssuanceTransactorSession struct {
	Contract     *IIssuanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IIssuanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IIssuanceRaw struct {
	Contract *IIssuance // Generic contract binding to access the raw methods on
}

// IIssuanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IIssuanceCallerRaw struct {
	Contract *IIssuanceCaller // Generic read-only contract binding to access the raw methods on
}

// IIssuanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IIssuanceTransactorRaw struct {
	Contract *IIssuanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIIssuance creates a new instance of IIssuance, bound to a specific deployed contract.
func NewIIssuance(address common.Address, backend bind.ContractBackend) (*IIssuance, error) {
	contract, err := bindIIssuance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IIssuance{IIssuanceCaller: IIssuanceCaller{contract: contract}, IIssuanceTransactor: IIssuanceTransactor{contract: contract}, IIssuanceFilterer: IIssuanceFilterer{contract: contract}}, nil
}

// NewIIssuanceCaller creates a new read-only instance of IIssuance, bound to a specific deployed contract.
func NewIIssuanceCaller(address common.Address, caller bind.ContractCaller) (*IIssuanceCaller, error) {
	contract, err := bindIIssuance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IIssuanceCaller{contract: contract}, nil
}

// NewIIssuanceTransactor creates a new write-only instance of IIssuance, bound to a specific deployed contract.
func NewIIssuanceTransactor(address common.Address, transactor bind.ContractTransactor) (*IIssuanceTransactor, error) {
	contract, err := bindIIssuance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IIssuanceTransactor{contract: contract}, nil
}

// NewIIssuanceFilterer creates a new log filterer instance of IIssuance, bound to a specific deployed contract.
func NewIIssuanceFilterer(address common.Address, filterer bind.ContractFilterer) (*IIssuanceFilterer, error) {
	contract, err := bindIIssuance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IIssuanceFilterer{contract: contract}, nil
}

// bindIIssuance binds a generic wrapper to an already deployed contract.
func bindIIssuance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IIssuanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IIssuance *IIssuanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IIssuance.Contract.IIssuanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IIssuance *IIssuanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IIssuance.Contract.IIssuanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IIssuance *IIssuanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IIssuance.Contract.IIssuanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IIssuance *IIssuanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IIssuance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IIssuance *IIssuanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IIssuance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IIssuance *IIssuanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IIssuance.Contract.contract.Transact(opts, method, params...)
}

// Issu is a paid mutator transaction binding the contract method 0xc862ea50.
//
// Solidity: function issu((uint64,uint64,uint64,uint256) ps) returns(uint256)
func (_IIssuance *IIssuanceTransactor) Issu(opts *bind.TransactOpts, ps IssuParams) (*types.Transaction, error) {
	return _IIssuance.contract.Transact(opts, "issu", ps)
}

// Issu is a paid mutator transaction binding the contract method 0xc862ea50.
//
// Solidity: function issu((uint64,uint64,uint64,uint256) ps) returns(uint256)
func (_IIssuance *IIssuanceSession) Issu(ps IssuParams) (*types.Transaction, error) {
	return _IIssuance.Contract.Issu(&_IIssuance.TransactOpts, ps)
}

// Issu is a paid mutator transaction binding the contract method 0xc862ea50.
//
// Solidity: function issu((uint64,uint64,uint64,uint256) ps) returns(uint256)
func (_IIssuance *IIssuanceTransactorSession) Issu(ps IssuParams) (*types.Transaction, error) {
	return _IIssuance.Contract.Issu(&_IIssuance.TransactOpts, ps)
}

// SetTP is a paid mutator transaction binding the contract method 0x11e65fc0.
//
// Solidity: function setTP(uint256 _add, uint256 _sub) returns()
func (_IIssuance *IIssuanceTransactor) SetTP(opts *bind.TransactOpts, _add *big.Int, _sub *big.Int) (*types.Transaction, error) {
	return _IIssuance.contract.Transact(opts, "setTP", _add, _sub)
}

// SetTP is a paid mutator transaction binding the contract method 0x11e65fc0.
//
// Solidity: function setTP(uint256 _add, uint256 _sub) returns()
func (_IIssuance *IIssuanceSession) SetTP(_add *big.Int, _sub *big.Int) (*types.Transaction, error) {
	return _IIssuance.Contract.SetTP(&_IIssuance.TransactOpts, _add, _sub)
}

// SetTP is a paid mutator transaction binding the contract method 0x11e65fc0.
//
// Solidity: function setTP(uint256 _add, uint256 _sub) returns()
func (_IIssuance *IIssuanceTransactorSession) SetTP(_add *big.Int, _sub *big.Int) (*types.Transaction, error) {
	return _IIssuance.Contract.SetTP(&_IIssuance.TransactOpts, _add, _sub)
}

// IRTokenMetaData contains all meta data concerning the IRToken contract.
var IRTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"addT\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"getTA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"getTI\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"isValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8ff8e15a": "addT(address)",
		"f7b7d6b6": "getTA(uint32)",
		"2df2685f": "getTI(address)",
		"3c7bdc19": "isValid(uint32)",
	},
}

// IRTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use IRTokenMetaData.ABI instead.
var IRTokenABI = IRTokenMetaData.ABI

// Deprecated: Use IRTokenMetaData.Sigs instead.
// IRTokenFuncSigs maps the 4-byte function signature to its string representation.
var IRTokenFuncSigs = IRTokenMetaData.Sigs

// IRToken is an auto generated Go binding around an Ethereum contract.
type IRToken struct {
	IRTokenCaller     // Read-only binding to the contract
	IRTokenTransactor // Write-only binding to the contract
	IRTokenFilterer   // Log filterer for contract events
}

// IRTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRTokenSession struct {
	Contract     *IRToken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRTokenCallerSession struct {
	Contract *IRTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IRTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRTokenTransactorSession struct {
	Contract     *IRTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IRTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRTokenRaw struct {
	Contract *IRToken // Generic contract binding to access the raw methods on
}

// IRTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRTokenCallerRaw struct {
	Contract *IRTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IRTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRTokenTransactorRaw struct {
	Contract *IRTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRToken creates a new instance of IRToken, bound to a specific deployed contract.
func NewIRToken(address common.Address, backend bind.ContractBackend) (*IRToken, error) {
	contract, err := bindIRToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRToken{IRTokenCaller: IRTokenCaller{contract: contract}, IRTokenTransactor: IRTokenTransactor{contract: contract}, IRTokenFilterer: IRTokenFilterer{contract: contract}}, nil
}

// NewIRTokenCaller creates a new read-only instance of IRToken, bound to a specific deployed contract.
func NewIRTokenCaller(address common.Address, caller bind.ContractCaller) (*IRTokenCaller, error) {
	contract, err := bindIRToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRTokenCaller{contract: contract}, nil
}

// NewIRTokenTransactor creates a new write-only instance of IRToken, bound to a specific deployed contract.
func NewIRTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IRTokenTransactor, error) {
	contract, err := bindIRToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRTokenTransactor{contract: contract}, nil
}

// NewIRTokenFilterer creates a new log filterer instance of IRToken, bound to a specific deployed contract.
func NewIRTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IRTokenFilterer, error) {
	contract, err := bindIRToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRTokenFilterer{contract: contract}, nil
}

// bindIRToken binds a generic wrapper to an already deployed contract.
func bindIRToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRToken *IRTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRToken.Contract.IRTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRToken *IRTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRToken.Contract.IRTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRToken *IRTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRToken.Contract.IRTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRToken *IRTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRToken *IRTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRToken *IRTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRToken.Contract.contract.Transact(opts, method, params...)
}

// GetTA is a free data retrieval call binding the contract method 0xf7b7d6b6.
//
// Solidity: function getTA(uint32 tIndex) view returns(address)
func (_IRToken *IRTokenCaller) GetTA(opts *bind.CallOpts, tIndex uint32) (common.Address, error) {
	var out []interface{}
	err := _IRToken.contract.Call(opts, &out, "getTA", tIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTA is a free data retrieval call binding the contract method 0xf7b7d6b6.
//
// Solidity: function getTA(uint32 tIndex) view returns(address)
func (_IRToken *IRTokenSession) GetTA(tIndex uint32) (common.Address, error) {
	return _IRToken.Contract.GetTA(&_IRToken.CallOpts, tIndex)
}

// GetTA is a free data retrieval call binding the contract method 0xf7b7d6b6.
//
// Solidity: function getTA(uint32 tIndex) view returns(address)
func (_IRToken *IRTokenCallerSession) GetTA(tIndex uint32) (common.Address, error) {
	return _IRToken.Contract.GetTA(&_IRToken.CallOpts, tIndex)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address t) view returns(uint32, bool)
func (_IRToken *IRTokenCaller) GetTI(opts *bind.CallOpts, t common.Address) (uint32, bool, error) {
	var out []interface{}
	err := _IRToken.contract.Call(opts, &out, "getTI", t)

	if err != nil {
		return *new(uint32), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address t) view returns(uint32, bool)
func (_IRToken *IRTokenSession) GetTI(t common.Address) (uint32, bool, error) {
	return _IRToken.Contract.GetTI(&_IRToken.CallOpts, t)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address t) view returns(uint32, bool)
func (_IRToken *IRTokenCallerSession) GetTI(t common.Address) (uint32, bool, error) {
	return _IRToken.Contract.GetTI(&_IRToken.CallOpts, t)
}

// IsValid is a free data retrieval call binding the contract method 0x3c7bdc19.
//
// Solidity: function isValid(uint32 tIndex) view returns(bool)
func (_IRToken *IRTokenCaller) IsValid(opts *bind.CallOpts, tIndex uint32) (bool, error) {
	var out []interface{}
	err := _IRToken.contract.Call(opts, &out, "isValid", tIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValid is a free data retrieval call binding the contract method 0x3c7bdc19.
//
// Solidity: function isValid(uint32 tIndex) view returns(bool)
func (_IRToken *IRTokenSession) IsValid(tIndex uint32) (bool, error) {
	return _IRToken.Contract.IsValid(&_IRToken.CallOpts, tIndex)
}

// IsValid is a free data retrieval call binding the contract method 0x3c7bdc19.
//
// Solidity: function isValid(uint32 tIndex) view returns(bool)
func (_IRToken *IRTokenCallerSession) IsValid(tIndex uint32) (bool, error) {
	return _IRToken.Contract.IsValid(&_IRToken.CallOpts, tIndex)
}

// AddT is a paid mutator transaction binding the contract method 0x8ff8e15a.
//
// Solidity: function addT(address t) returns(uint32 index)
func (_IRToken *IRTokenTransactor) AddT(opts *bind.TransactOpts, t common.Address) (*types.Transaction, error) {
	return _IRToken.contract.Transact(opts, "addT", t)
}

// AddT is a paid mutator transaction binding the contract method 0x8ff8e15a.
//
// Solidity: function addT(address t) returns(uint32 index)
func (_IRToken *IRTokenSession) AddT(t common.Address) (*types.Transaction, error) {
	return _IRToken.Contract.AddT(&_IRToken.TransactOpts, t)
}

// AddT is a paid mutator transaction binding the contract method 0x8ff8e15a.
//
// Solidity: function addT(address t) returns(uint32 index)
func (_IRToken *IRTokenTransactorSession) AddT(t common.Address) (*types.Transaction, error) {
	return _IRToken.Contract.AddT(&_IRToken.TransactOpts, t)
}

// IRoleMetaData contains all meta data concerning the IRole contract.
var IRoleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gIndex\",\"type\":\"uint64\"}],\"name\":\"addKeeperToGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"addProviderToGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_rType\",\"type\":\"uint8\"}],\"name\":\"checkIR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"checkT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"indexes\",\"type\":\"uint64[]\"},{\"internalType\":\"uint16\",\"name\":\"_level\",\"type\":\"uint16\"}],\"name\":\"createGroup\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getAddrGindex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getFsAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getGKNum\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getGroupInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ig\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ik\",\"type\":\"uint64\"}],\"name\":\"getGroupK\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"getRoleInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pledgeK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pledgeP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pledgePool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"registerKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"registerProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taddr\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"registerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isAdd\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"}],\"internalType\":\"structSGParams\",\"name\":\"ps\",\"type\":\"tuple\"}],\"name\":\"setGInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"kPledge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pPledge\",\"type\":\"uint256\"}],\"name\":\"setPledgeMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"withdrawFromFs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7977031c": "addKeeperToGroup(uint64,uint64)",
		"8ae21af3": "addProviderToGroup(uint64,uint64,bytes)",
		"cf8e99a8": "checkIR(uint64,uint8)",
		"de92e994": "checkT(uint32)",
		"f652391b": "createGroup(uint64[],uint16)",
		"9332aa6e": "getAddr(uint64)",
		"421795e5": "getAddrGindex(uint64)",
		"5f096376": "getFsAddr(uint64)",
		"429fb683": "getGKNum(uint64)",
		"4496c991": "getGroupInfo(uint64)",
		"64fe6290": "getGroupK(uint64,uint64)",
		"07483499": "getRoleInfo(address)",
		"a6ed590b": "pledgeK()",
		"8ba61d28": "pledgeP()",
		"de909560": "pledgePool()",
		"517985b0": "recharge(uint64,uint32,uint256,bytes)",
		"24b8fbf6": "register(address,bytes)",
		"10e35bbe": "registerKeeper(uint64,bytes,bytes)",
		"d57e8a4e": "registerProvider(uint64,bytes)",
		"09824a80": "registerToken(address)",
		"488cee1c": "registerUser(uint64,uint64,bytes,bytes)",
		"121ed07f": "setGInfo((uint64,bool,uint256,uint256))",
		"97948fda": "setPledgeMoney(uint256,uint256)",
		"d30d0ce5": "withdrawFromFs(uint64,uint32,uint256,bytes)",
	},
}

// IRoleABI is the input ABI used to generate the binding from.
// Deprecated: Use IRoleMetaData.ABI instead.
var IRoleABI = IRoleMetaData.ABI

// Deprecated: Use IRoleMetaData.Sigs instead.
// IRoleFuncSigs maps the 4-byte function signature to its string representation.
var IRoleFuncSigs = IRoleMetaData.Sigs

// IRole is an auto generated Go binding around an Ethereum contract.
type IRole struct {
	IRoleCaller     // Read-only binding to the contract
	IRoleTransactor // Write-only binding to the contract
	IRoleFilterer   // Log filterer for contract events
}

// IRoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRoleSession struct {
	Contract     *IRole            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRoleCallerSession struct {
	Contract *IRoleCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IRoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRoleTransactorSession struct {
	Contract     *IRoleTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRoleRaw struct {
	Contract *IRole // Generic contract binding to access the raw methods on
}

// IRoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRoleCallerRaw struct {
	Contract *IRoleCaller // Generic read-only contract binding to access the raw methods on
}

// IRoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRoleTransactorRaw struct {
	Contract *IRoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRole creates a new instance of IRole, bound to a specific deployed contract.
func NewIRole(address common.Address, backend bind.ContractBackend) (*IRole, error) {
	contract, err := bindIRole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRole{IRoleCaller: IRoleCaller{contract: contract}, IRoleTransactor: IRoleTransactor{contract: contract}, IRoleFilterer: IRoleFilterer{contract: contract}}, nil
}

// NewIRoleCaller creates a new read-only instance of IRole, bound to a specific deployed contract.
func NewIRoleCaller(address common.Address, caller bind.ContractCaller) (*IRoleCaller, error) {
	contract, err := bindIRole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRoleCaller{contract: contract}, nil
}

// NewIRoleTransactor creates a new write-only instance of IRole, bound to a specific deployed contract.
func NewIRoleTransactor(address common.Address, transactor bind.ContractTransactor) (*IRoleTransactor, error) {
	contract, err := bindIRole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRoleTransactor{contract: contract}, nil
}

// NewIRoleFilterer creates a new log filterer instance of IRole, bound to a specific deployed contract.
func NewIRoleFilterer(address common.Address, filterer bind.ContractFilterer) (*IRoleFilterer, error) {
	contract, err := bindIRole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRoleFilterer{contract: contract}, nil
}

// bindIRole binds a generic wrapper to an already deployed contract.
func bindIRole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRole *IRoleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRole.Contract.IRoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRole *IRoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRole.Contract.IRoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRole *IRoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRole.Contract.IRoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRole *IRoleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRole.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRole *IRoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRole.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRole *IRoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRole.Contract.contract.Transact(opts, method, params...)
}

// CheckIR is a free data retrieval call binding the contract method 0xcf8e99a8.
//
// Solidity: function checkIR(uint64 _index, uint8 _rType) view returns(address)
func (_IRole *IRoleCaller) CheckIR(opts *bind.CallOpts, _index uint64, _rType uint8) (common.Address, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "checkIR", _index, _rType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CheckIR is a free data retrieval call binding the contract method 0xcf8e99a8.
//
// Solidity: function checkIR(uint64 _index, uint8 _rType) view returns(address)
func (_IRole *IRoleSession) CheckIR(_index uint64, _rType uint8) (common.Address, error) {
	return _IRole.Contract.CheckIR(&_IRole.CallOpts, _index, _rType)
}

// CheckIR is a free data retrieval call binding the contract method 0xcf8e99a8.
//
// Solidity: function checkIR(uint64 _index, uint8 _rType) view returns(address)
func (_IRole *IRoleCallerSession) CheckIR(_index uint64, _rType uint8) (common.Address, error) {
	return _IRole.Contract.CheckIR(&_IRole.CallOpts, _index, _rType)
}

// CheckT is a free data retrieval call binding the contract method 0xde92e994.
//
// Solidity: function checkT(uint32 tIndex) view returns(address)
func (_IRole *IRoleCaller) CheckT(opts *bind.CallOpts, tIndex uint32) (common.Address, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "checkT", tIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CheckT is a free data retrieval call binding the contract method 0xde92e994.
//
// Solidity: function checkT(uint32 tIndex) view returns(address)
func (_IRole *IRoleSession) CheckT(tIndex uint32) (common.Address, error) {
	return _IRole.Contract.CheckT(&_IRole.CallOpts, tIndex)
}

// CheckT is a free data retrieval call binding the contract method 0xde92e994.
//
// Solidity: function checkT(uint32 tIndex) view returns(address)
func (_IRole *IRoleCallerSession) CheckT(tIndex uint32) (common.Address, error) {
	return _IRole.Contract.CheckT(&_IRole.CallOpts, tIndex)
}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 i) view returns(address)
func (_IRole *IRoleCaller) GetAddr(opts *bind.CallOpts, i uint64) (common.Address, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "getAddr", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 i) view returns(address)
func (_IRole *IRoleSession) GetAddr(i uint64) (common.Address, error) {
	return _IRole.Contract.GetAddr(&_IRole.CallOpts, i)
}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 i) view returns(address)
func (_IRole *IRoleCallerSession) GetAddr(i uint64) (common.Address, error) {
	return _IRole.Contract.GetAddr(&_IRole.CallOpts, i)
}

// GetAddrGindex is a free data retrieval call binding the contract method 0x421795e5.
//
// Solidity: function getAddrGindex(uint64 i) view returns(address, uint64)
func (_IRole *IRoleCaller) GetAddrGindex(opts *bind.CallOpts, i uint64) (common.Address, uint64, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "getAddrGindex", i)

	if err != nil {
		return *new(common.Address), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

// GetAddrGindex is a free data retrieval call binding the contract method 0x421795e5.
//
// Solidity: function getAddrGindex(uint64 i) view returns(address, uint64)
func (_IRole *IRoleSession) GetAddrGindex(i uint64) (common.Address, uint64, error) {
	return _IRole.Contract.GetAddrGindex(&_IRole.CallOpts, i)
}

// GetAddrGindex is a free data retrieval call binding the contract method 0x421795e5.
//
// Solidity: function getAddrGindex(uint64 i) view returns(address, uint64)
func (_IRole *IRoleCallerSession) GetAddrGindex(i uint64) (common.Address, uint64, error) {
	return _IRole.Contract.GetAddrGindex(&_IRole.CallOpts, i)
}

// GetFsAddr is a free data retrieval call binding the contract method 0x5f096376.
//
// Solidity: function getFsAddr(uint64 i) view returns(address)
func (_IRole *IRoleCaller) GetFsAddr(opts *bind.CallOpts, i uint64) (common.Address, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "getFsAddr", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFsAddr is a free data retrieval call binding the contract method 0x5f096376.
//
// Solidity: function getFsAddr(uint64 i) view returns(address)
func (_IRole *IRoleSession) GetFsAddr(i uint64) (common.Address, error) {
	return _IRole.Contract.GetFsAddr(&_IRole.CallOpts, i)
}

// GetFsAddr is a free data retrieval call binding the contract method 0x5f096376.
//
// Solidity: function getFsAddr(uint64 i) view returns(address)
func (_IRole *IRoleCallerSession) GetFsAddr(i uint64) (common.Address, error) {
	return _IRole.Contract.GetFsAddr(&_IRole.CallOpts, i)
}

// GetGKNum is a free data retrieval call binding the contract method 0x429fb683.
//
// Solidity: function getGKNum(uint64 i) view returns(uint64)
func (_IRole *IRoleCaller) GetGKNum(opts *bind.CallOpts, i uint64) (uint64, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "getGKNum", i)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetGKNum is a free data retrieval call binding the contract method 0x429fb683.
//
// Solidity: function getGKNum(uint64 i) view returns(uint64)
func (_IRole *IRoleSession) GetGKNum(i uint64) (uint64, error) {
	return _IRole.Contract.GetGKNum(&_IRole.CallOpts, i)
}

// GetGKNum is a free data retrieval call binding the contract method 0x429fb683.
//
// Solidity: function getGKNum(uint64 i) view returns(uint64)
func (_IRole *IRoleCallerSession) GetGKNum(i uint64) (uint64, error) {
	return _IRole.Contract.GetGKNum(&_IRole.CallOpts, i)
}

// GetGroupInfo is a free data retrieval call binding the contract method 0x4496c991.
//
// Solidity: function getGroupInfo(uint64 i) view returns(bool, bool, bool, uint16, uint256, uint256, address)
func (_IRole *IRoleCaller) GetGroupInfo(opts *bind.CallOpts, i uint64) (bool, bool, bool, uint16, *big.Int, *big.Int, common.Address, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "getGroupInfo", i)

	if err != nil {
		return *new(bool), *new(bool), *new(bool), *new(uint16), *new(*big.Int), *new(*big.Int), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)
	out3 := *abi.ConvertType(out[3], new(uint16)).(*uint16)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// GetGroupInfo is a free data retrieval call binding the contract method 0x4496c991.
//
// Solidity: function getGroupInfo(uint64 i) view returns(bool, bool, bool, uint16, uint256, uint256, address)
func (_IRole *IRoleSession) GetGroupInfo(i uint64) (bool, bool, bool, uint16, *big.Int, *big.Int, common.Address, error) {
	return _IRole.Contract.GetGroupInfo(&_IRole.CallOpts, i)
}

// GetGroupInfo is a free data retrieval call binding the contract method 0x4496c991.
//
// Solidity: function getGroupInfo(uint64 i) view returns(bool, bool, bool, uint16, uint256, uint256, address)
func (_IRole *IRoleCallerSession) GetGroupInfo(i uint64) (bool, bool, bool, uint16, *big.Int, *big.Int, common.Address, error) {
	return _IRole.Contract.GetGroupInfo(&_IRole.CallOpts, i)
}

// GetGroupK is a free data retrieval call binding the contract method 0x64fe6290.
//
// Solidity: function getGroupK(uint64 ig, uint64 ik) view returns(uint64)
func (_IRole *IRoleCaller) GetGroupK(opts *bind.CallOpts, ig uint64, ik uint64) (uint64, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "getGroupK", ig, ik)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetGroupK is a free data retrieval call binding the contract method 0x64fe6290.
//
// Solidity: function getGroupK(uint64 ig, uint64 ik) view returns(uint64)
func (_IRole *IRoleSession) GetGroupK(ig uint64, ik uint64) (uint64, error) {
	return _IRole.Contract.GetGroupK(&_IRole.CallOpts, ig, ik)
}

// GetGroupK is a free data retrieval call binding the contract method 0x64fe6290.
//
// Solidity: function getGroupK(uint64 ig, uint64 ik) view returns(uint64)
func (_IRole *IRoleCallerSession) GetGroupK(ig uint64, ik uint64) (uint64, error) {
	return _IRole.Contract.GetGroupK(&_IRole.CallOpts, ig, ik)
}

// GetRoleInfo is a free data retrieval call binding the contract method 0x07483499.
//
// Solidity: function getRoleInfo(address acc) view returns(bool, bool, uint8, uint64, uint64, bytes)
func (_IRole *IRoleCaller) GetRoleInfo(opts *bind.CallOpts, acc common.Address) (bool, bool, uint8, uint64, uint64, []byte, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "getRoleInfo", acc)

	if err != nil {
		return *new(bool), *new(bool), *new(uint8), *new(uint64), *new(uint64), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(uint8)).(*uint8)
	out3 := *abi.ConvertType(out[3], new(uint64)).(*uint64)
	out4 := *abi.ConvertType(out[4], new(uint64)).(*uint64)
	out5 := *abi.ConvertType(out[5], new([]byte)).(*[]byte)

	return out0, out1, out2, out3, out4, out5, err

}

// GetRoleInfo is a free data retrieval call binding the contract method 0x07483499.
//
// Solidity: function getRoleInfo(address acc) view returns(bool, bool, uint8, uint64, uint64, bytes)
func (_IRole *IRoleSession) GetRoleInfo(acc common.Address) (bool, bool, uint8, uint64, uint64, []byte, error) {
	return _IRole.Contract.GetRoleInfo(&_IRole.CallOpts, acc)
}

// GetRoleInfo is a free data retrieval call binding the contract method 0x07483499.
//
// Solidity: function getRoleInfo(address acc) view returns(bool, bool, uint8, uint64, uint64, bytes)
func (_IRole *IRoleCallerSession) GetRoleInfo(acc common.Address) (bool, bool, uint8, uint64, uint64, []byte, error) {
	return _IRole.Contract.GetRoleInfo(&_IRole.CallOpts, acc)
}

// PledgeK is a free data retrieval call binding the contract method 0xa6ed590b.
//
// Solidity: function pledgeK() view returns(uint256)
func (_IRole *IRoleCaller) PledgeK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "pledgeK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PledgeK is a free data retrieval call binding the contract method 0xa6ed590b.
//
// Solidity: function pledgeK() view returns(uint256)
func (_IRole *IRoleSession) PledgeK() (*big.Int, error) {
	return _IRole.Contract.PledgeK(&_IRole.CallOpts)
}

// PledgeK is a free data retrieval call binding the contract method 0xa6ed590b.
//
// Solidity: function pledgeK() view returns(uint256)
func (_IRole *IRoleCallerSession) PledgeK() (*big.Int, error) {
	return _IRole.Contract.PledgeK(&_IRole.CallOpts)
}

// PledgeP is a free data retrieval call binding the contract method 0x8ba61d28.
//
// Solidity: function pledgeP() view returns(uint256)
func (_IRole *IRoleCaller) PledgeP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "pledgeP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PledgeP is a free data retrieval call binding the contract method 0x8ba61d28.
//
// Solidity: function pledgeP() view returns(uint256)
func (_IRole *IRoleSession) PledgeP() (*big.Int, error) {
	return _IRole.Contract.PledgeP(&_IRole.CallOpts)
}

// PledgeP is a free data retrieval call binding the contract method 0x8ba61d28.
//
// Solidity: function pledgeP() view returns(uint256)
func (_IRole *IRoleCallerSession) PledgeP() (*big.Int, error) {
	return _IRole.Contract.PledgeP(&_IRole.CallOpts)
}

// PledgePool is a free data retrieval call binding the contract method 0xde909560.
//
// Solidity: function pledgePool() view returns(address)
func (_IRole *IRoleCaller) PledgePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "pledgePool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PledgePool is a free data retrieval call binding the contract method 0xde909560.
//
// Solidity: function pledgePool() view returns(address)
func (_IRole *IRoleSession) PledgePool() (common.Address, error) {
	return _IRole.Contract.PledgePool(&_IRole.CallOpts)
}

// PledgePool is a free data retrieval call binding the contract method 0xde909560.
//
// Solidity: function pledgePool() view returns(address)
func (_IRole *IRoleCallerSession) PledgePool() (common.Address, error) {
	return _IRole.Contract.PledgePool(&_IRole.CallOpts)
}

// AddKeeperToGroup is a paid mutator transaction binding the contract method 0x7977031c.
//
// Solidity: function addKeeperToGroup(uint64 _index, uint64 _gIndex) returns()
func (_IRole *IRoleTransactor) AddKeeperToGroup(opts *bind.TransactOpts, _index uint64, _gIndex uint64) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "addKeeperToGroup", _index, _gIndex)
}

// AddKeeperToGroup is a paid mutator transaction binding the contract method 0x7977031c.
//
// Solidity: function addKeeperToGroup(uint64 _index, uint64 _gIndex) returns()
func (_IRole *IRoleSession) AddKeeperToGroup(_index uint64, _gIndex uint64) (*types.Transaction, error) {
	return _IRole.Contract.AddKeeperToGroup(&_IRole.TransactOpts, _index, _gIndex)
}

// AddKeeperToGroup is a paid mutator transaction binding the contract method 0x7977031c.
//
// Solidity: function addKeeperToGroup(uint64 _index, uint64 _gIndex) returns()
func (_IRole *IRoleTransactorSession) AddKeeperToGroup(_index uint64, _gIndex uint64) (*types.Transaction, error) {
	return _IRole.Contract.AddKeeperToGroup(&_IRole.TransactOpts, _index, _gIndex)
}

// AddProviderToGroup is a paid mutator transaction binding the contract method 0x8ae21af3.
//
// Solidity: function addProviderToGroup(uint64 _index, uint64 _gIndex, bytes sign) returns()
func (_IRole *IRoleTransactor) AddProviderToGroup(opts *bind.TransactOpts, _index uint64, _gIndex uint64, sign []byte) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "addProviderToGroup", _index, _gIndex, sign)
}

// AddProviderToGroup is a paid mutator transaction binding the contract method 0x8ae21af3.
//
// Solidity: function addProviderToGroup(uint64 _index, uint64 _gIndex, bytes sign) returns()
func (_IRole *IRoleSession) AddProviderToGroup(_index uint64, _gIndex uint64, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.AddProviderToGroup(&_IRole.TransactOpts, _index, _gIndex, sign)
}

// AddProviderToGroup is a paid mutator transaction binding the contract method 0x8ae21af3.
//
// Solidity: function addProviderToGroup(uint64 _index, uint64 _gIndex, bytes sign) returns()
func (_IRole *IRoleTransactorSession) AddProviderToGroup(_index uint64, _gIndex uint64, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.AddProviderToGroup(&_IRole.TransactOpts, _index, _gIndex, sign)
}

// CreateGroup is a paid mutator transaction binding the contract method 0xf652391b.
//
// Solidity: function createGroup(uint64[] indexes, uint16 _level) returns(uint64)
func (_IRole *IRoleTransactor) CreateGroup(opts *bind.TransactOpts, indexes []uint64, _level uint16) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "createGroup", indexes, _level)
}

// CreateGroup is a paid mutator transaction binding the contract method 0xf652391b.
//
// Solidity: function createGroup(uint64[] indexes, uint16 _level) returns(uint64)
func (_IRole *IRoleSession) CreateGroup(indexes []uint64, _level uint16) (*types.Transaction, error) {
	return _IRole.Contract.CreateGroup(&_IRole.TransactOpts, indexes, _level)
}

// CreateGroup is a paid mutator transaction binding the contract method 0xf652391b.
//
// Solidity: function createGroup(uint64[] indexes, uint16 _level) returns(uint64)
func (_IRole *IRoleTransactorSession) CreateGroup(indexes []uint64, _level uint16) (*types.Transaction, error) {
	return _IRole.Contract.CreateGroup(&_IRole.TransactOpts, indexes, _level)
}

// Recharge is a paid mutator transaction binding the contract method 0x517985b0.
//
// Solidity: function recharge(uint64 uIndex, uint32 tokenIndex, uint256 money, bytes sign) payable returns()
func (_IRole *IRoleTransactor) Recharge(opts *bind.TransactOpts, uIndex uint64, tokenIndex uint32, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "recharge", uIndex, tokenIndex, money, sign)
}

// Recharge is a paid mutator transaction binding the contract method 0x517985b0.
//
// Solidity: function recharge(uint64 uIndex, uint32 tokenIndex, uint256 money, bytes sign) payable returns()
func (_IRole *IRoleSession) Recharge(uIndex uint64, tokenIndex uint32, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.Recharge(&_IRole.TransactOpts, uIndex, tokenIndex, money, sign)
}

// Recharge is a paid mutator transaction binding the contract method 0x517985b0.
//
// Solidity: function recharge(uint64 uIndex, uint32 tokenIndex, uint256 money, bytes sign) payable returns()
func (_IRole *IRoleTransactorSession) Recharge(uIndex uint64, tokenIndex uint32, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.Recharge(&_IRole.TransactOpts, uIndex, tokenIndex, money, sign)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address addr, bytes sign) returns(uint64)
func (_IRole *IRoleTransactor) Register(opts *bind.TransactOpts, addr common.Address, sign []byte) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "register", addr, sign)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address addr, bytes sign) returns(uint64)
func (_IRole *IRoleSession) Register(addr common.Address, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.Register(&_IRole.TransactOpts, addr, sign)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address addr, bytes sign) returns(uint64)
func (_IRole *IRoleTransactorSession) Register(addr common.Address, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.Register(&_IRole.TransactOpts, addr, sign)
}

// RegisterKeeper is a paid mutator transaction binding the contract method 0x10e35bbe.
//
// Solidity: function registerKeeper(uint64 index, bytes blsKey, bytes sign) returns()
func (_IRole *IRoleTransactor) RegisterKeeper(opts *bind.TransactOpts, index uint64, blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "registerKeeper", index, blsKey, sign)
}

// RegisterKeeper is a paid mutator transaction binding the contract method 0x10e35bbe.
//
// Solidity: function registerKeeper(uint64 index, bytes blsKey, bytes sign) returns()
func (_IRole *IRoleSession) RegisterKeeper(index uint64, blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.RegisterKeeper(&_IRole.TransactOpts, index, blsKey, sign)
}

// RegisterKeeper is a paid mutator transaction binding the contract method 0x10e35bbe.
//
// Solidity: function registerKeeper(uint64 index, bytes blsKey, bytes sign) returns()
func (_IRole *IRoleTransactorSession) RegisterKeeper(index uint64, blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.RegisterKeeper(&_IRole.TransactOpts, index, blsKey, sign)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0xd57e8a4e.
//
// Solidity: function registerProvider(uint64 index, bytes sign) returns()
func (_IRole *IRoleTransactor) RegisterProvider(opts *bind.TransactOpts, index uint64, sign []byte) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "registerProvider", index, sign)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0xd57e8a4e.
//
// Solidity: function registerProvider(uint64 index, bytes sign) returns()
func (_IRole *IRoleSession) RegisterProvider(index uint64, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.RegisterProvider(&_IRole.TransactOpts, index, sign)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0xd57e8a4e.
//
// Solidity: function registerProvider(uint64 index, bytes sign) returns()
func (_IRole *IRoleTransactorSession) RegisterProvider(index uint64, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.RegisterProvider(&_IRole.TransactOpts, index, sign)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address taddr) returns()
func (_IRole *IRoleTransactor) RegisterToken(opts *bind.TransactOpts, taddr common.Address) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "registerToken", taddr)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address taddr) returns()
func (_IRole *IRoleSession) RegisterToken(taddr common.Address) (*types.Transaction, error) {
	return _IRole.Contract.RegisterToken(&_IRole.TransactOpts, taddr)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address taddr) returns()
func (_IRole *IRoleTransactorSession) RegisterToken(taddr common.Address) (*types.Transaction, error) {
	return _IRole.Contract.RegisterToken(&_IRole.TransactOpts, taddr)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x488cee1c.
//
// Solidity: function registerUser(uint64 index, uint64 _gIndex, bytes blsKey, bytes sign) returns()
func (_IRole *IRoleTransactor) RegisterUser(opts *bind.TransactOpts, index uint64, _gIndex uint64, blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "registerUser", index, _gIndex, blsKey, sign)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x488cee1c.
//
// Solidity: function registerUser(uint64 index, uint64 _gIndex, bytes blsKey, bytes sign) returns()
func (_IRole *IRoleSession) RegisterUser(index uint64, _gIndex uint64, blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.RegisterUser(&_IRole.TransactOpts, index, _gIndex, blsKey, sign)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x488cee1c.
//
// Solidity: function registerUser(uint64 index, uint64 _gIndex, bytes blsKey, bytes sign) returns()
func (_IRole *IRoleTransactorSession) RegisterUser(index uint64, _gIndex uint64, blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.RegisterUser(&_IRole.TransactOpts, index, _gIndex, blsKey, sign)
}

// SetGInfo is a paid mutator transaction binding the contract method 0x121ed07f.
//
// Solidity: function setGInfo((uint64,bool,uint256,uint256) ps) returns()
func (_IRole *IRoleTransactor) SetGInfo(opts *bind.TransactOpts, ps SGParams) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "setGInfo", ps)
}

// SetGInfo is a paid mutator transaction binding the contract method 0x121ed07f.
//
// Solidity: function setGInfo((uint64,bool,uint256,uint256) ps) returns()
func (_IRole *IRoleSession) SetGInfo(ps SGParams) (*types.Transaction, error) {
	return _IRole.Contract.SetGInfo(&_IRole.TransactOpts, ps)
}

// SetGInfo is a paid mutator transaction binding the contract method 0x121ed07f.
//
// Solidity: function setGInfo((uint64,bool,uint256,uint256) ps) returns()
func (_IRole *IRoleTransactorSession) SetGInfo(ps SGParams) (*types.Transaction, error) {
	return _IRole.Contract.SetGInfo(&_IRole.TransactOpts, ps)
}

// SetPledgeMoney is a paid mutator transaction binding the contract method 0x97948fda.
//
// Solidity: function setPledgeMoney(uint256 kPledge, uint256 pPledge) returns()
func (_IRole *IRoleTransactor) SetPledgeMoney(opts *bind.TransactOpts, kPledge *big.Int, pPledge *big.Int) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "setPledgeMoney", kPledge, pPledge)
}

// SetPledgeMoney is a paid mutator transaction binding the contract method 0x97948fda.
//
// Solidity: function setPledgeMoney(uint256 kPledge, uint256 pPledge) returns()
func (_IRole *IRoleSession) SetPledgeMoney(kPledge *big.Int, pPledge *big.Int) (*types.Transaction, error) {
	return _IRole.Contract.SetPledgeMoney(&_IRole.TransactOpts, kPledge, pPledge)
}

// SetPledgeMoney is a paid mutator transaction binding the contract method 0x97948fda.
//
// Solidity: function setPledgeMoney(uint256 kPledge, uint256 pPledge) returns()
func (_IRole *IRoleTransactorSession) SetPledgeMoney(kPledge *big.Int, pPledge *big.Int) (*types.Transaction, error) {
	return _IRole.Contract.SetPledgeMoney(&_IRole.TransactOpts, kPledge, pPledge)
}

// WithdrawFromFs is a paid mutator transaction binding the contract method 0xd30d0ce5.
//
// Solidity: function withdrawFromFs(uint64 index, uint32 tIndex, uint256 amount, bytes sign) returns()
func (_IRole *IRoleTransactor) WithdrawFromFs(opts *bind.TransactOpts, index uint64, tIndex uint32, amount *big.Int, sign []byte) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "withdrawFromFs", index, tIndex, amount, sign)
}

// WithdrawFromFs is a paid mutator transaction binding the contract method 0xd30d0ce5.
//
// Solidity: function withdrawFromFs(uint64 index, uint32 tIndex, uint256 amount, bytes sign) returns()
func (_IRole *IRoleSession) WithdrawFromFs(index uint64, tIndex uint32, amount *big.Int, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.WithdrawFromFs(&_IRole.TransactOpts, index, tIndex, amount, sign)
}

// WithdrawFromFs is a paid mutator transaction binding the contract method 0xd30d0ce5.
//
// Solidity: function withdrawFromFs(uint64 index, uint32 tIndex, uint256 amount, bytes sign) returns()
func (_IRole *IRoleTransactorSession) WithdrawFromFs(index uint64, tIndex uint32, amount *big.Int, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.WithdrawFromFs(&_IRole.TransactOpts, index, tIndex, amount, sign)
}

// RecoverMetaData contains all meta data concerning the Recover contract.
var RecoverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"19045a25": "recover(bytes32,bytes)",
	},
	Bin: "0x6102a861003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c806319045a251461003a575b600080fd5b61004d610048366004610184565b610069565b6040516001600160a01b03909116815260200160405180910390f35b6000815160411461007c57506000610168565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156100c25760009350505050610168565b601b8160ff1610156100dc576100d981601b61023f565b90505b8060ff16601b141580156100f457508060ff16601c14155b156101055760009350505050610168565b60408051600081526020810180835288905260ff831691810191909152606081018490526080810183905260019060a0016020604051602081039080840390855afa158015610158573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561019757600080fd5b82359150602083013567ffffffffffffffff808211156101b657600080fd5b818501915085601f8301126101ca57600080fd5b8135818111156101dc576101dc61016e565b604051601f8201601f19908116603f011681019083821181831017156102045761020461016e565b8160405282815288602084870101111561021d57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b600060ff821660ff84168060ff0382111561026a57634e487b7160e01b600052601160045260246000fd5b01939250505056fea26469706673582212205d9f1d7211039fd5443ec28a75201eb5cc7f5e80053b93a182589f9091d0796664736f6c634300080c0033",
}

// RecoverABI is the input ABI used to generate the binding from.
// Deprecated: Use RecoverMetaData.ABI instead.
var RecoverABI = RecoverMetaData.ABI

// Deprecated: Use RecoverMetaData.Sigs instead.
// RecoverFuncSigs maps the 4-byte function signature to its string representation.
var RecoverFuncSigs = RecoverMetaData.Sigs

// RecoverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RecoverMetaData.Bin instead.
var RecoverBin = RecoverMetaData.Bin

// DeployRecover deploys a new Ethereum contract, binding an instance of Recover to it.
func DeployRecover(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Recover, error) {
	parsed, err := RecoverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RecoverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Recover{RecoverCaller: RecoverCaller{contract: contract}, RecoverTransactor: RecoverTransactor{contract: contract}, RecoverFilterer: RecoverFilterer{contract: contract}}, nil
}

// Recover is an auto generated Go binding around an Ethereum contract.
type Recover struct {
	RecoverCaller     // Read-only binding to the contract
	RecoverTransactor // Write-only binding to the contract
	RecoverFilterer   // Log filterer for contract events
}

// RecoverCaller is an auto generated read-only Go binding around an Ethereum contract.
type RecoverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RecoverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RecoverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RecoverSession struct {
	Contract     *Recover          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RecoverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RecoverCallerSession struct {
	Contract *RecoverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RecoverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RecoverTransactorSession struct {
	Contract     *RecoverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RecoverRaw is an auto generated low-level Go binding around an Ethereum contract.
type RecoverRaw struct {
	Contract *Recover // Generic contract binding to access the raw methods on
}

// RecoverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RecoverCallerRaw struct {
	Contract *RecoverCaller // Generic read-only contract binding to access the raw methods on
}

// RecoverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RecoverTransactorRaw struct {
	Contract *RecoverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRecover creates a new instance of Recover, bound to a specific deployed contract.
func NewRecover(address common.Address, backend bind.ContractBackend) (*Recover, error) {
	contract, err := bindRecover(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Recover{RecoverCaller: RecoverCaller{contract: contract}, RecoverTransactor: RecoverTransactor{contract: contract}, RecoverFilterer: RecoverFilterer{contract: contract}}, nil
}

// NewRecoverCaller creates a new read-only instance of Recover, bound to a specific deployed contract.
func NewRecoverCaller(address common.Address, caller bind.ContractCaller) (*RecoverCaller, error) {
	contract, err := bindRecover(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecoverCaller{contract: contract}, nil
}

// NewRecoverTransactor creates a new write-only instance of Recover, bound to a specific deployed contract.
func NewRecoverTransactor(address common.Address, transactor bind.ContractTransactor) (*RecoverTransactor, error) {
	contract, err := bindRecover(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecoverTransactor{contract: contract}, nil
}

// NewRecoverFilterer creates a new log filterer instance of Recover, bound to a specific deployed contract.
func NewRecoverFilterer(address common.Address, filterer bind.ContractFilterer) (*RecoverFilterer, error) {
	contract, err := bindRecover(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecoverFilterer{contract: contract}, nil
}

// bindRecover binds a generic wrapper to an already deployed contract.
func bindRecover(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RecoverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recover *RecoverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recover.Contract.RecoverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recover *RecoverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recover.Contract.RecoverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recover *RecoverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recover.Contract.RecoverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recover *RecoverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recover.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recover *RecoverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recover.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recover *RecoverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recover.Contract.contract.Transact(opts, method, params...)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) pure returns(address)
func (_Recover *RecoverCaller) Recover(opts *bind.CallOpts, hash [32]byte, signature []byte) (common.Address, error) {
	var out []interface{}
	err := _Recover.contract.Call(opts, &out, "recover", hash, signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) pure returns(address)
func (_Recover *RecoverSession) Recover(hash [32]byte, signature []byte) (common.Address, error) {
	return _Recover.Contract.Recover(&_Recover.CallOpts, hash, signature)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) pure returns(address)
func (_Recover *RecoverCallerSession) Recover(hash [32]byte, signature []byte) (common.Address, error) {
	return _Recover.Contract.Recover(&_Recover.CallOpts, hash, signature)
}

// RoleFSMetaData contains all meta data concerning the RoleFS contract.
var RoleFSMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"usign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"psign\",\"type\":\"bytes\"}],\"internalType\":\"structAOParams\",\"name\":\"ps\",\"type\":\"tuple\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nPIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pSign\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"kSigns\",\"type\":\"bytes[]\"}],\"name\":\"addRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"pAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"},{\"internalType\":\"uint64[]\",\"name\":\"kIndexes\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"internalType\":\"structPWParams\",\"name\":\"ps\",\"type\":\"tuple\"}],\"name\":\"proWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"iss\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rt\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"kIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"usign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"psign\",\"type\":\"bytes\"}],\"internalType\":\"structSOParams\",\"name\":\"ps\",\"type\":\"tuple\"}],\"name\":\"subOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nPIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pSign\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"kSigns\",\"type\":\"bytes[]\"}],\"name\":\"subRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0e09662b": "addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes))",
		"ffd0fd09": "addRepair(uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes[])",
		"edf26f5c": "proWithdraw((uint64,uint32,address,address,uint256,uint256,uint64[],bytes[]))",
		"6153c8c1": "setAddr(address,address,address)",
		"abfc55f1": "subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes))",
		"eb975779": "subRepair(uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes[])",
	},
	Bin: "0x608060405234801561001057600080fd5b50600280546001600160a01b03191633179055613a81806100326000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630e09662b146100675780636153c8c11461007c578063abfc55f11461008f578063eb975779146100a2578063edf26f5c146100b5578063ffd0fd09146100c8575b600080fd5b61007a610075366004612c3b565b6100db565b005b61007a61008a366004612d73565b610ae7565b61007a61009d366004612dbe565b610b64565b61007a6100b0366004612f78565b6112a2565b61007a6100c33660046130bf565b6118b9565b61007a6100d6366004612f78565b61205a565b80516001600160401b03166100f2576100f26131a8565b60208101516001600160401b031661010c5761010c6131a8565b805160009061011d906001906131d4565b905060006001836020015161013291906131d4565b60008054604051634999553760e11b81526001600160401b038616600482015292935090916001600160a01b0390911690639332aa6e90602401602060405180830381865afa158015610189573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101ad91906131fc565b60008054604051634999553760e11b81526001600160401b038616600482015292935090916001600160a01b0390911690639332aa6e90602401602060405180830381865afa158015610204573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061022891906131fc565b6000805460405163421795e560e01b81526001600160401b038816600482015292935090916001600160a01b039091169063421795e5906024016040805180830381865afa15801561027e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102a29190613220565b9050809150506000866000015187602001518860a0015189604001518a606001518b608001518c60c001518d60e001516040516020016102e998979695949392919061325a565b60408051601f198184030181529082905280516020909101206101008901516319045a2560e01b835290925060009173__$942d0120f12f58797d9427bd9464d63c7a$__916319045a259161034291869160040161331c565b602060405180830381865af415801561035f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061038391906131fc565b9050846001600160a01b0316816001600160a01b0316146103a6576103a66131a8565b6101208801516040516319045a2560e01b815273__$942d0120f12f58797d9427bd9464d63c7a$__916319045a25916103e391869160040161331c565b602060405180830381865af4158015610400573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061042491906131fc565b9050836001600160a01b0316816001600160a01b031614610447576104476131a8565b600088608001516001600160401b03161161048e5760405162461bcd60e51b815260206004820152600260248201526139bd60f11b60448201526064015b60405180910390fd5b87604001516001600160401b031688606001516001600160401b0316116104dc5760405162461bcd60e51b8152602060048201526002602482015261657360f01b6044820152606401610485565b60608801516001600160401b038116906104fa90620151809061333d565b6105079062015180613371565b6001600160401b0316146105425760405162461bcd60e51b8152602060048201526002602482015261746560f01b6044820152606401610485565b60005488516040516319f1d33560e31b81526001600160401b039091166004820152600160248201526001600160a01b039091169063cf8e99a890604401602060405180830381865afa15801561059d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c191906131fc565b5060005460208901516040516319f1d33560e31b81526001600160401b039091166004820152600260248201526001600160a01b039091169063cf8e99a890604401602060405180830381865afa158015610620573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061064491906131fc565b506000610661338a600001518b602001518c60c00151600161269c565b6000549092506001600160a01b03169050635f0963766106826001876131d4565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa1580156106c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ea91906131fc565b6001600160a01b0316630e09662b8a6040518263ffffffff1660e01b815260040161071591906133a0565b600060405180830381600087803b15801561072f57600080fd5b505af1158015610743573d6000803e3d6000fd5b505060408051608081018252600080825260208201819052918101829052606081019190915291506107729050565b6040808b01516001600160401b0390811683526060808d01518216602085015260808d01519091169183019190915260e08b01519082015260c08a015163ffffffff16610adb5760015460408051630c862ea560e41b815283516001600160401b039081166004830152602085015181166024830152918401519091166044820152606083015160648201526001600160a01b0390911690600090829063c862ea50906084016020604051808303816000875af1158015610837573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085b919061348f565b9050610893604051806080016040528060006001600160401b0316815260200160001515815260200160008152602001600081525090565b61089e6001866131d4565b6001600160401b0390811682526001602083015260808e01511660408083019190915260e08e01516060830152600054905163121ed07f60e01b81526001600160a01b039091169063121ed07f9061092b90849060040181516001600160401b03168152602080830151151590820152604080830151908201526060918201519181019190915260800190565b600060405180830381600087803b15801561094557600080fd5b505af1158015610959573d6000803e3d6000fd5b5050505081600014156109755750505050505050505050505050565b600354604051637bdbeb5b60e11b8152600060048201819052916001600160a01b03169063f7b7d6b690602401602060405180830381865afa1580156109bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109e391906131fc565b90506109ed612aa6565b816001600160a01b0316636bd1ce8960008054906101000a90046001600160a01b03166001600160a01b031663de9095606040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a4d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a7191906131fc565b86846040518463ffffffff1660e01b8152600401610a91939291906134a8565b6020604051808303816000875af1158015610ab0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ad49190613520565b5050505050505b50505050505050505050565b6002546001600160a01b03163314610b255760405162461bcd60e51b81526020600482015260016024820152602760f91b6044820152606401610485565b600180546001600160a01b039485166001600160a01b031991821617909155600080549385169382169390931790925560038054919093169116179055565b60208101516001600160401b0316610b7e57610b7e6131a8565b60408101516001600160401b0316610b9857610b986131a8565b600060018260200151610bab91906131d4565b9050600060018360400151610bc091906131d4565b60008054604051634999553760e11b81526001600160401b038616600482015292935090916001600160a01b0390911690639332aa6e90602401602060405180830381865afa158015610c17573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3b91906131fc565b60008054604051634999553760e11b81526001600160401b038616600482015292935090916001600160a01b0390911690639332aa6e90602401602060405180830381865afa158015610c92573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cb691906131fc565b6000805460405163421795e560e01b81526001600160401b038816600482015292935090916001600160a01b039091169063421795e5906024016040805180830381865afa158015610d0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d309190613220565b9050809150506000866020015187604001518860c0015189606001518a608001518b60a001518c60e001518d6101000151604051602001610d7898979695949392919061325a565b60408051601f198184030181529082905280516020909101206101208901516319045a2560e01b835290925060009173__$942d0120f12f58797d9427bd9464d63c7a$__916319045a2591610dd191869160040161331c565b602060405180830381865af4158015610dee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1291906131fc565b9050846001600160a01b0316816001600160a01b031614610e3557610e356131a8565b6101408801516040516319045a2560e01b815273__$942d0120f12f58797d9427bd9464d63c7a$__916319045a2591610e7291869160040161331c565b602060405180830381865af4158015610e8f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eb391906131fc565b9050836001600160a01b0316816001600160a01b031614610ed657610ed66131a8565b60005460208901516040516319f1d33560e31b81526001600160401b039091166004820152600160248201526001600160a01b039091169063cf8e99a890604401602060405180830381865afa158015610f34573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f5891906131fc565b5060005460408981015190516319f1d33560e31b81526001600160401b039091166004820152600260248201526001600160a01b039091169063cf8e99a890604401602060405180830381865afa158015610fb7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fdb91906131fc565b50600080610ff9338b602001518c604001518d60e00151600261269c565b9150915060008a60a001516001600160401b03161161103f5760405162461bcd60e51b8152602060048201526002602482015261455360f01b6044820152606401610485565b89606001516001600160401b03168a608001516001600160401b03161180156110755750428a608001516001600160401b031611155b6110a65760405162461bcd60e51b8152602060048201526002602482015261115560f21b6044820152606401610485565b6001600160401b0382168a526000546001600160a01b0316635f0963766110ce6001846131d4565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa158015611112573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061113691906131fc565b6001600160a01b031663abfc55f18b6040518263ffffffff1660e01b8152600401611161919061353b565b600060405180830381600087803b15801561117b57600080fd5b505af115801561118f573d6000803e3d6000fd5b505050506111c9604051806080016040528060006001600160401b0316815260200160001515815260200160008152602001600081525090565b6111d46001836131d4565b6001600160401b0390811682526000602083015260a08c01511660408201526101008b0151606082015260e08b015163ffffffff16611295576000546040805163121ed07f60e01b815283516001600160401b03166004820152602084015115156024820152908301516044820152606083015160648201526001600160a01b039091169063121ed07f90608401600060405180830381600087803b15801561127c57600080fd5b505af1158015611290573d6000803e3d6000fd5b505050505b5050505050505050505050565b6000546040516319f1d33560e31b81526001600160401b038c166004820152600260248201526001600160a01b039091169063cf8e99a890604401602060405180830381865afa1580156112fa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061131e91906131fc565b506000546040516319f1d33560e31b81526001600160401b038b166004820152600260248201526001600160a01b039091169063cf8e99a890604401602060405180830381865afa158015611377573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061139b91906131fc565b506000866001600160401b03161180156113c65750876001600160401b0316876001600160401b0316115b80156113db575042876001600160401b031611155b6113f75760405162461bcd60e51b815260040161048590613622565b600080611408338d8d89600261269c565b9150915060008c8b8b8b8b8b8b60405160200161148b979695949392919060c097881b6001600160c01b0319908116825296881b8716600882015294871b8616601086015292861b85166018850152941b909216602082015260e09290921b6001600160e01b0319166028830152602c820152607360f81b604c820152604d0190565b60405160208183030381529060405280519060200120905061154460008054906101000a90046001600160a01b03166001600160a01b0316639332aa6e60018f6114d591906131d4565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa158015611519573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061153d91906131fc565b82876129e1565b600080546003906001600160a01b031663429fb6836115646001876131d4565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa1580156115a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115cc919061363e565b6115d7906002613371565b6115e1919061333d565b6001600160401b03169050808551101561160d5760405162461bcd60e51b81526004016104859061365b565b60008060005b8751811015611769578473__$942d0120f12f58797d9427bd9464d63c7a$__6319045a2590918a848151811061164b5761164b613678565b60200260200101516040518363ffffffff1660e01b815260040161167092919061331c565b602060405180830381865af415801561168d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116b191906131fc565b60008054604051630748349960e01b81526001600160a01b03808516600483015293955091921690630748349990602401600060405180830381865afa1580156116ff573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611727919081019061368e565b50945050505050866001600160401b0316816001600160401b03161415611756578361175281613767565b9450505b50806117618161378e565b915050611613565b5082826001600160401b031610156117935760405162461bcd60e51b81526004016104859061365b565b50506000546001600160a01b03169150635f09637690506117b56001846131d4565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa1580156117f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061181d91906131fc565b604051637581851960e01b81526001600160401b038085166004830152808e166024830152808d166044830152808c166064830152808b166084830152891660a482015263ffffffff881660c482015260e481018790526001600160a01b03919091169063758185199061010401600060405180830381600087803b1580156118a557600080fd5b505af1158015610ad4573d6000803e3d6000fd5b80516001600160401b03166118d0576118d06131a8565b80516000906118e1906001906131d4565b6000805460405163421795e560e01b81526001600160401b0384166004820152929350909182916001600160a01b03169063421795e5906024016040805180830381865afa158015611937573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061195b9190613220565b8551602080880151608089015160a08a01516040519698509496506000956119ba9592939192910160c09490941b6001600160c01b031916845260e09290921b6001600160e01b0319166008840152600c830152602c820152604c0190565b60408051601f19818403018152919052805160209091012060c08601515160e0870151519192506000918082146119f3576119f36131a8565b60005b82816001600160401b03161015611b9b576000805460c08b015180516001600160a01b0390921691639332aa6e916001916001600160401b038716908110611a4057611a40613678565b6020026020010151611a5291906131d4565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa158015611a96573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611aba91906131fc565b905060008a60e0015190508673__$942d0120f12f58797d9427bd9464d63c7a$__6319045a25909183866001600160401b031681518110611afd57611afd613678565b60200260200101516040518363ffffffff1660e01b8152600401611b2292919061331c565b602060405180830381865af4158015611b3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b6391906131fc565b9550816001600160a01b0316866001600160a01b031614611b8657611b866131a8565b50508080611b9390613767565b9150506119f6565b5060008054604051630748349960e01b81526001600160a01b03898116600483015283928392911690630748349990602401600060405180830381865afa158015611bea573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611c12919081019061368e565b505050925092509250828015611c26575081155b8015611c3557508060ff166002145b611c515760405162461bcd60e51b815260040161048590613622565b600080546003906001600160a01b031663429fb683611c7160018d6131d4565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa158015611cb5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cd9919061363e565b611ce49060016137a9565b611cef906002613371565b611cf9919061333d565b9050806001600160401b0316851015611d395760405162461bcd60e51b8152602060048201526002602482015261534560f01b6044820152606401610485565b600080546001600160a01b0316634496c991611d5660018d6131d4565b6040516001600160e01b031960e084901b1681526001600160401b03909116600482015260240160e060405180830381865afa158015611d9a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dbe91906137d4565b5091955050505061ffff83168810159150611e0290505760405162461bcd60e51b8152602060048201526002602482015261534560f01b6044820152606401610485565b5050600080549093506001600160a01b03169150635f0963769050611e286001896131d4565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa158015611e6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e9091906131fc565b604051633b7c9bd760e21b81529091506001600160a01b0382169063edf26f5c90611ebf908c906004016138f4565b600060405180830381600087803b158015611ed957600080fd5b505af1158015611eed573d6000803e3d6000fd5b505050506000896020015163ffffffff1660001415610adb57895160208b0151604051637b31a24d60e01b81526001600160401b03909216600483015263ffffffff1660248201526001600160a01b03831690637b31a24d9060440161016060405180830381865afa158015611f67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f8b9190613998565b9091929394959697989950909192939495969798509091929394959697509091929394959650909192939495509091929394509091925090915090505080915050600160009054906101000a90046001600160a01b03166001600160a01b03166311e65fc08b60a00151836040518363ffffffff1660e01b815260040161201c929190918252602082015260400190565b600060405180830381600087803b15801561203657600080fd5b505af115801561204a573d6000803e3d6000fd5b5050505050505050505050505050565b6000546040516319f1d33560e31b81526001600160401b038c166004820152600260248201526001600160a01b039091169063cf8e99a890604401602060405180830381865afa1580156120b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120d691906131fc565b506000546040516319f1d33560e31b81526001600160401b038b166004820152600260248201526001600160a01b039091169063cf8e99a890604401602060405180830381865afa15801561212f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061215391906131fc565b506000866001600160401b031611801561217e5750876001600160401b0316876001600160401b0316115b61219a5760405162461bcd60e51b815260040161048590613622565b6000806121ab338d8d89600261269c565b9150915060008c8b8b8b8b8b8b60405160200161222e979695949392919060c097881b6001600160c01b0319908116825296881b8716600882015294871b8616601086015292861b85166018850152941b909216602082015260e09290921b6001600160e01b0319166028830152602c820152606160f81b604c820152604d0190565b60405160208183030381529060405280519060200120905061227860008054906101000a90046001600160a01b03166001600160a01b0316639332aa6e60018f6114d591906131d4565b600080546003906001600160a01b031663429fb6836122986001876131d4565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa1580156122dc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612300919061363e565b61230b906002613371565b612315919061333d565b6001600160401b0316905080855110156123415760405162461bcd60e51b81526004016104859061365b565b60008060005b875181101561249d578473__$942d0120f12f58797d9427bd9464d63c7a$__6319045a2590918a848151811061237f5761237f613678565b60200260200101516040518363ffffffff1660e01b81526004016123a492919061331c565b602060405180830381865af41580156123c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123e591906131fc565b60008054604051630748349960e01b81526001600160a01b03808516600483015293955091921690630748349990602401600060405180830381865afa158015612433573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261245b919081019061368e565b50945050505050866001600160401b0316816001600160401b0316141561248a578361248681613767565b9450505b50806124958161378e565b915050612347565b5082826001600160401b031610156124c75760405162461bcd60e51b81526004016104859061365b565b50506000546001600160a01b03169150635f09637690506124e96001846131d4565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa15801561252d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061255191906131fc565b604051630f60c7b360e01b81526001600160401b038085166004830152808f166024830152808e166044830152808d166064830152808c166084830152808b1660a4830152891660c482015263ffffffff881660e482015261010481018790526001600160a01b039190911690630f60c7b39061012401600060405180830381600087803b1580156125e257600080fd5b505af11580156125f6573d6000803e3d6000fd5b5050505063ffffffff861661268e576001546001600160a01b0316806311e65fc060006126238e8e6131d4565b612636906001600160401b03168a613a2c565b6040516001600160e01b031960e085901b16815260048101929092526024820152604401600060405180830381600087803b15801561267457600080fd5b505af1158015612688573d6000803e3d6000fd5b50505050505b505050505050505050505050565b600080546040516337a4ba6560e21b815263ffffffff8516600482015282916001600160a01b031690819063de92e99490602401602060405180830381865afa1580156126ed573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061271191906131fc565b50604051630748349960e01b81526001600160a01b038981166004830152600091829182918291829190871690630748349990602401600060405180830381865afa158015612764573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261278c919081019061368e565b50945094509450945094508860ff16600114156127e9578b6001600160401b0316826001600160401b0316146127e95760405162461bcd60e51b8152602060048201526002602482015261434560f01b6044820152606401610485565b8860ff1660021415612859578b6001600160401b0316826001600160401b03161461285957848015612819575083155b801561282857508260ff166003145b6128595760405162461bcd60e51b8152602060048201526002602482015261434560f01b6044820152606401610485565b6000866001600160a01b031663421795e560018f61287791906131d4565b6040516001600160e01b031960e084901b1681526001600160401b0390911660048201526024016040805180830381865afa1580156128ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128de9190613220565b9150506000876001600160a01b031663421795e560018f6128ff91906131d4565b6040516001600160e01b031960e084901b1681526001600160401b0390911660048201526024016040805180830381865afa158015612942573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129669190613220565b915050806001600160401b0316826001600160401b031614801561299b5750826001600160401b0316826001600160401b0316145b6129cc5760405162461bcd60e51b815260206004820152600260248201526111d160f21b6044820152606401610485565b50919d919c50909a5050505050505050505050565b6040516319045a2560e01b81526001600160a01b0384169073__$942d0120f12f58797d9427bd9464d63c7a$__906319045a2590612a25908690869060040161331c565b602060405180830381865af4158015612a42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a6691906131fc565b6001600160a01b031614612aa15760405162461bcd60e51b8152602060048201526002602482015261494360f01b6044820152606401610485565b505050565b6040518060a001604052806005905b6060815260200190600190039081612ab55790505090565b634e487b7160e01b600052604160045260246000fd5b60405161014081016001600160401b0381118282101715612b0657612b06612acd565b60405290565b60405161016081016001600160401b0381118282101715612b0657612b06612acd565b60405161010081016001600160401b0381118282101715612b0657612b06612acd565b604051601f8201601f191681016001600160401b0381118282101715612b7a57612b7a612acd565b604052919050565b6001600160401b0381168114612b9757600080fd5b50565b8035612ba581612b82565b919050565b803563ffffffff81168114612ba557600080fd5b60006001600160401b03821115612bd757612bd7612acd565b50601f01601f191660200190565b600082601f830112612bf657600080fd5b8135612c09612c0482612bbe565b612b52565b818152846020838601011115612c1e57600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215612c4d57600080fd5b81356001600160401b0380821115612c6457600080fd5b908301906101408286031215612c7957600080fd5b612c81612ae3565b612c8a83612b9a565b8152612c9860208401612b9a565b6020820152612ca960408401612b9a565b6040820152612cba60608401612b9a565b6060820152612ccb60808401612b9a565b6080820152612cdc60a08401612b9a565b60a0820152612ced60c08401612baa565b60c082015260e083013560e08201526101008084013583811115612d1057600080fd5b612d1c88828701612be5565b8284015250506101208084013583811115612d3657600080fd5b612d4288828701612be5565b918301919091525095945050505050565b6001600160a01b0381168114612b9757600080fd5b8035612ba581612d53565b600080600060608486031215612d8857600080fd5b8335612d9381612d53565b92506020840135612da381612d53565b91506040840135612db381612d53565b809150509250925092565b600060208284031215612dd057600080fd5b81356001600160401b0380821115612de757600080fd5b908301906101608286031215612dfc57600080fd5b612e04612b0c565b612e0d83612b9a565b8152612e1b60208401612b9a565b6020820152612e2c60408401612b9a565b6040820152612e3d60608401612b9a565b6060820152612e4e60808401612b9a565b6080820152612e5f60a08401612b9a565b60a0820152612e7060c08401612b9a565b60c0820152612e8160e08401612baa565b60e082015261010083810135908201526101208084013583811115612ea557600080fd5b612eb188828701612be5565b8284015250506101408084013583811115612d3657600080fd5b60006001600160401b03821115612ee457612ee4612acd565b5060051b60200190565b600082601f830112612eff57600080fd5b81356020612f0f612c0483612ecb565b82815260059290921b84018101918181019086841115612f2e57600080fd5b8286015b84811015612f6d5780356001600160401b03811115612f515760008081fd5b612f5f8986838b0101612be5565b845250918301918301612f32565b509695505050505050565b6000806000806000806000806000806101408b8d031215612f9857600080fd5b612fa18b612b9a565b9950612faf60208c01612b9a565b9850612fbd60408c01612b9a565b9750612fcb60608c01612b9a565b9650612fd960808c01612b9a565b9550612fe760a08c01612b9a565b9450612ff560c08c01612baa565b935060e08b013592506101008b01356001600160401b038082111561301957600080fd5b6130258e838f01612be5565b93506101208d013591508082111561303c57600080fd5b506130498d828e01612eee565b9150509295989b9194979a5092959850565b600082601f83011261306c57600080fd5b8135602061307c612c0483612ecb565b82815260059290921b8401810191818101908684111561309b57600080fd5b8286015b84811015612f6d5780356130b281612b82565b835291830191830161309f565b6000602082840312156130d157600080fd5b81356001600160401b03808211156130e857600080fd5b9083019061010082860312156130fd57600080fd5b613105612b2f565b61310e83612b9a565b815261311c60208401612baa565b602082015261312d60408401612d68565b604082015261313e60608401612d68565b60608201526080830135608082015260a083013560a082015260c08301358281111561316957600080fd5b6131758782860161305b565b60c08301525060e08301358281111561318d57600080fd5b61319987828601612eee565b60e08301525095945050505050565b634e487b7160e01b600052600160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001600160401b03838116908316818110156131f4576131f46131be565b039392505050565b60006020828403121561320e57600080fd5b815161321981612d53565b9392505050565b6000806040838503121561323357600080fd5b825161323e81612d53565b602084015190925061324f81612b82565b809150509250929050565b60c098891b6001600160c01b0319908116825297891b8816600882015295881b8716601087015293871b8616601886015291861b8516602085015290941b909216602882015260e09290921b6001600160e01b0319166030830152603482015260540190565b60005b838110156132db5781810151838201526020016132c3565b838111156132ea576000848401525b50505050565b600081518084526133088160208601602086016132c0565b601f01601f19169290920160200192915050565b82815260406020820152600061333560408301846132f0565b949350505050565b60006001600160401b038084168061336557634e487b7160e01b600052601260045260246000fd5b92169190910492915050565b60006001600160401b0380831681851681830481118215151615613397576133976131be565b02949350505050565b602081526133ba6020820183516001600160401b03169052565b600060208301516133d660408401826001600160401b03169052565b5060408301516001600160401b03811660608401525060608301516001600160401b03811660808401525060808301516001600160401b03811660a08401525060a08301516001600160401b03811660c08401525060c083015163ffffffff811660e08401525060e0830151610100838101919091528301516101406101208085018290526134696101608601846132f0565b90860151858203601f19018387015290925061348583826132f0565b9695505050505050565b6000602082840312156134a157600080fd5b5051919050565b6001600160a01b03841681526020808201849052606060408301819052600091610100840191840185845b600581101561350257605f198786030183526134f08583516132f0565b945091830191908301906001016134d3565b509298975050505050505050565b80518015158114612ba557600080fd5b60006020828403121561353257600080fd5b61321982613510565b602081526135556020820183516001600160401b03169052565b6000602083015161357160408401826001600160401b03169052565b5060408301516001600160401b03811660608401525060608301516001600160401b03811660808401525060808301516001600160401b03811660a08401525060a08301516001600160401b03811660c08401525060c08301516001600160401b03811660e08401525060e08301516101006135f48185018363ffffffff169052565b84015161012084810191909152840151610160610140808601829052919250906134696101808601846132f0565b602080825260029082015261049560f41b604082015260600190565b60006020828403121561365057600080fd5b815161321981612b82565b602080825260039082015262534e4560e81b604082015260600190565b634e487b7160e01b600052603260045260246000fd5b60008060008060008060c087890312156136a757600080fd5b6136b087613510565b95506136be60208801613510565b9450604087015160ff811681146136d457600080fd5b60608801519094506136e581612b82565b60808801519093506136f681612b82565b60a08801519092506001600160401b0381111561371257600080fd5b8701601f8101891361372357600080fd5b8051613731612c0482612bbe565b8181528a602083850101111561374657600080fd5b6137578260208301602086016132c0565b8093505050509295509295509295565b60006001600160401b0380831681811415613784576137846131be565b6001019392505050565b60006000198214156137a2576137a26131be565b5060010190565b60006001600160401b038083168185168083038211156137cb576137cb6131be565b01949350505050565b600080600080600080600060e0888a0312156137ef57600080fd5b6137f888613510565b965061380660208901613510565b955061381460408901613510565b9450606088015161ffff8116811461382b57600080fd5b809450506080880151925060a0880151915060c088015161384b81612d53565b8091505092959891949750929550565b600081518084526020808501945080840160005b838110156138945781516001600160401b03168752958201959082019060010161386f565b509495945050505050565b600081518084526020808501808196508360051b8101915082860160005b858110156138e75782840389526138d58483516132f0565b988501989350908401906001016138bd565b5091979650505050505050565b602081526001600160401b03825116602082015260006020830151613921604084018263ffffffff169052565b5060408301516001600160a01b03811660608401525060608301516001600160a01b038116608084015250608083015160a083015260a083015160c083015260c08301516101008060e085015261397c61012085018361385b565b915060e0850151601f198584030182860152613485838261389f565b60008060008060008060008060008060006101608c8e0312156139ba57600080fd5b8b516139c581612b82565b60208d0151909b506139d681612b82565b809a505060408c0151985060608c0151975060808c0151965060a08c0151955060c08c0151945060e08c015193506101008c015192506101208c015191506101408c015190509295989b509295989b9093969950565b6000816000190483118215151615613a4657613a466131be565b50029056fea26469706673582212208004b2b2f1634be6a3b52f55fb82bbc1b07355c353600e4188da12e90d6f439a64736f6c634300080c0033",
}

// RoleFSABI is the input ABI used to generate the binding from.
// Deprecated: Use RoleFSMetaData.ABI instead.
var RoleFSABI = RoleFSMetaData.ABI

// Deprecated: Use RoleFSMetaData.Sigs instead.
// RoleFSFuncSigs maps the 4-byte function signature to its string representation.
var RoleFSFuncSigs = RoleFSMetaData.Sigs

// RoleFSBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RoleFSMetaData.Bin instead.
var RoleFSBin = RoleFSMetaData.Bin

// DeployRoleFS deploys a new Ethereum contract, binding an instance of RoleFS to it.
func DeployRoleFS(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RoleFS, error) {
	parsed, err := RoleFSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	recoverAddr, _, _, _ := DeployRecover(auth, backend)
	RoleFSBin = strings.Replace(RoleFSBin, "__$942d0120f12f58797d9427bd9464d63c7a$__", recoverAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RoleFSBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RoleFS{RoleFSCaller: RoleFSCaller{contract: contract}, RoleFSTransactor: RoleFSTransactor{contract: contract}, RoleFSFilterer: RoleFSFilterer{contract: contract}}, nil
}

// RoleFS is an auto generated Go binding around an Ethereum contract.
type RoleFS struct {
	RoleFSCaller     // Read-only binding to the contract
	RoleFSTransactor // Write-only binding to the contract
	RoleFSFilterer   // Log filterer for contract events
}

// RoleFSCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoleFSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleFSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoleFSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleFSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoleFSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleFSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoleFSSession struct {
	Contract     *RoleFS           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoleFSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoleFSCallerSession struct {
	Contract *RoleFSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RoleFSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoleFSTransactorSession struct {
	Contract     *RoleFSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoleFSRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoleFSRaw struct {
	Contract *RoleFS // Generic contract binding to access the raw methods on
}

// RoleFSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoleFSCallerRaw struct {
	Contract *RoleFSCaller // Generic read-only contract binding to access the raw methods on
}

// RoleFSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoleFSTransactorRaw struct {
	Contract *RoleFSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoleFS creates a new instance of RoleFS, bound to a specific deployed contract.
func NewRoleFS(address common.Address, backend bind.ContractBackend) (*RoleFS, error) {
	contract, err := bindRoleFS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoleFS{RoleFSCaller: RoleFSCaller{contract: contract}, RoleFSTransactor: RoleFSTransactor{contract: contract}, RoleFSFilterer: RoleFSFilterer{contract: contract}}, nil
}

// NewRoleFSCaller creates a new read-only instance of RoleFS, bound to a specific deployed contract.
func NewRoleFSCaller(address common.Address, caller bind.ContractCaller) (*RoleFSCaller, error) {
	contract, err := bindRoleFS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoleFSCaller{contract: contract}, nil
}

// NewRoleFSTransactor creates a new write-only instance of RoleFS, bound to a specific deployed contract.
func NewRoleFSTransactor(address common.Address, transactor bind.ContractTransactor) (*RoleFSTransactor, error) {
	contract, err := bindRoleFS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoleFSTransactor{contract: contract}, nil
}

// NewRoleFSFilterer creates a new log filterer instance of RoleFS, bound to a specific deployed contract.
func NewRoleFSFilterer(address common.Address, filterer bind.ContractFilterer) (*RoleFSFilterer, error) {
	contract, err := bindRoleFS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoleFSFilterer{contract: contract}, nil
}

// bindRoleFS binds a generic wrapper to an already deployed contract.
func bindRoleFS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoleFSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoleFS *RoleFSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoleFS.Contract.RoleFSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoleFS *RoleFSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleFS.Contract.RoleFSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoleFS *RoleFSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoleFS.Contract.RoleFSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoleFS *RoleFSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoleFS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoleFS *RoleFSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleFS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoleFS *RoleFSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoleFS.Contract.contract.Transact(opts, method, params...)
}

// AddOrder is a paid mutator transaction binding the contract method 0x0e09662b.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes) ps) returns()
func (_RoleFS *RoleFSTransactor) AddOrder(opts *bind.TransactOpts, ps AOParams) (*types.Transaction, error) {
	return _RoleFS.contract.Transact(opts, "addOrder", ps)
}

// AddOrder is a paid mutator transaction binding the contract method 0x0e09662b.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes) ps) returns()
func (_RoleFS *RoleFSSession) AddOrder(ps AOParams) (*types.Transaction, error) {
	return _RoleFS.Contract.AddOrder(&_RoleFS.TransactOpts, ps)
}

// AddOrder is a paid mutator transaction binding the contract method 0x0e09662b.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes) ps) returns()
func (_RoleFS *RoleFSTransactorSession) AddOrder(ps AOParams) (*types.Transaction, error) {
	return _RoleFS.Contract.AddOrder(&_RoleFS.TransactOpts, ps)
}

// AddRepair is a paid mutator transaction binding the contract method 0xffd0fd09.
//
// Solidity: function addRepair(uint64 pIndex, uint64 nPIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sprice, bytes pSign, bytes[] kSigns) returns()
func (_RoleFS *RoleFSTransactor) AddRepair(opts *bind.TransactOpts, pIndex uint64, nPIndex uint64, _start uint64, end uint64, _size uint64, nonce uint64, tIndex uint32, sprice *big.Int, pSign []byte, kSigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.contract.Transact(opts, "addRepair", pIndex, nPIndex, _start, end, _size, nonce, tIndex, sprice, pSign, kSigns)
}

// AddRepair is a paid mutator transaction binding the contract method 0xffd0fd09.
//
// Solidity: function addRepair(uint64 pIndex, uint64 nPIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sprice, bytes pSign, bytes[] kSigns) returns()
func (_RoleFS *RoleFSSession) AddRepair(pIndex uint64, nPIndex uint64, _start uint64, end uint64, _size uint64, nonce uint64, tIndex uint32, sprice *big.Int, pSign []byte, kSigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.Contract.AddRepair(&_RoleFS.TransactOpts, pIndex, nPIndex, _start, end, _size, nonce, tIndex, sprice, pSign, kSigns)
}

// AddRepair is a paid mutator transaction binding the contract method 0xffd0fd09.
//
// Solidity: function addRepair(uint64 pIndex, uint64 nPIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sprice, bytes pSign, bytes[] kSigns) returns()
func (_RoleFS *RoleFSTransactorSession) AddRepair(pIndex uint64, nPIndex uint64, _start uint64, end uint64, _size uint64, nonce uint64, tIndex uint32, sprice *big.Int, pSign []byte, kSigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.Contract.AddRepair(&_RoleFS.TransactOpts, pIndex, nPIndex, _start, end, _size, nonce, tIndex, sprice, pSign, kSigns)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xedf26f5c.
//
// Solidity: function proWithdraw((uint64,uint32,address,address,uint256,uint256,uint64[],bytes[]) ps) returns()
func (_RoleFS *RoleFSTransactor) ProWithdraw(opts *bind.TransactOpts, ps PWParams) (*types.Transaction, error) {
	return _RoleFS.contract.Transact(opts, "proWithdraw", ps)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xedf26f5c.
//
// Solidity: function proWithdraw((uint64,uint32,address,address,uint256,uint256,uint64[],bytes[]) ps) returns()
func (_RoleFS *RoleFSSession) ProWithdraw(ps PWParams) (*types.Transaction, error) {
	return _RoleFS.Contract.ProWithdraw(&_RoleFS.TransactOpts, ps)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xedf26f5c.
//
// Solidity: function proWithdraw((uint64,uint32,address,address,uint256,uint256,uint64[],bytes[]) ps) returns()
func (_RoleFS *RoleFSTransactorSession) ProWithdraw(ps PWParams) (*types.Transaction, error) {
	return _RoleFS.Contract.ProWithdraw(&_RoleFS.TransactOpts, ps)
}

// SetAddr is a paid mutator transaction binding the contract method 0x6153c8c1.
//
// Solidity: function setAddr(address iss, address r, address rt) returns()
func (_RoleFS *RoleFSTransactor) SetAddr(opts *bind.TransactOpts, iss common.Address, r common.Address, rt common.Address) (*types.Transaction, error) {
	return _RoleFS.contract.Transact(opts, "setAddr", iss, r, rt)
}

// SetAddr is a paid mutator transaction binding the contract method 0x6153c8c1.
//
// Solidity: function setAddr(address iss, address r, address rt) returns()
func (_RoleFS *RoleFSSession) SetAddr(iss common.Address, r common.Address, rt common.Address) (*types.Transaction, error) {
	return _RoleFS.Contract.SetAddr(&_RoleFS.TransactOpts, iss, r, rt)
}

// SetAddr is a paid mutator transaction binding the contract method 0x6153c8c1.
//
// Solidity: function setAddr(address iss, address r, address rt) returns()
func (_RoleFS *RoleFSTransactorSession) SetAddr(iss common.Address, r common.Address, rt common.Address) (*types.Transaction, error) {
	return _RoleFS.Contract.SetAddr(&_RoleFS.TransactOpts, iss, r, rt)
}

// SubOrder is a paid mutator transaction binding the contract method 0xabfc55f1.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes) ps) returns()
func (_RoleFS *RoleFSTransactor) SubOrder(opts *bind.TransactOpts, ps SOParams) (*types.Transaction, error) {
	return _RoleFS.contract.Transact(opts, "subOrder", ps)
}

// SubOrder is a paid mutator transaction binding the contract method 0xabfc55f1.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes) ps) returns()
func (_RoleFS *RoleFSSession) SubOrder(ps SOParams) (*types.Transaction, error) {
	return _RoleFS.Contract.SubOrder(&_RoleFS.TransactOpts, ps)
}

// SubOrder is a paid mutator transaction binding the contract method 0xabfc55f1.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes) ps) returns()
func (_RoleFS *RoleFSTransactorSession) SubOrder(ps SOParams) (*types.Transaction, error) {
	return _RoleFS.Contract.SubOrder(&_RoleFS.TransactOpts, ps)
}

// SubRepair is a paid mutator transaction binding the contract method 0xeb975779.
//
// Solidity: function subRepair(uint64 pIndex, uint64 nPIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sPrice, bytes pSign, bytes[] kSigns) returns()
func (_RoleFS *RoleFSTransactor) SubRepair(opts *bind.TransactOpts, pIndex uint64, nPIndex uint64, _start uint64, end uint64, _size uint64, nonce uint64, tIndex uint32, sPrice *big.Int, pSign []byte, kSigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.contract.Transact(opts, "subRepair", pIndex, nPIndex, _start, end, _size, nonce, tIndex, sPrice, pSign, kSigns)
}

// SubRepair is a paid mutator transaction binding the contract method 0xeb975779.
//
// Solidity: function subRepair(uint64 pIndex, uint64 nPIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sPrice, bytes pSign, bytes[] kSigns) returns()
func (_RoleFS *RoleFSSession) SubRepair(pIndex uint64, nPIndex uint64, _start uint64, end uint64, _size uint64, nonce uint64, tIndex uint32, sPrice *big.Int, pSign []byte, kSigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.Contract.SubRepair(&_RoleFS.TransactOpts, pIndex, nPIndex, _start, end, _size, nonce, tIndex, sPrice, pSign, kSigns)
}

// SubRepair is a paid mutator transaction binding the contract method 0xeb975779.
//
// Solidity: function subRepair(uint64 pIndex, uint64 nPIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sPrice, bytes pSign, bytes[] kSigns) returns()
func (_RoleFS *RoleFSTransactorSession) SubRepair(pIndex uint64, nPIndex uint64, _start uint64, end uint64, _size uint64, nonce uint64, tIndex uint32, sPrice *big.Int, pSign []byte, kSigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.Contract.SubRepair(&_RoleFS.TransactOpts, pIndex, nPIndex, _start, end, _size, nonce, tIndex, sPrice, pSign, kSigns)
}
