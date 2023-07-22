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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"course\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"grade\",\"type\":\"string\"}],\"name\":\"Issued\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Certificates\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"course\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"grade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_course\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_grade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610be7806100606000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80639622c8361461003b578063d1e7be261461006e575b600080fd5b61005560048036038101906100509190610497565b61008a565b6040516100659493929190610554565b60405180910390f35b610088600480360381019061008391906106ea565b6102da565b005b60016020528060005260406000206000915090508060000180546100ad90610804565b80601f01602080910402602001604051908101604052809291908181526020018280546100d990610804565b80156101265780601f106100fb57610100808354040283529160200191610126565b820191906000526020600020905b81548152906001019060200180831161010957829003601f168201915b50505050509080600101805461013b90610804565b80601f016020809104026020016040519081016040528092919081815260200182805461016790610804565b80156101b45780601f10610189576101008083540402835291602001916101b4565b820191906000526020600020905b81548152906001019060200180831161019757829003601f168201915b5050505050908060020180546101c990610804565b80601f01602080910402602001604051908101604052809291908181526020018280546101f590610804565b80156102425780601f1061021757610100808354040283529160200191610242565b820191906000526020600020905b81548152906001019060200180831161022557829003601f168201915b50505050509080600301805461025790610804565b80601f016020809104026020016040519081016040528092919081815260200182805461028390610804565b80156102d05780601f106102a5576101008083540402835291602001916102d0565b820191906000526020600020905b8154815290600101906020018083116102b357829003601f168201915b5050505050905084565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610368576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035f90610881565b60405180910390fd5b6040518060800160405280858152602001848152602001838152602001828152506001600087815260200190815260200160002060008201518160000190816103b19190610a4d565b5060208201518160010190816103c79190610a4d565b5060408201518160020190816103dd9190610a4d565b5060608201518160030190816103f39190610a4d565b50905050826040516104059190610b5b565b60405180910390207fc7fc4858662996e8fe091a9941384ab745bfc51199d7c0a5e45a791ff945993c868460405161043e929190610b81565b60405180910390a25050505050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61047481610461565b811461047f57600080fd5b50565b6000813590506104918161046b565b92915050565b6000602082840312156104ad576104ac610457565b5b60006104bb84828501610482565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156104fe5780820151818401526020810190506104e3565b60008484015250505050565b6000601f19601f8301169050919050565b6000610526826104c4565b61053081856104cf565b93506105408185602086016104e0565b6105498161050a565b840191505092915050565b6000608082019050818103600083015261056e818761051b565b90508181036020830152610582818661051b565b90508181036040830152610596818561051b565b905081810360608301526105aa818461051b565b905095945050505050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6105f78261050a565b810181811067ffffffffffffffff82111715610616576106156105bf565b5b80604052505050565b600061062961044d565b905061063582826105ee565b919050565b600067ffffffffffffffff821115610655576106546105bf565b5b61065e8261050a565b9050602081019050919050565b82818337600083830152505050565b600061068d6106888461063a565b61061f565b9050828152602081018484840111156106a9576106a86105ba565b5b6106b484828561066b565b509392505050565b600082601f8301126106d1576106d06105b5565b5b81356106e184826020860161067a565b91505092915050565b600080600080600060a0868803121561070657610705610457565b5b600061071488828901610482565b955050602086013567ffffffffffffffff8111156107355761073461045c565b5b610741888289016106bc565b945050604086013567ffffffffffffffff8111156107625761076161045c565b5b61076e888289016106bc565b935050606086013567ffffffffffffffff81111561078f5761078e61045c565b5b61079b888289016106bc565b925050608086013567ffffffffffffffff8111156107bc576107bb61045c565b5b6107c8888289016106bc565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061081c57607f821691505b60208210810361082f5761082e6107d5565b5b50919050565b7f4163636573732044656e69656400000000000000000000000000000000000000600082015250565b600061086b600d836104cf565b915061087682610835565b602082019050919050565b6000602082019050818103600083015261089a8161085e565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026109037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826108c6565b61090d86836108c6565b95508019841693508086168417925050509392505050565b6000819050919050565b600061094a61094561094084610461565b610925565b610461565b9050919050565b6000819050919050565b6109648361092f565b61097861097082610951565b8484546108d3565b825550505050565b600090565b61098d610980565b61099881848461095b565b505050565b5b818110156109bc576109b1600082610985565b60018101905061099e565b5050565b601f821115610a01576109d2816108a1565b6109db846108b6565b810160208510156109ea578190505b6109fe6109f6856108b6565b83018261099d565b50505b505050565b600082821c905092915050565b6000610a2460001984600802610a06565b1980831691505092915050565b6000610a3d8383610a13565b9150826002028217905092915050565b610a56826104c4565b67ffffffffffffffff811115610a6f57610a6e6105bf565b5b610a798254610804565b610a848282856109c0565b600060209050601f831160018114610ab75760008415610aa5578287015190505b610aaf8582610a31565b865550610b17565b601f198416610ac5866108a1565b60005b82811015610aed57848901518255600182019150602085019450602081019050610ac8565b86831015610b0a5784890151610b06601f891682610a13565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b6000610b35826104c4565b610b3f8185610b1f565b9350610b4f8185602086016104e0565b80840191505092915050565b6000610b678284610b2a565b915081905092915050565b610b7b81610461565b82525050565b6000604082019050610b966000830185610b72565b8181036020830152610ba8818461051b565b9050939250505056fea2646970667358221220d32104daadc3676c16313433675333cb5b7e46ea6b03cdf06b05a9402e9ec52a64736f6c63430008120033",
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

