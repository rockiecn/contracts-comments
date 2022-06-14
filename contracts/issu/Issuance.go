// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package issu

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

// IssuParams is an auto generated low-level Go binding around an user-defined struct.
type IssuParams struct {
	Start  uint64
	End    uint64
	Size   uint64
	SPrice *big.Int
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

// IssuanceMetaData contains all meta data concerning the Issuance contract.
var IssuanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rfs\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIssuParams\",\"name\":\"ps\",\"type\":\"tuple\"}],\"name\":\"issu\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"issuRatio\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minRatio\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodTarget\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodTotalReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_add\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sub\",\"type\":\"uint256\"}],\"name\":\"setTP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spaceTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"subPMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"subSMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c862ea50": "issu((uint64,uint64,uint64,uint256))",
		"44178c25": "issuRatio()",
		"586fc5b5": "lastMint()",
		"86d8745b": "minRatio()",
		"988bda70": "mintLevel()",
		"030bb079": "periodTarget()",
		"517463e6": "periodTotalReward()",
		"a035b1fe": "price()",
		"11e65fc0": "setTP(uint256,uint256)",
		"949d225d": "size()",
		"43a2755c": "spaceTime()",
		"f0b44ab1": "subPMap(uint64)",
		"45fe29f6": "subSMap(uint64)",
		"e7b0f666": "totalPaid()",
		"e154adf5": "totalPay()",
	},
	Bin: "0x608060405234801561001057600080fd5b50604051610e50380380610e5083398101604081905261002f9161045d565b60004290508060028190555060016040518060600160405280603261ffff1681526020016564000000000081526020016283d6006001600160401b0316815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff1602179055506020820151816001015560408201518160020160006101000a8154816001600160401b0302191690836001600160401b03160217905550505060016040518060600160405280605061ffff1681526020016604000000000000815260200162c5c1006001600160401b0316815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff1602179055506020820151816001015560408201518160020160006101000a8154816001600160401b0302191690836001600160401b03160217905550505060016040518060600160405280606461ffff16815260200166c80000000000008152602001630107ac006001600160401b0316815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff1602179055506020820151816001015560408201518160020160006101000a8154816001600160401b0302191690836001600160401b03160217905550505060016040518060600160405280605061ffff16815260200167100000000000000081526020016283d6006001600160401b0316815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff1602179055506020820151816001015560408201518160020160006101000a8154816001600160401b0302191690836001600160401b03160217905550505060016040518060600160405280603261ffff1681526020016803200000000000000081526020016303c267006001600160401b0316815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff1602179055506020820151816001015560408201518160020160006101000a8154816001600160401b0302191690836001600160401b03160217905550505081600d60006101000a8154816001600160a01b0302191690836001600160a01b031602179055506003600a60026101000a81548161ffff021916908361ffff1602179055506a7c13bc4b2c133c560000006008819055506032600a60006101000a81548161ffff021916908361ffff160217905550505061048d565b60006020828403121561046f57600080fd5b81516001600160a01b038116811461048657600080fd5b9392505050565b6109b48061049c6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806386d8745b11610097578063c862ea5011610066578063c862ea50146101b6578063e154adf5146101c9578063e7b0f666146101d2578063f0b44ab1146101db57600080fd5b806386d8745b14610187578063949d225d1461019b578063988bda70146101a4578063a035b1fe146101ad57600080fd5b806344178c25116100d357806344178c251461013457806345fe29f614610155578063517463e614610175578063586fc5b51461017e57600080fd5b8063030bb079146100fa57806311e65fc01461011657806343a2755c1461012b575b600080fd5b61010360085481565b6040519081526020015b60405180910390f35b610129610124366004610795565b6101fb565b005b61010360055481565b600a546101429061ffff1681565b60405161ffff909116815260200161010d565b6101036101633660046107ce565b600c6020526000908152604090205481565b61010360095481565b61010360025481565b600a546101429062010000900461ffff1681565b61010360045481565b61010360005481565b61010360035481565b6101036101c43660046107f0565b61027a565b61010360065481565b61010360075481565b6101036101e93660046107ce565b600b6020526000908152604090205481565b600d546001600160a01b0316331461023e5760405162461bcd60e51b81526020600482015260016024820152602760f91b60448201526064015b60405180910390fd5b8160076000828254610250919061088c565b9091555050600754811161027657806007600082825461027091906108a4565b90915550505b5050565b600d546000906001600160a01b031633146102bb5760405162461bcd60e51b81526020600482015260016024820152602760f91b6044820152606401610235565b60608201516020808401516001600160401b03166000908152600b9091526040812080549091906102ed90849061088c565b9250508190555081604001516001600160401b0316600c600084602001516001600160401b03166001600160401b03168152602001908152602001600020600082825461033a919061088c565b90915550506002544290620151809061035390836108a4565b111561036c57600254610369906201518061088c565b90505b60006002548261037c91906108a4565b905060008160035461038e91906108bb565b905061039d62015180846108f0565b620151806002546103ae91906108f0565b101561046a5760006103c362015180856108f0565b6103d090620151806108bb565b6001600160401b0381166000908152600b6020526040902054909150801561042f5760006103fe83876108a4565b61040890836108bb565b905061041481856108a4565b9350816003600082825461042891906108a4565b9091555050505b6001600160401b0382166000908152600c6020526040902054801561046657806004600082825461046091906108a4565b90915550505b5050505b806007600082825461047c919061088c565b90915550508451602086015160009161049491610904565b6001600160401b031686604001516001600160401b03166104b591906108bb565b905080600560008282546104c9919061088c565b9091555050855160208701516000916104e191610904565b6001600160401b031687606001516104f991906108bb565b9050806006600082825461050d919061088c565b9250508190555086604001516001600160401b031660046000828254610533919061088c565b909155505060608701516003805460009061054f90849061088c565b90915550508261056b5750505060029190915550600092915050565b6000545b6001548110156106805760006001828154811061058e5761058e61092c565b90600052602060002090600302016001015490506004548110156105b157506004545b6000600183815481106105c6576105c661092c565b60009182526020909120600260039092020101546005546001600160401b03909116915081906105f79084906108f0565b106106645761060783600161088c565b6000819055506001600054815481106106225761062261092c565b6000918252602090912060039091020154600a805461ffff191661ffff909216919091179055600854610657906002906108f0565b600855600060095561066b565b5050610680565b5050808061067890610942565b91505061056f565b50600a546000906106959061ffff16856108bb565b90506106a26064826108f0565b9050856002819055506000816009546106bb919061088c565b9050600854811180156106db5750600a5461ffff62010000820481169116115b156107875760006009546008546106f291906108a4565b9050600061070082856108a4565b905061070d6002826108f0565b905061071981856108a4565b9350836009600082825461072d919061088c565b9091555050600a546107459060029061ffff1661095d565b600a805461ffff191661ffff9283169081179182905562010000909104909116111561078457600a805462010000810461ffff1661ffff199091161790555b50505b50955050505050505b919050565b600080604083850312156107a857600080fd5b50508035926020909101359150565b80356001600160401b038116811461079057600080fd5b6000602082840312156107e057600080fd5b6107e9826107b7565b9392505050565b60006080828403121561080257600080fd5b604051608081018181106001600160401b038211171561083257634e487b7160e01b600052604160045260246000fd5b60405261083e836107b7565b815261084c602084016107b7565b602082015261085d604084016107b7565b6040820152606083013560608201528091505092915050565b634e487b7160e01b600052601160045260246000fd5b6000821982111561089f5761089f610876565b500190565b6000828210156108b6576108b6610876565b500390565b60008160001904831182151516156108d5576108d5610876565b500290565b634e487b7160e01b600052601260045260246000fd5b6000826108ff576108ff6108da565b500490565b60006001600160401b038381169083168181101561092457610924610876565b039392505050565b634e487b7160e01b600052603260045260246000fd5b600060001982141561095657610956610876565b5060010190565b600061ffff80841680610972576109726108da565b9216919091049291505056fea2646970667358221220172c955c2790c6facc63faf73a0abc85b063563d4fbe46f5dd9c4377527fa23e64736f6c634300080c0033",
}

