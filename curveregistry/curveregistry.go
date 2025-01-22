// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curveregistry

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

// CurveregistryMetaData contains all meta data concerning the Curveregistry contract.
var CurveregistryMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"CommitNewAdmin\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_address_provider\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_registry_handler\",\"inputs\":[{\"name\":\"_registry_handler\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"update_registry_handler\",\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"},{\"name\":\"_registry_handler\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_registry_handlers_from_pool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[10]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_base_registry\",\"inputs\":[{\"name\":\"registry_handler\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pools_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_admin_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_admin_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_base_pool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_base_pool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_indices\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_indices\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_fees\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[10]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_fees\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[10]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauge\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauge\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"gauge_idx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauge\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"gauge_idx\",\"type\":\"uint256\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauge_type\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauge_type\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"gauge_idx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauge_type\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"gauge_idx\",\"type\":\"uint256\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_lp_token\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_lp_token\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_n_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_n_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_n_underlying_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_n_underlying_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_asset_type\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_asset_type\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_from_lp_token\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_from_lp_token\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_params\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[20]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_params\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[20]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_name\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_name\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price_from_lp_token\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price_from_lp_token\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_meta\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_meta\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_registered\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_registered\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_list\",\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"address_provider\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_registry\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"registry_length\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// CurveregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveregistryMetaData.ABI instead.
var CurveregistryABI = CurveregistryMetaData.ABI

// Curveregistry is an auto generated Go binding around an Ethereum contract.
type Curveregistry struct {
	CurveregistryCaller     // Read-only binding to the contract
	CurveregistryTransactor // Write-only binding to the contract
	CurveregistryFilterer   // Log filterer for contract events
}

// CurveregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveregistrySession struct {
	Contract     *Curveregistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveregistryCallerSession struct {
	Contract *CurveregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CurveregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveregistryTransactorSession struct {
	Contract     *CurveregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CurveregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveregistryRaw struct {
	Contract *Curveregistry // Generic contract binding to access the raw methods on
}

// CurveregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveregistryCallerRaw struct {
	Contract *CurveregistryCaller // Generic read-only contract binding to access the raw methods on
}

// CurveregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveregistryTransactorRaw struct {
	Contract *CurveregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveregistry creates a new instance of Curveregistry, bound to a specific deployed contract.
func NewCurveregistry(address common.Address, backend bind.ContractBackend) (*Curveregistry, error) {
	contract, err := bindCurveregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Curveregistry{CurveregistryCaller: CurveregistryCaller{contract: contract}, CurveregistryTransactor: CurveregistryTransactor{contract: contract}, CurveregistryFilterer: CurveregistryFilterer{contract: contract}}, nil
}

// NewCurveregistryCaller creates a new read-only instance of Curveregistry, bound to a specific deployed contract.
func NewCurveregistryCaller(address common.Address, caller bind.ContractCaller) (*CurveregistryCaller, error) {
	contract, err := bindCurveregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveregistryCaller{contract: contract}, nil
}

// NewCurveregistryTransactor creates a new write-only instance of Curveregistry, bound to a specific deployed contract.
func NewCurveregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveregistryTransactor, error) {
	contract, err := bindCurveregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveregistryTransactor{contract: contract}, nil
}

// NewCurveregistryFilterer creates a new log filterer instance of Curveregistry, bound to a specific deployed contract.
func NewCurveregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveregistryFilterer, error) {
	contract, err := bindCurveregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveregistryFilterer{contract: contract}, nil
}

// bindCurveregistry binds a generic wrapper to an already deployed contract.
func bindCurveregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurveregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Curveregistry *CurveregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Curveregistry.Contract.CurveregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Curveregistry *CurveregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curveregistry.Contract.CurveregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Curveregistry *CurveregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Curveregistry.Contract.CurveregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Curveregistry *CurveregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Curveregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Curveregistry *CurveregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curveregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Curveregistry *CurveregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Curveregistry.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_Curveregistry *CurveregistryCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "address_provider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_Curveregistry *CurveregistrySession) AddressProvider() (common.Address, error) {
	return _Curveregistry.Contract.AddressProvider(&_Curveregistry.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_Curveregistry *CurveregistryCallerSession) AddressProvider() (common.Address, error) {
	return _Curveregistry.Contract.AddressProvider(&_Curveregistry.CallOpts)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_Curveregistry *CurveregistryCaller) FindPoolForCoins(opts *bind.CallOpts, _from common.Address, _to common.Address) (common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "find_pool_for_coins", _from, _to)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_Curveregistry *CurveregistrySession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _Curveregistry.Contract.FindPoolForCoins(&_Curveregistry.CallOpts, _from, _to)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_Curveregistry *CurveregistryCallerSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _Curveregistry.Contract.FindPoolForCoins(&_Curveregistry.CallOpts, _from, _to)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_Curveregistry *CurveregistryCaller) FindPoolForCoins0(opts *bind.CallOpts, _from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "find_pool_for_coins0", _from, _to, i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_Curveregistry *CurveregistrySession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _Curveregistry.Contract.FindPoolForCoins0(&_Curveregistry.CallOpts, _from, _to, i)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_Curveregistry *CurveregistryCallerSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _Curveregistry.Contract.FindPoolForCoins0(&_Curveregistry.CallOpts, _from, _to, i)
}

// FindPoolsForCoins is a free data retrieval call binding the contract method 0xa064072b.
//
// Solidity: function find_pools_for_coins(address _from, address _to) view returns(address[])
func (_Curveregistry *CurveregistryCaller) FindPoolsForCoins(opts *bind.CallOpts, _from common.Address, _to common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "find_pools_for_coins", _from, _to)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// FindPoolsForCoins is a free data retrieval call binding the contract method 0xa064072b.
//
// Solidity: function find_pools_for_coins(address _from, address _to) view returns(address[])
func (_Curveregistry *CurveregistrySession) FindPoolsForCoins(_from common.Address, _to common.Address) ([]common.Address, error) {
	return _Curveregistry.Contract.FindPoolsForCoins(&_Curveregistry.CallOpts, _from, _to)
}

// FindPoolsForCoins is a free data retrieval call binding the contract method 0xa064072b.
//
// Solidity: function find_pools_for_coins(address _from, address _to) view returns(address[])
func (_Curveregistry *CurveregistryCallerSession) FindPoolsForCoins(_from common.Address, _to common.Address) ([]common.Address, error) {
	return _Curveregistry.Contract.FindPoolsForCoins(&_Curveregistry.CallOpts, _from, _to)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[8])
func (_Curveregistry *CurveregistryCaller) GetAdminBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_admin_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[8])
func (_Curveregistry *CurveregistrySession) GetAdminBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetAdminBalances(&_Curveregistry.CallOpts, _pool)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[8])
func (_Curveregistry *CurveregistryCallerSession) GetAdminBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetAdminBalances(&_Curveregistry.CallOpts, _pool)
}

// GetAdminBalances0 is a free data retrieval call binding the contract method 0xc0bf4cbf.
//
// Solidity: function get_admin_balances(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Curveregistry *CurveregistryCaller) GetAdminBalances0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_admin_balances0", _pool, _handler_id)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetAdminBalances0 is a free data retrieval call binding the contract method 0xc0bf4cbf.
//
// Solidity: function get_admin_balances(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Curveregistry *CurveregistrySession) GetAdminBalances0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetAdminBalances0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetAdminBalances0 is a free data retrieval call binding the contract method 0xc0bf4cbf.
//
// Solidity: function get_admin_balances(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Curveregistry *CurveregistryCallerSession) GetAdminBalances0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetAdminBalances0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[8])
func (_Curveregistry *CurveregistryCaller) GetBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[8])
func (_Curveregistry *CurveregistrySession) GetBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetBalances(&_Curveregistry.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[8])
func (_Curveregistry *CurveregistryCallerSession) GetBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetBalances(&_Curveregistry.CallOpts, _pool)
}

