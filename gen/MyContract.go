// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MyContract

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

// MyContractMetaData contains all meta data concerning the MyContract contract.
var MyContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Counter\",\"type\":\"uint256\"}],\"name\":\"Event\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"eventCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_event\",\"type\":\"uint256\"}],\"name\":\"produceEvents\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610186806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806352c55db91461003b578063bd72bea414610056575b600080fd5b61004460005481565b60405190815260200160405180910390f35b6100696100643660046100f0565b61006b565b005b6000805461007a90600161011f565b90505b8160005461008b919061011f565b81116100d6576040518181527f510e730eb6600b4c67d51768c6996795863364461fee983d92d5e461f209c7cf9060200160405180910390a1806100ce81610137565b91505061007d565b50806000808282546100e8919061011f565b909155505050565b60006020828403121561010257600080fd5b5035919050565b634e487b7160e01b600052601160045260246000fd5b6000821982111561013257610132610109565b500190565b60006001820161014957610149610109565b506001019056fea2646970667358221220739ad4ff70da63d3f7b7d793ac1ce9665d5200ebad0329fd99610a949b09866e64736f6c634300080d0033",
}

// MyContractABI is the input ABI used to generate the binding from.
// Deprecated: Use MyContractMetaData.ABI instead.
var MyContractABI = MyContractMetaData.ABI

// MyContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyContractMetaData.Bin instead.
var MyContractBin = MyContractMetaData.Bin

// DeployMyContract deploys a new Ethereum contract, binding an instance of MyContract to it.
func DeployMyContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MyContract, error) {
	parsed, err := MyContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyContract{MyContractCaller: MyContractCaller{contract: contract}, MyContractTransactor: MyContractTransactor{contract: contract}, MyContractFilterer: MyContractFilterer{contract: contract}}, nil
}

// MyContract is an auto generated Go binding around an Ethereum contract.
type MyContract struct {
	MyContractCaller     // Read-only binding to the contract
	MyContractTransactor // Write-only binding to the contract
	MyContractFilterer   // Log filterer for contract events
}

// MyContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyContractSession struct {
	Contract     *MyContract       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyContractCallerSession struct {
	Contract *MyContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MyContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyContractTransactorSession struct {
	Contract     *MyContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MyContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyContractRaw struct {
	Contract *MyContract // Generic contract binding to access the raw methods on
}

// MyContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyContractCallerRaw struct {
	Contract *MyContractCaller // Generic read-only contract binding to access the raw methods on
}

// MyContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyContractTransactorRaw struct {
	Contract *MyContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyContract creates a new instance of MyContract, bound to a specific deployed contract.
func NewMyContract(address common.Address, backend bind.ContractBackend) (*MyContract, error) {
	contract, err := bindMyContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyContract{MyContractCaller: MyContractCaller{contract: contract}, MyContractTransactor: MyContractTransactor{contract: contract}, MyContractFilterer: MyContractFilterer{contract: contract}}, nil
}

// NewMyContractCaller creates a new read-only instance of MyContract, bound to a specific deployed contract.
func NewMyContractCaller(address common.Address, caller bind.ContractCaller) (*MyContractCaller, error) {
	contract, err := bindMyContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyContractCaller{contract: contract}, nil
}

// NewMyContractTransactor creates a new write-only instance of MyContract, bound to a specific deployed contract.
func NewMyContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MyContractTransactor, error) {
	contract, err := bindMyContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyContractTransactor{contract: contract}, nil
}

// NewMyContractFilterer creates a new log filterer instance of MyContract, bound to a specific deployed contract.
func NewMyContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MyContractFilterer, error) {
	contract, err := bindMyContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyContractFilterer{contract: contract}, nil
}

// bindMyContract binds a generic wrapper to an already deployed contract.
func bindMyContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyContract *MyContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyContract.Contract.MyContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyContract *MyContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.Contract.MyContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyContract *MyContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyContract.Contract.MyContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyContract *MyContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyContract *MyContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyContract *MyContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyContract.Contract.contract.Transact(opts, method, params...)
}