// IssuanceABI is the input ABI used to generate the binding from.
// Deprecated: Use IssuanceMetaData.ABI instead.
var IssuanceABI = IssuanceMetaData.ABI

// Deprecated: Use IssuanceMetaData.Sigs instead.
// IssuanceFuncSigs maps the 4-byte function signature to its string representation.
var IssuanceFuncSigs = IssuanceMetaData.Sigs

// IssuanceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IssuanceMetaData.Bin instead.
var IssuanceBin = IssuanceMetaData.Bin

// DeployIssuance deploys a new Ethereum contract, binding an instance of Issuance to it.
func DeployIssuance(auth *bind.TransactOpts, backend bind.ContractBackend, rfs common.Address) (common.Address, *types.Transaction, *Issuance, error) {
	parsed, err := IssuanceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IssuanceBin), backend, rfs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Issuance{IssuanceCaller: IssuanceCaller{contract: contract}, IssuanceTransactor: IssuanceTransactor{contract: contract}, IssuanceFilterer: IssuanceFilterer{contract: contract}}, nil
}

// Issuance is an auto generated Go binding around an Ethereum contract.
type Issuance struct {
	IssuanceCaller     // Read-only binding to the contract
	IssuanceTransactor // Write-only binding to the contract
	IssuanceFilterer   // Log filterer for contract events
}

// IssuanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IssuanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IssuanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IssuanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IssuanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IssuanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IssuanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IssuanceSession struct {
	Contract     *Issuance         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IssuanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IssuanceCallerSession struct {
	Contract *IssuanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IssuanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IssuanceTransactorSession struct {
	Contract     *IssuanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IssuanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IssuanceRaw struct {
	Contract *Issuance // Generic contract binding to access the raw methods on
}

// IssuanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IssuanceCallerRaw struct {
	Contract *IssuanceCaller // Generic read-only contract binding to access the raw methods on
}

// IssuanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IssuanceTransactorRaw struct {
	Contract *IssuanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIssuance creates a new instance of Issuance, bound to a specific deployed contract.
func NewIssuance(address common.Address, backend bind.ContractBackend) (*Issuance, error) {
	contract, err := bindIssuance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Issuance{IssuanceCaller: IssuanceCaller{contract: contract}, IssuanceTransactor: IssuanceTransactor{contract: contract}, IssuanceFilterer: IssuanceFilterer{contract: contract}}, nil
}

// NewIssuanceCaller creates a new read-only instance of Issuance, bound to a specific deployed contract.
func NewIssuanceCaller(address common.Address, caller bind.ContractCaller) (*IssuanceCaller, error) {
	contract, err := bindIssuance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IssuanceCaller{contract: contract}, nil
}

// NewIssuanceTransactor creates a new write-only instance of Issuance, bound to a specific deployed contract.
func NewIssuanceTransactor(address common.Address, transactor bind.ContractTransactor) (*IssuanceTransactor, error) {
	contract, err := bindIssuance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IssuanceTransactor{contract: contract}, nil
}

// NewIssuanceFilterer creates a new log filterer instance of Issuance, bound to a specific deployed contract.
func NewIssuanceFilterer(address common.Address, filterer bind.ContractFilterer) (*IssuanceFilterer, error) {
	contract, err := bindIssuance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IssuanceFilterer{contract: contract}, nil
}

// bindIssuance binds a generic wrapper to an already deployed contract.
func bindIssuance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IssuanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Issuance *IssuanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Issuance.Contract.IssuanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Issuance *IssuanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Issuance.Contract.IssuanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Issuance *IssuanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Issuance.Contract.IssuanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Issuance *IssuanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Issuance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Issuance *IssuanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Issuance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Issuance *IssuanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Issuance.Contract.contract.Transact(opts, method, params...)
}

