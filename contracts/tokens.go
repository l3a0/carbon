package contracts

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// CBATSymbol is the CBAT symbol.
	CBATSymbol = "CBAT"

	// CDAISymbol is the CDAI symbol.
	CDAISymbol = "CDAI"

	// CETHSymbol is the CETH symbol.
	CETHSymbol = "CETH"

	// CREPSymbol is the CREP symbol.
	CREPSymbol = "CREP"

	// CSAISymbol is the CSAI symbol.
	CSAISymbol = "CSAI"

	// CUSDCSymbol is the CUSDC symbol.
	CUSDCSymbol = "CUSDC"

	// CWBTCSymbol is the CWBTC symbol.
	CWBTCSymbol = "CWBTC"

	// CZRXSymbol is the CZRX symbol.
	CZRXSymbol = "CZRX"
)

// Token represents a token contract.
type Token interface {
	Name(opts *bind.CallOpts) (string, error)
	FilterBorrowEvents(opts *bind.FilterOpts) (TokenBorrowIterator, error)
}

// TokenBorrowIterator provides a mechanism to iterate over a token's Borrow events.
type TokenBorrowIterator interface {
	Next() bool
	GetEvent() TokenBorrow
}

// TokenAddresses contains Compound Token addresses.
var TokenAddresses = map[string]common.Address{
	CBATSymbol:  common.HexToAddress("0x6c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e"),
	CDAISymbol:  common.HexToAddress("0x5d3a536e4d6dbd6114cc1ead35777bab948e3643"),
	CETHSymbol:  common.HexToAddress("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5"),
	CREPSymbol:  common.HexToAddress("0x158079ee67fce2f58472a96584a73c7ab9ac95c1"),
	CSAISymbol:  common.HexToAddress("0xf5dce57282a584d2746faf1593d3121fcac444dc"),
	CUSDCSymbol: common.HexToAddress("0x39AA39c021dfbaE8faC545936693aC917d5E7563"),
	CWBTCSymbol: common.HexToAddress("0xc11b1268c1a384e55c48c2391d8d480264a3a7f4"),
	CZRXSymbol:  common.HexToAddress("0xb3319f5d18bc0d84dd1b4825dcde5d5f7266d407"),
}

// TokenBorrow represents a borrow event.
type TokenBorrow interface {
	GetBorrower() common.Address
	GetBorrowAmount() *big.Int
	GetAccountBorrows() *big.Int
	GetTotalBorrows() *big.Int
}

// NewToken creates a new token contract.
func NewToken(tokenSymbol string, ethClient *ethclient.Client) Token {
	var token Token
	var err error

	switch tokenSymbol {
	case CBATSymbol:
		token, err = NewCBAT(TokenAddresses[CBATSymbol], ethClient)
	case CDAISymbol:
		token, err = NewCDAI(TokenAddresses[CDAISymbol], ethClient)
	case CETHSymbol:
		token, err = NewCETH(TokenAddresses[CETHSymbol], ethClient)
	case CREPSymbol:
		token, err = NewCREP(TokenAddresses[CREPSymbol], ethClient)
	case CSAISymbol:
		token, err = NewCSAI(TokenAddresses[CSAISymbol], ethClient)
	case CUSDCSymbol:
		token, err = NewCUSDC(TokenAddresses[CUSDCSymbol], ethClient)
	case CWBTCSymbol:
		token, err = NewCWBTC(TokenAddresses[CWBTCSymbol], ethClient)
	case CZRXSymbol:
		token, err = NewCZRX(TokenAddresses[CZRXSymbol], ethClient)
	}

	if err != nil {
		log.Fatalf("Failed to instantiate %#v Token contract: %#v", tokenSymbol, err)
	}

	return token
}
