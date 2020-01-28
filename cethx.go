package main

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// GetBorrower returns the borrower.
func (b *CETHBorrow) GetBorrower() common.Address {
	return b.Borrower
}

// GetBorrowAmount returns the borrow amount.
func (b *CETHBorrow) GetBorrowAmount() *big.Int {
	return b.BorrowAmount
}

// GetAccountBorrows returns the account borrow amount.
func (b *CETHBorrow) GetAccountBorrows() *big.Int {
	return b.AccountBorrows
}

// GetTotalBorrows returns the contract borrow amount.
func (b *CETHBorrow) GetTotalBorrows() *big.Int {
	return b.TotalBorrows
}

// FilterBorrowEvents returns the borrow events.
func (b *CETHFilterer) FilterBorrowEvents(opts *bind.FilterOpts) TokenBorrowIterator {
	iter, err := b.FilterBorrow(opts)

	if err != nil {
		log.Fatalf("Failed to call FilterBorrow: %#v", err)
	}

	return iter
}

// GetEvent returns the Event containing the contract specifics and raw log.
func (b *CETHBorrowIterator) GetEvent() TokenBorrow {
	return b.Event
}
