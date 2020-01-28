// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ComptrollerABI is the input ABI used to generate the binding from.
const ComptrollerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isComptroller\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"unitroller\",\"type\":\"address\"}],\"name\":\"_become\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cToken\",\"type\":\"address\"},{\"name\":\"payer\",\"type\":\"address\"},{\"name\":\"borrower\",\"type\":\"address\"},{\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"name\":\"borrowerIndex\",\"type\":\"uint256\"}],\"name\":\"repayBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cToken\",\"type\":\"address\"},{\"name\":\"payer\",\"type\":\"address\"},{\"name\":\"borrower\",\"type\":\"address\"},{\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pauseGuardian\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setSeizePaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCloseFactor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cToken\",\"type\":\"address\"},{\"name\":\"minter\",\"type\":\"address\"},{\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"mintVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"name\":\"liquidator\",\"type\":\"address\"},{\"name\":\"borrower\",\"type\":\"address\"},{\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"liquidationIncentiveMantissa\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cToken\",\"type\":\"address\"},{\"name\":\"minter\",\"type\":\"address\"},{\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"_setLiquidationIncentive\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cToken\",\"type\":\"address\"},{\"name\":\"redeemer\",\"type\":\"address\"},{\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"_setPriceOracle\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setBorrowPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cToken\",\"type\":\"address\"},{\"name\":\"borrower\",\"type\":\"address\"},{\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintGuardianPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLiquidity\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"_setPauseGuardian\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"name\":\"liquidator\",\"type\":\"address\"},{\"name\":\"borrower\",\"type\":\"address\"},{\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cToken\",\"type\":\"address\"},{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"name\":\"liquidator\",\"type\":\"address\"},{\"name\":\"borrower\",\"type\":\"address\"},{\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transferGuardianPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"markets\",\"outputs\":[{\"name\":\"isListed\",\"type\":\"bool\"},{\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setTransferPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"checkMembership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxAssets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowGuardianPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setMintPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"_supportMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAssetsIn\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seizeGuardianPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptrollerImplementation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cToken\",\"type\":\"address\"},{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cTokens\",\"type\":\"address[]\"}],\"name\":\"enterMarkets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateCalculateSeizeTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"name\":\"liquidator\",\"type\":\"address\"},{\"name\":\"borrower\",\"type\":\"address\"},{\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newMaxAssets\",\"type\":\"uint256\"}],\"name\":\"_setMaxAssets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cToken\",\"type\":\"address\"},{\"name\":\"borrower\",\"type\":\"address\"},{\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountAssets\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingComptrollerImplementation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cToken\",\"type\":\"address\"},{\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCollateralFactor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closeFactorMantissa\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cToken\",\"type\":\"address\"},{\"name\":\"redeemer\",\"type\":\"address\"},{\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cTokenAddress\",\"type\":\"address\"}],\"name\":\"exitMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"MarketListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketExited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldCloseFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCloseFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldCollateralFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCollateralFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldLiquidationIncentiveMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"NewLiquidationIncentive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldMaxAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newMaxAssets\",\"type\":\"uint256\"}],\"name\":\"NewMaxAssets\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldPriceOracle\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"NewPriceOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldPauseGuardian\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"NewPauseGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"}]"

// Comptroller is an auto generated Go binding around an Ethereum contract.
type Comptroller struct {
	ComptrollerCaller     // Read-only binding to the contract
	ComptrollerTransactor // Write-only binding to the contract
	ComptrollerFilterer   // Log filterer for contract events
}

// ComptrollerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComptrollerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComptrollerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComptrollerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComptrollerSession struct {
	Contract     *Comptroller      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ComptrollerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComptrollerCallerSession struct {
	Contract *ComptrollerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ComptrollerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComptrollerTransactorSession struct {
	Contract     *ComptrollerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ComptrollerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComptrollerRaw struct {
	Contract *Comptroller // Generic contract binding to access the raw methods on
}

// ComptrollerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComptrollerCallerRaw struct {
	Contract *ComptrollerCaller // Generic read-only contract binding to access the raw methods on
}

// ComptrollerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComptrollerTransactorRaw struct {
	Contract *ComptrollerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComptroller creates a new instance of Comptroller, bound to a specific deployed contract.
func NewComptroller(address common.Address, backend bind.ContractBackend) (*Comptroller, error) {
	contract, err := bindComptroller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Comptroller{ComptrollerCaller: ComptrollerCaller{contract: contract}, ComptrollerTransactor: ComptrollerTransactor{contract: contract}, ComptrollerFilterer: ComptrollerFilterer{contract: contract}}, nil
}

// NewComptrollerCaller creates a new read-only instance of Comptroller, bound to a specific deployed contract.
func NewComptrollerCaller(address common.Address, caller bind.ContractCaller) (*ComptrollerCaller, error) {
	contract, err := bindComptroller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComptrollerCaller{contract: contract}, nil
}

// NewComptrollerTransactor creates a new write-only instance of Comptroller, bound to a specific deployed contract.
func NewComptrollerTransactor(address common.Address, transactor bind.ContractTransactor) (*ComptrollerTransactor, error) {
	contract, err := bindComptroller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComptrollerTransactor{contract: contract}, nil
}

// NewComptrollerFilterer creates a new log filterer instance of Comptroller, bound to a specific deployed contract.
func NewComptrollerFilterer(address common.Address, filterer bind.ContractFilterer) (*ComptrollerFilterer, error) {
	contract, err := bindComptroller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComptrollerFilterer{contract: contract}, nil
}

// bindComptroller binds a generic wrapper to an already deployed contract.
func bindComptroller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ComptrollerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Comptroller *ComptrollerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Comptroller.Contract.ComptrollerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Comptroller *ComptrollerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comptroller.Contract.ComptrollerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Comptroller *ComptrollerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Comptroller.Contract.ComptrollerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Comptroller *ComptrollerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Comptroller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Comptroller *ComptrollerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comptroller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Comptroller *ComptrollerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Comptroller.Contract.contract.Transact(opts, method, params...)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) constant returns(address)
func (_Comptroller *ComptrollerCaller) AccountAssets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Comptroller.contract.Call(opts, out, "accountAssets", arg0, arg1)
	return *ret0, err
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) constant returns(address)
func (_Comptroller *ComptrollerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Comptroller.Contract.AccountAssets(&_Comptroller.CallOpts, arg0, arg1)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) constant returns(address)
func (_Comptroller *ComptrollerCallerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Comptroller.Contract.AccountAssets(&_Comptroller.CallOpts, arg0, arg1)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Comptroller *ComptrollerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Comptroller.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Comptroller *ComptrollerSession) Admin() (common.Address, error) {
	return _Comptroller.Contract.Admin(&_Comptroller.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Comptroller *ComptrollerCallerSession) Admin() (common.Address, error) {
	return _Comptroller.Contract.Admin(&_Comptroller.CallOpts)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x9530f644.
//
// Solidity: function borrowGuardianPaused() constant returns(bool)
func (_Comptroller *ComptrollerCaller) BorrowGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Comptroller.contract.Call(opts, out, "borrowGuardianPaused")
	return *ret0, err
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x9530f644.
//
// Solidity: function borrowGuardianPaused() constant returns(bool)
func (_Comptroller *ComptrollerSession) BorrowGuardianPaused() (bool, error) {
	return _Comptroller.Contract.BorrowGuardianPaused(&_Comptroller.CallOpts)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x9530f644.
//
// Solidity: function borrowGuardianPaused() constant returns(bool)
func (_Comptroller *ComptrollerCallerSession) BorrowGuardianPaused() (bool, error) {
	return _Comptroller.Contract.BorrowGuardianPaused(&_Comptroller.CallOpts)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address cToken) constant returns(bool)
func (_Comptroller *ComptrollerCaller) CheckMembership(opts *bind.CallOpts, account common.Address, cToken common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Comptroller.contract.Call(opts, out, "checkMembership", account, cToken)
	return *ret0, err
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address cToken) constant returns(bool)
func (_Comptroller *ComptrollerSession) CheckMembership(account common.Address, cToken common.Address) (bool, error) {
	return _Comptroller.Contract.CheckMembership(&_Comptroller.CallOpts, account, cToken)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address cToken) constant returns(bool)
func (_Comptroller *ComptrollerCallerSession) CheckMembership(account common.Address, cToken common.Address) (bool, error) {
	return _Comptroller.Contract.CheckMembership(&_Comptroller.CallOpts, account, cToken)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() constant returns(uint256)
func (_Comptroller *ComptrollerCaller) CloseFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Comptroller.contract.Call(opts, out, "closeFactorMantissa")
	return *ret0, err
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() constant returns(uint256)
func (_Comptroller *ComptrollerSession) CloseFactorMantissa() (*big.Int, error) {
	return _Comptroller.Contract.CloseFactorMantissa(&_Comptroller.CallOpts)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() constant returns(uint256)
func (_Comptroller *ComptrollerCallerSession) CloseFactorMantissa() (*big.Int, error) {
	return _Comptroller.Contract.CloseFactorMantissa(&_Comptroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() constant returns(address)
func (_Comptroller *ComptrollerCaller) ComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Comptroller.contract.Call(opts, out, "comptrollerImplementation")
	return *ret0, err
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() constant returns(address)
func (_Comptroller *ComptrollerSession) ComptrollerImplementation() (common.Address, error) {
	return _Comptroller.Contract.ComptrollerImplementation(&_Comptroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() constant returns(address)
func (_Comptroller *ComptrollerCallerSession) ComptrollerImplementation() (common.Address, error) {
	return _Comptroller.Contract.ComptrollerImplementation(&_Comptroller.CallOpts)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) constant returns(uint256, uint256, uint256)
func (_Comptroller *ComptrollerCaller) GetAccountLiquidity(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Comptroller.contract.Call(opts, out, "getAccountLiquidity", account)
	return *ret0, *ret1, *ret2, err
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) constant returns(uint256, uint256, uint256)
func (_Comptroller *ComptrollerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Comptroller.Contract.GetAccountLiquidity(&_Comptroller.CallOpts, account)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) constant returns(uint256, uint256, uint256)
func (_Comptroller *ComptrollerCallerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Comptroller.Contract.GetAccountLiquidity(&_Comptroller.CallOpts, account)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) constant returns(address[])
func (_Comptroller *ComptrollerCaller) GetAssetsIn(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Comptroller.contract.Call(opts, out, "getAssetsIn", account)
	return *ret0, err
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) constant returns(address[])
func (_Comptroller *ComptrollerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _Comptroller.Contract.GetAssetsIn(&_Comptroller.CallOpts, account)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) constant returns(address[])
func (_Comptroller *ComptrollerCallerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _Comptroller.Contract.GetAssetsIn(&_Comptroller.CallOpts, account)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() constant returns(bool)
func (_Comptroller *ComptrollerCaller) IsComptroller(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Comptroller.contract.Call(opts, out, "isComptroller")
	return *ret0, err
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() constant returns(bool)
func (_Comptroller *ComptrollerSession) IsComptroller() (bool, error) {
	return _Comptroller.Contract.IsComptroller(&_Comptroller.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() constant returns(bool)
func (_Comptroller *ComptrollerCallerSession) IsComptroller() (bool, error) {
	return _Comptroller.Contract.IsComptroller(&_Comptroller.CallOpts)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 repayAmount) constant returns(uint256, uint256)
func (_Comptroller *ComptrollerCaller) LiquidateCalculateSeizeTokens(opts *bind.CallOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, repayAmount *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Comptroller.contract.Call(opts, out, "liquidateCalculateSeizeTokens", cTokenBorrowed, cTokenCollateral, repayAmount)
	return *ret0, *ret1, err
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 repayAmount) constant returns(uint256, uint256)
func (_Comptroller *ComptrollerSession) LiquidateCalculateSeizeTokens(cTokenBorrowed common.Address, cTokenCollateral common.Address, repayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Comptroller.Contract.LiquidateCalculateSeizeTokens(&_Comptroller.CallOpts, cTokenBorrowed, cTokenCollateral, repayAmount)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 repayAmount) constant returns(uint256, uint256)
func (_Comptroller *ComptrollerCallerSession) LiquidateCalculateSeizeTokens(cTokenBorrowed common.Address, cTokenCollateral common.Address, repayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Comptroller.Contract.LiquidateCalculateSeizeTokens(&_Comptroller.CallOpts, cTokenBorrowed, cTokenCollateral, repayAmount)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() constant returns(uint256)
func (_Comptroller *ComptrollerCaller) LiquidationIncentiveMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Comptroller.contract.Call(opts, out, "liquidationIncentiveMantissa")
	return *ret0, err
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() constant returns(uint256)
func (_Comptroller *ComptrollerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _Comptroller.Contract.LiquidationIncentiveMantissa(&_Comptroller.CallOpts)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() constant returns(uint256)
func (_Comptroller *ComptrollerCallerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _Comptroller.Contract.LiquidationIncentiveMantissa(&_Comptroller.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) constant returns(bool isListed, uint256 collateralFactorMantissa)
func (_Comptroller *ComptrollerCaller) Markets(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
}, error) {
	ret := new(struct {
		IsListed                 bool
		CollateralFactorMantissa *big.Int
	})
	out := ret
	err := _Comptroller.contract.Call(opts, out, "markets", arg0)
	return *ret, err
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) constant returns(bool isListed, uint256 collateralFactorMantissa)
func (_Comptroller *ComptrollerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
}, error) {
	return _Comptroller.Contract.Markets(&_Comptroller.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) constant returns(bool isListed, uint256 collateralFactorMantissa)
func (_Comptroller *ComptrollerCallerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
}, error) {
	return _Comptroller.Contract.Markets(&_Comptroller.CallOpts, arg0)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() constant returns(uint256)
func (_Comptroller *ComptrollerCaller) MaxAssets(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Comptroller.contract.Call(opts, out, "maxAssets")
	return *ret0, err
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() constant returns(uint256)
func (_Comptroller *ComptrollerSession) MaxAssets() (*big.Int, error) {
	return _Comptroller.Contract.MaxAssets(&_Comptroller.CallOpts)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() constant returns(uint256)
func (_Comptroller *ComptrollerCallerSession) MaxAssets() (*big.Int, error) {
	return _Comptroller.Contract.MaxAssets(&_Comptroller.CallOpts)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x5dce0515.
//
// Solidity: function mintGuardianPaused() constant returns(bool)
func (_Comptroller *ComptrollerCaller) MintGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Comptroller.contract.Call(opts, out, "mintGuardianPaused")
	return *ret0, err
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x5dce0515.
//
// Solidity: function mintGuardianPaused() constant returns(bool)
func (_Comptroller *ComptrollerSession) MintGuardianPaused() (bool, error) {
	return _Comptroller.Contract.MintGuardianPaused(&_Comptroller.CallOpts)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x5dce0515.
//
// Solidity: function mintGuardianPaused() constant returns(bool)
func (_Comptroller *ComptrollerCallerSession) MintGuardianPaused() (bool, error) {
	return _Comptroller.Contract.MintGuardianPaused(&_Comptroller.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_Comptroller *ComptrollerCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Comptroller.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_Comptroller *ComptrollerSession) Oracle() (common.Address, error) {
	return _Comptroller.Contract.Oracle(&_Comptroller.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_Comptroller *ComptrollerCallerSession) Oracle() (common.Address, error) {
	return _Comptroller.Contract.Oracle(&_Comptroller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() constant returns(address)
func (_Comptroller *ComptrollerCaller) PauseGuardian(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Comptroller.contract.Call(opts, out, "pauseGuardian")
	return *ret0, err
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() constant returns(address)
func (_Comptroller *ComptrollerSession) PauseGuardian() (common.Address, error) {
	return _Comptroller.Contract.PauseGuardian(&_Comptroller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() constant returns(address)
func (_Comptroller *ComptrollerCallerSession) PauseGuardian() (common.Address, error) {
	return _Comptroller.Contract.PauseGuardian(&_Comptroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_Comptroller *ComptrollerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Comptroller.contract.Call(opts, out, "pendingAdmin")
	return *ret0, err
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_Comptroller *ComptrollerSession) PendingAdmin() (common.Address, error) {
	return _Comptroller.Contract.PendingAdmin(&_Comptroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_Comptroller *ComptrollerCallerSession) PendingAdmin() (common.Address, error) {
	return _Comptroller.Contract.PendingAdmin(&_Comptroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() constant returns(address)
func (_Comptroller *ComptrollerCaller) PendingComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Comptroller.contract.Call(opts, out, "pendingComptrollerImplementation")
	return *ret0, err
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() constant returns(address)
func (_Comptroller *ComptrollerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Comptroller.Contract.PendingComptrollerImplementation(&_Comptroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() constant returns(address)
func (_Comptroller *ComptrollerCallerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Comptroller.Contract.PendingComptrollerImplementation(&_Comptroller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() constant returns(bool)
func (_Comptroller *ComptrollerCaller) SeizeGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Comptroller.contract.Call(opts, out, "seizeGuardianPaused")
	return *ret0, err
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() constant returns(bool)
func (_Comptroller *ComptrollerSession) SeizeGuardianPaused() (bool, error) {
	return _Comptroller.Contract.SeizeGuardianPaused(&_Comptroller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() constant returns(bool)
func (_Comptroller *ComptrollerCallerSession) SeizeGuardianPaused() (bool, error) {
	return _Comptroller.Contract.SeizeGuardianPaused(&_Comptroller.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() constant returns(bool)
func (_Comptroller *ComptrollerCaller) TransferGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Comptroller.contract.Call(opts, out, "transferGuardianPaused")
	return *ret0, err
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() constant returns(bool)
func (_Comptroller *ComptrollerSession) TransferGuardianPaused() (bool, error) {
	return _Comptroller.Contract.TransferGuardianPaused(&_Comptroller.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() constant returns(bool)
func (_Comptroller *ComptrollerCallerSession) TransferGuardianPaused() (bool, error) {
	return _Comptroller.Contract.TransferGuardianPaused(&_Comptroller.CallOpts)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Comptroller *ComptrollerTransactor) Become(opts *bind.TransactOpts, unitroller common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_become", unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Comptroller *ComptrollerSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.Become(&_Comptroller.TransactOpts, unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Comptroller *ComptrollerTransactorSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.Become(&_Comptroller.TransactOpts, unitroller)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x56133fc8.
//
// Solidity: function _setBorrowPaused(bool state) returns(bool)
func (_Comptroller *ComptrollerTransactor) SetBorrowPaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setBorrowPaused", state)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x56133fc8.
//
// Solidity: function _setBorrowPaused(bool state) returns(bool)
func (_Comptroller *ComptrollerSession) SetBorrowPaused(state bool) (*types.Transaction, error) {
	return _Comptroller.Contract.SetBorrowPaused(&_Comptroller.TransactOpts, state)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x56133fc8.
//
// Solidity: function _setBorrowPaused(bool state) returns(bool)
func (_Comptroller *ComptrollerTransactorSession) SetBorrowPaused(state bool) (*types.Transaction, error) {
	return _Comptroller.Contract.SetBorrowPaused(&_Comptroller.TransactOpts, state)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetCloseFactor(opts *bind.TransactOpts, newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setCloseFactor", newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Comptroller *ComptrollerSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetCloseFactor(&_Comptroller.TransactOpts, newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetCloseFactor(&_Comptroller.TransactOpts, newCloseFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address cToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetCollateralFactor(opts *bind.TransactOpts, cToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setCollateralFactor", cToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address cToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Comptroller *ComptrollerSession) SetCollateralFactor(cToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetCollateralFactor(&_Comptroller.TransactOpts, cToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address cToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetCollateralFactor(cToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetCollateralFactor(&_Comptroller.TransactOpts, cToken, newCollateralFactorMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetLiquidationIncentive(opts *bind.TransactOpts, newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setLiquidationIncentive", newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Comptroller *ComptrollerSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetLiquidationIncentive(&_Comptroller.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetLiquidationIncentive(&_Comptroller.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetMaxAssets is a paid mutator transaction binding the contract method 0xd9226ced.
//
// Solidity: function _setMaxAssets(uint256 newMaxAssets) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetMaxAssets(opts *bind.TransactOpts, newMaxAssets *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setMaxAssets", newMaxAssets)
}

// SetMaxAssets is a paid mutator transaction binding the contract method 0xd9226ced.
//
// Solidity: function _setMaxAssets(uint256 newMaxAssets) returns(uint256)
func (_Comptroller *ComptrollerSession) SetMaxAssets(newMaxAssets *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetMaxAssets(&_Comptroller.TransactOpts, newMaxAssets)
}

// SetMaxAssets is a paid mutator transaction binding the contract method 0xd9226ced.
//
// Solidity: function _setMaxAssets(uint256 newMaxAssets) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetMaxAssets(newMaxAssets *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetMaxAssets(&_Comptroller.TransactOpts, newMaxAssets)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x9845f280.
//
// Solidity: function _setMintPaused(bool state) returns(bool)
func (_Comptroller *ComptrollerTransactor) SetMintPaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setMintPaused", state)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x9845f280.
//
// Solidity: function _setMintPaused(bool state) returns(bool)
func (_Comptroller *ComptrollerSession) SetMintPaused(state bool) (*types.Transaction, error) {
	return _Comptroller.Contract.SetMintPaused(&_Comptroller.TransactOpts, state)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x9845f280.
//
// Solidity: function _setMintPaused(bool state) returns(bool)
func (_Comptroller *ComptrollerTransactorSession) SetMintPaused(state bool) (*types.Transaction, error) {
	return _Comptroller.Contract.SetMintPaused(&_Comptroller.TransactOpts, state)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetPauseGuardian(opts *bind.TransactOpts, newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setPauseGuardian", newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Comptroller *ComptrollerSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetPauseGuardian(&_Comptroller.TransactOpts, newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetPauseGuardian(&_Comptroller.TransactOpts, newPauseGuardian)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetPriceOracle(opts *bind.TransactOpts, newOracle common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setPriceOracle", newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Comptroller *ComptrollerSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetPriceOracle(&_Comptroller.TransactOpts, newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetPriceOracle(&_Comptroller.TransactOpts, newOracle)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool state) returns(bool)
func (_Comptroller *ComptrollerTransactor) SetSeizePaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setSeizePaused", state)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool state) returns(bool)
func (_Comptroller *ComptrollerSession) SetSeizePaused(state bool) (*types.Transaction, error) {
	return _Comptroller.Contract.SetSeizePaused(&_Comptroller.TransactOpts, state)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool state) returns(bool)
func (_Comptroller *ComptrollerTransactorSession) SetSeizePaused(state bool) (*types.Transaction, error) {
	return _Comptroller.Contract.SetSeizePaused(&_Comptroller.TransactOpts, state)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool state) returns(bool)
func (_Comptroller *ComptrollerTransactor) SetTransferPaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setTransferPaused", state)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool state) returns(bool)
func (_Comptroller *ComptrollerSession) SetTransferPaused(state bool) (*types.Transaction, error) {
	return _Comptroller.Contract.SetTransferPaused(&_Comptroller.TransactOpts, state)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool state) returns(bool)
func (_Comptroller *ComptrollerTransactorSession) SetTransferPaused(state bool) (*types.Transaction, error) {
	return _Comptroller.Contract.SetTransferPaused(&_Comptroller.TransactOpts, state)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address cToken) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SupportMarket(opts *bind.TransactOpts, cToken common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_supportMarket", cToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address cToken) returns(uint256)
func (_Comptroller *ComptrollerSession) SupportMarket(cToken common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SupportMarket(&_Comptroller.TransactOpts, cToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address cToken) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SupportMarket(cToken common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SupportMarket(&_Comptroller.TransactOpts, cToken)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactor) BorrowAllowed(opts *bind.TransactOpts, cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "borrowAllowed", cToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Comptroller *ComptrollerSession) BorrowAllowed(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.BorrowAllowed(&_Comptroller.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) BorrowAllowed(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.BorrowAllowed(&_Comptroller.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_Comptroller *ComptrollerTransactor) BorrowVerify(opts *bind.TransactOpts, cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "borrowVerify", cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_Comptroller *ComptrollerSession) BorrowVerify(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.BorrowVerify(&_Comptroller.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_Comptroller *ComptrollerTransactorSession) BorrowVerify(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.BorrowVerify(&_Comptroller.TransactOpts, cToken, borrower, borrowAmount)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_Comptroller *ComptrollerTransactor) EnterMarkets(opts *bind.TransactOpts, cTokens []common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "enterMarkets", cTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_Comptroller *ComptrollerSession) EnterMarkets(cTokens []common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.EnterMarkets(&_Comptroller.TransactOpts, cTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_Comptroller *ComptrollerTransactorSession) EnterMarkets(cTokens []common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.EnterMarkets(&_Comptroller.TransactOpts, cTokens)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cTokenAddress) returns(uint256)
func (_Comptroller *ComptrollerTransactor) ExitMarket(opts *bind.TransactOpts, cTokenAddress common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "exitMarket", cTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cTokenAddress) returns(uint256)
func (_Comptroller *ComptrollerSession) ExitMarket(cTokenAddress common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.ExitMarket(&_Comptroller.TransactOpts, cTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cTokenAddress) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) ExitMarket(cTokenAddress common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.ExitMarket(&_Comptroller.TransactOpts, cTokenAddress)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactor) LiquidateBorrowAllowed(opts *bind.TransactOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "liquidateBorrowAllowed", cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Comptroller *ComptrollerSession) LiquidateBorrowAllowed(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.LiquidateBorrowAllowed(&_Comptroller.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) LiquidateBorrowAllowed(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.LiquidateBorrowAllowed(&_Comptroller.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount, uint256 seizeTokens) returns()
func (_Comptroller *ComptrollerTransactor) LiquidateBorrowVerify(opts *bind.TransactOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "liquidateBorrowVerify", cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount, uint256 seizeTokens) returns()
func (_Comptroller *ComptrollerSession) LiquidateBorrowVerify(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.LiquidateBorrowVerify(&_Comptroller.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount, uint256 seizeTokens) returns()
func (_Comptroller *ComptrollerTransactorSession) LiquidateBorrowVerify(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.LiquidateBorrowVerify(&_Comptroller.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount, seizeTokens)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactor) MintAllowed(opts *bind.TransactOpts, cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "mintAllowed", cToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_Comptroller *ComptrollerSession) MintAllowed(cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.MintAllowed(&_Comptroller.TransactOpts, cToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) MintAllowed(cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.MintAllowed(&_Comptroller.TransactOpts, cToken, minter, mintAmount)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 mintAmount, uint256 mintTokens) returns()
func (_Comptroller *ComptrollerTransactor) MintVerify(opts *bind.TransactOpts, cToken common.Address, minter common.Address, mintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "mintVerify", cToken, minter, mintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 mintAmount, uint256 mintTokens) returns()
func (_Comptroller *ComptrollerSession) MintVerify(cToken common.Address, minter common.Address, mintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.MintVerify(&_Comptroller.TransactOpts, cToken, minter, mintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 mintAmount, uint256 mintTokens) returns()
func (_Comptroller *ComptrollerTransactorSession) MintVerify(cToken common.Address, minter common.Address, mintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.MintVerify(&_Comptroller.TransactOpts, cToken, minter, mintAmount, mintTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Comptroller *ComptrollerTransactor) RedeemAllowed(opts *bind.TransactOpts, cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "redeemAllowed", cToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Comptroller *ComptrollerSession) RedeemAllowed(cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RedeemAllowed(&_Comptroller.TransactOpts, cToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) RedeemAllowed(cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RedeemAllowed(&_Comptroller.TransactOpts, cToken, redeemer, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Comptroller *ComptrollerTransactor) RedeemVerify(opts *bind.TransactOpts, cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "redeemVerify", cToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Comptroller *ComptrollerSession) RedeemVerify(cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RedeemVerify(&_Comptroller.TransactOpts, cToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Comptroller *ComptrollerTransactorSession) RedeemVerify(cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RedeemVerify(&_Comptroller.TransactOpts, cToken, redeemer, redeemAmount, redeemTokens)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactor) RepayBorrowAllowed(opts *bind.TransactOpts, cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "repayBorrowAllowed", cToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Comptroller *ComptrollerSession) RepayBorrowAllowed(cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RepayBorrowAllowed(&_Comptroller.TransactOpts, cToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) RepayBorrowAllowed(cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RepayBorrowAllowed(&_Comptroller.TransactOpts, cToken, payer, borrower, repayAmount)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 repayAmount, uint256 borrowerIndex) returns()
func (_Comptroller *ComptrollerTransactor) RepayBorrowVerify(opts *bind.TransactOpts, cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "repayBorrowVerify", cToken, payer, borrower, repayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 repayAmount, uint256 borrowerIndex) returns()
func (_Comptroller *ComptrollerSession) RepayBorrowVerify(cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RepayBorrowVerify(&_Comptroller.TransactOpts, cToken, payer, borrower, repayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 repayAmount, uint256 borrowerIndex) returns()
func (_Comptroller *ComptrollerTransactorSession) RepayBorrowVerify(cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RepayBorrowVerify(&_Comptroller.TransactOpts, cToken, payer, borrower, repayAmount, borrowerIndex)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SeizeAllowed(opts *bind.TransactOpts, cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "seizeAllowed", cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Comptroller *ComptrollerSession) SeizeAllowed(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SeizeAllowed(&_Comptroller.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SeizeAllowed(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SeizeAllowed(&_Comptroller.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Comptroller *ComptrollerTransactor) SeizeVerify(opts *bind.TransactOpts, cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "seizeVerify", cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Comptroller *ComptrollerSession) SeizeVerify(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SeizeVerify(&_Comptroller.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Comptroller *ComptrollerTransactorSession) SeizeVerify(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SeizeVerify(&_Comptroller.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Comptroller *ComptrollerTransactor) TransferAllowed(opts *bind.TransactOpts, cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "transferAllowed", cToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Comptroller *ComptrollerSession) TransferAllowed(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.TransferAllowed(&_Comptroller.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) TransferAllowed(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.TransferAllowed(&_Comptroller.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_Comptroller *ComptrollerTransactor) TransferVerify(opts *bind.TransactOpts, cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "transferVerify", cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_Comptroller *ComptrollerSession) TransferVerify(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.TransferVerify(&_Comptroller.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_Comptroller *ComptrollerTransactorSession) TransferVerify(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.TransferVerify(&_Comptroller.TransactOpts, cToken, src, dst, transferTokens)
}

// ComptrollerActionPausedIterator is returned from FilterActionPaused and is used to iterate over the raw logs and unpacked data for ActionPaused events raised by the Comptroller contract.
type ComptrollerActionPausedIterator struct {
	Event *ComptrollerActionPaused // Event containing the contract specifics and raw log

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
func (it *ComptrollerActionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerActionPaused)
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
		it.Event = new(ComptrollerActionPaused)
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
func (it *ComptrollerActionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerActionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerActionPaused represents a ActionPaused event raised by the Comptroller contract.
type ComptrollerActionPaused struct {
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused is a free log retrieval operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Comptroller *ComptrollerFilterer) FilterActionPaused(opts *bind.FilterOpts) (*ComptrollerActionPausedIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return &ComptrollerActionPausedIterator{contract: _Comptroller.contract, event: "ActionPaused", logs: logs, sub: sub}, nil
}

// WatchActionPaused is a free log subscription operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Comptroller *ComptrollerFilterer) WatchActionPaused(opts *bind.WatchOpts, sink chan<- *ComptrollerActionPaused) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerActionPaused)
				if err := _Comptroller.contract.UnpackLog(event, "ActionPaused", log); err != nil {
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

// ParseActionPaused is a log parse operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Comptroller *ComptrollerFilterer) ParseActionPaused(log types.Log) (*ComptrollerActionPaused, error) {
	event := new(ComptrollerActionPaused)
	if err := _Comptroller.contract.UnpackLog(event, "ActionPaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Comptroller contract.
type ComptrollerFailureIterator struct {
	Event *ComptrollerFailure // Event containing the contract specifics and raw log

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
func (it *ComptrollerFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerFailure)
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
		it.Event = new(ComptrollerFailure)
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
func (it *ComptrollerFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerFailure represents a Failure event raised by the Comptroller contract.
type ComptrollerFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Comptroller *ComptrollerFilterer) FilterFailure(opts *bind.FilterOpts) (*ComptrollerFailureIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &ComptrollerFailureIterator{contract: _Comptroller.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Comptroller *ComptrollerFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *ComptrollerFailure) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerFailure)
				if err := _Comptroller.contract.UnpackLog(event, "Failure", log); err != nil {
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

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Comptroller *ComptrollerFilterer) ParseFailure(log types.Log) (*ComptrollerFailure, error) {
	event := new(ComptrollerFailure)
	if err := _Comptroller.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerMarketEnteredIterator is returned from FilterMarketEntered and is used to iterate over the raw logs and unpacked data for MarketEntered events raised by the Comptroller contract.
type ComptrollerMarketEnteredIterator struct {
	Event *ComptrollerMarketEntered // Event containing the contract specifics and raw log

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
func (it *ComptrollerMarketEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerMarketEntered)
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
		it.Event = new(ComptrollerMarketEntered)
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
func (it *ComptrollerMarketEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerMarketEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerMarketEntered represents a MarketEntered event raised by the Comptroller contract.
type ComptrollerMarketEntered struct {
	CToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketEntered is a free log retrieval operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address cToken, address account)
func (_Comptroller *ComptrollerFilterer) FilterMarketEntered(opts *bind.FilterOpts) (*ComptrollerMarketEnteredIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return &ComptrollerMarketEnteredIterator{contract: _Comptroller.contract, event: "MarketEntered", logs: logs, sub: sub}, nil
}

// WatchMarketEntered is a free log subscription operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address cToken, address account)
func (_Comptroller *ComptrollerFilterer) WatchMarketEntered(opts *bind.WatchOpts, sink chan<- *ComptrollerMarketEntered) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerMarketEntered)
				if err := _Comptroller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
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

// ParseMarketEntered is a log parse operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address cToken, address account)
func (_Comptroller *ComptrollerFilterer) ParseMarketEntered(log types.Log) (*ComptrollerMarketEntered, error) {
	event := new(ComptrollerMarketEntered)
	if err := _Comptroller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerMarketExitedIterator is returned from FilterMarketExited and is used to iterate over the raw logs and unpacked data for MarketExited events raised by the Comptroller contract.
type ComptrollerMarketExitedIterator struct {
	Event *ComptrollerMarketExited // Event containing the contract specifics and raw log

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
func (it *ComptrollerMarketExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerMarketExited)
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
		it.Event = new(ComptrollerMarketExited)
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
func (it *ComptrollerMarketExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerMarketExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerMarketExited represents a MarketExited event raised by the Comptroller contract.
type ComptrollerMarketExited struct {
	CToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketExited is a free log retrieval operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address cToken, address account)
func (_Comptroller *ComptrollerFilterer) FilterMarketExited(opts *bind.FilterOpts) (*ComptrollerMarketExitedIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return &ComptrollerMarketExitedIterator{contract: _Comptroller.contract, event: "MarketExited", logs: logs, sub: sub}, nil
}

// WatchMarketExited is a free log subscription operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address cToken, address account)
func (_Comptroller *ComptrollerFilterer) WatchMarketExited(opts *bind.WatchOpts, sink chan<- *ComptrollerMarketExited) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerMarketExited)
				if err := _Comptroller.contract.UnpackLog(event, "MarketExited", log); err != nil {
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

// ParseMarketExited is a log parse operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address cToken, address account)
func (_Comptroller *ComptrollerFilterer) ParseMarketExited(log types.Log) (*ComptrollerMarketExited, error) {
	event := new(ComptrollerMarketExited)
	if err := _Comptroller.contract.UnpackLog(event, "MarketExited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerMarketListedIterator is returned from FilterMarketListed and is used to iterate over the raw logs and unpacked data for MarketListed events raised by the Comptroller contract.
type ComptrollerMarketListedIterator struct {
	Event *ComptrollerMarketListed // Event containing the contract specifics and raw log

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
func (it *ComptrollerMarketListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerMarketListed)
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
		it.Event = new(ComptrollerMarketListed)
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
func (it *ComptrollerMarketListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerMarketListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerMarketListed represents a MarketListed event raised by the Comptroller contract.
type ComptrollerMarketListed struct {
	CToken common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketListed is a free log retrieval operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address cToken)
func (_Comptroller *ComptrollerFilterer) FilterMarketListed(opts *bind.FilterOpts) (*ComptrollerMarketListedIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return &ComptrollerMarketListedIterator{contract: _Comptroller.contract, event: "MarketListed", logs: logs, sub: sub}, nil
}

// WatchMarketListed is a free log subscription operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address cToken)
func (_Comptroller *ComptrollerFilterer) WatchMarketListed(opts *bind.WatchOpts, sink chan<- *ComptrollerMarketListed) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerMarketListed)
				if err := _Comptroller.contract.UnpackLog(event, "MarketListed", log); err != nil {
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

// ParseMarketListed is a log parse operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address cToken)
func (_Comptroller *ComptrollerFilterer) ParseMarketListed(log types.Log) (*ComptrollerMarketListed, error) {
	event := new(ComptrollerMarketListed)
	if err := _Comptroller.contract.UnpackLog(event, "MarketListed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerNewCloseFactorIterator is returned from FilterNewCloseFactor and is used to iterate over the raw logs and unpacked data for NewCloseFactor events raised by the Comptroller contract.
type ComptrollerNewCloseFactorIterator struct {
	Event *ComptrollerNewCloseFactor // Event containing the contract specifics and raw log

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
func (it *ComptrollerNewCloseFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewCloseFactor)
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
		it.Event = new(ComptrollerNewCloseFactor)
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
func (it *ComptrollerNewCloseFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewCloseFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewCloseFactor represents a NewCloseFactor event raised by the Comptroller contract.
type ComptrollerNewCloseFactor struct {
	OldCloseFactorMantissa *big.Int
	NewCloseFactorMantissa *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewCloseFactor is a free log retrieval operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Comptroller *ComptrollerFilterer) FilterNewCloseFactor(opts *bind.FilterOpts) (*ComptrollerNewCloseFactorIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewCloseFactorIterator{contract: _Comptroller.contract, event: "NewCloseFactor", logs: logs, sub: sub}, nil
}

// WatchNewCloseFactor is a free log subscription operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Comptroller *ComptrollerFilterer) WatchNewCloseFactor(opts *bind.WatchOpts, sink chan<- *ComptrollerNewCloseFactor) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewCloseFactor)
				if err := _Comptroller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
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

// ParseNewCloseFactor is a log parse operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Comptroller *ComptrollerFilterer) ParseNewCloseFactor(log types.Log) (*ComptrollerNewCloseFactor, error) {
	event := new(ComptrollerNewCloseFactor)
	if err := _Comptroller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerNewCollateralFactorIterator is returned from FilterNewCollateralFactor and is used to iterate over the raw logs and unpacked data for NewCollateralFactor events raised by the Comptroller contract.
type ComptrollerNewCollateralFactorIterator struct {
	Event *ComptrollerNewCollateralFactor // Event containing the contract specifics and raw log

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
func (it *ComptrollerNewCollateralFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewCollateralFactor)
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
		it.Event = new(ComptrollerNewCollateralFactor)
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
func (it *ComptrollerNewCollateralFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewCollateralFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewCollateralFactor represents a NewCollateralFactor event raised by the Comptroller contract.
type ComptrollerNewCollateralFactor struct {
	CToken                      common.Address
	OldCollateralFactorMantissa *big.Int
	NewCollateralFactorMantissa *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterNewCollateralFactor is a free log retrieval operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address cToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Comptroller *ComptrollerFilterer) FilterNewCollateralFactor(opts *bind.FilterOpts) (*ComptrollerNewCollateralFactorIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewCollateralFactorIterator{contract: _Comptroller.contract, event: "NewCollateralFactor", logs: logs, sub: sub}, nil
}

// WatchNewCollateralFactor is a free log subscription operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address cToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Comptroller *ComptrollerFilterer) WatchNewCollateralFactor(opts *bind.WatchOpts, sink chan<- *ComptrollerNewCollateralFactor) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewCollateralFactor)
				if err := _Comptroller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
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

// ParseNewCollateralFactor is a log parse operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address cToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Comptroller *ComptrollerFilterer) ParseNewCollateralFactor(log types.Log) (*ComptrollerNewCollateralFactor, error) {
	event := new(ComptrollerNewCollateralFactor)
	if err := _Comptroller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerNewLiquidationIncentiveIterator is returned from FilterNewLiquidationIncentive and is used to iterate over the raw logs and unpacked data for NewLiquidationIncentive events raised by the Comptroller contract.
type ComptrollerNewLiquidationIncentiveIterator struct {
	Event *ComptrollerNewLiquidationIncentive // Event containing the contract specifics and raw log

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
func (it *ComptrollerNewLiquidationIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewLiquidationIncentive)
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
		it.Event = new(ComptrollerNewLiquidationIncentive)
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
func (it *ComptrollerNewLiquidationIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewLiquidationIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewLiquidationIncentive represents a NewLiquidationIncentive event raised by the Comptroller contract.
type ComptrollerNewLiquidationIncentive struct {
	OldLiquidationIncentiveMantissa *big.Int
	NewLiquidationIncentiveMantissa *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterNewLiquidationIncentive is a free log retrieval operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Comptroller *ComptrollerFilterer) FilterNewLiquidationIncentive(opts *bind.FilterOpts) (*ComptrollerNewLiquidationIncentiveIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewLiquidationIncentiveIterator{contract: _Comptroller.contract, event: "NewLiquidationIncentive", logs: logs, sub: sub}, nil
}

// WatchNewLiquidationIncentive is a free log subscription operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Comptroller *ComptrollerFilterer) WatchNewLiquidationIncentive(opts *bind.WatchOpts, sink chan<- *ComptrollerNewLiquidationIncentive) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewLiquidationIncentive)
				if err := _Comptroller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
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

// ParseNewLiquidationIncentive is a log parse operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Comptroller *ComptrollerFilterer) ParseNewLiquidationIncentive(log types.Log) (*ComptrollerNewLiquidationIncentive, error) {
	event := new(ComptrollerNewLiquidationIncentive)
	if err := _Comptroller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerNewMaxAssetsIterator is returned from FilterNewMaxAssets and is used to iterate over the raw logs and unpacked data for NewMaxAssets events raised by the Comptroller contract.
type ComptrollerNewMaxAssetsIterator struct {
	Event *ComptrollerNewMaxAssets // Event containing the contract specifics and raw log

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
func (it *ComptrollerNewMaxAssetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewMaxAssets)
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
		it.Event = new(ComptrollerNewMaxAssets)
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
func (it *ComptrollerNewMaxAssetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewMaxAssetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewMaxAssets represents a NewMaxAssets event raised by the Comptroller contract.
type ComptrollerNewMaxAssets struct {
	OldMaxAssets *big.Int
	NewMaxAssets *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewMaxAssets is a free log retrieval operation binding the contract event 0x7093cf1eb653f749c3ff531d6df7f92764536a7fa0d13530cd26e070780c32ea.
//
// Solidity: event NewMaxAssets(uint256 oldMaxAssets, uint256 newMaxAssets)
func (_Comptroller *ComptrollerFilterer) FilterNewMaxAssets(opts *bind.FilterOpts) (*ComptrollerNewMaxAssetsIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewMaxAssets")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewMaxAssetsIterator{contract: _Comptroller.contract, event: "NewMaxAssets", logs: logs, sub: sub}, nil
}

// WatchNewMaxAssets is a free log subscription operation binding the contract event 0x7093cf1eb653f749c3ff531d6df7f92764536a7fa0d13530cd26e070780c32ea.
//
// Solidity: event NewMaxAssets(uint256 oldMaxAssets, uint256 newMaxAssets)
func (_Comptroller *ComptrollerFilterer) WatchNewMaxAssets(opts *bind.WatchOpts, sink chan<- *ComptrollerNewMaxAssets) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewMaxAssets")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewMaxAssets)
				if err := _Comptroller.contract.UnpackLog(event, "NewMaxAssets", log); err != nil {
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

// ParseNewMaxAssets is a log parse operation binding the contract event 0x7093cf1eb653f749c3ff531d6df7f92764536a7fa0d13530cd26e070780c32ea.
//
// Solidity: event NewMaxAssets(uint256 oldMaxAssets, uint256 newMaxAssets)
func (_Comptroller *ComptrollerFilterer) ParseNewMaxAssets(log types.Log) (*ComptrollerNewMaxAssets, error) {
	event := new(ComptrollerNewMaxAssets)
	if err := _Comptroller.contract.UnpackLog(event, "NewMaxAssets", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerNewPauseGuardianIterator is returned from FilterNewPauseGuardian and is used to iterate over the raw logs and unpacked data for NewPauseGuardian events raised by the Comptroller contract.
type ComptrollerNewPauseGuardianIterator struct {
	Event *ComptrollerNewPauseGuardian // Event containing the contract specifics and raw log

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
func (it *ComptrollerNewPauseGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewPauseGuardian)
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
		it.Event = new(ComptrollerNewPauseGuardian)
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
func (it *ComptrollerNewPauseGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewPauseGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewPauseGuardian represents a NewPauseGuardian event raised by the Comptroller contract.
type ComptrollerNewPauseGuardian struct {
	OldPauseGuardian common.Address
	NewPauseGuardian common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPauseGuardian is a free log retrieval operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Comptroller *ComptrollerFilterer) FilterNewPauseGuardian(opts *bind.FilterOpts) (*ComptrollerNewPauseGuardianIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewPauseGuardianIterator{contract: _Comptroller.contract, event: "NewPauseGuardian", logs: logs, sub: sub}, nil
}

// WatchNewPauseGuardian is a free log subscription operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Comptroller *ComptrollerFilterer) WatchNewPauseGuardian(opts *bind.WatchOpts, sink chan<- *ComptrollerNewPauseGuardian) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewPauseGuardian)
				if err := _Comptroller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
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

// ParseNewPauseGuardian is a log parse operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Comptroller *ComptrollerFilterer) ParseNewPauseGuardian(log types.Log) (*ComptrollerNewPauseGuardian, error) {
	event := new(ComptrollerNewPauseGuardian)
	if err := _Comptroller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerNewPriceOracleIterator is returned from FilterNewPriceOracle and is used to iterate over the raw logs and unpacked data for NewPriceOracle events raised by the Comptroller contract.
type ComptrollerNewPriceOracleIterator struct {
	Event *ComptrollerNewPriceOracle // Event containing the contract specifics and raw log

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
func (it *ComptrollerNewPriceOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewPriceOracle)
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
		it.Event = new(ComptrollerNewPriceOracle)
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
func (it *ComptrollerNewPriceOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewPriceOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewPriceOracle represents a NewPriceOracle event raised by the Comptroller contract.
type ComptrollerNewPriceOracle struct {
	OldPriceOracle common.Address
	NewPriceOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewPriceOracle is a free log retrieval operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Comptroller *ComptrollerFilterer) FilterNewPriceOracle(opts *bind.FilterOpts) (*ComptrollerNewPriceOracleIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewPriceOracleIterator{contract: _Comptroller.contract, event: "NewPriceOracle", logs: logs, sub: sub}, nil
}

// WatchNewPriceOracle is a free log subscription operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Comptroller *ComptrollerFilterer) WatchNewPriceOracle(opts *bind.WatchOpts, sink chan<- *ComptrollerNewPriceOracle) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewPriceOracle)
				if err := _Comptroller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
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

// ParseNewPriceOracle is a log parse operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Comptroller *ComptrollerFilterer) ParseNewPriceOracle(log types.Log) (*ComptrollerNewPriceOracle, error) {
	event := new(ComptrollerNewPriceOracle)
	if err := _Comptroller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
		return nil, err
	}
	return event, nil
}
