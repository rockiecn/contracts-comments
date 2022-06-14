// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pledgepool

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

// SGParams is an auto generated low-level Go binding around an user-defined struct.
type SGParams struct {
	Index  uint64
	IsAdd  bool
	Size   *big.Int
	SPrice *big.Int
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

// IPledgePoolMetaData contains all meta data concerning the IPledgePool contract.
var IPledgePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"getPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8859344c": "addToken(address,uint32)",
		"761966b7": "getBalance(uint64,uint32)",
		"c3b19ccc": "getPledge(uint32)",
		"364e4bf5": "pledge(uint64,uint256,bytes)",
		"32704298": "withdraw(uint64,uint32,uint256,bytes)",
	},
}

// IPledgePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use IPledgePoolMetaData.ABI instead.
var IPledgePoolABI = IPledgePoolMetaData.ABI

// Deprecated: Use IPledgePoolMetaData.Sigs instead.
// IPledgePoolFuncSigs maps the 4-byte function signature to its string representation.
var IPledgePoolFuncSigs = IPledgePoolMetaData.Sigs

// IPledgePool is an auto generated Go binding around an Ethereum contract.
type IPledgePool struct {
	IPledgePoolCaller     // Read-only binding to the contract
	IPledgePoolTransactor // Write-only binding to the contract
	IPledgePoolFilterer   // Log filterer for contract events
}

// IPledgePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPledgePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPledgePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPledgePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPledgePoolSession struct {
	Contract     *IPledgePool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPledgePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPledgePoolCallerSession struct {
	Contract *IPledgePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IPledgePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPledgePoolTransactorSession struct {
	Contract     *IPledgePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IPledgePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPledgePoolRaw struct {
	Contract *IPledgePool // Generic contract binding to access the raw methods on
}

// IPledgePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPledgePoolCallerRaw struct {
	Contract *IPledgePoolCaller // Generic read-only contract binding to access the raw methods on
}

// IPledgePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPledgePoolTransactorRaw struct {
	Contract *IPledgePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPledgePool creates a new instance of IPledgePool, bound to a specific deployed contract.
func NewIPledgePool(address common.Address, backend bind.ContractBackend) (*IPledgePool, error) {
	contract, err := bindIPledgePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPledgePool{IPledgePoolCaller: IPledgePoolCaller{contract: contract}, IPledgePoolTransactor: IPledgePoolTransactor{contract: contract}, IPledgePoolFilterer: IPledgePoolFilterer{contract: contract}}, nil
}

// NewIPledgePoolCaller creates a new read-only instance of IPledgePool, bound to a specific deployed contract.
func NewIPledgePoolCaller(address common.Address, caller bind.ContractCaller) (*IPledgePoolCaller, error) {
	contract, err := bindIPledgePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPledgePoolCaller{contract: contract}, nil
}

// NewIPledgePoolTransactor creates a new write-only instance of IPledgePool, bound to a specific deployed contract.
func NewIPledgePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*IPledgePoolTransactor, error) {
	contract, err := bindIPledgePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPledgePoolTransactor{contract: contract}, nil
}

// NewIPledgePoolFilterer creates a new log filterer instance of IPledgePool, bound to a specific deployed contract.
func NewIPledgePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*IPledgePoolFilterer, error) {
	contract, err := bindIPledgePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPledgePoolFilterer{contract: contract}, nil
}

