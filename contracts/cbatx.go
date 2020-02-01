package contracts

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// GetBorrower returns the borrower.
func (b *CBATBorrow) GetBorrower() common.Address {
	return b.Borrower
}

// GetBorrowAmount returns the borrow amount.
func (b *CBATBorrow) GetBorrowAmount() *big.Int {
	return b.BorrowAmount
}

// GetAccountBorrows returns the account borrow amount.
func (b *CBATBorrow) GetAccountBorrows() *big.Int {
	return b.AccountBorrows
}

// GetTotalBorrows returns the contract borrow amount.
func (b *CBATBorrow) GetTotalBorrows() *big.Int {
	return b.TotalBorrows
}

// GetBlockNumber returns the block number of the event.
func (b *CBATBorrow) GetBlockNumber() uint64 {
	return b.Raw.BlockNumber
}

// FilterBorrowEvents returns the borrow events.
func (b *CBATFilterer) FilterBorrowEvents(opts *bind.FilterOpts) (TokenBorrowIterator, error) {
	iter, err := b.FilterBorrow(opts)

	if err != nil {
		log.Fatalf("Failed to call FilterBorrow: %#v", err)
	}

	return iter, err
}

// GetEvent returns the Event containing the contract specifics and raw log.
func (b *CBATBorrowIterator) GetEvent() TokenBorrow {
	return b.Event
}
