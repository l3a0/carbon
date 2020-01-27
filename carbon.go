package main

import (
    "crypto/tls"
    "fmt"
    "log"
    // "math"
    // "math/big"
    "net"
    "os"
    "time"

    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"

    "gopkg.in/mgo.v2"
)

type Account struct {
    address string
}

func main() {
    // client, err := ethclient.Dial("/home/l3a0/.ethereum/geth.ipc")
    client, err := ethclient.Dial("https://mainnet.infura.io")

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("we have a connection")
    // fmt.Printf("client: %#v\n", client)

    // 0x6C34811807ad578802c0122D8A36bDF48c17d12C
    // addressHex := "0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5"
    // address := common.HexToAddress(addressHex)
    // fmt.Printf("address: %v\n", address)
    // fmt.Printf("address: %+v\n", address)
    // fmt.Printf("address: %#v\n", address)
    // fmt.Printf("address: %T\n", address)
    // fmt.Println("address: ", address.Hex())
    // fmt.Println("address: ", address.Hash().Hex())
    // fmt.Println("address: ", address.Bytes())

    // balance, err := client.BalanceAt(context.Background(), address, nil)

    // if err != nil {
    //     log.Fatal(err)
    // }
    
    // fmt.Println("balance: ", balance)

    // fbalance := new(big.Float)
    // fbalance.SetString(balance.String())
    // // wei / 10^18
    // ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
    
    // fmt.Println("ethValue: ", ethValue)

    // // cUSDC smart contract address
    address := common.HexToAddress("0x39AA39c021dfbaE8faC545936693aC917d5E7563")

    // fmt.Println("address: ", address.Hex())

    // bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block

    // if err != nil {
    //     log.Fatal(err)
    // }

    // isContract := len(bytecode) > 0

    // fmt.Printf("is contract: %v\n", isContract)

    token, err := NewCUSDC(address, client)

    if err != nil {
        log.Fatalf("Failed to instantiate a Token contract: %#v", err)
    }
    
    name, err := token.Name(nil)

    if err != nil {
        log.Fatalf("Failed to retrieve token name: %#v", err)
    }
    
    fmt.Println("Token name:", name)
    // fmt.Printf("token: %#v\n", token)

    it, err := token.FilterBorrow(nil)

    if err != nil {
        log.Fatalf("Failed to call FilterBorrow: %#v", err)
    }

    // fmt.Printf("it: %#v\n", it)

    if !it.done && it.fail == nil {
        // fmt.Println("!it.done && it.fail == nil")

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
            fmt.Printf("Can't connect, go error %#v\n", err)
            os.Exit(1)
        }

        // fmt.Printf("connected to cosmos db session: %#v\n", session)

        defer session.Close()

        // get collection
        // accounts := session.DB(database).C("accounts")
        accounts := make(map[string]*Account)

        fmt.Printf("opened accounts collection: %#v\n", accounts)

        for index := 0; it.Next(); index++ {
            borrowEvent := it.Event

            account := &Account{
                address: borrowEvent.Borrower.Hex(),
            }

            // fmt.Printf("account: %#v\n", account)
            // fmt.Printf("account: %T\n", account)
            // fmt.Printf("Borrow[%d]: Borrower: %#v BorrowAmount: %#v AccountBorrows: %#v TotalBorrows: %#v\n", index, borrowEvent.Borrower.Hex(), borrowEvent.BorrowAmount, borrowEvent.AccountBorrows, borrowEvent.TotalBorrows)

            accounts[account.address] = account
        }

        fmt.Printf("len(accounts): %#v\n", len(accounts))
    }
}