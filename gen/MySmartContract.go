// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MySmartContract

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

// MySmartContractMetaData contains all meta data concerning the MySmartContract contract.
var MySmartContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Counter\",\"type\":\"uint256\"}],\"name\":\"Event\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"eventCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_event\",\"type\":\"uint256\"}],\"name\":\"produceEvents\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610186806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806352c55db91461003b578063bd72bea414610056575b600080fd5b61004460005481565b60405190815260200160405180910390f35b6100696100643660046100f0565b61006b565b005b6000805461007a90600161011f565b90505b8160005461008b919061011f565b81116100d6576040518181527f510e730eb6600b4c67d51768c6996795863364461fee983d92d5e461f209c7cf9060200160405180910390a1806100ce81610137565b91505061007d565b50806000808282546100e8919061011f565b909155505050565b60006020828403121561010257600080fd5b5035919050565b634e487b7160e01b600052601160045260246000fd5b6000821982111561013257610132610109565b500190565b60006001820161014957610149610109565b506001019056fea264697066735822122023f818d4057aac95a2d617ceda31529109c6a1f983957dbde173d32debb82c1664736f6c634300080d0033",
}

// MySmartContractABI is the input ABI used to generate the binding from.
// Deprecated: Use MySmartContractMetaData.ABI instead.
var MySmartContractABI = MySmartContractMetaData.ABI

// MySmartContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MySmartContractMetaData.Bin instead.
var MySmartContractBin = MySmartContractMetaData.Bin

// DeployMySmartContract deploys a new Ethereum contract, binding an instance of MySmartContract to it.
func DeployMySmartContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MySmartContract, error) {
	parsed, err := MySmartContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MySmartContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MySmartContract{MySmartContractCaller: MySmartContractCaller{contract: contract}, MySmartContractTransactor: MySmartContractTransactor{contract: contract}, MySmartContractFilterer: MySmartContractFilterer{contract: contract}}, nil
}

// MySmartContract is an auto generated Go binding around an Ethereum contract.
type MySmartContract struct {
	MySmartContractCaller     // Read-only binding to the contract
	MySmartContractTransactor // Write-only binding to the contract
	MySmartContractFilterer   // Log filterer for contract events
}

// MySmartContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MySmartContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MySmartContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MySmartContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MySmartContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MySmartContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MySmartContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MySmartContractSession struct {
	Contract     *MySmartContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MySmartContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MySmartContractCallerSession struct {
	Contract *MySmartContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MySmartContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MySmartContractTransactorSession struct {
	Contract     *MySmartContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MySmartContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MySmartContractRaw struct {
	Contract *MySmartContract // Generic contract binding to access the raw methods on
}

// MySmartContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MySmartContractCallerRaw struct {
	Contract *MySmartContractCaller // Generic read-only contract binding to access the raw methods on
}

// MySmartContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MySmartContractTransactorRaw struct {
	Contract *MySmartContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMySmartContract creates a new instance of MySmartContract, bound to a specific deployed contract.
func NewMySmartContract(address common.Address, backend bind.ContractBackend) (*MySmartContract, error) {
	contract, err := bindMySmartContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MySmartContract{MySmartContractCaller: MySmartContractCaller{contract: contract}, MySmartContractTransactor: MySmartContractTransactor{contract: contract}, MySmartContractFilterer: MySmartContractFilterer{contract: contract}}, nil
}

// NewMySmartContractCaller creates a new read-only instance of MySmartContract, bound to a specific deployed contract.
func NewMySmartContractCaller(address common.Address, caller bind.ContractCaller) (*MySmartContractCaller, error) {
	contract, err := bindMySmartContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MySmartContractCaller{contract: contract}, nil
}

// NewMySmartContractTransactor creates a new write-only instance of MySmartContract, bound to a specific deployed contract.
func NewMySmartContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MySmartContractTransactor, error) {
	contract, err := bindMySmartContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MySmartContractTransactor{contract: contract}, nil
}

// NewMySmartContractFilterer creates a new log filterer instance of MySmartContract, bound to a specific deployed contract.
func NewMySmartContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MySmartContractFilterer, error) {
	contract, err := bindMySmartContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MySmartContractFilterer{contract: contract}, nil
}

// bindMySmartContract binds a generic wrapper to an already deployed contract.
func bindMySmartContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MySmartContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MySmartContract *MySmartContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MySmartContract.Contract.MySmartContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MySmartContract *MySmartContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MySmartContract.Contract.MySmartContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MySmartContract *MySmartContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MySmartContract.Contract.MySmartContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MySmartContract *MySmartContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MySmartContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MySmartContract *MySmartContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MySmartContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MySmartContract *MySmartContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MySmartContract.Contract.contract.Transact(opts, method, params...)
}

