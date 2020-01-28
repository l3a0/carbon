package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"math/big"
	"net"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"gopkg.in/mgo.v2"
)

// Account represents an account with debt.
type Account struct {
	address string
}

// Token represents a token contract.
type Token interface {
	Name(opts *bind.CallOpts) (string, error)
	FilterBorrowEvents(opts *bind.FilterOpts) TokenBorrowIterator
}

// TokenBorrowIterator provides a mechanism to iterate over a token's Borrow events.
type TokenBorrowIterator interface {
	Next() bool
	GetEvent() TokenBorrow
}

// TokenBorrow represents a borrow event.
type TokenBorrow interface {
	GetBorrower() common.Address
	GetBorrowAmount() *big.Int
	GetAccountBorrows() *big.Int
	GetTotalBorrows() *big.Int
}

// CBATSymbol is the CBAT symbol.
const CBATSymbol = "CBAT"

// CDAISymbol is the CDAI symbol.
const CDAISymbol = "CDAI"

// CETHSymbol is the CETH symbol.
const CETHSymbol = "CETH"

// CREPSymbol is the CREP symbol.
const CREPSymbol = "CREP"

// CSAISymbol is the CSAI symbol.
const CSAISymbol = "CSAI"

// CUSDCSymbol is the CUSDC symbol.
const CUSDCSymbol = "CUSDC"

// CWBTCSymbol is the CWBTC symbol.
const CWBTCSymbol = "CWBTC"

// CZRXSymbol is the CZRX symbol.
const CZRXSymbol = "CZRX"

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
func NewToken(tokenSymbol string, ethClient *ethclient.Client) Token {
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
		log.Fatalf("Failed to instantiate %#v Token contract: %#v", tokenSymbol, err)
	}

	return token
}

func main() {

	// ethClient, err := ethclient.Dial("/home/l3a0/.ethereum/geth.ipc")
	ethClient, err := ethclient.Dial("https://mainnet.infura.io")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")

	accounts := make(map[string]*Account)

	fmt.Println("opened accounts collection")

	for tokenSymbol, tokenAddress := range tokenAddresses {
		fmt.Printf("Processing accounts for token %#v @ %#v\n", tokenSymbol, tokenAddress.Hex())

		token := NewToken(tokenSymbol, ethClient)

		if token != nil {
			name, err := token.Name(nil)

			if err != nil {
				log.Fatalf("Failed to retrieve %#v token name: %#v", tokenSymbol, err)
			}

			fmt.Printf("Initialized token %#v (%#v)\n", name, tokenSymbol)

			filterOptions := &bind.FilterOpts{Start: 0, End: nil, Context: nil}
			iter := token.FilterBorrowEvents(filterOptions)

			if iter != nil {
				for i := 0; iter.Next(); i++ {
					borrowEvent := iter.GetEvent()

					if borrowEvent != nil {
						account := &Account{
							address: borrowEvent.GetBorrower().Hex(),
						}

						// fmt.Printf("account: %#v\n", account)
						// fmt.Printf("account: %T\n", account)
						// fmt.Printf("Borrow[%d]: Borrower: %#v BorrowAmount: %#v AccountBorrows: %#v TotalBorrows: %#v\n", i, borrowEvent.GetBorrower().Hex(), borrowEvent.GetBorrowAmount(), borrowEvent.GetAccountBorrows(), borrowEvent.GetTotalBorrows())

						accounts[account.address] = account
					}
				}
			}
		}
	}

	fmt.Printf("len(accounts): %#v\n", len(accounts))

	database := "bao-blockchain"
	username := "bao-blockchain"
	password := ""

	// DialInfo holds options for establishing a session with Azure Cosmos DB for MongoDB API account.
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{"bao-blockchain.mongo.cosmos.azure.com:10255"}, // Get HOST + PORT
		Timeout:  10 * time.Second,
		Database: database,
		Username: username,
		Password: password,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}

	// Create a session which maintains a pool of socket connections
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Fatalf("Can't connect, go error %#v\n", err)
	}

	defer session.Close()

	// CBATToken, err := NewCBAT(tokenAddresses["CBAT"], ethClient)

	// if err != nil {
	// 	log.Fatalf("Failed to instantiate a Token contract: %#v", err)
	// }

	// name, err = CBATToken.Name(nil)

	// if err != nil {
	// 	log.Fatalf("Failed to retrieve token name: %#v", err)
	// }

	// fmt.Println("Token name:", name)

	// CBATIterator, err := CBATToken.FilterBorrow(nil)

	// if err != nil {
	// 	log.Fatalf("Failed to call FilterBorrow: %#v", err)
	// }

	// for i := 0; CBATIterator.Next(); i++ {
	// 	borrowEvent := CBATIterator.Event

	// 	account := &Account{
	// 		address: borrowEvent.Borrower.Hex(),
	// 	}

	// 	// fmt.Printf("account: %#v\n", account)
	// 	// fmt.Printf("account: %T\n", account)
	// 	// fmt.Printf("Borrow[%d]: Borrower: %#v BorrowAmount: %#v AccountBorrows: %#v TotalBorrows: %#v\n", index, borrowEvent.Borrower.Hex(), borrowEvent.BorrowAmount, borrowEvent.AccountBorrows, borrowEvent.TotalBorrows)

	// 	accounts[account.address] = account
	// }

	// fmt.Printf("len(accounts): %#v\n", len(accounts))

	// comptrollerAddress := common.HexToAddress("0x3d9819210a31b4961b30ef54be2aed79b9c9cd3b")
	// comptroller, err := NewComptroller(comptrollerAddress, ethClient)

	// if err != nil {
	//     log.Fatalf("Failed to load comptroller: %#v\n", err)
	// }

	// fmt.Printf("comptroller: %#v\n", comptroller)

	// i := 0

	// for key, element := range accounts {
	//     if i > 5 {
	//         break
	//     }

	//     fmt.Println("Key:", key, "=>", "Element:", element)

	//     error, liquidity, shortfall, err := comptroller.GetAccountLiquidity(nil, common.HexToAddress(element.address))

	//     if err != nil {
	//         log.Fatalf("Failed to GetAccountLiquidity: %#v\n", err)
	//     }

	//     fmt.Printf("error: %#v liquidity: %#v shortfall: %#v\n", error, liquidity, shortfall)
	//     i++
	// }
	// }
}