// bindIPledgePool binds a generic wrapper to an already deployed contract.
func bindIPledgePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPledgePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPledgePool *IPledgePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPledgePool.Contract.IPledgePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPledgePool *IPledgePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPledgePool.Contract.IPledgePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPledgePool *IPledgePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPledgePool.Contract.IPledgePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPledgePool *IPledgePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPledgePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPledgePool *IPledgePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPledgePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPledgePool *IPledgePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPledgePool.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x761966b7.
//
// Solidity: function getBalance(uint64 index, uint32 tIndex) view returns(uint256)
func (_IPledgePool *IPledgePoolCaller) GetBalance(opts *bind.CallOpts, index uint64, tIndex uint32) (*big.Int, error) {
	var out []interface{}
	err := _IPledgePool.contract.Call(opts, &out, "getBalance", index, tIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x761966b7.
//
// Solidity: function getBalance(uint64 index, uint32 tIndex) view returns(uint256)
func (_IPledgePool *IPledgePoolSession) GetBalance(index uint64, tIndex uint32) (*big.Int, error) {
	return _IPledgePool.Contract.GetBalance(&_IPledgePool.CallOpts, index, tIndex)
}

// GetBalance is a free data retrieval call binding the contract method 0x761966b7.
//
// Solidity: function getBalance(uint64 index, uint32 tIndex) view returns(uint256)
func (_IPledgePool *IPledgePoolCallerSession) GetBalance(index uint64, tIndex uint32) (*big.Int, error) {
	return _IPledgePool.Contract.GetBalance(&_IPledgePool.CallOpts, index, tIndex)
}

// GetPledge is a free data retrieval call binding the contract method 0xc3b19ccc.
//
// Solidity: function getPledge(uint32 tIndex) view returns(uint256)
func (_IPledgePool *IPledgePoolCaller) GetPledge(opts *bind.CallOpts, tIndex uint32) (*big.Int, error) {
	var out []interface{}
	err := _IPledgePool.contract.Call(opts, &out, "getPledge", tIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPledge is a free data retrieval call binding the contract method 0xc3b19ccc.
//
// Solidity: function getPledge(uint32 tIndex) view returns(uint256)
func (_IPledgePool *IPledgePoolSession) GetPledge(tIndex uint32) (*big.Int, error) {
	return _IPledgePool.Contract.GetPledge(&_IPledgePool.CallOpts, tIndex)
}

// GetPledge is a free data retrieval call binding the contract method 0xc3b19ccc.
//
// Solidity: function getPledge(uint32 tIndex) view returns(uint256)
func (_IPledgePool *IPledgePoolCallerSession) GetPledge(tIndex uint32) (*big.Int, error) {
	return _IPledgePool.Contract.GetPledge(&_IPledgePool.CallOpts, tIndex)
}

// AddToken is a paid mutator transaction binding the contract method 0x8859344c.
//
// Solidity: function addToken(address tAddr, uint32 tIndex) returns()
func (_IPledgePool *IPledgePoolTransactor) AddToken(opts *bind.TransactOpts, tAddr common.Address, tIndex uint32) (*types.Transaction, error) {
	return _IPledgePool.contract.Transact(opts, "addToken", tAddr, tIndex)
}

// AddToken is a paid mutator transaction binding the contract method 0x8859344c.
//
// Solidity: function addToken(address tAddr, uint32 tIndex) returns()
func (_IPledgePool *IPledgePoolSession) AddToken(tAddr common.Address, tIndex uint32) (*types.Transaction, error) {
	return _IPledgePool.Contract.AddToken(&_IPledgePool.TransactOpts, tAddr, tIndex)
}

// AddToken is a paid mutator transaction binding the contract method 0x8859344c.
//
// Solidity: function addToken(address tAddr, uint32 tIndex) returns()
func (_IPledgePool *IPledgePoolTransactorSession) AddToken(tAddr common.Address, tIndex uint32) (*types.Transaction, error) {
	return _IPledgePool.Contract.AddToken(&_IPledgePool.TransactOpts, tAddr, tIndex)
}

// Pledge is a paid mutator transaction binding the contract method 0x364e4bf5.
//
// Solidity: function pledge(uint64 index, uint256 money, bytes sign) payable returns()
func (_IPledgePool *IPledgePoolTransactor) Pledge(opts *bind.TransactOpts, index uint64, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _IPledgePool.contract.Transact(opts, "pledge", index, money, sign)
}

// Pledge is a paid mutator transaction binding the contract method 0x364e4bf5.
//
// Solidity: function pledge(uint64 index, uint256 money, bytes sign) payable returns()
func (_IPledgePool *IPledgePoolSession) Pledge(index uint64, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _IPledgePool.Contract.Pledge(&_IPledgePool.TransactOpts, index, money, sign)
}

// Pledge is a paid mutator transaction binding the contract method 0x364e4bf5.
//
// Solidity: function pledge(uint64 index, uint256 money, bytes sign) payable returns()
func (_IPledgePool *IPledgePoolTransactorSession) Pledge(index uint64, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _IPledgePool.Contract.Pledge(&_IPledgePool.TransactOpts, index, money, sign)
}

// Withdraw is a paid mutator transaction binding the contract method 0x32704298.
//
// Solidity: function withdraw(uint64 index, uint32 tIndex, uint256 money, bytes sign) returns()
func (_IPledgePool *IPledgePoolTransactor) Withdraw(opts *bind.TransactOpts, index uint64, tIndex uint32, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _IPledgePool.contract.Transact(opts, "withdraw", index, tIndex, money, sign)
}

// Withdraw is a paid mutator transaction binding the contract method 0x32704298.
//
// Solidity: function withdraw(uint64 index, uint32 tIndex, uint256 money, bytes sign) returns()
func (_IPledgePool *IPledgePoolSession) Withdraw(index uint64, tIndex uint32, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _IPledgePool.Contract.Withdraw(&_IPledgePool.TransactOpts, index, tIndex, money, sign)
}

// Withdraw is a paid mutator transaction binding the contract method 0x32704298.
//
// Solidity: function withdraw(uint64 index, uint32 tIndex, uint256 money, bytes sign) returns()
func (_IPledgePool *IPledgePoolTransactorSession) Withdraw(index uint64, tIndex uint32, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _IPledgePool.Contract.Withdraw(&_IPledgePool.TransactOpts, index, tIndex, money, sign)
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

// PledgePoolMetaData contains all meta data concerning the PledgePool contract.
var PledgePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"primeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"Pledge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"getPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"role\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"8859344c": "addToken(address,uint32)",
		"761966b7": "getBalance(uint64,uint32)",
		"c3b19ccc": "getPledge(uint32)",
		"364e4bf5": "pledge(uint64,uint256,bytes)",
		"40c65f72": "rToken()",
		"10055c1d": "role()",
		"c21a43e4": "totalPledge()",
		"32704298": "withdraw(uint64,uint32,uint256,bytes)",
	},
	Bin: "0x60806040523480156200001157600080fd5b5060405162001d3d38038062001d3d8339810160408190526200003491620001a1565b600080546001600160a01b038481166001600160a01b031992831617835560018054858316931692909217909155604080518082018252601281527162616c616e63654f6628616464726573732960701b602091820152815130602480830191909152835180830390910181526044909101835290810180516001600160e01b03166370a0823160e01b17905290518392871691620000d391620001eb565b600060405180830381855afa9150503d806000811462000110576040519150601f19603f3d011682016040523d82523d6000602084013e62000115565b606091505b50915091508180156200012a57506020815110155b6200013457600080fd5b6000818060200190518101906200014c919062000229565b6000805260036020527f3617319a054d772f909f7c479a2cebe5066e836a939412e32403c99029b92f00555062000243945050505050565b80516001600160a01b03811681146200019c57600080fd5b919050565b600080600060608486031215620001b757600080fd5b620001c28462000184565b9250620001d26020850162000184565b9150620001e26040850162000184565b90509250925092565b6000825160005b818110156200020e5760208186018101518583015201620001f2565b818111156200021e576000828501525b509190910192915050565b6000602082840312156200023c57600080fd5b5051919050565b611aea80620002536000396000f3fe60806040526004361061007f5760003560e01c8063761966b71161004e578063761966b71461011d5780638859344c1461014b578063c21a43e41461016b578063c3b19ccc1461018157600080fd5b806310055c1d1461008b57806332704298146100c8578063364e4bf5146100ea57806340c65f72146100fd57600080fd5b3661008657005b600080fd5b34801561009757600080fd5b506001546100ab906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156100d457600080fd5b506100e86100e336600461169c565b6101a1565b005b6100e86100f8366004611707565b610930565b34801561010957600080fd5b506000546100ab906001600160a01b031681565b34801561012957600080fd5b5061013d61013836600461175f565b610e3e565b6040519081526020016100bf565b34801561015757600080fd5b506100e86101663660046117ad565b611091565b34801561017757600080fd5b5061013d60025481565b34801561018d57600080fd5b5061013d61019c3660046117cb565b61110c565b600180546000916001600160a01b0390911690639332aa6e906101c490886117fe565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa158015610208573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061022c9190611826565b600154604051630748349960e01b81526001600160a01b03808416600483015292935060009283928392911690630748349990602401600060405180830381865afa15801561027f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526102a79190810190611884565b505050925092505081156102e65760405162461bcd60e51b81526020600482015260016024820152602160f91b60448201526064015b60405180910390fd5b336001600160a01b03851614610357576040516bffffffffffffffffffffffff193360601b1660208201526001600160e01b031960e089901b166034820152603881018790526000906058016040516020818303038152906040528051906020012090506103558582886111b8565b505b8060ff16600214156103dd57600160009054906101000a90046001600160a01b03166001600160a01b0316638ba61d286040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103da919061195d565b92505b8060ff166003141561046357600160009054906101000a90046001600160a01b03166001600160a01b031663a6ed590b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561043c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610460919061195d565b92505b505060008054604051637bdbeb5b60e11b815263ffffffff881660048201526001600160a01b039091169063f7b7d6b690602401602060405180830381865afa1580156104b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d89190611826565b90506001600160a01b0381166105005760405162461bcd60e51b81526004016102dd90611976565b6001600160401b03871660009081526004602090815260408083208380528252918290208251808401909352805483526001015490820181905215801561054657508051155b15610554575050505061092a565b60208101516002541561063d5760008060009054906101000a90046001600160a01b03166001600160a01b031663b04c6c686040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105d99190611992565b905060005b8163ffffffff168163ffffffff16101561063a5763ffffffff8a161580159061061357508063ffffffff168a63ffffffff1614155b1561061d57610628565b61062883828d61127d565b80610632816119af565b9150506105de565b50505b6001600160401b038916600090815260046020908152604080832063ffffffff8c168452825291829020825180840190935280548084526001909101549183019190915290925015801561069357506020820151155b156106a257505050505061092a565b602082015163ffffffff89166106cf57848111156106cb576106c485826119d3565b90506106cf565b5060005b6000881180156106de57508088105b156106e65750865b80156108765760405163a9059cbb60e01b81526001600160a01b0387811660048301526024820183905285169063a9059cbb906044016020604051808303816000875af115801561073b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061075f91906119ea565b5063ffffffff89166000908152600360205260409020600101548110156107b35763ffffffff8916600090815260036020526040812060010180548392906107a89084906119d3565b909155506107cd9050565b63ffffffff89166000908152600360205260408120600101555b6001600160401b038a16600090815260046020908152604080832063ffffffff8d168452909152902060010154811015610848576001600160401b038a16600090815260046020908152604080832063ffffffff8d1684529091528120600101805483929061083d9084906119d3565b909155506108769050565b6001600160401b038a16600090815260046020908152604080832063ffffffff8d1684529091528120600101555b63ffffffff89166108e0576001600160401b038a16600090815260046020908152604080832063ffffffff8d16845290915281206001015460028054919290916108c1908490611a05565b9250508190555081600260008282546108da91906119d3565b90915550505b856001600160a01b03167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243648260405161091b91815260200190565b60405180910390a25050505050505b50505050565b6001805433916001600160a01b03909116906000908290639332aa6e9061095790896117fe565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa15801561099b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109bf9190611826565b604051630748349960e01b81526001600160a01b038083166004830152919250600091841690630748349990602401600060405180830381865afa158015610a0b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a339190810190611884565b505050509150508015610a6c5760405162461bcd60e51b81526020600482015260016024820152602160f91b60448201526064016102dd565b816001600160a01b0316846001600160a01b031614610ae2576040516bffffffffffffffffffffffff19606086901b1660208201526034810187905265706c6564676560d01b6054820152600090605a01604051602081830303815290604052805190602001209050610ae08382886111b8565b505b6001600160401b038716600090815260046020818152604080842084805282528084208151808301835281548152600190910154818401819052855483516316098d8d60e31b815293519296919591946001600160a01b039091169363b04c6c68938184019390918290030181865afa158015610b63573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b879190611992565b905060005b8163ffffffff168163ffffffff161015610c85576001600160401b038b16600090815260046020908152604080832063ffffffff851684528252918290208251808401909352805480845260019091015491830191909152909450158015610bf657506020840151155b15610c6857600360008263ffffffff1663ffffffff16815260200190815260200160002060000154600460008d6001600160401b03166001600160401b0316815260200190815260200160002060008363ffffffff1663ffffffff168152602001908152602001600020600001819055505b610c7383828d61127d565b80610c7d816119af565b915050610b8c565b5060008054604051637bdbeb5b60e11b815260048101929092526001600160a01b03169063f7b7d6b690602401602060405180830381865afa158015610ccf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cf39190611826565b6040516323b872dd60e01b81526001600160a01b038781166004830152306024830152604482018c905291909116906323b872dd906064016020604051808303816000875af1158015610d4a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d6e91906119ea565b50600080805260036020527f3617319a054d772f909f7c479a2cebe5066e836a939412e32403c99029b92f0080548b9290610daa908490611a05565b90915550506001600160401b038a166000908152600460209081526040808320838052909152812060010180548b9290610de5908490611a05565b925050819055508860026000828254610dfe9190611a05565b90915550506040518981526001600160a01b038616907f5e91ea8ea1c46300eb761859be01d7b16d44389ef91e03a163a87413cbf55b959060200161091b565b60008054604051637bdbeb5b60e11b815263ffffffff8416600482015282916001600160a01b03169063f7b7d6b690602401602060405180830381865afa158015610e8d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eb19190611826565b90506001600160a01b038116610ed95760405162461bcd60e51b81526004016102dd90611976565b6001600160401b03841660009081526004602090815260408083208380528252918290208251808401909352805480845260019091015491830191909152158015610f2657506020810151155b15610f365760009250505061108b565b60208082015163ffffffff861660009081526003835260408082208151808301909252805480835260019091015494820194909452919290610f78863061148e565b9050826020015181610f8a91906119d3565b9050600081118015610f9e57506000600254115b15610fcd57600254610fb482633b9aca00611a1d565b610fbe9190611a3c565b9050610fca8183611a05565b91505b6001600160401b038916600090815260046020908152604080832063ffffffff8c168452825291829020825180840190935280548084526001909101549183019190915215801561102057506020810151155b1561104557633b9aca006110348685611a1d565b61103e9190611a3c565b9250611081565b805161105190846119d3565b9250633b9aca006110628685611a1d565b61106c9190611a3c565b925080602001518361107e9190611a05565b92505b5090955050505050505b92915050565b6001546001600160a01b031633146110cf5760405162461bcd60e51b81526020600482015260016024820152602760f91b60448201526064016102dd565b63ffffffff81166000908152600360205260408120556110ef823061148e565b63ffffffff90911660009081526003602052604090206001015550565b60008054604051637bdbeb5b60e11b815263ffffffff8416600482015282916001600160a01b03169063f7b7d6b690602401602060405180830381865afa15801561115b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061117f9190611826565b90506001600160a01b0381166111a75760405162461bcd60e51b81526004016102dd90611976565b6111b1813061148e565b9392505050565b6040516319045a2560e01b81526001600160a01b0384169073__$942d0120f12f58797d9427bd9464d63c7a$__906319045a25906111fc9086908690600401611a5e565b602060405180830381865af4158015611219573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061123d9190611826565b6001600160a01b0316146112785760405162461bcd60e51b8152602060048201526002602482015261494360f01b60448201526064016102dd565b505050565b60008054604051637bdbeb5b60e11b815263ffffffff851660048201526112f8916001600160a01b03169063f7b7d6b690602401602060405180830381865afa1580156112ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112f29190611826565b3061148e565b63ffffffff8416600090815260036020526040812060010154919250908211156113425763ffffffff841660009081526003602052604090206001015461133f90836119d3565b90505b60008111801561135457506000600254115b156113a45760025461136a82633b9aca00611a1d565b6113749190611a3c565b63ffffffff851660009081526003602052604081208054929350839290919061139e908490611a05565b90915550505b63ffffffff84166000818152600360208181526040808420600181018890556001600160401b038916855260048352818520958552948252832054919052915490916113ef916119d3565b9050633b9aca006114008783611a1d565b61140a9190611a3c565b6001600160401b038516600090815260046020908152604080832063ffffffff8a16845290915281206001018054929350839290919061144b908490611a05565b9091555050505063ffffffff929092166000818152600360209081526040808320546001600160401b039590951683526004825280832093835292905220555050565b604080518082018252601281527162616c616e63654f6628616464726573732960701b60209182015281516001600160a01b03848116602480840191909152845180840390910181526044909201845291810180516001600160e01b03166370a0823160e01b1790529151600092839283929087169161150d91611a98565b600060405180830381855afa9150503d8060008114611548576040519150601f19603f3d011682016040523d82523d6000602084013e61154d565b606091505b509150915081801561156157506020815110155b6115925760405162461bcd60e51b8152602060048201526002602482015261434560f01b60448201526064016102dd565b808060200190518101906115a6919061195d565b95945050505050565b6001600160401b03811681146115c457600080fd5b50565b63ffffffff811681146115c457600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715611617576116176115d9565b604052919050565b60006001600160401b03821115611638576116386115d9565b50601f01601f191660200190565b600082601f83011261165757600080fd5b813561166a6116658261161f565b6115ef565b81815284602083860101111561167f57600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080608085870312156116b257600080fd5b84356116bd816115af565b935060208501356116cd816115c7565b92506040850135915060608501356001600160401b038111156116ef57600080fd5b6116fb87828801611646565b91505092959194509250565b60008060006060848603121561171c57600080fd5b8335611727816115af565b92506020840135915060408401356001600160401b0381111561174957600080fd5b61175586828701611646565b9150509250925092565b6000806040838503121561177257600080fd5b823561177d816115af565b9150602083013561178d816115c7565b809150509250929050565b6001600160a01b03811681146115c457600080fd5b600080604083850312156117c057600080fd5b823561177d81611798565b6000602082840312156117dd57600080fd5b81356111b1816115c7565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b038381169083168181101561181e5761181e6117e8565b039392505050565b60006020828403121561183857600080fd5b81516111b181611798565b8051801515811461185357600080fd5b919050565b60005b8381101561187357818101518382015260200161185b565b8381111561092a5750506000910152565b60008060008060008060c0878903121561189d57600080fd5b6118a687611843565b95506118b460208801611843565b9450604087015160ff811681146118ca57600080fd5b60608801519094506118db816115af565b60808801519093506118ec816115af565b60a08801519092506001600160401b0381111561190857600080fd5b8701601f8101891361191957600080fd5b80516119276116658261161f565b8181528a602083850101111561193c57600080fd5b61194d826020830160208601611858565b8093505050509295509295509295565b60006020828403121561196f57600080fd5b5051919050565b602080825260029082015261544560f01b604082015260600190565b6000602082840312156119a457600080fd5b81516111b1816115c7565b600063ffffffff808316818114156119c9576119c96117e8565b6001019392505050565b6000828210156119e5576119e56117e8565b500390565b6000602082840312156119fc57600080fd5b6111b182611843565b60008219821115611a1857611a186117e8565b500190565b6000816000190483118215151615611a3757611a376117e8565b500290565b600082611a5957634e487b7160e01b600052601260045260246000fd5b500490565b8281526040602082015260008251806040840152611a83816060850160208701611858565b601f01601f1916919091016060019392505050565b60008251611aaa818460208701611858565b919091019291505056fea26469706673582212208e356b39282cceb32efae8e814e9a7f80d587410b277b01e012e6351222ecc7d64736f6c634300080c0033",
}

// PledgePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PledgePoolMetaData.ABI instead.
var PledgePoolABI = PledgePoolMetaData.ABI

// Deprecated: Use PledgePoolMetaData.Sigs instead.
// PledgePoolFuncSigs maps the 4-byte function signature to its string representation.
var PledgePoolFuncSigs = PledgePoolMetaData.Sigs

// PledgePoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PledgePoolMetaData.Bin instead.
var PledgePoolBin = PledgePoolMetaData.Bin

// DeployPledgePool deploys a new Ethereum contract, binding an instance of PledgePool to it.
func DeployPledgePool(auth *bind.TransactOpts, backend bind.ContractBackend, primeToken common.Address, _rToken common.Address, r common.Address) (common.Address, *types.Transaction, *PledgePool, error) {
	parsed, err := PledgePoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	recoverAddr, _, _, _ := DeployRecover(auth, backend)
	PledgePoolBin = strings.Replace(PledgePoolBin, "__$942d0120f12f58797d9427bd9464d63c7a$__", recoverAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PledgePoolBin), backend, primeToken, _rToken, r)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PledgePool{PledgePoolCaller: PledgePoolCaller{contract: contract}, PledgePoolTransactor: PledgePoolTransactor{contract: contract}, PledgePoolFilterer: PledgePoolFilterer{contract: contract}}, nil
}

// PledgePool is an auto generated Go binding around an Ethereum contract.
type PledgePool struct {
	PledgePoolCaller     // Read-only binding to the contract
	PledgePoolTransactor // Write-only binding to the contract
	PledgePoolFilterer   // Log filterer for contract events
}

// PledgePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PledgePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PledgePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PledgePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PledgePoolSession struct {
	Contract     *PledgePool       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PledgePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PledgePoolCallerSession struct {
	Contract *PledgePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PledgePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PledgePoolTransactorSession struct {
	Contract     *PledgePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PledgePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PledgePoolRaw struct {
	Contract *PledgePool // Generic contract binding to access the raw methods on
}

// PledgePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PledgePoolCallerRaw struct {
	Contract *PledgePoolCaller // Generic read-only contract binding to access the raw methods on
}

// PledgePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PledgePoolTransactorRaw struct {
	Contract *PledgePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPledgePool creates a new instance of PledgePool, bound to a specific deployed contract.
func NewPledgePool(address common.Address, backend bind.ContractBackend) (*PledgePool, error) {
	contract, err := bindPledgePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PledgePool{PledgePoolCaller: PledgePoolCaller{contract: contract}, PledgePoolTransactor: PledgePoolTransactor{contract: contract}, PledgePoolFilterer: PledgePoolFilterer{contract: contract}}, nil
}

// NewPledgePoolCaller creates a new read-only instance of PledgePool, bound to a specific deployed contract.
func NewPledgePoolCaller(address common.Address, caller bind.ContractCaller) (*PledgePoolCaller, error) {
	contract, err := bindPledgePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PledgePoolCaller{contract: contract}, nil
}

// NewPledgePoolTransactor creates a new write-only instance of PledgePool, bound to a specific deployed contract.
func NewPledgePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PledgePoolTransactor, error) {
	contract, err := bindPledgePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTransactor{contract: contract}, nil
}

// NewPledgePoolFilterer creates a new log filterer instance of PledgePool, bound to a specific deployed contract.
func NewPledgePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PledgePoolFilterer, error) {
	contract, err := bindPledgePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PledgePoolFilterer{contract: contract}, nil
}

// bindPledgePool binds a generic wrapper to an already deployed contract.
func bindPledgePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PledgePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PledgePool *PledgePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PledgePool.Contract.PledgePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PledgePool *PledgePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PledgePool.Contract.PledgePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PledgePool *PledgePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PledgePool.Contract.PledgePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PledgePool *PledgePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PledgePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PledgePool *PledgePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PledgePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PledgePool *PledgePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PledgePool.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x761966b7.
//
// Solidity: function getBalance(uint64 index, uint32 tIndex) view returns(uint256)
func (_PledgePool *PledgePoolCaller) GetBalance(opts *bind.CallOpts, index uint64, tIndex uint32) (*big.Int, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "getBalance", index, tIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x761966b7.
//
// Solidity: function getBalance(uint64 index, uint32 tIndex) view returns(uint256)
func (_PledgePool *PledgePoolSession) GetBalance(index uint64, tIndex uint32) (*big.Int, error) {
	return _PledgePool.Contract.GetBalance(&_PledgePool.CallOpts, index, tIndex)
}

// GetBalance is a free data retrieval call binding the contract method 0x761966b7.
//
// Solidity: function getBalance(uint64 index, uint32 tIndex) view returns(uint256)
func (_PledgePool *PledgePoolCallerSession) GetBalance(index uint64, tIndex uint32) (*big.Int, error) {
	return _PledgePool.Contract.GetBalance(&_PledgePool.CallOpts, index, tIndex)
}

// GetPledge is a free data retrieval call binding the contract method 0xc3b19ccc.
//
// Solidity: function getPledge(uint32 tIndex) view returns(uint256)
func (_PledgePool *PledgePoolCaller) GetPledge(opts *bind.CallOpts, tIndex uint32) (*big.Int, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "getPledge", tIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPledge is a free data retrieval call binding the contract method 0xc3b19ccc.
//
// Solidity: function getPledge(uint32 tIndex) view returns(uint256)
func (_PledgePool *PledgePoolSession) GetPledge(tIndex uint32) (*big.Int, error) {
	return _PledgePool.Contract.GetPledge(&_PledgePool.CallOpts, tIndex)
}

// GetPledge is a free data retrieval call binding the contract method 0xc3b19ccc.
//
// Solidity: function getPledge(uint32 tIndex) view returns(uint256)
func (_PledgePool *PledgePoolCallerSession) GetPledge(tIndex uint32) (*big.Int, error) {
	return _PledgePool.Contract.GetPledge(&_PledgePool.CallOpts, tIndex)
}

// RToken is a free data retrieval call binding the contract method 0x40c65f72.
//
// Solidity: function rToken() view returns(address)
func (_PledgePool *PledgePoolCaller) RToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "rToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RToken is a free data retrieval call binding the contract method 0x40c65f72.
//
// Solidity: function rToken() view returns(address)
func (_PledgePool *PledgePoolSession) RToken() (common.Address, error) {
	return _PledgePool.Contract.RToken(&_PledgePool.CallOpts)
}

// RToken is a free data retrieval call binding the contract method 0x40c65f72.
//
// Solidity: function rToken() view returns(address)
func (_PledgePool *PledgePoolCallerSession) RToken() (common.Address, error) {
	return _PledgePool.Contract.RToken(&_PledgePool.CallOpts)
}

// Role is a free data retrieval call binding the contract method 0x10055c1d.
//
// Solidity: function role() view returns(address)
func (_PledgePool *PledgePoolCaller) Role(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "role")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Role is a free data retrieval call binding the contract method 0x10055c1d.
//
// Solidity: function role() view returns(address)
func (_PledgePool *PledgePoolSession) Role() (common.Address, error) {
	return _PledgePool.Contract.Role(&_PledgePool.CallOpts)
}

// Role is a free data retrieval call binding the contract method 0x10055c1d.
//
// Solidity: function role() view returns(address)
func (_PledgePool *PledgePoolCallerSession) Role() (common.Address, error) {
	return _PledgePool.Contract.Role(&_PledgePool.CallOpts)
}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_PledgePool *PledgePoolCaller) TotalPledge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "totalPledge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_PledgePool *PledgePoolSession) TotalPledge() (*big.Int, error) {
	return _PledgePool.Contract.TotalPledge(&_PledgePool.CallOpts)
}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_PledgePool *PledgePoolCallerSession) TotalPledge() (*big.Int, error) {
	return _PledgePool.Contract.TotalPledge(&_PledgePool.CallOpts)
}