// EventCounter is a free data retrieval call binding the contract method 0x52c55db9.
//
// Solidity: function eventCounter() view returns(uint256)
func (_MySmartContract *MySmartContractCaller) EventCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MySmartContract.contract.Call(opts, &out, "eventCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EventCounter is a free data retrieval call binding the contract method 0x52c55db9.
//
// Solidity: function eventCounter() view returns(uint256)
func (_MySmartContract *MySmartContractSession) EventCounter() (*big.Int, error) {
	return _MySmartContract.Contract.EventCounter(&_MySmartContract.CallOpts)
}

// EventCounter is a free data retrieval call binding the contract method 0x52c55db9.
//
// Solidity: function eventCounter() view returns(uint256)
func (_MySmartContract *MySmartContractCallerSession) EventCounter() (*big.Int, error) {
	return _MySmartContract.Contract.EventCounter(&_MySmartContract.CallOpts)
}

// ProduceEvents is a paid mutator transaction binding the contract method 0xbd72bea4.
//
// Solidity: function produceEvents(uint256 _event) returns()
func (_MySmartContract *MySmartContractTransactor) ProduceEvents(opts *bind.TransactOpts, _event *big.Int) (*types.Transaction, error) {
	return _MySmartContract.contract.Transact(opts, "produceEvents", _event)
}

// ProduceEvents is a paid mutator transaction binding the contract method 0xbd72bea4.
//
// Solidity: function produceEvents(uint256 _event) returns()
func (_MySmartContract *MySmartContractSession) ProduceEvents(_event *big.Int) (*types.Transaction, error) {
	return _MySmartContract.Contract.ProduceEvents(&_MySmartContract.TransactOpts, _event)
}

// ProduceEvents is a paid mutator transaction binding the contract method 0xbd72bea4.
//
// Solidity: function produceEvents(uint256 _event) returns()
func (_MySmartContract *MySmartContractTransactorSession) ProduceEvents(_event *big.Int) (*types.Transaction, error) {
	return _MySmartContract.Contract.ProduceEvents(&_MySmartContract.TransactOpts, _event)
}

// MySmartContractEventIterator is returned from FilterEvent and is used to iterate over the raw logs and unpacked data for Event events raised by the MySmartContract contract.
type MySmartContractEventIterator struct {
	Event *MySmartContractEvent // Event containing the contract specifics and raw log

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
func (it *MySmartContractEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MySmartContractEvent)
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
		it.Event = new(MySmartContractEvent)
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
func (it *MySmartContractEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MySmartContractEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MySmartContractEvent represents a Event event raised by the MySmartContract contract.
type MySmartContractEvent struct {
	Counter *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEvent is a free log retrieval operation binding the contract event 0x510e730eb6600b4c67d51768c6996795863364461fee983d92d5e461f209c7cf.
//
// Solidity: event Event(uint256 Counter)
func (_MySmartContract *MySmartContractFilterer) FilterEvent(opts *bind.FilterOpts) (*MySmartContractEventIterator, error) {

	logs, sub, err := _MySmartContract.contract.FilterLogs(opts, "Event")
	if err != nil {
		return nil, err
	}
	return &MySmartContractEventIterator{contract: _MySmartContract.contract, event: "Event", logs: logs, sub: sub}, nil
}

// WatchEvent is a free log subscription operation binding the contract event 0x510e730eb6600b4c67d51768c6996795863364461fee983d92d5e461f209c7cf.
//
// Solidity: event Event(uint256 Counter)
func (_MySmartContract *MySmartContractFilterer) WatchEvent(opts *bind.WatchOpts, sink chan<- *MySmartContractEvent) (event.Subscription, error) {

	logs, sub, err := _MySmartContract.contract.WatchLogs(opts, "Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MySmartContractEvent)
				if err := _MySmartContract.contract.UnpackLog(event, "Event", log); err != nil {
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

// ParseEvent is a log parse operation binding the contract event 0x510e730eb6600b4c67d51768c6996795863364461fee983d92d5e461f209c7cf.
//
// Solidity: event Event(uint256 Counter)
func (_MySmartContract *MySmartContractFilterer) ParseEvent(log types.Log) (*MySmartContractEvent, error) {
	event := new(MySmartContractEvent)
	if err := _MySmartContract.contract.UnpackLog(event, "Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
