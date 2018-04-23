// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// AddrSetABI is the input ABI used to generate the binding from.
const AddrSetABI = "[]"

// AddrSetBin is the compiled bytecode used for deploying new contracts.
const AddrSetBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146060604052600080fd00a165627a7a7230582046c72a62cd98aba474a608a422fc144df02b9461b8c7732d167cb4fb1fc7539c0029`

// DeployAddrSet deploys a new Ethereum contract, binding an instance of AddrSet to it.
func DeployAddrSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddrSet, error) {
	parsed, err := abi.JSON(strings.NewReader(AddrSetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddrSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddrSet{AddrSetCaller: AddrSetCaller{contract: contract}, AddrSetTransactor: AddrSetTransactor{contract: contract}, AddrSetFilterer: AddrSetFilterer{contract: contract}}, nil
}

// AddrSet is an auto generated Go binding around an Ethereum contract.
type AddrSet struct {
	AddrSetCaller     // Read-only binding to the contract
	AddrSetTransactor // Write-only binding to the contract
	AddrSetFilterer   // Log filterer for contract events
}

// AddrSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddrSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddrSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddrSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddrSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddrSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddrSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddrSetSession struct {
	Contract     *AddrSet          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddrSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddrSetCallerSession struct {
	Contract *AddrSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddrSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddrSetTransactorSession struct {
	Contract     *AddrSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddrSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddrSetRaw struct {
	Contract *AddrSet // Generic contract binding to access the raw methods on
}

// AddrSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddrSetCallerRaw struct {
	Contract *AddrSetCaller // Generic read-only contract binding to access the raw methods on
}

// AddrSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddrSetTransactorRaw struct {
	Contract *AddrSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddrSet creates a new instance of AddrSet, bound to a specific deployed contract.
func NewAddrSet(address common.Address, backend bind.ContractBackend) (*AddrSet, error) {
	contract, err := bindAddrSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddrSet{AddrSetCaller: AddrSetCaller{contract: contract}, AddrSetTransactor: AddrSetTransactor{contract: contract}, AddrSetFilterer: AddrSetFilterer{contract: contract}}, nil
}

// NewAddrSetCaller creates a new read-only instance of AddrSet, bound to a specific deployed contract.
func NewAddrSetCaller(address common.Address, caller bind.ContractCaller) (*AddrSetCaller, error) {
	contract, err := bindAddrSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddrSetCaller{contract: contract}, nil
}

// NewAddrSetTransactor creates a new write-only instance of AddrSet, bound to a specific deployed contract.
func NewAddrSetTransactor(address common.Address, transactor bind.ContractTransactor) (*AddrSetTransactor, error) {
	contract, err := bindAddrSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddrSetTransactor{contract: contract}, nil
}

// NewAddrSetFilterer creates a new log filterer instance of AddrSet, bound to a specific deployed contract.
func NewAddrSetFilterer(address common.Address, filterer bind.ContractFilterer) (*AddrSetFilterer, error) {
	contract, err := bindAddrSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddrSetFilterer{contract: contract}, nil
}

// bindAddrSet binds a generic wrapper to an already deployed contract.
func bindAddrSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddrSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddrSet *AddrSetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddrSet.Contract.AddrSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddrSet *AddrSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddrSet.Contract.AddrSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddrSet *AddrSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddrSet.Contract.AddrSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddrSet *AddrSetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddrSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddrSet *AddrSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddrSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddrSet *AddrSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddrSet.Contract.contract.Transact(opts, method, params...)
}

// KYCABI is the input ABI used to generate the binding from.
const KYCABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"approveAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"registerProvider\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"kycStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeProvider\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"suspendAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ProviderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ProviderRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"by\",\"type\":\"address\"}],\"name\":\"AddrApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"by\",\"type\":\"address\"}],\"name\":\"AddrSuspended\",\"type\":\"event\"}]"

// KYCBin is the compiled bytecode used for deploying new contracts.
const KYCBin = `0x606060405260008054600160a060020a033316600160a060020a03199091161790556105d4806100306000396000f30060606040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663055273c981146100925780630e260016146100b357806330ccebb5146100d257806346cc599e146101155780638a355a57146101345780638da5cb5b14610153578063996a6f8214610182578063f2fde38b146101a1575b600080fd5b341561009d57600080fd5b6100b1600160a060020a03600435166101c0565b005b34156100be57600080fd5b6100b1600160a060020a0360043516610299565b34156100dd57600080fd5b6100f1600160a060020a036004351661030a565b6040518082600281111561010157fe5b60ff16815260200191505060405180910390f35b341561012057600080fd5b6100f1600160a060020a0360043516610328565b341561013f57600080fd5b6100b1600160a060020a036004351661033d565b341561015e57600080fd5b6101666103ae565b604051600160a060020a03909116815260200160405180910390f35b341561018d57600080fd5b6100b1600160a060020a03600435166103bd565b34156101ac57600080fd5b6100b1600160a060020a0360043516610498565b6000805433600160a060020a03908116911614806101e457506101e46001336104e2565b15156101ef57600080fd5b50600160a060020a03811660009081526002602052604090205460ff16600181600281111561021a57fe5b141561022557600080fd5b600160a060020a038216600090815260026020526040902080546001919060ff19168280021790555033600160a060020a03167fa3673b71ec0beba775defcf8c7ad027536fdbac996023d594b5efe0c4181acd083604051600160a060020a03909116815260200160405180910390a25050565b60005433600160a060020a039081169116146102b457600080fd5b6102bf600182610505565b15156102ca57600080fd5b7fae9c2c6481964847714ce58f65a7f6dcc41d0d8394449bacdf161b5920c4744a81604051600160a060020a03909116815260200160405180910390a150565b600160a060020a031660009081526002602052604090205460ff1690565b60026020526000908152604090205460ff1681565b60005433600160a060020a0390811691161461035857600080fd5b610363600182610558565b151561036e57600080fd5b7f1589f8555933761a3cff8aa925061be3b46e2dd43f621322ab611d300f62b1d981604051600160a060020a03909116815260200160405180910390a150565b600054600160a060020a031681565b6000805433600160a060020a03908116911614806103e157506103e16001336104e2565b15156103ec57600080fd5b50600160a060020a03811660009081526002602081905260409091205460ff169081600281111561041957fe5b141561042457600080fd5b600160a060020a0382166000908152600260208190526040909120805460ff1916600183021790555033600160a060020a03167fc17dc8dc1039ea489e8aa5d0bd8bfc32c1afb87b98959710df9646cd80c331cb83604051600160a060020a03909116815260200160405180910390a25050565b60005433600160a060020a039081169116146104b357600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600160a060020a03811660009081526020839052604090205460ff165b92915050565b600160a060020a03811660009081526020839052604081205460ff161561052e575060006104ff565b50600160a060020a0316600090815260209190915260409020805460ff1916600190811790915590565b600160a060020a03811660009081526020839052604081205460ff161515610582575060006104ff565b50600160a060020a0316600090815260209190915260409020805460ff191690556001905600a165627a7a72305820c2e85bc8935a34b12ab42f9235cb1a58d5068a58247ff6ec1d88738ffe1c79700029`

// DeployKYC deploys a new Ethereum contract, binding an instance of KYC to it.
func DeployKYC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KYC, error) {
	parsed, err := abi.JSON(strings.NewReader(KYCABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KYCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KYC{KYCCaller: KYCCaller{contract: contract}, KYCTransactor: KYCTransactor{contract: contract}, KYCFilterer: KYCFilterer{contract: contract}}, nil
}

// KYC is an auto generated Go binding around an Ethereum contract.
type KYC struct {
	KYCCaller     // Read-only binding to the contract
	KYCTransactor // Write-only binding to the contract
	KYCFilterer   // Log filterer for contract events
}

// KYCCaller is an auto generated read-only Go binding around an Ethereum contract.
type KYCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KYCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KYCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KYCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KYCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KYCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KYCSession struct {
	Contract     *KYC              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KYCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KYCCallerSession struct {
	Contract *KYCCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// KYCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KYCTransactorSession struct {
	Contract     *KYCTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KYCRaw is an auto generated low-level Go binding around an Ethereum contract.
type KYCRaw struct {
	Contract *KYC // Generic contract binding to access the raw methods on
}

// KYCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KYCCallerRaw struct {
	Contract *KYCCaller // Generic read-only contract binding to access the raw methods on
}

// KYCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KYCTransactorRaw struct {
	Contract *KYCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKYC creates a new instance of KYC, bound to a specific deployed contract.
func NewKYC(address common.Address, backend bind.ContractBackend) (*KYC, error) {
	contract, err := bindKYC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KYC{KYCCaller: KYCCaller{contract: contract}, KYCTransactor: KYCTransactor{contract: contract}, KYCFilterer: KYCFilterer{contract: contract}}, nil
}

// NewKYCCaller creates a new read-only instance of KYC, bound to a specific deployed contract.
func NewKYCCaller(address common.Address, caller bind.ContractCaller) (*KYCCaller, error) {
	contract, err := bindKYC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KYCCaller{contract: contract}, nil
}

// NewKYCTransactor creates a new write-only instance of KYC, bound to a specific deployed contract.
func NewKYCTransactor(address common.Address, transactor bind.ContractTransactor) (*KYCTransactor, error) {
	contract, err := bindKYC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KYCTransactor{contract: contract}, nil
}

// NewKYCFilterer creates a new log filterer instance of KYC, bound to a specific deployed contract.
func NewKYCFilterer(address common.Address, filterer bind.ContractFilterer) (*KYCFilterer, error) {
	contract, err := bindKYC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KYCFilterer{contract: contract}, nil
}

// bindKYC binds a generic wrapper to an already deployed contract.
func bindKYC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KYCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KYC *KYCRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KYC.Contract.KYCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KYC *KYCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KYC.Contract.KYCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KYC *KYCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KYC.Contract.KYCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KYC *KYCCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KYC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KYC *KYCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KYC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KYC *KYCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KYC.Contract.contract.Transact(opts, method, params...)
}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(addr address) constant returns(uint8)
func (_KYC *KYCCaller) GetStatus(opts *bind.CallOpts, addr common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _KYC.contract.Call(opts, out, "getStatus", addr)
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(addr address) constant returns(uint8)
func (_KYC *KYCSession) GetStatus(addr common.Address) (uint8, error) {
	return _KYC.Contract.GetStatus(&_KYC.CallOpts, addr)
}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(addr address) constant returns(uint8)
func (_KYC *KYCCallerSession) GetStatus(addr common.Address) (uint8, error) {
	return _KYC.Contract.GetStatus(&_KYC.CallOpts, addr)
}

// KycStatus is a free data retrieval call binding the contract method 0x46cc599e.
//
// Solidity: function kycStatus( address) constant returns(uint8)
func (_KYC *KYCCaller) KycStatus(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _KYC.contract.Call(opts, out, "kycStatus", arg0)
	return *ret0, err
}

// KycStatus is a free data retrieval call binding the contract method 0x46cc599e.
//
// Solidity: function kycStatus( address) constant returns(uint8)
func (_KYC *KYCSession) KycStatus(arg0 common.Address) (uint8, error) {
	return _KYC.Contract.KycStatus(&_KYC.CallOpts, arg0)
}

// KycStatus is a free data retrieval call binding the contract method 0x46cc599e.
//
// Solidity: function kycStatus( address) constant returns(uint8)
func (_KYC *KYCCallerSession) KycStatus(arg0 common.Address) (uint8, error) {
	return _KYC.Contract.KycStatus(&_KYC.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_KYC *KYCCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KYC.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_KYC *KYCSession) Owner() (common.Address, error) {
	return _KYC.Contract.Owner(&_KYC.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_KYC *KYCCallerSession) Owner() (common.Address, error) {
	return _KYC.Contract.Owner(&_KYC.CallOpts)
}

// ApproveAddr is a paid mutator transaction binding the contract method 0x055273c9.
//
// Solidity: function approveAddr(addr address) returns()
func (_KYC *KYCTransactor) ApproveAddr(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _KYC.contract.Transact(opts, "approveAddr", addr)
}

// ApproveAddr is a paid mutator transaction binding the contract method 0x055273c9.
//
// Solidity: function approveAddr(addr address) returns()
func (_KYC *KYCSession) ApproveAddr(addr common.Address) (*types.Transaction, error) {
	return _KYC.Contract.ApproveAddr(&_KYC.TransactOpts, addr)
}

// ApproveAddr is a paid mutator transaction binding the contract method 0x055273c9.
//
// Solidity: function approveAddr(addr address) returns()
func (_KYC *KYCTransactorSession) ApproveAddr(addr common.Address) (*types.Transaction, error) {
	return _KYC.Contract.ApproveAddr(&_KYC.TransactOpts, addr)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0x0e260016.
//
// Solidity: function registerProvider(addr address) returns()
func (_KYC *KYCTransactor) RegisterProvider(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _KYC.contract.Transact(opts, "registerProvider", addr)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0x0e260016.
//
// Solidity: function registerProvider(addr address) returns()
func (_KYC *KYCSession) RegisterProvider(addr common.Address) (*types.Transaction, error) {
	return _KYC.Contract.RegisterProvider(&_KYC.TransactOpts, addr)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0x0e260016.
//
// Solidity: function registerProvider(addr address) returns()
func (_KYC *KYCTransactorSession) RegisterProvider(addr common.Address) (*types.Transaction, error) {
	return _KYC.Contract.RegisterProvider(&_KYC.TransactOpts, addr)
}

// RemoveProvider is a paid mutator transaction binding the contract method 0x8a355a57.
//
// Solidity: function removeProvider(addr address) returns()
func (_KYC *KYCTransactor) RemoveProvider(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _KYC.contract.Transact(opts, "removeProvider", addr)
}

// RemoveProvider is a paid mutator transaction binding the contract method 0x8a355a57.
//
// Solidity: function removeProvider(addr address) returns()
func (_KYC *KYCSession) RemoveProvider(addr common.Address) (*types.Transaction, error) {
	return _KYC.Contract.RemoveProvider(&_KYC.TransactOpts, addr)
}

// RemoveProvider is a paid mutator transaction binding the contract method 0x8a355a57.
//
// Solidity: function removeProvider(addr address) returns()
func (_KYC *KYCTransactorSession) RemoveProvider(addr common.Address) (*types.Transaction, error) {
	return _KYC.Contract.RemoveProvider(&_KYC.TransactOpts, addr)
}

// SuspendAddr is a paid mutator transaction binding the contract method 0x996a6f82.
//
// Solidity: function suspendAddr(addr address) returns()
func (_KYC *KYCTransactor) SuspendAddr(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _KYC.contract.Transact(opts, "suspendAddr", addr)
}

// SuspendAddr is a paid mutator transaction binding the contract method 0x996a6f82.
//
// Solidity: function suspendAddr(addr address) returns()
func (_KYC *KYCSession) SuspendAddr(addr common.Address) (*types.Transaction, error) {
	return _KYC.Contract.SuspendAddr(&_KYC.TransactOpts, addr)
}

// SuspendAddr is a paid mutator transaction binding the contract method 0x996a6f82.
//
// Solidity: function suspendAddr(addr address) returns()
func (_KYC *KYCTransactorSession) SuspendAddr(addr common.Address) (*types.Transaction, error) {
	return _KYC.Contract.SuspendAddr(&_KYC.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_KYC *KYCTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KYC.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_KYC *KYCSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KYC.Contract.TransferOwnership(&_KYC.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_KYC *KYCTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KYC.Contract.TransferOwnership(&_KYC.TransactOpts, newOwner)
}

// KYCAddrApprovedIterator is returned from FilterAddrApproved and is used to iterate over the raw logs and unpacked data for AddrApproved events raised by the KYC contract.
type KYCAddrApprovedIterator struct {
	Event *KYCAddrApproved // Event containing the contract specifics and raw log

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
func (it *KYCAddrApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCAddrApproved)
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
		it.Event = new(KYCAddrApproved)
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
func (it *KYCAddrApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCAddrApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCAddrApproved represents a AddrApproved event raised by the KYC contract.
type KYCAddrApproved struct {
	Addr common.Address
	By   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrApproved is a free log retrieval operation binding the contract event 0xa3673b71ec0beba775defcf8c7ad027536fdbac996023d594b5efe0c4181acd0.
//
// Solidity: event AddrApproved(addr address, by indexed address)
func (_KYC *KYCFilterer) FilterAddrApproved(opts *bind.FilterOpts, by []common.Address) (*KYCAddrApprovedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _KYC.contract.FilterLogs(opts, "AddrApproved", byRule)
	if err != nil {
		return nil, err
	}
	return &KYCAddrApprovedIterator{contract: _KYC.contract, event: "AddrApproved", logs: logs, sub: sub}, nil
}

// WatchAddrApproved is a free log subscription operation binding the contract event 0xa3673b71ec0beba775defcf8c7ad027536fdbac996023d594b5efe0c4181acd0.
//
// Solidity: event AddrApproved(addr address, by indexed address)
func (_KYC *KYCFilterer) WatchAddrApproved(opts *bind.WatchOpts, sink chan<- *KYCAddrApproved, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _KYC.contract.WatchLogs(opts, "AddrApproved", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCAddrApproved)
				if err := _KYC.contract.UnpackLog(event, "AddrApproved", log); err != nil {
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

// KYCAddrSuspendedIterator is returned from FilterAddrSuspended and is used to iterate over the raw logs and unpacked data for AddrSuspended events raised by the KYC contract.
type KYCAddrSuspendedIterator struct {
	Event *KYCAddrSuspended // Event containing the contract specifics and raw log

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
func (it *KYCAddrSuspendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCAddrSuspended)
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
		it.Event = new(KYCAddrSuspended)
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
func (it *KYCAddrSuspendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCAddrSuspendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCAddrSuspended represents a AddrSuspended event raised by the KYC contract.
type KYCAddrSuspended struct {
	Addr common.Address
	By   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrSuspended is a free log retrieval operation binding the contract event 0xc17dc8dc1039ea489e8aa5d0bd8bfc32c1afb87b98959710df9646cd80c331cb.
//
// Solidity: event AddrSuspended(addr address, by indexed address)
func (_KYC *KYCFilterer) FilterAddrSuspended(opts *bind.FilterOpts, by []common.Address) (*KYCAddrSuspendedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _KYC.contract.FilterLogs(opts, "AddrSuspended", byRule)
	if err != nil {
		return nil, err
	}
	return &KYCAddrSuspendedIterator{contract: _KYC.contract, event: "AddrSuspended", logs: logs, sub: sub}, nil
}

// WatchAddrSuspended is a free log subscription operation binding the contract event 0xc17dc8dc1039ea489e8aa5d0bd8bfc32c1afb87b98959710df9646cd80c331cb.
//
// Solidity: event AddrSuspended(addr address, by indexed address)
func (_KYC *KYCFilterer) WatchAddrSuspended(opts *bind.WatchOpts, sink chan<- *KYCAddrSuspended, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _KYC.contract.WatchLogs(opts, "AddrSuspended", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCAddrSuspended)
				if err := _KYC.contract.UnpackLog(event, "AddrSuspended", log); err != nil {
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

// KYCProviderAddedIterator is returned from FilterProviderAdded and is used to iterate over the raw logs and unpacked data for ProviderAdded events raised by the KYC contract.
type KYCProviderAddedIterator struct {
	Event *KYCProviderAdded // Event containing the contract specifics and raw log

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
func (it *KYCProviderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCProviderAdded)
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
		it.Event = new(KYCProviderAdded)
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
func (it *KYCProviderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCProviderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCProviderAdded represents a ProviderAdded event raised by the KYC contract.
type KYCProviderAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterProviderAdded is a free log retrieval operation binding the contract event 0xae9c2c6481964847714ce58f65a7f6dcc41d0d8394449bacdf161b5920c4744a.
//
// Solidity: event ProviderAdded(addr address)
func (_KYC *KYCFilterer) FilterProviderAdded(opts *bind.FilterOpts) (*KYCProviderAddedIterator, error) {

	logs, sub, err := _KYC.contract.FilterLogs(opts, "ProviderAdded")
	if err != nil {
		return nil, err
	}
	return &KYCProviderAddedIterator{contract: _KYC.contract, event: "ProviderAdded", logs: logs, sub: sub}, nil
}

// WatchProviderAdded is a free log subscription operation binding the contract event 0xae9c2c6481964847714ce58f65a7f6dcc41d0d8394449bacdf161b5920c4744a.
//
// Solidity: event ProviderAdded(addr address)
func (_KYC *KYCFilterer) WatchProviderAdded(opts *bind.WatchOpts, sink chan<- *KYCProviderAdded) (event.Subscription, error) {

	logs, sub, err := _KYC.contract.WatchLogs(opts, "ProviderAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCProviderAdded)
				if err := _KYC.contract.UnpackLog(event, "ProviderAdded", log); err != nil {
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

// KYCProviderRemovedIterator is returned from FilterProviderRemoved and is used to iterate over the raw logs and unpacked data for ProviderRemoved events raised by the KYC contract.
type KYCProviderRemovedIterator struct {
	Event *KYCProviderRemoved // Event containing the contract specifics and raw log

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
func (it *KYCProviderRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KYCProviderRemoved)
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
		it.Event = new(KYCProviderRemoved)
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
func (it *KYCProviderRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KYCProviderRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KYCProviderRemoved represents a ProviderRemoved event raised by the KYC contract.
type KYCProviderRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterProviderRemoved is a free log retrieval operation binding the contract event 0x1589f8555933761a3cff8aa925061be3b46e2dd43f621322ab611d300f62b1d9.
//
// Solidity: event ProviderRemoved(addr address)
func (_KYC *KYCFilterer) FilterProviderRemoved(opts *bind.FilterOpts) (*KYCProviderRemovedIterator, error) {

	logs, sub, err := _KYC.contract.FilterLogs(opts, "ProviderRemoved")
	if err != nil {
		return nil, err
	}
	return &KYCProviderRemovedIterator{contract: _KYC.contract, event: "ProviderRemoved", logs: logs, sub: sub}, nil
}

// WatchProviderRemoved is a free log subscription operation binding the contract event 0x1589f8555933761a3cff8aa925061be3b46e2dd43f621322ab611d300f62b1d9.
//
// Solidity: event ProviderRemoved(addr address)
func (_KYC *KYCFilterer) WatchProviderRemoved(opts *bind.WatchOpts, sink chan<- *KYCProviderRemoved) (event.Subscription, error) {

	logs, sub, err := _KYC.contract.WatchLogs(opts, "ProviderRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KYCProviderRemoved)
				if err := _KYC.contract.UnpackLog(event, "ProviderRemoved", log); err != nil {
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

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OwnedBin is the compiled bytecode used for deploying new contracts.
const OwnedBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556101668061003b6000396000f30060606040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063f2fde38b1461008c575b600080fd5b341561005b57600080fd5b6100636100ba565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561009757600080fd5b6100b873ffffffffffffffffffffffffffffffffffffffff600435166100d6565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146100fe57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff929092169190911790555600a165627a7a7230582092b2d38ce301b18c6013549e3567413e150d2de2a6569d02724196fec221fc3d0029`