// AddToken is a paid mutator transaction binding the contract method 0x8859344c.
//
// Solidity: function addToken(address tAddr, uint32 tIndex) returns()
func (_PledgePool *PledgePoolTransactor) AddToken(opts *bind.TransactOpts, tAddr common.Address, tIndex uint32) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "addToken", tAddr, tIndex)
}

// AddToken is a paid mutator transaction binding the contract method 0x8859344c.
//
// Solidity: function addToken(address tAddr, uint32 tIndex) returns()
func (_PledgePool *PledgePoolSession) AddToken(tAddr common.Address, tIndex uint32) (*types.Transaction, error) {
	return _PledgePool.Contract.AddToken(&_PledgePool.TransactOpts, tAddr, tIndex)
}

// AddToken is a paid mutator transaction binding the contract method 0x8859344c.
//
// Solidity: function addToken(address tAddr, uint32 tIndex) returns()
func (_PledgePool *PledgePoolTransactorSession) AddToken(tAddr common.Address, tIndex uint32) (*types.Transaction, error) {
	return _PledgePool.Contract.AddToken(&_PledgePool.TransactOpts, tAddr, tIndex)
}

// Pledge is a paid mutator transaction binding the contract method 0x364e4bf5.
//
// Solidity: function pledge(uint64 index, uint256 money, bytes sign) payable returns()
func (_PledgePool *PledgePoolTransactor) Pledge(opts *bind.TransactOpts, index uint64, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "pledge", index, money, sign)
}

