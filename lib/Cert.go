// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lib

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
	_ = abi.ConvertType
)

// CertMetaData contains all meta data concerning the Cert contract.
var CertMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"course\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"grade\",\"type\":\"string\"}],\"name\":\"Issued\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"Certificates\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"course\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"grade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_course\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_grade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610c0c806100606000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80630dfb1e861461003b5780639d3becb214610057575b600080fd5b610055600480360381019061005091906105c8565b61008a565b005b610071600480360381019061006c91906106cf565b610208565b6040516100819493929190610797565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610118576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161010f90610844565b60405180910390fd5b60405180608001604052808581526020018481526020018381526020018281525060018660405161014991906108a0565b9081526020016040518091039020600082015181600001908161016c9190610acd565b5060208201518160010190816101829190610acd565b5060408201518160020190816101989190610acd565b5060608201518160030190816101ae9190610acd565b50905050826040516101c091906108a0565b60405180910390207ff34394c565949d2aff2cc92a2ca9216303b89a22f9953a2e86034ba2f7ee3cdd86846040516101f9929190610b9f565b60405180910390a25050505050565b600181805160208101820180518482526020830160208501208183528095505050505050600091509050806000018054610241906108e6565b80601f016020809104026020016040519081016040528092919081815260200182805461026d906108e6565b80156102ba5780601f1061028f576101008083540402835291602001916102ba565b820191906000526020600020905b81548152906001019060200180831161029d57829003601f168201915b5050505050908060010180546102cf906108e6565b80601f01602080910402602001604051908101604052809291908181526020018280546102fb906108e6565b80156103485780601f1061031d57610100808354040283529160200191610348565b820191906000526020600020905b81548152906001019060200180831161032b57829003601f168201915b50505050509080600201805461035d906108e6565b80601f0160208091040260200160405190810160405280929190818152602001828054610389906108e6565b80156103d65780601f106103ab576101008083540402835291602001916103d6565b820191906000526020600020905b8154815290600101906020018083116103b957829003601f168201915b5050505050908060030180546103eb906108e6565b80601f0160208091040260200160405190810160405280929190818152602001828054610417906108e6565b80156104645780601f1061043957610100808354040283529160200191610464565b820191906000526020600020905b81548152906001019060200180831161044757829003601f168201915b5050505050905084565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6104d58261048c565b810181811067ffffffffffffffff821117156104f4576104f361049d565b5b80604052505050565b600061050761046e565b905061051382826104cc565b919050565b600067ffffffffffffffff8211156105335761053261049d565b5b61053c8261048c565b9050602081019050919050565b82818337600083830152505050565b600061056b61056684610518565b6104fd565b90508281526020810184848401111561058757610586610487565b5b610592848285610549565b509392505050565b600082601f8301126105af576105ae610482565b5b81356105bf848260208601610558565b91505092915050565b600080600080600060a086880312156105e4576105e3610478565b5b600086013567ffffffffffffffff8111156106025761060161047d565b5b61060e8882890161059a565b955050602086013567ffffffffffffffff81111561062f5761062e61047d565b5b61063b8882890161059a565b945050604086013567ffffffffffffffff81111561065c5761065b61047d565b5b6106688882890161059a565b935050606086013567ffffffffffffffff8111156106895761068861047d565b5b6106958882890161059a565b925050608086013567ffffffffffffffff8111156106b6576106b561047d565b5b6106c28882890161059a565b9150509295509295909350565b6000602082840312156106e5576106e4610478565b5b600082013567ffffffffffffffff8111156107035761070261047d565b5b61070f8482850161059a565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610752578082015181840152602081019050610737565b60008484015250505050565b600061076982610718565b6107738185610723565b9350610783818560208601610734565b61078c8161048c565b840191505092915050565b600060808201905081810360008301526107b1818761075e565b905081810360208301526107c5818661075e565b905081810360408301526107d9818561075e565b905081810360608301526107ed818461075e565b905095945050505050565b7f4163636573732044656e69656400000000000000000000000000000000000000600082015250565b600061082e600d83610723565b9150610839826107f8565b602082019050919050565b6000602082019050818103600083015261085d81610821565b9050919050565b600081905092915050565b600061087a82610718565b6108848185610864565b9350610894818560208601610734565b80840191505092915050565b60006108ac828461086f565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806108fe57607f821691505b602082108103610911576109106108b7565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026109797fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261093c565b610983868361093c565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006109ca6109c56109c08461099b565b6109a5565b61099b565b9050919050565b6000819050919050565b6109e4836109af565b6109f86109f0826109d1565b848454610949565b825550505050565b600090565b610a0d610a00565b610a188184846109db565b505050565b5b81811015610a3c57610a31600082610a05565b600181019050610a1e565b5050565b601f821115610a8157610a5281610917565b610a5b8461092c565b81016020851015610a6a578190505b610a7e610a768561092c565b830182610a1d565b50505b505050565b600082821c905092915050565b6000610aa460001984600802610a86565b1980831691505092915050565b6000610abd8383610a93565b9150826002028217905092915050565b610ad682610718565b67ffffffffffffffff811115610aef57610aee61049d565b5b610af982546108e6565b610b04828285610a40565b600060209050601f831160018114610b375760008415610b25578287015190505b610b2f8582610ab1565b865550610b97565b601f198416610b4586610917565b60005b82811015610b6d57848901518255600182019150602085019450602081019050610b48565b86831015610b8a5784890151610b86601f891682610a93565b8355505b6001600288020188555050505b505050505050565b60006040820190508181036000830152610bb9818561075e565b90508181036020830152610bcd818461075e565b9050939250505056fea26469706673582212200a0fd118df64eaad3b27d1109b4070a9265f53e95e74916d43bd9c06faf1d26a64736f6c63430008120033",
}