// DeployOwned deploys a new Ethereum contract, binding an instance of Owned to it.
func DeployOwned(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Owned, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// Owned is an auto generated Go binding around an Ethereum contract.
type Owned struct {
	OwnedCaller     // Read-only binding to the contract
	OwnedTransactor // Write-only binding to the contract
	OwnedFilterer   // Log filterer for contract events
}

// OwnedCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnedSession struct {
	Contract     *Owned            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnedCallerSession struct {
	Contract *OwnedCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnedTransactorSession struct {
	Contract     *OwnedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnedRaw struct {
	Contract *Owned // Generic contract binding to access the raw methods on
}

// OwnedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnedCallerRaw struct {
	Contract *OwnedCaller // Generic read-only contract binding to access the raw methods on
}

// OwnedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnedTransactorRaw struct {
	Contract *OwnedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwned creates a new instance of Owned, bound to a specific deployed contract.
func NewOwned(address common.Address, backend bind.ContractBackend) (*Owned, error) {
	contract, err := bindOwned(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// NewOwnedCaller creates a new read-only instance of Owned, bound to a specific deployed contract.
func NewOwnedCaller(address common.Address, caller bind.ContractCaller) (*OwnedCaller, error) {
	contract, err := bindOwned(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedCaller{contract: contract}, nil
}

// NewOwnedTransactor creates a new write-only instance of Owned, bound to a specific deployed contract.
func NewOwnedTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnedTransactor, error) {
	contract, err := bindOwned(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedTransactor{contract: contract}, nil
}

// NewOwnedFilterer creates a new log filterer instance of Owned, bound to a specific deployed contract.
func NewOwnedFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnedFilterer, error) {
	contract, err := bindOwned(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnedFilterer{contract: contract}, nil
}

// bindOwned binds a generic wrapper to an already deployed contract.
func bindOwned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.OwnedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Owned.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCallerSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Owned *OwnedTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Owned *OwnedSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.TransferOwnership(&_Owned.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Owned *OwnedTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.TransferOwnership(&_Owned.TransactOpts, newOwner)
}
