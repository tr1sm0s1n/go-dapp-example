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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_course\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_grade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"course\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"grade\",\"type\":\"string\"}],\"name\":\"Issued\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"Certificates\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"course\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"grade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610bad8061005b5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c80630dfb1e86146100385780639d3becb214610054575b5f5ffd5b610052600480360381019061004d91906105ac565b610087565b005b61006e600480360381019061006991906106af565b610203565b60405161007e9493929190610756565b60405180910390f35b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610115576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161010c906107ff565b60405180910390fd5b6040518060800160405280858152602001848152602001838152602001828152506001866040516101469190610857565b90815260200160405180910390205f820151815f0190816101679190610a73565b50602082015181600101908161017d9190610a73565b5060408201518160020190816101939190610a73565b5060608201518160030190816101a99190610a73565b50905050826040516101bb9190610857565b60405180910390207ff34394c565949d2aff2cc92a2ca9216303b89a22f9953a2e86034ba2f7ee3cdd86846040516101f4929190610b42565b60405180910390a25050505050565b6001818051602081018201805184825260208301602085012081835280955050505050505f91509050805f01805461023a9061089a565b80601f01602080910402602001604051908101604052809291908181526020018280546102669061089a565b80156102b15780601f10610288576101008083540402835291602001916102b1565b820191905f5260205f20905b81548152906001019060200180831161029457829003601f168201915b5050505050908060010180546102c69061089a565b80601f01602080910402602001604051908101604052809291908181526020018280546102f29061089a565b801561033d5780601f106103145761010080835404028352916020019161033d565b820191905f5260205f20905b81548152906001019060200180831161032057829003601f168201915b5050505050908060020180546103529061089a565b80601f016020809104026020016040519081016040528092919081815260200182805461037e9061089a565b80156103c95780601f106103a0576101008083540402835291602001916103c9565b820191905f5260205f20905b8154815290600101906020018083116103ac57829003601f168201915b5050505050908060030180546103de9061089a565b80601f016020809104026020016040519081016040528092919081815260200182805461040a9061089a565b80156104555780601f1061042c57610100808354040283529160200191610455565b820191905f5260205f20905b81548152906001019060200180831161043857829003601f168201915b5050505050905084565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6104be82610478565b810181811067ffffffffffffffff821117156104dd576104dc610488565b5b80604052505050565b5f6104ef61045f565b90506104fb82826104b5565b919050565b5f67ffffffffffffffff82111561051a57610519610488565b5b61052382610478565b9050602081019050919050565b828183375f83830152505050565b5f61055061054b84610500565b6104e6565b90508281526020810184848401111561056c5761056b610474565b5b610577848285610530565b509392505050565b5f82601f83011261059357610592610470565b5b81356105a384826020860161053e565b91505092915050565b5f5f5f5f5f60a086880312156105c5576105c4610468565b5b5f86013567ffffffffffffffff8111156105e2576105e161046c565b5b6105ee8882890161057f565b955050602086013567ffffffffffffffff81111561060f5761060e61046c565b5b61061b8882890161057f565b945050604086013567ffffffffffffffff81111561063c5761063b61046c565b5b6106488882890161057f565b935050606086013567ffffffffffffffff8111156106695761066861046c565b5b6106758882890161057f565b925050608086013567ffffffffffffffff8111156106965761069561046c565b5b6106a28882890161057f565b9150509295509295909350565b5f602082840312156106c4576106c3610468565b5b5f82013567ffffffffffffffff8111156106e1576106e061046c565b5b6106ed8482850161057f565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610728826106f6565b6107328185610700565b9350610742818560208601610710565b61074b81610478565b840191505092915050565b5f6080820190508181035f83015261076e818761071e565b90508181036020830152610782818661071e565b90508181036040830152610796818561071e565b905081810360608301526107aa818461071e565b905095945050505050565b7f4163636573732044656e696564000000000000000000000000000000000000005f82015250565b5f6107e9600d83610700565b91506107f4826107b5565b602082019050919050565b5f6020820190508181035f830152610816816107dd565b9050919050565b5f81905092915050565b5f610831826106f6565b61083b818561081d565b935061084b818560208601610710565b80840191505092915050565b5f6108628284610827565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806108b157607f821691505b6020821081036108c4576108c361086d565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026109267fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826108eb565b61093086836108eb565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61097461096f61096a84610948565b610951565b610948565b9050919050565b5f819050919050565b61098d8361095a565b6109a16109998261097b565b8484546108f7565b825550505050565b5f5f905090565b6109b86109a9565b6109c3818484610984565b505050565b5b818110156109e6576109db5f826109b0565b6001810190506109c9565b5050565b601f821115610a2b576109fc816108ca565b610a05846108dc565b81016020851015610a14578190505b610a28610a20856108dc565b8301826109c8565b50505b505050565b5f82821c905092915050565b5f610a4b5f1984600802610a30565b1980831691505092915050565b5f610a638383610a3c565b9150826002028217905092915050565b610a7c826106f6565b67ffffffffffffffff811115610a9557610a94610488565b5b610a9f825461089a565b610aaa8282856109ea565b5f60209050601f831160018114610adb575f8415610ac9578287015190505b610ad38582610a58565b865550610b3a565b601f198416610ae9866108ca565b5f5b82811015610b1057848901518255600182019150602085019450602081019050610aeb565b86831015610b2d5784890151610b29601f891682610a3c565b8355505b6001600288020188555050505b505050505050565b5f6040820190508181035f830152610b5a818561071e565b90508181036020830152610b6e818461071e565b9050939250505056fea26469706673582212207fb455001f16499250109c2937d79fbf96f4372c155f3151d1292947eb1d574a64736f6c634300081c0033",
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