// Pledge is a paid mutator transaction binding the contract method 0x364e4bf5.
//
// Solidity: function pledge(uint64 index, uint256 money, bytes sign) payable returns()
func (_PledgePool *PledgePoolSession) Pledge(index uint64, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _PledgePool.Contract.Pledge(&_PledgePool.TransactOpts, index, money, sign)
}

// Pledge is a paid mutator transaction binding the contract method 0x364e4bf5.
//
// Solidity: function pledge(uint64 index, uint256 money, bytes sign) payable returns()
func (_PledgePool *PledgePoolTransactorSession) Pledge(index uint64, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _PledgePool.Contract.Pledge(&_PledgePool.TransactOpts, index, money, sign)
}

// Withdraw is a paid mutator transaction binding the contract method 0x32704298.
//
// Solidity: function withdraw(uint64 index, uint32 tIndex, uint256 money, bytes sign) returns()
func (_PledgePool *PledgePoolTransactor) Withdraw(opts *bind.TransactOpts, index uint64, tIndex uint32, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "withdraw", index, tIndex, money, sign)
}

// Withdraw is a paid mutator transaction binding the contract method 0x32704298.
//
// Solidity: function withdraw(uint64 index, uint32 tIndex, uint256 money, bytes sign) returns()
func (_PledgePool *PledgePoolSession) Withdraw(index uint64, tIndex uint32, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _PledgePool.Contract.Withdraw(&_PledgePool.TransactOpts, index, tIndex, money, sign)
}