// Certificates is a free data retrieval call binding the contract method 0x9622c836.
//
// Solidity: function Certificates(uint256 ) view returns(string name, string course, string grade, string date)
func (_Cert *CertCaller) Certificates(opts *bind.CallOpts, arg0 *big.Int) (struct {
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

// Certificates is a free data retrieval call binding the contract method 0x9622c836.
//
// Solidity: function Certificates(uint256 ) view returns(string name, string course, string grade, string date)
func (_Cert *CertSession) Certificates(arg0 *big.Int) (struct {
	Name   string
	Course string
	Grade  string
	Date   string
}, error) {
	return _Cert.Contract.Certificates(&_Cert.CallOpts, arg0)
}

// Certificates is a free data retrieval call binding the contract method 0x9622c836.
//
// Solidity: function Certificates(uint256 ) view returns(string name, string course, string grade, string date)
func (_Cert *CertCallerSession) Certificates(arg0 *big.Int) (struct {
	Name   string
	Course string
	Grade  string
	Date   string
}, error) {
	return _Cert.Contract.Certificates(&_Cert.CallOpts, arg0)
}

// Issue is a paid mutator transaction binding the contract method 0xd1e7be26.
//
// Solidity: function issue(uint256 _id, string _name, string _course, string _grade, string _date) returns()
func (_Cert *CertTransactor) Issue(opts *bind.TransactOpts, _id *big.Int, _name string, _course string, _grade string, _date string) (*types.Transaction, error) {
	return _Cert.contract.Transact(opts, "issue", _id, _name, _course, _grade, _date)
}

// Issue is a paid mutator transaction binding the contract method 0xd1e7be26.
//
// Solidity: function issue(uint256 _id, string _name, string _course, string _grade, string _date) returns()
func (_Cert *CertSession) Issue(_id *big.Int, _name string, _course string, _grade string, _date string) (*types.Transaction, error) {
	return _Cert.Contract.Issue(&_Cert.TransactOpts, _id, _name, _course, _grade, _date)
}

// Issue is a paid mutator transaction binding the contract method 0xd1e7be26.
//
// Solidity: function issue(uint256 _id, string _name, string _course, string _grade, string _date) returns()
func (_Cert *CertTransactorSession) Issue(_id *big.Int, _name string, _course string, _grade string, _date string) (*types.Transaction, error) {
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
	Id     *big.Int
	Grade  string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssued is a free log retrieval operation binding the contract event 0xc7fc4858662996e8fe091a9941384ab745bfc51199d7c0a5e45a791ff945993c.
//
// Solidity: event Issued(string indexed course, uint256 id, string grade)
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

// WatchIssued is a free log subscription operation binding the contract event 0xc7fc4858662996e8fe091a9941384ab745bfc51199d7c0a5e45a791ff945993c.
//
// Solidity: event Issued(string indexed course, uint256 id, string grade)
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

// ParseIssued is a log parse operation binding the contract event 0xc7fc4858662996e8fe091a9941384ab745bfc51199d7c0a5e45a791ff945993c.
//
// Solidity: event Issued(string indexed course, uint256 id, string grade)
func (_Cert *CertFilterer) ParseIssued(log types.Log) (*CertIssued, error) {
	event := new(CertIssued)
	if err := _Cert.contract.UnpackLog(event, "Issued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
