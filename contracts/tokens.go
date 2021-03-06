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

// TokenBorrow represents a borrow event.
type TokenBorrow interface {
	GetBorrower() common.Address
	GetBorrowAmount() *big.Int
	GetAccountBorrows() *big.Int
	GetTotalBorrows() *big.Int
	GetBlockNumber() uint64
}

// TokenBorrowIterator provides a mechanism to iterate over a token's Borrow events.
type TokenBorrowIterator interface {
	Next() bool
	GetEvent() TokenBorrow
}

// TokensProvider provides a mechanism to initialize token contracts.
type TokensProvider interface {
	GetTokens() map[string]Token
	GetAddresses() map[string]common.Address
}

// TokenContracts maintains token contract state.
type TokenContracts struct {
	ethClient *ethclient.Client
	tokens    map[string]Token
}

// MockToken is used for testing.
type MockToken struct {
	TokenBorrowIterator TokenBorrowIterator
}

// MockTokenContracts maintains token contract state.
type MockTokenContracts struct {
	Contracts map[string]Token
	Addresses map[string]common.Address
}

// MockTokenBorrowIterator provides a mechanism to iterate over a token's Borrow events.
type MockTokenBorrowIterator struct {
	BorrowEvents []TokenBorrow
	Index        int
}

// TokenAddresses contains Compound Token addresses.
var tokenAddresses = map[string]common.Address{
	CBATSymbol:  common.HexToAddress("0x6c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e"),
	CDAISymbol:  common.HexToAddress("0x5d3a536e4d6dbd6114cc1ead35777bab948e3643"),
	CETHSymbol:  common.HexToAddress("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5"),
	CREPSymbol:  common.HexToAddress("0x158079ee67fce2f58472a96584a73c7ab9ac95c1"),
	CSAISymbol:  common.HexToAddress("0xf5dce57282a584d2746faf1593d3121fcac444dc"),
	CUSDCSymbol: common.HexToAddress("0x39AA39c021dfbaE8faC545936693aC917d5E7563"),
	CWBTCSymbol: common.HexToAddress("0xc11b1268c1a384e55c48c2391d8d480264a3a7f4"),
	CZRXSymbol:  common.HexToAddress("0xb3319f5d18bc0d84dd1b4825dcde5d5f7266d407"),
}

// NewToken creates a new token contract.
func NewToken(tokenSymbol string, ethClient *ethclient.Client, logger *log.Logger) (Token, error) {
	var token Token
	var err error

	switch tokenSymbol {
	case CBATSymbol:
		token, err = NewCBAT(tokenAddresses[CBATSymbol], ethClient)
	case CDAISymbol:
		token, err = NewCDAI(tokenAddresses[CDAISymbol], ethClient)
	case CETHSymbol:
		token, err = NewCETH(tokenAddresses[CETHSymbol], ethClient)
	case CREPSymbol:
		token, err = NewCREP(tokenAddresses[CREPSymbol], ethClient)
	case CSAISymbol:
		token, err = NewCSAI(tokenAddresses[CSAISymbol], ethClient)
	case CUSDCSymbol:
		token, err = NewCUSDC(tokenAddresses[CUSDCSymbol], ethClient)
	case CWBTCSymbol:
		token, err = NewCWBTC(tokenAddresses[CWBTCSymbol], ethClient)
	case CZRXSymbol:
		token, err = NewCZRX(tokenAddresses[CZRXSymbol], ethClient)
	}

	if err != nil {
		logger.Printf("Failed to instantiate %#v Token contract: %#v", tokenSymbol, err)
	}

	return token, err
}

// NewTokenContracts creates a new TokenContracts.
func NewTokenContracts(ethClient *ethclient.Client, logger *log.Logger, tokenFactory func(string, *ethclient.Client, *log.Logger) (Token, error)) (TokensProvider, error) {
	tokens := make(map[string]Token)
	tokenContracts := &TokenContracts{ethClient: ethClient, tokens: tokens}
	for tokenSymbol := range tokenAddresses {
		token, err := tokenFactory(tokenSymbol, ethClient, logger)
		if err != nil {
			return nil, err
		}
		if token != nil {
			name, err := token.Name(nil)
			if err != nil {
				logger.Fatalf("Failed to retrieve %#v token name: %#v", tokenSymbol, err)
			}
			tokens[tokenSymbol] = token
			logger.Printf("Initialized token %#v (%#v)\n", name, tokenSymbol)
		}
	}
	logger.Printf("Initialized tokens collection: %#v\n", tokens)
	return tokenContracts, nil
}

// GetTokens returns token contracts.
func (c *TokenContracts) GetTokens() map[string]Token {
	return c.tokens
}

// GetAddresses returns token contract addresses.
func (c *TokenContracts) GetAddresses() map[string]common.Address {
	return tokenAddresses
}

// GetTokens returns token contracts.
func (c *MockTokenContracts) GetTokens() map[string]Token {
	return c.Contracts
}

// GetAddresses returns token contract addresses.
func (c *MockTokenContracts) GetAddresses() map[string]common.Address {
	return c.Addresses
}

// Name returns the token name
func (t *MockToken) Name(opts *bind.CallOpts) (string, error) {
	return "MockToken", nil
}

// FilterBorrowEvents returns the borrow events.
func (t *MockToken) FilterBorrowEvents(opts *bind.FilterOpts) (TokenBorrowIterator, error) {
	return t.TokenBorrowIterator, nil
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (i *MockTokenBorrowIterator) Next() bool {
	i.Index++
	return !(i.Index > len(i.BorrowEvents))
}

// GetEvent returns the Event containing the contract specifics and raw log.
func (i *MockTokenBorrowIterator) GetEvent() TokenBorrow {
	if i.BorrowEvents == nil || len(i.BorrowEvents) == 0 {
		return nil
	}
	return i.BorrowEvents[i.Index-1]
}
