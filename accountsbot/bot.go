package accountsbot

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/google/uuid"

	"gopkg.in/cenkalti/backoff.v2"

	c "github.com/l3a0/carbon/contracts"
	m "github.com/l3a0/carbon/models"
)

// AccountsBot maintains state for accounts with debt.
type AccountsBot struct {
	id        uuid.UUID
	ethClient *ethclient.Client
}

// NewAccountsBot creates a new AccountsBot.
func NewAccountsBot(id uuid.UUID, ethClient *ethclient.Client) *AccountsBot {
	return &AccountsBot{
		id:        id,
		ethClient: ethClient,
	}
}

func (b AccountsBot) String() string {
	return fmt.Sprintf("AccountsBot(%v)", b.id)
}

// Work puts the bot to work.
func (b *AccountsBot) Work(statusChannel chan int) {
	log.Printf("%v working...\n", b)

	accounts := make(map[string]*m.Account)

	log.Printf("opened accounts collection: %#v\n", accounts)

	for tokenSymbol, tokenAddress := range c.TokenAddresses {
		log.Printf("Processing accounts for token %#v @ %#v\n", tokenSymbol, tokenAddress.Hex())

		token := c.NewToken(tokenSymbol, b.ethClient)

		if token != nil {
			name, err := token.Name(nil)

			if err != nil {
				log.Fatalf("Failed to retrieve %#v token name: %#v", tokenSymbol, err)
			}

			log.Printf("Initialized token %#v (%#v)\n", name, tokenSymbol)

			filterOptions := &bind.FilterOpts{Start: 0, End: nil, Context: nil}

			var iter c.TokenBorrowIterator

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

			parseAccounts(iter, tokenSymbol, accounts)
		}
	}

	log.Printf("len(accounts): %#v\n", len(accounts))

	statusChannel <- 0
}

func parseAccounts(iter c.TokenBorrowIterator, tokenSymbol string, accounts map[string]*m.Account) {
	if iter != nil {
		log.Printf("Parsing accounts...\n")
		for i := 0; iter.Next(); i++ {
			borrowEvent := iter.GetEvent()
			if borrowEvent != nil {
				address := borrowEvent.GetBorrower().Hex()
				borrows := borrowEvent.GetAccountBorrows()
				account, ok := accounts[address]
				if !ok && borrows.Cmp(big.NewInt(0)) == 1 {
					account = &m.Account{
						Address: address,
						Borrows: make(map[string]*big.Int),
					}
					account.Borrows[tokenSymbol] = borrows
					accounts[account.Address] = account
					// fmt.Printf("Added account: %#v. Borrowed %#v (%#v)\n", account.address, account.borrows[tokenSymbol], tokenSymbol)
				} else if ok && borrows.Cmp(big.NewInt(0)) == 1 {
					account.Borrows[tokenSymbol] = borrows
					// fmt.Printf("Updated account: %#v. Borrowed %#v (%#v)\n", account.address, account.borrows[tokenSymbol], tokenSymbol)
				} else if ok && borrows.Cmp(big.NewInt(0)) < 1 {
					account.Borrows[tokenSymbol] = borrows
					// check if all token borrows for the account are 0.
					accountEmpty := true
					for _, value := range account.Borrows {
						if value.Cmp(big.NewInt(0)) == 1 {
							accountEmpty = false
						}
					}
					if accountEmpty {
						delete(accounts, address)
						log.Printf("Deleted account: %#v. Balance: %#v (%#v)\n", account.Address, borrows, tokenSymbol)
					}
				}
			}
		}
	}
}