// GetBalances0 is a free data retrieval call binding the contract method 0xaa85169c.
//
// Solidity: function get_balances(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Curveregistry *CurveregistryCaller) GetBalances0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_balances0", _pool, _handler_id)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetBalances0 is a free data retrieval call binding the contract method 0xaa85169c.
//
// Solidity: function get_balances(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Curveregistry *CurveregistrySession) GetBalances0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetBalances0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetBalances0 is a free data retrieval call binding the contract method 0xaa85169c.
//
// Solidity: function get_balances(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Curveregistry *CurveregistryCallerSession) GetBalances0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetBalances0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetBasePool is a free data retrieval call binding the contract method 0x6f20d6dd.
//
// Solidity: function get_base_pool(address _pool) view returns(address)
func (_Curveregistry *CurveregistryCaller) GetBasePool(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_base_pool", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBasePool is a free data retrieval call binding the contract method 0x6f20d6dd.
//
// Solidity: function get_base_pool(address _pool) view returns(address)
func (_Curveregistry *CurveregistrySession) GetBasePool(_pool common.Address) (common.Address, error) {
	return _Curveregistry.Contract.GetBasePool(&_Curveregistry.CallOpts, _pool)
}

// GetBasePool is a free data retrieval call binding the contract method 0x6f20d6dd.
//
// Solidity: function get_base_pool(address _pool) view returns(address)
func (_Curveregistry *CurveregistryCallerSession) GetBasePool(_pool common.Address) (common.Address, error) {
	return _Curveregistry.Contract.GetBasePool(&_Curveregistry.CallOpts, _pool)
}

// GetBasePool0 is a free data retrieval call binding the contract method 0xa54e3ade.
//
// Solidity: function get_base_pool(address _pool, uint256 _handler_id) view returns(address)
func (_Curveregistry *CurveregistryCaller) GetBasePool0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_base_pool0", _pool, _handler_id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBasePool0 is a free data retrieval call binding the contract method 0xa54e3ade.
//
// Solidity: function get_base_pool(address _pool, uint256 _handler_id) view returns(address)
func (_Curveregistry *CurveregistrySession) GetBasePool0(_pool common.Address, _handler_id *big.Int) (common.Address, error) {
	return _Curveregistry.Contract.GetBasePool0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetBasePool0 is a free data retrieval call binding the contract method 0xa54e3ade.
//
// Solidity: function get_base_pool(address _pool, uint256 _handler_id) view returns(address)
func (_Curveregistry *CurveregistryCallerSession) GetBasePool0(_pool common.Address, _handler_id *big.Int) (common.Address, error) {
	return _Curveregistry.Contract.GetBasePool0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetBaseRegistry is a free data retrieval call binding the contract method 0x84e1710d.
//
// Solidity: function get_base_registry(address registry_handler) view returns(address)
func (_Curveregistry *CurveregistryCaller) GetBaseRegistry(opts *bind.CallOpts, registry_handler common.Address) (common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_base_registry", registry_handler)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBaseRegistry is a free data retrieval call binding the contract method 0x84e1710d.
//
// Solidity: function get_base_registry(address registry_handler) view returns(address)
func (_Curveregistry *CurveregistrySession) GetBaseRegistry(registry_handler common.Address) (common.Address, error) {
	return _Curveregistry.Contract.GetBaseRegistry(&_Curveregistry.CallOpts, registry_handler)
}

// GetBaseRegistry is a free data retrieval call binding the contract method 0x84e1710d.
//
// Solidity: function get_base_registry(address registry_handler) view returns(address)
func (_Curveregistry *CurveregistryCallerSession) GetBaseRegistry(registry_handler common.Address) (common.Address, error) {
	return _Curveregistry.Contract.GetBaseRegistry(&_Curveregistry.CallOpts, registry_handler)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_Curveregistry *CurveregistryCaller) GetCoinIndices(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_coin_indices", _pool, _from, _to)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_Curveregistry *CurveregistrySession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	return _Curveregistry.Contract.GetCoinIndices(&_Curveregistry.CallOpts, _pool, _from, _to)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_Curveregistry *CurveregistryCallerSession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	return _Curveregistry.Contract.GetCoinIndices(&_Curveregistry.CallOpts, _pool, _from, _to)
}

// GetCoinIndices0 is a free data retrieval call binding the contract method 0x63fb3ddb.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to, uint256 _handler_id) view returns(int128, int128, bool)
func (_Curveregistry *CurveregistryCaller) GetCoinIndices0(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address, _handler_id *big.Int) (*big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_coin_indices0", _pool, _from, _to, _handler_id)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetCoinIndices0 is a free data retrieval call binding the contract method 0x63fb3ddb.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to, uint256 _handler_id) view returns(int128, int128, bool)
func (_Curveregistry *CurveregistrySession) GetCoinIndices0(_pool common.Address, _from common.Address, _to common.Address, _handler_id *big.Int) (*big.Int, *big.Int, bool, error) {
	return _Curveregistry.Contract.GetCoinIndices0(&_Curveregistry.CallOpts, _pool, _from, _to, _handler_id)
}

// GetCoinIndices0 is a free data retrieval call binding the contract method 0x63fb3ddb.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to, uint256 _handler_id) view returns(int128, int128, bool)
func (_Curveregistry *CurveregistryCallerSession) GetCoinIndices0(_pool common.Address, _from common.Address, _to common.Address, _handler_id *big.Int) (*big.Int, *big.Int, bool, error) {
	return _Curveregistry.Contract.GetCoinIndices0(&_Curveregistry.CallOpts, _pool, _from, _to, _handler_id)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[8])
func (_Curveregistry *CurveregistryCaller) GetCoins(opts *bind.CallOpts, _pool common.Address) ([8]common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_coins", _pool)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[8])
func (_Curveregistry *CurveregistrySession) GetCoins(_pool common.Address) ([8]common.Address, error) {
	return _Curveregistry.Contract.GetCoins(&_Curveregistry.CallOpts, _pool)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[8])
func (_Curveregistry *CurveregistryCallerSession) GetCoins(_pool common.Address) ([8]common.Address, error) {
	return _Curveregistry.Contract.GetCoins(&_Curveregistry.CallOpts, _pool)
}

// GetCoins0 is a free data retrieval call binding the contract method 0x6ebe51fc.
//
// Solidity: function get_coins(address _pool, uint256 _handler_id) view returns(address[8])
func (_Curveregistry *CurveregistryCaller) GetCoins0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) ([8]common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_coins0", _pool, _handler_id)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetCoins0 is a free data retrieval call binding the contract method 0x6ebe51fc.
//
// Solidity: function get_coins(address _pool, uint256 _handler_id) view returns(address[8])
func (_Curveregistry *CurveregistrySession) GetCoins0(_pool common.Address, _handler_id *big.Int) ([8]common.Address, error) {
	return _Curveregistry.Contract.GetCoins0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetCoins0 is a free data retrieval call binding the contract method 0x6ebe51fc.
//
// Solidity: function get_coins(address _pool, uint256 _handler_id) view returns(address[8])
func (_Curveregistry *CurveregistryCallerSession) GetCoins0(_pool common.Address, _handler_id *big.Int) ([8]common.Address, error) {
	return _Curveregistry.Contract.GetCoins0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[8])
func (_Curveregistry *CurveregistryCaller) GetDecimals(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_decimals", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[8])
func (_Curveregistry *CurveregistrySession) GetDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetDecimals(&_Curveregistry.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[8])
func (_Curveregistry *CurveregistryCallerSession) GetDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetDecimals(&_Curveregistry.CallOpts, _pool)
}

// GetDecimals0 is a free data retrieval call binding the contract method 0x403f502f.
//
// Solidity: function get_decimals(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Curveregistry *CurveregistryCaller) GetDecimals0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_decimals0", _pool, _handler_id)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetDecimals0 is a free data retrieval call binding the contract method 0x403f502f.
//
// Solidity: function get_decimals(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Curveregistry *CurveregistrySession) GetDecimals0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetDecimals0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetDecimals0 is a free data retrieval call binding the contract method 0x403f502f.
//
// Solidity: function get_decimals(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Curveregistry *CurveregistryCallerSession) GetDecimals0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetDecimals0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256[10])
func (_Curveregistry *CurveregistryCaller) GetFees(opts *bind.CallOpts, _pool common.Address) ([10]*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_fees", _pool)

	if err != nil {
		return *new([10]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([10]*big.Int)).(*[10]*big.Int)

	return out0, err

}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256[10])
func (_Curveregistry *CurveregistrySession) GetFees(_pool common.Address) ([10]*big.Int, error) {
	return _Curveregistry.Contract.GetFees(&_Curveregistry.CallOpts, _pool)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256[10])
func (_Curveregistry *CurveregistryCallerSession) GetFees(_pool common.Address) ([10]*big.Int, error) {
	return _Curveregistry.Contract.GetFees(&_Curveregistry.CallOpts, _pool)
}

// GetFees0 is a free data retrieval call binding the contract method 0x0ed5a427.
//
// Solidity: function get_fees(address _pool, uint256 _handler_id) view returns(uint256[10])
func (_Curveregistry *CurveregistryCaller) GetFees0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) ([10]*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_fees0", _pool, _handler_id)

	if err != nil {
		return *new([10]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([10]*big.Int)).(*[10]*big.Int)

	return out0, err

}

// GetFees0 is a free data retrieval call binding the contract method 0x0ed5a427.
//
// Solidity: function get_fees(address _pool, uint256 _handler_id) view returns(uint256[10])
func (_Curveregistry *CurveregistrySession) GetFees0(_pool common.Address, _handler_id *big.Int) ([10]*big.Int, error) {
	return _Curveregistry.Contract.GetFees0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetFees0 is a free data retrieval call binding the contract method 0x0ed5a427.
//
// Solidity: function get_fees(address _pool, uint256 _handler_id) view returns(uint256[10])
func (_Curveregistry *CurveregistryCallerSession) GetFees0(_pool common.Address, _handler_id *big.Int) ([10]*big.Int, error) {
	return _Curveregistry.Contract.GetFees0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_Curveregistry *CurveregistryCaller) GetGauge(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_gauge", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_Curveregistry *CurveregistrySession) GetGauge(_pool common.Address) (common.Address, error) {
	return _Curveregistry.Contract.GetGauge(&_Curveregistry.CallOpts, _pool)
}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_Curveregistry *CurveregistryCallerSession) GetGauge(_pool common.Address) (common.Address, error) {
	return _Curveregistry.Contract.GetGauge(&_Curveregistry.CallOpts, _pool)
}

// GetGauge0 is a free data retrieval call binding the contract method 0xe4081220.
//
// Solidity: function get_gauge(address _pool, uint256 gauge_idx) view returns(address)
func (_Curveregistry *CurveregistryCaller) GetGauge0(opts *bind.CallOpts, _pool common.Address, gauge_idx *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_gauge0", _pool, gauge_idx)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGauge0 is a free data retrieval call binding the contract method 0xe4081220.
//
// Solidity: function get_gauge(address _pool, uint256 gauge_idx) view returns(address)
func (_Curveregistry *CurveregistrySession) GetGauge0(_pool common.Address, gauge_idx *big.Int) (common.Address, error) {
	return _Curveregistry.Contract.GetGauge0(&_Curveregistry.CallOpts, _pool, gauge_idx)
}

// GetGauge0 is a free data retrieval call binding the contract method 0xe4081220.
//
// Solidity: function get_gauge(address _pool, uint256 gauge_idx) view returns(address)
func (_Curveregistry *CurveregistryCallerSession) GetGauge0(_pool common.Address, gauge_idx *big.Int) (common.Address, error) {
	return _Curveregistry.Contract.GetGauge0(&_Curveregistry.CallOpts, _pool, gauge_idx)
}

// GetGauge1 is a free data retrieval call binding the contract method 0x773fb7e3.
//
// Solidity: function get_gauge(address _pool, uint256 gauge_idx, uint256 _handler_id) view returns(address)
func (_Curveregistry *CurveregistryCaller) GetGauge1(opts *bind.CallOpts, _pool common.Address, gauge_idx *big.Int, _handler_id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_gauge1", _pool, gauge_idx, _handler_id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGauge1 is a free data retrieval call binding the contract method 0x773fb7e3.
//
// Solidity: function get_gauge(address _pool, uint256 gauge_idx, uint256 _handler_id) view returns(address)
func (_Curveregistry *CurveregistrySession) GetGauge1(_pool common.Address, gauge_idx *big.Int, _handler_id *big.Int) (common.Address, error) {
	return _Curveregistry.Contract.GetGauge1(&_Curveregistry.CallOpts, _pool, gauge_idx, _handler_id)
}

// GetGauge1 is a free data retrieval call binding the contract method 0x773fb7e3.
//
// Solidity: function get_gauge(address _pool, uint256 gauge_idx, uint256 _handler_id) view returns(address)
func (_Curveregistry *CurveregistryCallerSession) GetGauge1(_pool common.Address, gauge_idx *big.Int, _handler_id *big.Int) (common.Address, error) {
	return _Curveregistry.Contract.GetGauge1(&_Curveregistry.CallOpts, _pool, gauge_idx, _handler_id)
}

// GetGaugeType is a free data retrieval call binding the contract method 0x25fa5d13.
//
// Solidity: function get_gauge_type(address _pool) view returns(int128)
func (_Curveregistry *CurveregistryCaller) GetGaugeType(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_gauge_type", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGaugeType is a free data retrieval call binding the contract method 0x25fa5d13.
//
// Solidity: function get_gauge_type(address _pool) view returns(int128)
func (_Curveregistry *CurveregistrySession) GetGaugeType(_pool common.Address) (*big.Int, error) {
	return _Curveregistry.Contract.GetGaugeType(&_Curveregistry.CallOpts, _pool)
}

// GetGaugeType is a free data retrieval call binding the contract method 0x25fa5d13.
//
// Solidity: function get_gauge_type(address _pool) view returns(int128)
func (_Curveregistry *CurveregistryCallerSession) GetGaugeType(_pool common.Address) (*big.Int, error) {
	return _Curveregistry.Contract.GetGaugeType(&_Curveregistry.CallOpts, _pool)
}

// GetGaugeType0 is a free data retrieval call binding the contract method 0x7c51db55.
//
// Solidity: function get_gauge_type(address _pool, uint256 gauge_idx) view returns(int128)
func (_Curveregistry *CurveregistryCaller) GetGaugeType0(opts *bind.CallOpts, _pool common.Address, gauge_idx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_gauge_type0", _pool, gauge_idx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGaugeType0 is a free data retrieval call binding the contract method 0x7c51db55.
//
// Solidity: function get_gauge_type(address _pool, uint256 gauge_idx) view returns(int128)
func (_Curveregistry *CurveregistrySession) GetGaugeType0(_pool common.Address, gauge_idx *big.Int) (*big.Int, error) {
	return _Curveregistry.Contract.GetGaugeType0(&_Curveregistry.CallOpts, _pool, gauge_idx)
}

// GetGaugeType0 is a free data retrieval call binding the contract method 0x7c51db55.
//
// Solidity: function get_gauge_type(address _pool, uint256 gauge_idx) view returns(int128)
func (_Curveregistry *CurveregistryCallerSession) GetGaugeType0(_pool common.Address, gauge_idx *big.Int) (*big.Int, error) {
	return _Curveregistry.Contract.GetGaugeType0(&_Curveregistry.CallOpts, _pool, gauge_idx)
}

// GetGaugeType1 is a free data retrieval call binding the contract method 0xf8b661e2.
//
// Solidity: function get_gauge_type(address _pool, uint256 gauge_idx, uint256 _handler_id) view returns(int128)
func (_Curveregistry *CurveregistryCaller) GetGaugeType1(opts *bind.CallOpts, _pool common.Address, gauge_idx *big.Int, _handler_id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_gauge_type1", _pool, gauge_idx, _handler_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGaugeType1 is a free data retrieval call binding the contract method 0xf8b661e2.
//
// Solidity: function get_gauge_type(address _pool, uint256 gauge_idx, uint256 _handler_id) view returns(int128)
func (_Curveregistry *CurveregistrySession) GetGaugeType1(_pool common.Address, gauge_idx *big.Int, _handler_id *big.Int) (*big.Int, error) {
	return _Curveregistry.Contract.GetGaugeType1(&_Curveregistry.CallOpts, _pool, gauge_idx, _handler_id)
}

// GetGaugeType1 is a free data retrieval call binding the contract method 0xf8b661e2.
//
// Solidity: function get_gauge_type(address _pool, uint256 gauge_idx, uint256 _handler_id) view returns(int128)
func (_Curveregistry *CurveregistryCallerSession) GetGaugeType1(_pool common.Address, gauge_idx *big.Int, _handler_id *big.Int) (*big.Int, error) {
	return _Curveregistry.Contract.GetGaugeType1(&_Curveregistry.CallOpts, _pool, gauge_idx, _handler_id)
}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address _pool) view returns(address)
func (_Curveregistry *CurveregistryCaller) GetLpToken(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_lp_token", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address _pool) view returns(address)
func (_Curveregistry *CurveregistrySession) GetLpToken(_pool common.Address) (common.Address, error) {
	return _Curveregistry.Contract.GetLpToken(&_Curveregistry.CallOpts, _pool)
}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address _pool) view returns(address)
func (_Curveregistry *CurveregistryCallerSession) GetLpToken(_pool common.Address) (common.Address, error) {
	return _Curveregistry.Contract.GetLpToken(&_Curveregistry.CallOpts, _pool)
}

// GetLpToken0 is a free data retrieval call binding the contract method 0x0881715f.
//
// Solidity: function get_lp_token(address _pool, uint256 _handler_id) view returns(address)
func (_Curveregistry *CurveregistryCaller) GetLpToken0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_lp_token0", _pool, _handler_id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLpToken0 is a free data retrieval call binding the contract method 0x0881715f.
//
// Solidity: function get_lp_token(address _pool, uint256 _handler_id) view returns(address)
func (_Curveregistry *CurveregistrySession) GetLpToken0(_pool common.Address, _handler_id *big.Int) (common.Address, error) {
	return _Curveregistry.Contract.GetLpToken0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetLpToken0 is a free data retrieval call binding the contract method 0x0881715f.
//
// Solidity: function get_lp_token(address _pool, uint256 _handler_id) view returns(address)
func (_Curveregistry *CurveregistryCallerSession) GetLpToken0(_pool common.Address, _handler_id *big.Int) (common.Address, error) {
	return _Curveregistry.Contract.GetLpToken0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_Curveregistry *CurveregistryCaller) GetNCoins(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_n_coins", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_Curveregistry *CurveregistrySession) GetNCoins(_pool common.Address) (*big.Int, error) {
	return _Curveregistry.Contract.GetNCoins(&_Curveregistry.CallOpts, _pool)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_Curveregistry *CurveregistryCallerSession) GetNCoins(_pool common.Address) (*big.Int, error) {
	return _Curveregistry.Contract.GetNCoins(&_Curveregistry.CallOpts, _pool)
}

// GetNCoins0 is a free data retrieval call binding the contract method 0x11d81279.
//
// Solidity: function get_n_coins(address _pool, uint256 _handler_id) view returns(uint256)
func (_Curveregistry *CurveregistryCaller) GetNCoins0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_n_coins0", _pool, _handler_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNCoins0 is a free data retrieval call binding the contract method 0x11d81279.
//
// Solidity: function get_n_coins(address _pool, uint256 _handler_id) view returns(uint256)
func (_Curveregistry *CurveregistrySession) GetNCoins0(_pool common.Address, _handler_id *big.Int) (*big.Int, error) {
	return _Curveregistry.Contract.GetNCoins0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetNCoins0 is a free data retrieval call binding the contract method 0x11d81279.
//
// Solidity: function get_n_coins(address _pool, uint256 _handler_id) view returns(uint256)
func (_Curveregistry *CurveregistryCallerSession) GetNCoins0(_pool common.Address, _handler_id *big.Int) (*big.Int, error) {
	return _Curveregistry.Contract.GetNCoins0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetNUnderlyingCoins is a free data retrieval call binding the contract method 0x0a700c08.
//
// Solidity: function get_n_underlying_coins(address _pool) view returns(uint256)
func (_Curveregistry *CurveregistryCaller) GetNUnderlyingCoins(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_n_underlying_coins", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNUnderlyingCoins is a free data retrieval call binding the contract method 0x0a700c08.
//
// Solidity: function get_n_underlying_coins(address _pool) view returns(uint256)
func (_Curveregistry *CurveregistrySession) GetNUnderlyingCoins(_pool common.Address) (*big.Int, error) {
	return _Curveregistry.Contract.GetNUnderlyingCoins(&_Curveregistry.CallOpts, _pool)
}

// GetNUnderlyingCoins is a free data retrieval call binding the contract method 0x0a700c08.
//
// Solidity: function get_n_underlying_coins(address _pool) view returns(uint256)
func (_Curveregistry *CurveregistryCallerSession) GetNUnderlyingCoins(_pool common.Address) (*big.Int, error) {
	return _Curveregistry.Contract.GetNUnderlyingCoins(&_Curveregistry.CallOpts, _pool)
}

// GetNUnderlyingCoins0 is a free data retrieval call binding the contract method 0xdecdf68f.
//
// Solidity: function get_n_underlying_coins(address _pool, uint256 _handler_id) view returns(uint256)
func (_Curveregistry *CurveregistryCaller) GetNUnderlyingCoins0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_n_underlying_coins0", _pool, _handler_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNUnderlyingCoins0 is a free data retrieval call binding the contract method 0xdecdf68f.
//
// Solidity: function get_n_underlying_coins(address _pool, uint256 _handler_id) view returns(uint256)
func (_Curveregistry *CurveregistrySession) GetNUnderlyingCoins0(_pool common.Address, _handler_id *big.Int) (*big.Int, error) {
	return _Curveregistry.Contract.GetNUnderlyingCoins0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetNUnderlyingCoins0 is a free data retrieval call binding the contract method 0xdecdf68f.
//
// Solidity: function get_n_underlying_coins(address _pool, uint256 _handler_id) view returns(uint256)
func (_Curveregistry *CurveregistryCallerSession) GetNUnderlyingCoins0(_pool common.Address, _handler_id *big.Int) (*big.Int, error) {
	return _Curveregistry.Contract.GetNUnderlyingCoins0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_Curveregistry *CurveregistryCaller) GetPoolAssetType(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_pool_asset_type", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_Curveregistry *CurveregistrySession) GetPoolAssetType(_pool common.Address) (*big.Int, error) {
	return _Curveregistry.Contract.GetPoolAssetType(&_Curveregistry.CallOpts, _pool)
}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_Curveregistry *CurveregistryCallerSession) GetPoolAssetType(_pool common.Address) (*big.Int, error) {
	return _Curveregistry.Contract.GetPoolAssetType(&_Curveregistry.CallOpts, _pool)
}

// GetPoolAssetType0 is a free data retrieval call binding the contract method 0x90d1dd2d.
//
// Solidity: function get_pool_asset_type(address _pool, uint256 _handler_id) view returns(uint256)
func (_Curveregistry *CurveregistryCaller) GetPoolAssetType0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_pool_asset_type0", _pool, _handler_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolAssetType0 is a free data retrieval call binding the contract method 0x90d1dd2d.
//
// Solidity: function get_pool_asset_type(address _pool, uint256 _handler_id) view returns(uint256)
func (_Curveregistry *CurveregistrySession) GetPoolAssetType0(_pool common.Address, _handler_id *big.Int) (*big.Int, error) {
	return _Curveregistry.Contract.GetPoolAssetType0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetPoolAssetType0 is a free data retrieval call binding the contract method 0x90d1dd2d.
//
// Solidity: function get_pool_asset_type(address _pool, uint256 _handler_id) view returns(uint256)
func (_Curveregistry *CurveregistryCallerSession) GetPoolAssetType0(_pool common.Address, _handler_id *big.Int) (*big.Int, error) {
	return _Curveregistry.Contract.GetPoolAssetType0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetPoolFromLpToken is a free data retrieval call binding the contract method 0xbdf475c3.
//
// Solidity: function get_pool_from_lp_token(address _token) view returns(address)
func (_Curveregistry *CurveregistryCaller) GetPoolFromLpToken(opts *bind.CallOpts, _token common.Address) (common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_pool_from_lp_token", _token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolFromLpToken is a free data retrieval call binding the contract method 0xbdf475c3.
//
// Solidity: function get_pool_from_lp_token(address _token) view returns(address)
func (_Curveregistry *CurveregistrySession) GetPoolFromLpToken(_token common.Address) (common.Address, error) {
	return _Curveregistry.Contract.GetPoolFromLpToken(&_Curveregistry.CallOpts, _token)
}

// GetPoolFromLpToken is a free data retrieval call binding the contract method 0xbdf475c3.
//
// Solidity: function get_pool_from_lp_token(address _token) view returns(address)
func (_Curveregistry *CurveregistryCallerSession) GetPoolFromLpToken(_token common.Address) (common.Address, error) {
	return _Curveregistry.Contract.GetPoolFromLpToken(&_Curveregistry.CallOpts, _token)
}

// GetPoolFromLpToken0 is a free data retrieval call binding the contract method 0x36678b80.
//
// Solidity: function get_pool_from_lp_token(address _token, uint256 _handler_id) view returns(address)
func (_Curveregistry *CurveregistryCaller) GetPoolFromLpToken0(opts *bind.CallOpts, _token common.Address, _handler_id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_pool_from_lp_token0", _token, _handler_id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolFromLpToken0 is a free data retrieval call binding the contract method 0x36678b80.
//
// Solidity: function get_pool_from_lp_token(address _token, uint256 _handler_id) view returns(address)
func (_Curveregistry *CurveregistrySession) GetPoolFromLpToken0(_token common.Address, _handler_id *big.Int) (common.Address, error) {
	return _Curveregistry.Contract.GetPoolFromLpToken0(&_Curveregistry.CallOpts, _token, _handler_id)
}

// GetPoolFromLpToken0 is a free data retrieval call binding the contract method 0x36678b80.
//
// Solidity: function get_pool_from_lp_token(address _token, uint256 _handler_id) view returns(address)
func (_Curveregistry *CurveregistryCallerSession) GetPoolFromLpToken0(_token common.Address, _handler_id *big.Int) (common.Address, error) {
	return _Curveregistry.Contract.GetPoolFromLpToken0(&_Curveregistry.CallOpts, _token, _handler_id)
}

// GetPoolName is a free data retrieval call binding the contract method 0x5c911741.
//
// Solidity: function get_pool_name(address _pool) view returns(string)
func (_Curveregistry *CurveregistryCaller) GetPoolName(opts *bind.CallOpts, _pool common.Address) (string, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_pool_name", _pool)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPoolName is a free data retrieval call binding the contract method 0x5c911741.
//
// Solidity: function get_pool_name(address _pool) view returns(string)
func (_Curveregistry *CurveregistrySession) GetPoolName(_pool common.Address) (string, error) {
	return _Curveregistry.Contract.GetPoolName(&_Curveregistry.CallOpts, _pool)
}

// GetPoolName is a free data retrieval call binding the contract method 0x5c911741.
//
// Solidity: function get_pool_name(address _pool) view returns(string)
func (_Curveregistry *CurveregistryCallerSession) GetPoolName(_pool common.Address) (string, error) {
	return _Curveregistry.Contract.GetPoolName(&_Curveregistry.CallOpts, _pool)
}

// GetPoolName0 is a free data retrieval call binding the contract method 0x92234f45.
//
// Solidity: function get_pool_name(address _pool, uint256 _handler_id) view returns(string)
func (_Curveregistry *CurveregistryCaller) GetPoolName0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) (string, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_pool_name0", _pool, _handler_id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPoolName0 is a free data retrieval call binding the contract method 0x92234f45.
//
// Solidity: function get_pool_name(address _pool, uint256 _handler_id) view returns(string)
func (_Curveregistry *CurveregistrySession) GetPoolName0(_pool common.Address, _handler_id *big.Int) (string, error) {
	return _Curveregistry.Contract.GetPoolName0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetPoolName0 is a free data retrieval call binding the contract method 0x92234f45.
//
// Solidity: function get_pool_name(address _pool, uint256 _handler_id) view returns(string)
func (_Curveregistry *CurveregistryCallerSession) GetPoolName0(_pool common.Address, _handler_id *big.Int) (string, error) {
	return _Curveregistry.Contract.GetPoolName0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetPoolParams is a free data retrieval call binding the contract method 0x688532aa.
//
// Solidity: function get_pool_params(address _pool) view returns(uint256[20])
func (_Curveregistry *CurveregistryCaller) GetPoolParams(opts *bind.CallOpts, _pool common.Address) ([20]*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_pool_params", _pool)

	if err != nil {
		return *new([20]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([20]*big.Int)).(*[20]*big.Int)

	return out0, err

}

// GetPoolParams is a free data retrieval call binding the contract method 0x688532aa.
//
// Solidity: function get_pool_params(address _pool) view returns(uint256[20])
func (_Curveregistry *CurveregistrySession) GetPoolParams(_pool common.Address) ([20]*big.Int, error) {
	return _Curveregistry.Contract.GetPoolParams(&_Curveregistry.CallOpts, _pool)
}

// GetPoolParams is a free data retrieval call binding the contract method 0x688532aa.
//
// Solidity: function get_pool_params(address _pool) view returns(uint256[20])
func (_Curveregistry *CurveregistryCallerSession) GetPoolParams(_pool common.Address) ([20]*big.Int, error) {
	return _Curveregistry.Contract.GetPoolParams(&_Curveregistry.CallOpts, _pool)
}

// GetPoolParams0 is a free data retrieval call binding the contract method 0x7a65d2dd.
//
// Solidity: function get_pool_params(address _pool, uint256 _handler_id) view returns(uint256[20])
func (_Curveregistry *CurveregistryCaller) GetPoolParams0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) ([20]*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_pool_params0", _pool, _handler_id)

	if err != nil {
		return *new([20]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([20]*big.Int)).(*[20]*big.Int)

	return out0, err

}

// GetPoolParams0 is a free data retrieval call binding the contract method 0x7a65d2dd.
//
// Solidity: function get_pool_params(address _pool, uint256 _handler_id) view returns(uint256[20])
func (_Curveregistry *CurveregistrySession) GetPoolParams0(_pool common.Address, _handler_id *big.Int) ([20]*big.Int, error) {
	return _Curveregistry.Contract.GetPoolParams0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetPoolParams0 is a free data retrieval call binding the contract method 0x7a65d2dd.
//
// Solidity: function get_pool_params(address _pool, uint256 _handler_id) view returns(uint256[20])
func (_Curveregistry *CurveregistryCallerSession) GetPoolParams0(_pool common.Address, _handler_id *big.Int) ([20]*big.Int, error) {
	return _Curveregistry.Contract.GetPoolParams0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetRegistry is a free data retrieval call binding the contract method 0x913d9b4d.
//
// Solidity: function get_registry(uint256 arg0) view returns(address)
func (_Curveregistry *CurveregistryCaller) GetRegistry(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_registry", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRegistry is a free data retrieval call binding the contract method 0x913d9b4d.
//
// Solidity: function get_registry(uint256 arg0) view returns(address)
func (_Curveregistry *CurveregistrySession) GetRegistry(arg0 *big.Int) (common.Address, error) {
	return _Curveregistry.Contract.GetRegistry(&_Curveregistry.CallOpts, arg0)
}

// GetRegistry is a free data retrieval call binding the contract method 0x913d9b4d.
//
// Solidity: function get_registry(uint256 arg0) view returns(address)
func (_Curveregistry *CurveregistryCallerSession) GetRegistry(arg0 *big.Int) (common.Address, error) {
	return _Curveregistry.Contract.GetRegistry(&_Curveregistry.CallOpts, arg0)
}

// GetRegistryHandlersFromPool is a free data retrieval call binding the contract method 0x308d1b6d.
//
// Solidity: function get_registry_handlers_from_pool(address _pool) view returns(address[10])
func (_Curveregistry *CurveregistryCaller) GetRegistryHandlersFromPool(opts *bind.CallOpts, _pool common.Address) ([10]common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_registry_handlers_from_pool", _pool)

	if err != nil {
		return *new([10]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([10]common.Address)).(*[10]common.Address)

	return out0, err

}

// GetRegistryHandlersFromPool is a free data retrieval call binding the contract method 0x308d1b6d.
//
// Solidity: function get_registry_handlers_from_pool(address _pool) view returns(address[10])
func (_Curveregistry *CurveregistrySession) GetRegistryHandlersFromPool(_pool common.Address) ([10]common.Address, error) {
	return _Curveregistry.Contract.GetRegistryHandlersFromPool(&_Curveregistry.CallOpts, _pool)
}

// GetRegistryHandlersFromPool is a free data retrieval call binding the contract method 0x308d1b6d.
//
// Solidity: function get_registry_handlers_from_pool(address _pool) view returns(address[10])
func (_Curveregistry *CurveregistryCallerSession) GetRegistryHandlersFromPool(_pool common.Address) ([10]common.Address, error) {
	return _Curveregistry.Contract.GetRegistryHandlersFromPool(&_Curveregistry.CallOpts, _pool)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_Curveregistry *CurveregistryCaller) GetUnderlyingBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_underlying_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_Curveregistry *CurveregistrySession) GetUnderlyingBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetUnderlyingBalances(&_Curveregistry.CallOpts, _pool)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_Curveregistry *CurveregistryCallerSession) GetUnderlyingBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetUnderlyingBalances(&_Curveregistry.CallOpts, _pool)
}

// GetUnderlyingBalances0 is a free data retrieval call binding the contract method 0x876112e6.
//
// Solidity: function get_underlying_balances(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Curveregistry *CurveregistryCaller) GetUnderlyingBalances0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_underlying_balances0", _pool, _handler_id)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingBalances0 is a free data retrieval call binding the contract method 0x876112e6.
//
// Solidity: function get_underlying_balances(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Curveregistry *CurveregistrySession) GetUnderlyingBalances0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetUnderlyingBalances0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetUnderlyingBalances0 is a free data retrieval call binding the contract method 0x876112e6.
//
// Solidity: function get_underlying_balances(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Curveregistry *CurveregistryCallerSession) GetUnderlyingBalances0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetUnderlyingBalances0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_Curveregistry *CurveregistryCaller) GetUnderlyingCoins(opts *bind.CallOpts, _pool common.Address) ([8]common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_underlying_coins", _pool)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_Curveregistry *CurveregistrySession) GetUnderlyingCoins(_pool common.Address) ([8]common.Address, error) {
	return _Curveregistry.Contract.GetUnderlyingCoins(&_Curveregistry.CallOpts, _pool)
}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_Curveregistry *CurveregistryCallerSession) GetUnderlyingCoins(_pool common.Address) ([8]common.Address, error) {
	return _Curveregistry.Contract.GetUnderlyingCoins(&_Curveregistry.CallOpts, _pool)
}

// GetUnderlyingCoins0 is a free data retrieval call binding the contract method 0x2fc90bcf.
//
// Solidity: function get_underlying_coins(address _pool, uint256 _handler_id) view returns(address[8])
func (_Curveregistry *CurveregistryCaller) GetUnderlyingCoins0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) ([8]common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_underlying_coins0", _pool, _handler_id)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetUnderlyingCoins0 is a free data retrieval call binding the contract method 0x2fc90bcf.
//
// Solidity: function get_underlying_coins(address _pool, uint256 _handler_id) view returns(address[8])
func (_Curveregistry *CurveregistrySession) GetUnderlyingCoins0(_pool common.Address, _handler_id *big.Int) ([8]common.Address, error) {
	return _Curveregistry.Contract.GetUnderlyingCoins0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetUnderlyingCoins0 is a free data retrieval call binding the contract method 0x2fc90bcf.
//
// Solidity: function get_underlying_coins(address _pool, uint256 _handler_id) view returns(address[8])
func (_Curveregistry *CurveregistryCallerSession) GetUnderlyingCoins0(_pool common.Address, _handler_id *big.Int) ([8]common.Address, error) {
	return _Curveregistry.Contract.GetUnderlyingCoins0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_Curveregistry *CurveregistryCaller) GetUnderlyingDecimals(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_underlying_decimals", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_Curveregistry *CurveregistrySession) GetUnderlyingDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetUnderlyingDecimals(&_Curveregistry.CallOpts, _pool)
}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_Curveregistry *CurveregistryCallerSession) GetUnderlyingDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetUnderlyingDecimals(&_Curveregistry.CallOpts, _pool)
}

// GetUnderlyingDecimals0 is a free data retrieval call binding the contract method 0x622d64f4.
//
// Solidity: function get_underlying_decimals(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Curveregistry *CurveregistryCaller) GetUnderlyingDecimals0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_underlying_decimals0", _pool, _handler_id)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingDecimals0 is a free data retrieval call binding the contract method 0x622d64f4.
//
// Solidity: function get_underlying_decimals(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Curveregistry *CurveregistrySession) GetUnderlyingDecimals0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetUnderlyingDecimals0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetUnderlyingDecimals0 is a free data retrieval call binding the contract method 0x622d64f4.
//
// Solidity: function get_underlying_decimals(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Curveregistry *CurveregistryCallerSession) GetUnderlyingDecimals0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Curveregistry.Contract.GetUnderlyingDecimals0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_Curveregistry *CurveregistryCaller) GetVirtualPriceFromLpToken(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_virtual_price_from_lp_token", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_Curveregistry *CurveregistrySession) GetVirtualPriceFromLpToken(_token common.Address) (*big.Int, error) {
	return _Curveregistry.Contract.GetVirtualPriceFromLpToken(&_Curveregistry.CallOpts, _token)
}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_Curveregistry *CurveregistryCallerSession) GetVirtualPriceFromLpToken(_token common.Address) (*big.Int, error) {
	return _Curveregistry.Contract.GetVirtualPriceFromLpToken(&_Curveregistry.CallOpts, _token)
}

// GetVirtualPriceFromLpToken0 is a free data retrieval call binding the contract method 0x4460f66f.
//
// Solidity: function get_virtual_price_from_lp_token(address _token, uint256 _handler_id) view returns(uint256)
func (_Curveregistry *CurveregistryCaller) GetVirtualPriceFromLpToken0(opts *bind.CallOpts, _token common.Address, _handler_id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "get_virtual_price_from_lp_token0", _token, _handler_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPriceFromLpToken0 is a free data retrieval call binding the contract method 0x4460f66f.
//
// Solidity: function get_virtual_price_from_lp_token(address _token, uint256 _handler_id) view returns(uint256)
func (_Curveregistry *CurveregistrySession) GetVirtualPriceFromLpToken0(_token common.Address, _handler_id *big.Int) (*big.Int, error) {
	return _Curveregistry.Contract.GetVirtualPriceFromLpToken0(&_Curveregistry.CallOpts, _token, _handler_id)
}

// GetVirtualPriceFromLpToken0 is a free data retrieval call binding the contract method 0x4460f66f.
//
// Solidity: function get_virtual_price_from_lp_token(address _token, uint256 _handler_id) view returns(uint256)
func (_Curveregistry *CurveregistryCallerSession) GetVirtualPriceFromLpToken0(_token common.Address, _handler_id *big.Int) (*big.Int, error) {
	return _Curveregistry.Contract.GetVirtualPriceFromLpToken0(&_Curveregistry.CallOpts, _token, _handler_id)
}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_Curveregistry *CurveregistryCaller) IsMeta(opts *bind.CallOpts, _pool common.Address) (bool, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "is_meta", _pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_Curveregistry *CurveregistrySession) IsMeta(_pool common.Address) (bool, error) {
	return _Curveregistry.Contract.IsMeta(&_Curveregistry.CallOpts, _pool)
}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_Curveregistry *CurveregistryCallerSession) IsMeta(_pool common.Address) (bool, error) {
	return _Curveregistry.Contract.IsMeta(&_Curveregistry.CallOpts, _pool)
}

// IsMeta0 is a free data retrieval call binding the contract method 0xa64485a1.
//
// Solidity: function is_meta(address _pool, uint256 _handler_id) view returns(bool)
func (_Curveregistry *CurveregistryCaller) IsMeta0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) (bool, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "is_meta0", _pool, _handler_id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMeta0 is a free data retrieval call binding the contract method 0xa64485a1.
//
// Solidity: function is_meta(address _pool, uint256 _handler_id) view returns(bool)
func (_Curveregistry *CurveregistrySession) IsMeta0(_pool common.Address, _handler_id *big.Int) (bool, error) {
	return _Curveregistry.Contract.IsMeta0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// IsMeta0 is a free data retrieval call binding the contract method 0xa64485a1.
//
// Solidity: function is_meta(address _pool, uint256 _handler_id) view returns(bool)
func (_Curveregistry *CurveregistryCallerSession) IsMeta0(_pool common.Address, _handler_id *big.Int) (bool, error) {
	return _Curveregistry.Contract.IsMeta0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// IsRegistered is a free data retrieval call binding the contract method 0x619ea806.
//
// Solidity: function is_registered(address _pool) view returns(bool)
func (_Curveregistry *CurveregistryCaller) IsRegistered(opts *bind.CallOpts, _pool common.Address) (bool, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "is_registered", _pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0x619ea806.
//
// Solidity: function is_registered(address _pool) view returns(bool)
func (_Curveregistry *CurveregistrySession) IsRegistered(_pool common.Address) (bool, error) {
	return _Curveregistry.Contract.IsRegistered(&_Curveregistry.CallOpts, _pool)
}

// IsRegistered is a free data retrieval call binding the contract method 0x619ea806.
//
// Solidity: function is_registered(address _pool) view returns(bool)
func (_Curveregistry *CurveregistryCallerSession) IsRegistered(_pool common.Address) (bool, error) {
	return _Curveregistry.Contract.IsRegistered(&_Curveregistry.CallOpts, _pool)
}

// IsRegistered0 is a free data retrieval call binding the contract method 0xc9c4f3ec.
//
// Solidity: function is_registered(address _pool, uint256 _handler_id) view returns(bool)
func (_Curveregistry *CurveregistryCaller) IsRegistered0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) (bool, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "is_registered0", _pool, _handler_id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered0 is a free data retrieval call binding the contract method 0xc9c4f3ec.
//
// Solidity: function is_registered(address _pool, uint256 _handler_id) view returns(bool)
func (_Curveregistry *CurveregistrySession) IsRegistered0(_pool common.Address, _handler_id *big.Int) (bool, error) {
	return _Curveregistry.Contract.IsRegistered0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// IsRegistered0 is a free data retrieval call binding the contract method 0xc9c4f3ec.
//
// Solidity: function is_registered(address _pool, uint256 _handler_id) view returns(bool)
func (_Curveregistry *CurveregistryCallerSession) IsRegistered0(_pool common.Address, _handler_id *big.Int) (bool, error) {
	return _Curveregistry.Contract.IsRegistered0(&_Curveregistry.CallOpts, _pool, _handler_id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Curveregistry *CurveregistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Curveregistry *CurveregistrySession) Owner() (common.Address, error) {
	return _Curveregistry.Contract.Owner(&_Curveregistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Curveregistry *CurveregistryCallerSession) Owner() (common.Address, error) {
	return _Curveregistry.Contract.Owner(&_Curveregistry.CallOpts)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_Curveregistry *CurveregistryCaller) PoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "pool_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_Curveregistry *CurveregistrySession) PoolCount() (*big.Int, error) {
	return _Curveregistry.Contract.PoolCount(&_Curveregistry.CallOpts)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_Curveregistry *CurveregistryCallerSession) PoolCount() (*big.Int, error) {
	return _Curveregistry.Contract.PoolCount(&_Curveregistry.CallOpts)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 _index) view returns(address)
func (_Curveregistry *CurveregistryCaller) PoolList(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "pool_list", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 _index) view returns(address)
func (_Curveregistry *CurveregistrySession) PoolList(_index *big.Int) (common.Address, error) {
	return _Curveregistry.Contract.PoolList(&_Curveregistry.CallOpts, _index)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 _index) view returns(address)
func (_Curveregistry *CurveregistryCallerSession) PoolList(_index *big.Int) (common.Address, error) {
	return _Curveregistry.Contract.PoolList(&_Curveregistry.CallOpts, _index)
}

// RegistryLength is a free data retrieval call binding the contract method 0x083297d2.
//
// Solidity: function registry_length() view returns(uint256)
func (_Curveregistry *CurveregistryCaller) RegistryLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curveregistry.contract.Call(opts, &out, "registry_length")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegistryLength is a free data retrieval call binding the contract method 0x083297d2.
//
// Solidity: function registry_length() view returns(uint256)
func (_Curveregistry *CurveregistrySession) RegistryLength() (*big.Int, error) {
	return _Curveregistry.Contract.RegistryLength(&_Curveregistry.CallOpts)
}

// RegistryLength is a free data retrieval call binding the contract method 0x083297d2.
//
// Solidity: function registry_length() view returns(uint256)
func (_Curveregistry *CurveregistryCallerSession) RegistryLength() (*big.Int, error) {
	return _Curveregistry.Contract.RegistryLength(&_Curveregistry.CallOpts)
}

// AddRegistryHandler is a paid mutator transaction binding the contract method 0x22f595c8.
//
// Solidity: function add_registry_handler(address _registry_handler) returns()
func (_Curveregistry *CurveregistryTransactor) AddRegistryHandler(opts *bind.TransactOpts, _registry_handler common.Address) (*types.Transaction, error) {
	return _Curveregistry.contract.Transact(opts, "add_registry_handler", _registry_handler)
}

// AddRegistryHandler is a paid mutator transaction binding the contract method 0x22f595c8.
//
// Solidity: function add_registry_handler(address _registry_handler) returns()
func (_Curveregistry *CurveregistrySession) AddRegistryHandler(_registry_handler common.Address) (*types.Transaction, error) {
	return _Curveregistry.Contract.AddRegistryHandler(&_Curveregistry.TransactOpts, _registry_handler)
}

// AddRegistryHandler is a paid mutator transaction binding the contract method 0x22f595c8.
//
// Solidity: function add_registry_handler(address _registry_handler) returns()
func (_Curveregistry *CurveregistryTransactorSession) AddRegistryHandler(_registry_handler common.Address) (*types.Transaction, error) {
	return _Curveregistry.Contract.AddRegistryHandler(&_Curveregistry.TransactOpts, _registry_handler)
}

// UpdateRegistryHandler is a paid mutator transaction binding the contract method 0x0dbdc9ff.
//
// Solidity: function update_registry_handler(uint256 _index, address _registry_handler) returns()
func (_Curveregistry *CurveregistryTransactor) UpdateRegistryHandler(opts *bind.TransactOpts, _index *big.Int, _registry_handler common.Address) (*types.Transaction, error) {
	return _Curveregistry.contract.Transact(opts, "update_registry_handler", _index, _registry_handler)
}

// UpdateRegistryHandler is a paid mutator transaction binding the contract method 0x0dbdc9ff.
//
// Solidity: function update_registry_handler(uint256 _index, address _registry_handler) returns()
func (_Curveregistry *CurveregistrySession) UpdateRegistryHandler(_index *big.Int, _registry_handler common.Address) (*types.Transaction, error) {
	return _Curveregistry.Contract.UpdateRegistryHandler(&_Curveregistry.TransactOpts, _index, _registry_handler)
}

// UpdateRegistryHandler is a paid mutator transaction binding the contract method 0x0dbdc9ff.
//
// Solidity: function update_registry_handler(uint256 _index, address _registry_handler) returns()
func (_Curveregistry *CurveregistryTransactorSession) UpdateRegistryHandler(_index *big.Int, _registry_handler common.Address) (*types.Transaction, error) {
	return _Curveregistry.Contract.UpdateRegistryHandler(&_Curveregistry.TransactOpts, _index, _registry_handler)
}

// CurveregistryCommitNewAdminIterator is returned from FilterCommitNewAdmin and is used to iterate over the raw logs and unpacked data for CommitNewAdmin events raised by the Curveregistry contract.
type CurveregistryCommitNewAdminIterator struct {
	Event *CurveregistryCommitNewAdmin // Event containing the contract specifics and raw log

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
func (it *CurveregistryCommitNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveregistryCommitNewAdmin)
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
		it.Event = new(CurveregistryCommitNewAdmin)
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
func (it *CurveregistryCommitNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveregistryCommitNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveregistryCommitNewAdmin represents a CommitNewAdmin event raised by the Curveregistry contract.
type CurveregistryCommitNewAdmin struct {
	Deadline *big.Int
	Admin    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewAdmin is a free log retrieval operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Curveregistry *CurveregistryFilterer) FilterCommitNewAdmin(opts *bind.FilterOpts, deadline []*big.Int, admin []common.Address) (*CurveregistryCommitNewAdminIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Curveregistry.contract.FilterLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &CurveregistryCommitNewAdminIterator{contract: _Curveregistry.contract, event: "CommitNewAdmin", logs: logs, sub: sub}, nil
}

// WatchCommitNewAdmin is a free log subscription operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Curveregistry *CurveregistryFilterer) WatchCommitNewAdmin(opts *bind.WatchOpts, sink chan<- *CurveregistryCommitNewAdmin, deadline []*big.Int, admin []common.Address) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Curveregistry.contract.WatchLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveregistryCommitNewAdmin)
				if err := _Curveregistry.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
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

// ParseCommitNewAdmin is a log parse operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Curveregistry *CurveregistryFilterer) ParseCommitNewAdmin(log types.Log) (*CurveregistryCommitNewAdmin, error) {
	event := new(CurveregistryCommitNewAdmin)
	if err := _Curveregistry.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveregistryNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Curveregistry contract.
type CurveregistryNewAdminIterator struct {
	Event *CurveregistryNewAdmin // Event containing the contract specifics and raw log

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
func (it *CurveregistryNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveregistryNewAdmin)
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
		it.Event = new(CurveregistryNewAdmin)
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
func (it *CurveregistryNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveregistryNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveregistryNewAdmin represents a NewAdmin event raised by the Curveregistry contract.
type CurveregistryNewAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Curveregistry *CurveregistryFilterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address) (*CurveregistryNewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Curveregistry.contract.FilterLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return &CurveregistryNewAdminIterator{contract: _Curveregistry.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Curveregistry *CurveregistryFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *CurveregistryNewAdmin, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Curveregistry.contract.WatchLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveregistryNewAdmin)
				if err := _Curveregistry.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Curveregistry *CurveregistryFilterer) ParseNewAdmin(log types.Log) (*CurveregistryNewAdmin, error) {
	event := new(CurveregistryNewAdmin)
	if err := _Curveregistry.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