// Withdraw is a paid mutator transaction binding the contract method 0x32704298.
//
// Solidity: function withdraw(uint64 index, uint32 tIndex, uint256 money, bytes sign) returns()
func (_PledgePool *PledgePoolTransactorSession) Withdraw(index uint64, tIndex uint32, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _PledgePool.Contract.Withdraw(&_PledgePool.TransactOpts, index, tIndex, money, sign)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PledgePool *PledgePoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PledgePool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PledgePool *PledgePoolSession) Receive() (*types.Transaction, error) {
	return _PledgePool.Contract.Receive(&_PledgePool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PledgePool *PledgePoolTransactorSession) Receive() (*types.Transaction, error) {
	return _PledgePool.Contract.Receive(&_PledgePool.TransactOpts)
}

// PledgePoolPledgeIterator is returned from FilterPledge and is used to iterate over the raw logs and unpacked data for Pledge events raised by the PledgePool contract.
type PledgePoolPledgeIterator struct {
	Event *PledgePoolPledge // Event containing the contract specifics and raw log

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
func (it *PledgePoolPledgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolPledge)
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
		it.Event = new(PledgePoolPledge)
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
func (it *PledgePoolPledgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolPledgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolPledge represents a Pledge event raised by the PledgePool contract.
type PledgePoolPledge struct {
	From  common.Address
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPledge is a free log retrieval operation binding the contract event 0x5e91ea8ea1c46300eb761859be01d7b16d44389ef91e03a163a87413cbf55b95.
//
// Solidity: event Pledge(address indexed from, uint256 money)
func (_PledgePool *PledgePoolFilterer) FilterPledge(opts *bind.FilterOpts, from []common.Address) (*PledgePoolPledgeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "Pledge", fromRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolPledgeIterator{contract: _PledgePool.contract, event: "Pledge", logs: logs, sub: sub}, nil
}

// WatchPledge is a free log subscription operation binding the contract event 0x5e91ea8ea1c46300eb761859be01d7b16d44389ef91e03a163a87413cbf55b95.
//
// Solidity: event Pledge(address indexed from, uint256 money)
func (_PledgePool *PledgePoolFilterer) WatchPledge(opts *bind.WatchOpts, sink chan<- *PledgePoolPledge, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "Pledge", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolPledge)
				if err := _PledgePool.contract.UnpackLog(event, "Pledge", log); err != nil {
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

// ParsePledge is a log parse operation binding the contract event 0x5e91ea8ea1c46300eb761859be01d7b16d44389ef91e03a163a87413cbf55b95.
//
// Solidity: event Pledge(address indexed from, uint256 money)
func (_PledgePool *PledgePoolFilterer) ParsePledge(log types.Log) (*PledgePoolPledge, error) {
	event := new(PledgePoolPledge)
	if err := _PledgePool.contract.UnpackLog(event, "Pledge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the PledgePool contract.
type PledgePoolWithdrawIterator struct {
	Event *PledgePoolWithdraw // Event containing the contract specifics and raw log

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
func (it *PledgePoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolWithdraw)
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
		it.Event = new(PledgePoolWithdraw)
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
func (it *PledgePoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolWithdraw represents a Withdraw event raised by the PledgePool contract.
type PledgePoolWithdraw struct {
	From  common.Address
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed from, uint256 money)
func (_PledgePool *PledgePoolFilterer) FilterWithdraw(opts *bind.FilterOpts, from []common.Address) (*PledgePoolWithdrawIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "Withdraw", fromRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolWithdrawIterator{contract: _PledgePool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed from, uint256 money)
func (_PledgePool *PledgePoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PledgePoolWithdraw, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "Withdraw", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolWithdraw)
				if err := _PledgePool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed from, uint256 money)
func (_PledgePool *PledgePoolFilterer) ParseWithdraw(log types.Log) (*PledgePoolWithdraw, error) {
	event := new(PledgePoolWithdraw)
	if err := _PledgePool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RTokenMetaData contains all meta data concerning the RToken contract.
var RTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"AddT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"addT\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"getTA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"getTI\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTNum\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"isValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8ff8e15a": "addT(address)",
		"f7b7d6b6": "getTA(uint32)",
		"2df2685f": "getTI(address)",
		"b04c6c68": "getTNum()",
		"3c7bdc19": "isValid(uint32)",
	},
	Bin: "0x608060405234801561001057600080fd5b50600280546001600160a01b03191633179055610518806100326000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632df2685f1461005c5780633c7bdc19146100905780638ff8e15a146100b3578063b04c6c68146100db578063f7b7d6b6146100e3575b600080fd5b61006f61006a36600461042b565b61010e565b6040805163ffffffff90931683529015156020830152015b60405180910390f35b6100a361009e36600461045b565b6101e2565b6040519015158152602001610087565b6100c66100c136600461042b565b610259565b60405163ffffffff9091168152602001610087565b6000546100c6565b6100f66100f136600461045b565b61037c565b6040516001600160a01b039091168152602001610087565b6001600160a01b038116600090815260016020526040808220549051633c7bdc1960e01b815263ffffffff9091166004820181905282913090633c7bdc1990602401602060405180830381865afa15801561016d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101919190610481565b156101d85760008163ffffffff16815481106101af576101af6104a3565b6000918252602090912001546001600160a01b03858116911614156101d8579360019350915050565b9360009350915050565b6000805463ffffffff8316108015610244575060016000808463ffffffff1681548110610211576102116104a3565b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff64010000000090910416155b1561025157506001919050565b506000919050565b6002546000906001600160a01b0316331461029f5760405162461bcd60e51b81526020600482015260026024820152614e4f60f01b604482015260640160405180910390fd5b6102a8826103c1565b156102cf57506001600160a01b031660009081526001602052604090205463ffffffff1690565b60008054600180820183557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563820180546001600160a01b0319166001600160a01b0387169081179091558084526020918252604093849020805463ffffffff191663ffffffff851690811790915584519182529181019190915290917f7a5e6bb234636aa6f5c8428d056a65a5c9ec9d25638a69ad3bd3e362e64a8de6910160405180910390a192915050565b6000805463ffffffff831610156102515760008263ffffffff16815481106103a6576103a66104a3565b6000918252602090912001546001600160a01b031692915050565b6000805b60005481101561042257826001600160a01b0316600082815481106103ec576103ec6104a3565b6000918252602090912001546001600160a01b031614156104105750600192915050565b8061041a816104b9565b9150506103c5565b50600092915050565b60006020828403121561043d57600080fd5b81356001600160a01b038116811461045457600080fd5b9392505050565b60006020828403121561046d57600080fd5b813563ffffffff8116811461045457600080fd5b60006020828403121561049357600080fd5b8151801515811461045457600080fd5b634e487b7160e01b600052603260045260246000fd5b60006000198214156104db57634e487b7160e01b600052601160045260246000fd5b506001019056fea2646970667358221220bc7636b85fde76450de38627f9ca6d60fe8062016a1b555c902cc6955a14bc0764736f6c634300080c0033",
}

// RTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use RTokenMetaData.ABI instead.
var RTokenABI = RTokenMetaData.ABI

// Deprecated: Use RTokenMetaData.Sigs instead.
// RTokenFuncSigs maps the 4-byte function signature to its string representation.
var RTokenFuncSigs = RTokenMetaData.Sigs

// RTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RTokenMetaData.Bin instead.
var RTokenBin = RTokenMetaData.Bin

// DeployRToken deploys a new Ethereum contract, binding an instance of RToken to it.
func DeployRToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RToken, error) {
	parsed, err := RTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RToken{RTokenCaller: RTokenCaller{contract: contract}, RTokenTransactor: RTokenTransactor{contract: contract}, RTokenFilterer: RTokenFilterer{contract: contract}}, nil
}

// RToken is an auto generated Go binding around an Ethereum contract.
type RToken struct {
	RTokenCaller     // Read-only binding to the contract
	RTokenTransactor // Write-only binding to the contract
	RTokenFilterer   // Log filterer for contract events
}

// RTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type RTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RTokenSession struct {
	Contract     *RToken           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RTokenCallerSession struct {
	Contract *RTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RTokenTransactorSession struct {
	Contract     *RTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type RTokenRaw struct {
	Contract *RToken // Generic contract binding to access the raw methods on
}

// RTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RTokenCallerRaw struct {
	Contract *RTokenCaller // Generic read-only contract binding to access the raw methods on
}

// RTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RTokenTransactorRaw struct {
	Contract *RTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRToken creates a new instance of RToken, bound to a specific deployed contract.
func NewRToken(address common.Address, backend bind.ContractBackend) (*RToken, error) {
	contract, err := bindRToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RToken{RTokenCaller: RTokenCaller{contract: contract}, RTokenTransactor: RTokenTransactor{contract: contract}, RTokenFilterer: RTokenFilterer{contract: contract}}, nil
}

// NewRTokenCaller creates a new read-only instance of RToken, bound to a specific deployed contract.
func NewRTokenCaller(address common.Address, caller bind.ContractCaller) (*RTokenCaller, error) {
	contract, err := bindRToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RTokenCaller{contract: contract}, nil
}

// NewRTokenTransactor creates a new write-only instance of RToken, bound to a specific deployed contract.
func NewRTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*RTokenTransactor, error) {
	contract, err := bindRToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RTokenTransactor{contract: contract}, nil
}

// NewRTokenFilterer creates a new log filterer instance of RToken, bound to a specific deployed contract.
func NewRTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*RTokenFilterer, error) {
	contract, err := bindRToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RTokenFilterer{contract: contract}, nil
}

// bindRToken binds a generic wrapper to an already deployed contract.
func bindRToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RToken *RTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RToken.Contract.RTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RToken *RTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RToken.Contract.RTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RToken *RTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RToken.Contract.RTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RToken *RTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RToken *RTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RToken *RTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RToken.Contract.contract.Transact(opts, method, params...)
}

// GetTA is a free data retrieval call binding the contract method 0xf7b7d6b6.
//
// Solidity: function getTA(uint32 tIndex) view returns(address)
func (_RToken *RTokenCaller) GetTA(opts *bind.CallOpts, tIndex uint32) (common.Address, error) {
	var out []interface{}
	err := _RToken.contract.Call(opts, &out, "getTA", tIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTA is a free data retrieval call binding the contract method 0xf7b7d6b6.
//
// Solidity: function getTA(uint32 tIndex) view returns(address)
func (_RToken *RTokenSession) GetTA(tIndex uint32) (common.Address, error) {
	return _RToken.Contract.GetTA(&_RToken.CallOpts, tIndex)
}

// GetTA is a free data retrieval call binding the contract method 0xf7b7d6b6.
//
// Solidity: function getTA(uint32 tIndex) view returns(address)
func (_RToken *RTokenCallerSession) GetTA(tIndex uint32) (common.Address, error) {
	return _RToken.Contract.GetTA(&_RToken.CallOpts, tIndex)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address t) view returns(uint32, bool)
func (_RToken *RTokenCaller) GetTI(opts *bind.CallOpts, t common.Address) (uint32, bool, error) {
	var out []interface{}
	err := _RToken.contract.Call(opts, &out, "getTI", t)

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
func (_RToken *RTokenSession) GetTI(t common.Address) (uint32, bool, error) {
	return _RToken.Contract.GetTI(&_RToken.CallOpts, t)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address t) view returns(uint32, bool)
func (_RToken *RTokenCallerSession) GetTI(t common.Address) (uint32, bool, error) {
	return _RToken.Contract.GetTI(&_RToken.CallOpts, t)
}

// GetTNum is a free data retrieval call binding the contract method 0xb04c6c68.
//
// Solidity: function getTNum() view returns(uint32)
func (_RToken *RTokenCaller) GetTNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _RToken.contract.Call(opts, &out, "getTNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTNum is a free data retrieval call binding the contract method 0xb04c6c68.
//
// Solidity: function getTNum() view returns(uint32)
func (_RToken *RTokenSession) GetTNum() (uint32, error) {
	return _RToken.Contract.GetTNum(&_RToken.CallOpts)
}

// GetTNum is a free data retrieval call binding the contract method 0xb04c6c68.
//
// Solidity: function getTNum() view returns(uint32)
func (_RToken *RTokenCallerSession) GetTNum() (uint32, error) {
	return _RToken.Contract.GetTNum(&_RToken.CallOpts)
}

// IsValid is a free data retrieval call binding the contract method 0x3c7bdc19.
//
// Solidity: function isValid(uint32 tIndex) view returns(bool)
func (_RToken *RTokenCaller) IsValid(opts *bind.CallOpts, tIndex uint32) (bool, error) {
	var out []interface{}
	err := _RToken.contract.Call(opts, &out, "isValid", tIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValid is a free data retrieval call binding the contract method 0x3c7bdc19.
//
// Solidity: function isValid(uint32 tIndex) view returns(bool)
func (_RToken *RTokenSession) IsValid(tIndex uint32) (bool, error) {
	return _RToken.Contract.IsValid(&_RToken.CallOpts, tIndex)
}

// IsValid is a free data retrieval call binding the contract method 0x3c7bdc19.
//
// Solidity: function isValid(uint32 tIndex) view returns(bool)
func (_RToken *RTokenCallerSession) IsValid(tIndex uint32) (bool, error) {
	return _RToken.Contract.IsValid(&_RToken.CallOpts, tIndex)
}

// AddT is a paid mutator transaction binding the contract method 0x8ff8e15a.
//
// Solidity: function addT(address t) returns(uint32 index)
func (_RToken *RTokenTransactor) AddT(opts *bind.TransactOpts, t common.Address) (*types.Transaction, error) {
	return _RToken.contract.Transact(opts, "addT", t)
}

// AddT is a paid mutator transaction binding the contract method 0x8ff8e15a.
//
// Solidity: function addT(address t) returns(uint32 index)
func (_RToken *RTokenSession) AddT(t common.Address) (*types.Transaction, error) {
	return _RToken.Contract.AddT(&_RToken.TransactOpts, t)
}

// AddT is a paid mutator transaction binding the contract method 0x8ff8e15a.
//
// Solidity: function addT(address t) returns(uint32 index)
func (_RToken *RTokenTransactorSession) AddT(t common.Address) (*types.Transaction, error) {
	return _RToken.Contract.AddT(&_RToken.TransactOpts, t)
}

// RTokenAddTIterator is returned from FilterAddT and is used to iterate over the raw logs and unpacked data for AddT events raised by the RToken contract.
type RTokenAddTIterator struct {
	Event *RTokenAddT // Event containing the contract specifics and raw log

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
func (it *RTokenAddTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTokenAddT)
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
		it.Event = new(RTokenAddT)
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
func (it *RTokenAddTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTokenAddTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTokenAddT represents a AddT event raised by the RToken contract.
type RTokenAddT struct {
	T      common.Address
	TIndex uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddT is a free log retrieval operation binding the contract event 0x7a5e6bb234636aa6f5c8428d056a65a5c9ec9d25638a69ad3bd3e362e64a8de6.
//
// Solidity: event AddT(address t, uint32 tIndex)
func (_RToken *RTokenFilterer) FilterAddT(opts *bind.FilterOpts) (*RTokenAddTIterator, error) {

	logs, sub, err := _RToken.contract.FilterLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return &RTokenAddTIterator{contract: _RToken.contract, event: "AddT", logs: logs, sub: sub}, nil
}

// WatchAddT is a free log subscription operation binding the contract event 0x7a5e6bb234636aa6f5c8428d056a65a5c9ec9d25638a69ad3bd3e362e64a8de6.
//
// Solidity: event AddT(address t, uint32 tIndex)
func (_RToken *RTokenFilterer) WatchAddT(opts *bind.WatchOpts, sink chan<- *RTokenAddT) (event.Subscription, error) {

	logs, sub, err := _RToken.contract.WatchLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTokenAddT)
				if err := _RToken.contract.UnpackLog(event, "AddT", log); err != nil {
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

// ParseAddT is a log parse operation binding the contract event 0x7a5e6bb234636aa6f5c8428d056a65a5c9ec9d25638a69ad3bd3e362e64a8de6.
//
// Solidity: event AddT(address t, uint32 tIndex)
func (_RToken *RTokenFilterer) ParseAddT(log types.Log) (*RTokenAddT, error) {
	event := new(RTokenAddT)
	if err := _RToken.contract.UnpackLog(event, "AddT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