// IssuRatio is a free data retrieval call binding the contract method 0x44178c25.
//
// Solidity: function issuRatio() view returns(uint16)
func (_Issuance *IssuanceCaller) IssuRatio(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "issuRatio")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// IssuRatio is a free data retrieval call binding the contract method 0x44178c25.
//
// Solidity: function issuRatio() view returns(uint16)
func (_Issuance *IssuanceSession) IssuRatio() (uint16, error) {
	return _Issuance.Contract.IssuRatio(&_Issuance.CallOpts)
}

// IssuRatio is a free data retrieval call binding the contract method 0x44178c25.
//
// Solidity: function issuRatio() view returns(uint16)
func (_Issuance *IssuanceCallerSession) IssuRatio() (uint16, error) {
	return _Issuance.Contract.IssuRatio(&_Issuance.CallOpts)
}

// LastMint is a free data retrieval call binding the contract method 0x586fc5b5.
//
// Solidity: function lastMint() view returns(uint256)
func (_Issuance *IssuanceCaller) LastMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "lastMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastMint is a free data retrieval call binding the contract method 0x586fc5b5.
//
// Solidity: function lastMint() view returns(uint256)
func (_Issuance *IssuanceSession) LastMint() (*big.Int, error) {
	return _Issuance.Contract.LastMint(&_Issuance.CallOpts)
}

// LastMint is a free data retrieval call binding the contract method 0x586fc5b5.
//
// Solidity: function lastMint() view returns(uint256)
func (_Issuance *IssuanceCallerSession) LastMint() (*big.Int, error) {
	return _Issuance.Contract.LastMint(&_Issuance.CallOpts)
}

// MinRatio is a free data retrieval call binding the contract method 0x86d8745b.
//
// Solidity: function minRatio() view returns(uint16)
func (_Issuance *IssuanceCaller) MinRatio(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "minRatio")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MinRatio is a free data retrieval call binding the contract method 0x86d8745b.
//
// Solidity: function minRatio() view returns(uint16)
func (_Issuance *IssuanceSession) MinRatio() (uint16, error) {
	return _Issuance.Contract.MinRatio(&_Issuance.CallOpts)
}

// MinRatio is a free data retrieval call binding the contract method 0x86d8745b.
//
// Solidity: function minRatio() view returns(uint16)
func (_Issuance *IssuanceCallerSession) MinRatio() (uint16, error) {
	return _Issuance.Contract.MinRatio(&_Issuance.CallOpts)
}

// MintLevel is a free data retrieval call binding the contract method 0x988bda70.
//
// Solidity: function mintLevel() view returns(uint256)
func (_Issuance *IssuanceCaller) MintLevel(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "mintLevel")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintLevel is a free data retrieval call binding the contract method 0x988bda70.
//
// Solidity: function mintLevel() view returns(uint256)
func (_Issuance *IssuanceSession) MintLevel() (*big.Int, error) {
	return _Issuance.Contract.MintLevel(&_Issuance.CallOpts)
}

// MintLevel is a free data retrieval call binding the contract method 0x988bda70.
//
// Solidity: function mintLevel() view returns(uint256)
func (_Issuance *IssuanceCallerSession) MintLevel() (*big.Int, error) {
	return _Issuance.Contract.MintLevel(&_Issuance.CallOpts)
}

// PeriodTarget is a free data retrieval call binding the contract method 0x030bb079.
//
// Solidity: function periodTarget() view returns(uint256)
func (_Issuance *IssuanceCaller) PeriodTarget(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "periodTarget")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodTarget is a free data retrieval call binding the contract method 0x030bb079.
//
// Solidity: function periodTarget() view returns(uint256)
func (_Issuance *IssuanceSession) PeriodTarget() (*big.Int, error) {
	return _Issuance.Contract.PeriodTarget(&_Issuance.CallOpts)
}

// PeriodTarget is a free data retrieval call binding the contract method 0x030bb079.
//
// Solidity: function periodTarget() view returns(uint256)
func (_Issuance *IssuanceCallerSession) PeriodTarget() (*big.Int, error) {
	return _Issuance.Contract.PeriodTarget(&_Issuance.CallOpts)
}

