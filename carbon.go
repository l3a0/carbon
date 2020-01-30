package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"gopkg.in/cenkalti/backoff.v2"

	"github.com/l3a0/carbon/contracts"
)

// Account represents an account with debt.
type Account struct {
	address string
	borrows map[string]*big.Int
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

	for tokenSymbol, tokenAddress := range contracts.TokenAddresses {
		fmt.Printf("Processing accounts for token %#v @ %#v\n", tokenSymbol, tokenAddress.Hex())

		token := contracts.NewToken(tokenSymbol, ethClient)

		if token != nil {
			name, err := token.Name(nil)

			if err != nil {
				log.Fatalf("Failed to retrieve %#v token name: %#v", tokenSymbol, err)
			}

			fmt.Printf("Initialized token %#v (%#v)\n", name, tokenSymbol)

			filterOptions := &bind.FilterOpts{Start: 0, End: nil, Context: nil}

			var iter contracts.TokenBorrowIterator

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

	// database := "bao-blockchain"
	// username := "bao-blockchain"
	// password := ""

	// // DialInfo holds options for establishing a session with Azure Cosmos DB for MongoDB API account.
	// dialInfo := &mgo.DialInfo{
	// 	Addrs:    []string{"bao-blockchain.mongo.cosmos.azure.com:10255"}, // Get HOST + PORT
	// 	Timeout:  10 * time.Second,
	// 	Database: database,
	// 	Username: username,
	// 	Password: password,
	// 	DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
	// 		return tls.Dial("tcp", addr.String(), &tls.Config{})
	// 	},
	// }

	// Create a session which maintains a pool of socket connections
	// session, err := mgo.DialWithInfo(dialInfo)

	// if err != nil {
	// 	log.Fatalf("Can't connect, go error %#v\n", err)
	// }

	// defer session.Close()

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