// CertABI is the input ABI used to generate the binding from.
// Deprecated: Use CertMetaData.ABI instead.
var CertABI = CertMetaData.ABI

// CertBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CertMetaData.Bin instead.
var CertBin = CertMetaData.Bin

// DeployCert deploys a new Ethereum contract, binding an instance of Cert to it.
func DeployCert(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Cert, error) {
	parsed, err := CertMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CertBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cert{CertCaller: CertCaller{contract: contract}, CertTransactor: CertTransactor{contract: contract}, CertFilterer: CertFilterer{contract: contract}}, nil
}

// Cert is an auto generated Go binding around an Ethereum contract.
type Cert struct {
	CertCaller     // Read-only binding to the contract
	CertTransactor // Write-only binding to the contract
	CertFilterer   // Log filterer for contract events
}

// CertCaller is an auto generated read-only Go binding around an Ethereum contract.
type CertCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CertTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CertFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CertSession struct {
	Contract     *Cert             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CertCallerSession struct {
	Contract *CertCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CertTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CertTransactorSession struct {
	Contract     *CertTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertRaw is an auto generated low-level Go binding around an Ethereum contract.
type CertRaw struct {
	Contract *Cert // Generic contract binding to access the raw methods on
}

// CertCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CertCallerRaw struct {
	Contract *CertCaller // Generic read-only contract binding to access the raw methods on
}

// CertTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CertTransactorRaw struct {
	Contract *CertTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCert creates a new instance of Cert, bound to a specific deployed contract.
func NewCert(address common.Address, backend bind.ContractBackend) (*Cert, error) {
	contract, err := bindCert(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cert{CertCaller: CertCaller{contract: contract}, CertTransactor: CertTransactor{contract: contract}, CertFilterer: CertFilterer{contract: contract}}, nil
}

// NewCertCaller creates a new read-only instance of Cert, bound to a specific deployed contract.
func NewCertCaller(address common.Address, caller bind.ContractCaller) (*CertCaller, error) {
	contract, err := bindCert(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertCaller{contract: contract}, nil
}

// NewCertTransactor creates a new write-only instance of Cert, bound to a specific deployed contract.
func NewCertTransactor(address common.Address, transactor bind.ContractTransactor) (*CertTransactor, error) {
	contract, err := bindCert(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertTransactor{contract: contract}, nil
}

// NewCertFilterer creates a new log filterer instance of Cert, bound to a specific deployed contract.
func NewCertFilterer(address common.Address, filterer bind.ContractFilterer) (*CertFilterer, error) {
	contract, err := bindCert(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertFilterer{contract: contract}, nil
}

// bindCert binds a generic wrapper to an already deployed contract.
func bindCert(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CertMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cert *CertRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cert.Contract.CertCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cert *CertRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cert.Contract.CertTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cert *CertRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cert.Contract.CertTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cert *CertCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cert.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cert *CertTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cert.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cert *CertTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cert.Contract.contract.Transact(opts, method, params...)
}

// Certificates is a free data retrieval call binding the contract method 0x9d3becb2.
//
// Solidity: function Certificates(string ) view returns(string name, string course, string grade, string date)
func (_Cert *CertCaller) Certificates(opts *bind.CallOpts, arg0 string) (struct {
	Name   string
	Course string
	Grade  string
	Date   string
}, error) {
	var out []interface{}
	err := _Cert.contract.Call(opts, &out, "Certificates", arg0)

	outstruct := new(struct {
		Name   string
		Course string
		Grade  string
		Date   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Course = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Grade = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Date = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// Certificates is a free data retrieval call binding the contract method 0x9d3becb2.
//
// Solidity: function Certificates(string ) view returns(string name, string course, string grade, string date)
func (_Cert *CertSession) Certificates(arg0 string) (struct {
	Name   string
	Course string
	Grade  string
	Date   string
}, error) {
	return _Cert.Contract.Certificates(&_Cert.CallOpts, arg0)
}

// Certificates is a free data retrieval call binding the contract method 0x9d3becb2.
//
// Solidity: function Certificates(string ) view returns(string name, string course, string grade, string date)
func (_Cert *CertCallerSession) Certificates(arg0 string) (struct {
	Name   string
	Course string
	Grade  string
	Date   string
}, error) {
	return _Cert.Contract.Certificates(&_Cert.CallOpts, arg0)
}

// Issue is a paid mutator transaction binding the contract method 0x0dfb1e86.
//
// Solidity: function issue(string _id, string _name, string _course, string _grade, string _date) returns()
func (_Cert *CertTransactor) Issue(opts *bind.TransactOpts, _id string, _name string, _course string, _grade string, _date string) (*types.Transaction, error) {
	return _Cert.contract.Transact(opts, "issue", _id, _name, _course, _grade, _date)
}

// Issue is a paid mutator transaction binding the contract method 0x0dfb1e86.
//
// Solidity: function issue(string _id, string _name, string _course, string _grade, string _date) returns()
func (_Cert *CertSession) Issue(_id string, _name string, _course string, _grade string, _date string) (*types.Transaction, error) {
	return _Cert.Contract.Issue(&_Cert.TransactOpts, _id, _name, _course, _grade, _date)
}

// Issue is a paid mutator transaction binding the contract method 0x0dfb1e86.
//
// Solidity: function issue(string _id, string _name, string _course, string _grade, string _date) returns()
func (_Cert *CertTransactorSession) Issue(_id string, _name string, _course string, _grade string, _date string) (*types.Transaction, error) {
	return _Cert.Contract.Issue(&_Cert.TransactOpts, _id, _name, _course, _grade, _date)
}

// CertIssuedIterator is returned from FilterIssued and is used to iterate over the raw logs and unpacked data for Issued events raised by the Cert contract.
type CertIssuedIterator struct {
	Event *CertIssued // Event containing the contract specifics and raw log

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
func (it *CertIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertIssued)
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
		it.Event = new(CertIssued)
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
func (it *CertIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertIssued represents a Issued event raised by the Cert contract.
type CertIssued struct {
	Course common.Hash
	Id     string
	Grade  string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssued is a free log retrieval operation binding the contract event 0xf34394c565949d2aff2cc92a2ca9216303b89a22f9953a2e86034ba2f7ee3cdd.
//
// Solidity: event Issued(string indexed course, string id, string grade)
func (_Cert *CertFilterer) FilterIssued(opts *bind.FilterOpts, course []string) (*CertIssuedIterator, error) {

	var courseRule []interface{}
	for _, courseItem := range course {
		courseRule = append(courseRule, courseItem)
	}

	logs, sub, err := _Cert.contract.FilterLogs(opts, "Issued", courseRule)
	if err != nil {
		return nil, err
	}
	return &CertIssuedIterator{contract: _Cert.contract, event: "Issued", logs: logs, sub: sub}, nil
}

// WatchIssued is a free log subscription operation binding the contract event 0xf34394c565949d2aff2cc92a2ca9216303b89a22f9953a2e86034ba2f7ee3cdd.
//
// Solidity: event Issued(string indexed course, string id, string grade)
func (_Cert *CertFilterer) WatchIssued(opts *bind.WatchOpts, sink chan<- *CertIssued, course []string) (event.Subscription, error) {

	var courseRule []interface{}
	for _, courseItem := range course {
		courseRule = append(courseRule, courseItem)
	}

	logs, sub, err := _Cert.contract.WatchLogs(opts, "Issued", courseRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertIssued)
				if err := _Cert.contract.UnpackLog(event, "Issued", log); err != nil {
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

// ParseIssued is a log parse operation binding the contract event 0xf34394c565949d2aff2cc92a2ca9216303b89a22f9953a2e86034ba2f7ee3cdd.
//
// Solidity: event Issued(string indexed course, string id, string grade)
func (_Cert *CertFilterer) ParseIssued(log types.Log) (*CertIssued, error) {
	event := new(CertIssued)
	if err := _Cert.contract.UnpackLog(event, "Issued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
