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

	"gopkg.in/cenkalti/backoff.v2"
	"gopkg.in/mgo.v2"
)

// Account represents an account with debt.
type Account struct {
	address string
	borrows map[string]*big.Int
}

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

// TokenBorrow represents a borrow event.
type TokenBorrow interface {
	GetBorrower() common.Address
	GetBorrowAmount() *big.Int
	GetAccountBorrows() *big.Int
	GetTotalBorrows() *big.Int
}

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

			var iter TokenBorrowIterator

			// An operation that may fail.
			operation := func() error {
				iter, err = token.FilterBorrowEvents(filterOptions)

				if err != nil {
					return err
				}

				return nil
			}

			err = backoff.Retry(operation, backoff.NewExponentialBackOff())

			if err != nil {
				log.Fatalf("Failed to FilterBorrowEvents for token %#v: %#v", tokenSymbol, err)
			}

			if iter != nil {
				fmt.Printf("Parsing accounts...\n")
				for i := 0; iter.Next(); i++ {
					borrowEvent := iter.GetEvent()
					if borrowEvent != nil {
						address := borrowEvent.GetBorrower().Hex()
						borrows := borrowEvent.GetAccountBorrows()
						account, ok := accounts[address]
						if !ok && borrows.Cmp(big.NewInt(0)) == 1 {
							account = &Account{
								address: address,
								borrows: make(map[string]*big.Int),
							}
							account.borrows[tokenSymbol] = borrows
							accounts[account.address] = account
							// fmt.Printf("Added account: %#v. Borrowed %#v (%#v)\n", account.address, account.borrows[tokenSymbol], tokenSymbol)
						} else if ok && borrows.Cmp(big.NewInt(0)) == 1 {
							account.borrows[tokenSymbol] = borrows
							// fmt.Printf("Updated account: %#v. Borrowed %#v (%#v)\n", account.address, account.borrows[tokenSymbol], tokenSymbol)
						} else if ok && borrows.Cmp(big.NewInt(0)) < 1 {
							account.borrows[tokenSymbol] = borrows
							// check if all token borrows for the account are 0.
							accountEmpty := true
							for _, value := range account.borrows {
								if value.Cmp(big.NewInt(0)) == 1 {
									accountEmpty = false
								}
							}
							if accountEmpty {
								delete(accounts, address)
								fmt.Printf("Deleted account: %#v. Balance: %#v (%#v)\n", account.address, borrows, tokenSymbol)
							}
						}
						// fmt.Printf("account: %#v\n", account)
						// fmt.Printf("account: %T\n", account)
						// fmt.Printf("Borrow[%d]: Borrower: %#v BorrowAmount: %#v AccountBorrows: %#v TotalBorrows: %#v\n", i, borrowEvent.GetBorrower().Hex(), borrowEvent.GetBorrowAmount(), borrowEvent.GetAccountBorrows(), borrowEvent.GetTotalBorrows())
					}
				}

			}
		}
	}

	fmt.Printf("len(accounts): %#v\n", len(accounts))

	// for _, account := range accounts {
	// 	fmt.Printf("Account(%#v)\n", account.address)
	// 	for tokenSymbol, tokenBorrow := range account.borrows {
	// 		fmt.Printf("Borrow: %#v (%#v)\n", tokenSymbol, tokenBorrow)
	// 	}
	// }

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