// EventCounter is a free data retrieval call binding the contract method 0x52c55db9.
//
// Solidity: function eventCounter() view returns(uint256)
func (_MyContract *MyContractCaller) EventCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyContract.contract.Call(opts, &out, "eventCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EventCounter is a free data retrieval call binding the contract method 0x52c55db9.
//
// Solidity: function eventCounter() view returns(uint256)
func (_MyContract *MyContractSession) EventCounter() (*big.Int, error) {
	return _MyContract.Contract.EventCounter(&_MyContract.CallOpts)
}

// EventCounter is a free data retrieval call binding the contract method 0x52c55db9.
//
// Solidity: function eventCounter() view returns(uint256)
func (_MyContract *MyContractCallerSession) EventCounter() (*big.Int, error) {
	return _MyContract.Contract.EventCounter(&_MyContract.CallOpts)
}

// ProduceEvents is a paid mutator transaction binding the contract method 0xbd72bea4.
//
// Solidity: function produceEvents(uint256 _event) returns()
func (_MyContract *MyContractTransactor) ProduceEvents(opts *bind.TransactOpts, _event *big.Int) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "produceEvents", _event)
}

// ProduceEvents is a paid mutator transaction binding the contract method 0xbd72bea4.
//
// Solidity: function produceEvents(uint256 _event) returns()
func (_MyContract *MyContractSession) ProduceEvents(_event *big.Int) (*types.Transaction, error) {
	return _MyContract.Contract.ProduceEvents(&_MyContract.TransactOpts, _event)
}

// ProduceEvents is a paid mutator transaction binding the contract method 0xbd72bea4.
//
// Solidity: function produceEvents(uint256 _event) returns()
func (_MyContract *MyContractTransactorSession) ProduceEvents(_event *big.Int) (*types.Transaction, error) {
	return _MyContract.Contract.ProduceEvents(&_MyContract.TransactOpts, _event)
}

// MyContractEventIterator is returned from FilterEvent and is used to iterate over the raw logs and unpacked data for Event events raised by the MyContract contract.
type MyContractEventIterator struct {
	Event *MyContractEvent // Event containing the contract specifics and raw log

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
func (it *MyContractEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyContractEvent)
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
		it.Event = new(MyContractEvent)
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
func (it *MyContractEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyContractEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyContractEvent represents a Event event raised by the MyContract contract.
type MyContractEvent struct {
	Counter *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEvent is a free log retrieval operation binding the contract event 0x510e730eb6600b4c67d51768c6996795863364461fee983d92d5e461f209c7cf.
//
// Solidity: event Event(uint256 Counter)
func (_MyContract *MyContractFilterer) FilterEvent(opts *bind.FilterOpts) (*MyContractEventIterator, error) {

	logs, sub, err := _MyContract.contract.FilterLogs(opts, "Event")
	if err != nil {
		return nil, err
	}
	return &MyContractEventIterator{contract: _MyContract.contract, event: "Event", logs: logs, sub: sub}, nil
}

// WatchEvent is a free log subscription operation binding the contract event 0x510e730eb6600b4c67d51768c6996795863364461fee983d92d5e461f209c7cf.
//
// Solidity: event Event(uint256 Counter)
func (_MyContract *MyContractFilterer) WatchEvent(opts *bind.WatchOpts, sink chan<- *MyContractEvent) (event.Subscription, error) {

	logs, sub, err := _MyContract.contract.WatchLogs(opts, "Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyContractEvent)
				if err := _MyContract.contract.UnpackLog(event, "Event", log); err != nil {
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
func (_MyContract *MyContractFilterer) ParseEvent(log types.Log) (*MyContractEvent, error) {
	event := new(MyContractEvent)
	if err := _MyContract.contract.UnpackLog(event, "Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