// PeriodTotalReward is a free data retrieval call binding the contract method 0x517463e6.
//
// Solidity: function periodTotalReward() view returns(uint256)
func (_Issuance *IssuanceCaller) PeriodTotalReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "periodTotalReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodTotalReward is a free data retrieval call binding the contract method 0x517463e6.
//
// Solidity: function periodTotalReward() view returns(uint256)
func (_Issuance *IssuanceSession) PeriodTotalReward() (*big.Int, error) {
	return _Issuance.Contract.PeriodTotalReward(&_Issuance.CallOpts)
}

// PeriodTotalReward is a free data retrieval call binding the contract method 0x517463e6.
//
// Solidity: function periodTotalReward() view returns(uint256)
func (_Issuance *IssuanceCallerSession) PeriodTotalReward() (*big.Int, error) {
	return _Issuance.Contract.PeriodTotalReward(&_Issuance.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Issuance *IssuanceCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Issuance *IssuanceSession) Price() (*big.Int, error) {
	return _Issuance.Contract.Price(&_Issuance.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Issuance *IssuanceCallerSession) Price() (*big.Int, error) {
	return _Issuance.Contract.Price(&_Issuance.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_Issuance *IssuanceCaller) Size(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "size")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_Issuance *IssuanceSession) Size() (*big.Int, error) {
	return _Issuance.Contract.Size(&_Issuance.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_Issuance *IssuanceCallerSession) Size() (*big.Int, error) {
	return _Issuance.Contract.Size(&_Issuance.CallOpts)
}

// SpaceTime is a free data retrieval call binding the contract method 0x43a2755c.
//
// Solidity: function spaceTime() view returns(uint256)
func (_Issuance *IssuanceCaller) SpaceTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "spaceTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpaceTime is a free data retrieval call binding the contract method 0x43a2755c.
//
// Solidity: function spaceTime() view returns(uint256)
func (_Issuance *IssuanceSession) SpaceTime() (*big.Int, error) {
	return _Issuance.Contract.SpaceTime(&_Issuance.CallOpts)
}

// SpaceTime is a free data retrieval call binding the contract method 0x43a2755c.
//
// Solidity: function spaceTime() view returns(uint256)
func (_Issuance *IssuanceCallerSession) SpaceTime() (*big.Int, error) {
	return _Issuance.Contract.SpaceTime(&_Issuance.CallOpts)
}

// SubPMap is a free data retrieval call binding the contract method 0xf0b44ab1.
//
// Solidity: function subPMap(uint64 ) view returns(uint256)
func (_Issuance *IssuanceCaller) SubPMap(opts *bind.CallOpts, arg0 uint64) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "subPMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubPMap is a free data retrieval call binding the contract method 0xf0b44ab1.
//
// Solidity: function subPMap(uint64 ) view returns(uint256)
func (_Issuance *IssuanceSession) SubPMap(arg0 uint64) (*big.Int, error) {
	return _Issuance.Contract.SubPMap(&_Issuance.CallOpts, arg0)
}

// SubPMap is a free data retrieval call binding the contract method 0xf0b44ab1.
//
// Solidity: function subPMap(uint64 ) view returns(uint256)
func (_Issuance *IssuanceCallerSession) SubPMap(arg0 uint64) (*big.Int, error) {
	return _Issuance.Contract.SubPMap(&_Issuance.CallOpts, arg0)
}

// SubSMap is a free data retrieval call binding the contract method 0x45fe29f6.
//
// Solidity: function subSMap(uint64 ) view returns(uint256)
func (_Issuance *IssuanceCaller) SubSMap(opts *bind.CallOpts, arg0 uint64) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "subSMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubSMap is a free data retrieval call binding the contract method 0x45fe29f6.
//
// Solidity: function subSMap(uint64 ) view returns(uint256)
func (_Issuance *IssuanceSession) SubSMap(arg0 uint64) (*big.Int, error) {
	return _Issuance.Contract.SubSMap(&_Issuance.CallOpts, arg0)
}

// SubSMap is a free data retrieval call binding the contract method 0x45fe29f6.
//
// Solidity: function subSMap(uint64 ) view returns(uint256)
func (_Issuance *IssuanceCallerSession) SubSMap(arg0 uint64) (*big.Int, error) {
	return _Issuance.Contract.SubSMap(&_Issuance.CallOpts, arg0)
}

// TotalPaid is a free data retrieval call binding the contract method 0xe7b0f666.
//
// Solidity: function totalPaid() view returns(uint256)
func (_Issuance *IssuanceCaller) TotalPaid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "totalPaid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPaid is a free data retrieval call binding the contract method 0xe7b0f666.
//
// Solidity: function totalPaid() view returns(uint256)
func (_Issuance *IssuanceSession) TotalPaid() (*big.Int, error) {
	return _Issuance.Contract.TotalPaid(&_Issuance.CallOpts)
}

// TotalPaid is a free data retrieval call binding the contract method 0xe7b0f666.
//
// Solidity: function totalPaid() view returns(uint256)
func (_Issuance *IssuanceCallerSession) TotalPaid() (*big.Int, error) {
	return _Issuance.Contract.TotalPaid(&_Issuance.CallOpts)
}

// TotalPay is a free data retrieval call binding the contract method 0xe154adf5.
//
// Solidity: function totalPay() view returns(uint256)
func (_Issuance *IssuanceCaller) TotalPay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "totalPay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPay is a free data retrieval call binding the contract method 0xe154adf5.
//
// Solidity: function totalPay() view returns(uint256)
func (_Issuance *IssuanceSession) TotalPay() (*big.Int, error) {
	return _Issuance.Contract.TotalPay(&_Issuance.CallOpts)
}

// TotalPay is a free data retrieval call binding the contract method 0xe154adf5.
//
// Solidity: function totalPay() view returns(uint256)
func (_Issuance *IssuanceCallerSession) TotalPay() (*big.Int, error) {
	return _Issuance.Contract.TotalPay(&_Issuance.CallOpts)
}

// Issu is a paid mutator transaction binding the contract method 0xc862ea50.
//
// Solidity: function issu((uint64,uint64,uint64,uint256) ps) returns(uint256)
func (_Issuance *IssuanceTransactor) Issu(opts *bind.TransactOpts, ps IssuParams) (*types.Transaction, error) {
	return _Issuance.contract.Transact(opts, "issu", ps)
}

// Issu is a paid mutator transaction binding the contract method 0xc862ea50.
//
// Solidity: function issu((uint64,uint64,uint64,uint256) ps) returns(uint256)
func (_Issuance *IssuanceSession) Issu(ps IssuParams) (*types.Transaction, error) {
	return _Issuance.Contract.Issu(&_Issuance.TransactOpts, ps)
}

// Issu is a paid mutator transaction binding the contract method 0xc862ea50.
//
// Solidity: function issu((uint64,uint64,uint64,uint256) ps) returns(uint256)
func (_Issuance *IssuanceTransactorSession) Issu(ps IssuParams) (*types.Transaction, error) {
	return _Issuance.Contract.Issu(&_Issuance.TransactOpts, ps)
}

// SetTP is a paid mutator transaction binding the contract method 0x11e65fc0.
//
// Solidity: function setTP(uint256 _add, uint256 _sub) returns()
func (_Issuance *IssuanceTransactor) SetTP(opts *bind.TransactOpts, _add *big.Int, _sub *big.Int) (*types.Transaction, error) {
	return _Issuance.contract.Transact(opts, "setTP", _add, _sub)
}

// SetTP is a paid mutator transaction binding the contract method 0x11e65fc0.
//
// Solidity: function setTP(uint256 _add, uint256 _sub) returns()
func (_Issuance *IssuanceSession) SetTP(_add *big.Int, _sub *big.Int) (*types.Transaction, error) {
	return _Issuance.Contract.SetTP(&_Issuance.TransactOpts, _add, _sub)
}

// SetTP is a paid mutator transaction binding the contract method 0x11e65fc0.
//
// Solidity: function setTP(uint256 _add, uint256 _sub) returns()
func (_Issuance *IssuanceTransactorSession) SetTP(_add *big.Int, _sub *big.Int) (*types.Transaction, error) {
	return _Issuance.Contract.SetTP(&_Issuance.TransactOpts, _add, _sub)
}
