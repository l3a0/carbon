package models

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/l3a0/carbon/contracts"
)

// ComptrollerService is responsible for interacting with Comptroller contract.
type ComptrollerService interface {
	GetAccountLiquidity(opts *bind.CallOpts, account common.Address) (errorCode *big.Int, liquidity *big.Int, shortfall *big.Int, err error)
}

// Comptroller contains blockchain client and state.
type Comptroller struct {
	logger   *log.Logger
	address  common.Address
	contract *contracts.Comptroller
}

// MockComptroller enables unit testing.
type MockComptroller struct {
}

// NewComptrollerService creates a new ComptrollerService
func NewComptrollerService(logger *log.Logger, ethClient *ethclient.Client) ComptrollerService {
	address := common.HexToAddress("0x3d9819210a31b4961b30ef54be2aed79b9c9cd3b")
	contract, err := contracts.NewComptroller(address, ethClient)
	if err != nil {
		logger.Panicf("Failed to load comptroller: %v\n", err)
	}
	return &Comptroller{
		logger:   logger,
		address:  address,
		contract: contract,
	}
}

// GetAccountLiquidity returns the account's liquidity.
func (service *Comptroller) GetAccountLiquidity(opts *bind.CallOpts, account common.Address) (errorCode *big.Int, liquidity *big.Int, shortfall *big.Int, err error) {
	return service.contract.GetAccountLiquidity(opts, account)
}

// GetAccountLiquidity returns the account's liquidity.
func (service *MockComptroller) GetAccountLiquidity(opts *bind.CallOpts, account common.Address) (errorCode *big.Int, liquidity *big.Int, shortfall *big.Int, err error) {
	return common.Big0, common.Big1, common.Big0, nil
}
