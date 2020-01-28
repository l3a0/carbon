package main

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// GetBorrower returns the borrower.
func (b *CUSDCBorrow) GetBorrower() common.Address {
	return b.Borrower
}

// GetBorrowAmount returns the borrow amount.
func (b *CUSDCBorrow) GetBorrowAmount() *big.Int {
	return b.BorrowAmount
}

// GetAccountBorrows returns the account borrow amount.
func (b *CUSDCBorrow) GetAccountBorrows() *big.Int {
	return b.AccountBorrows
}

// GetTotalBorrows returns the contract borrow amount.
func (b *CUSDCBorrow) GetTotalBorrows() *big.Int {
	return b.TotalBorrows
}

// FilterBorrowEvents returns the borrow events.
func (b *CUSDCFilterer) FilterBorrowEvents(opts *bind.FilterOpts) TokenBorrowIterator {
	iter, err := b.FilterBorrow(opts)

	if err != nil {
		log.Fatalf("Failed to call FilterBorrow: %#v", err)
	}

	return iter
}

// GetEvent returns the Event containing the contract specifics and raw log.
func (b *CUSDCBorrowIterator) GetEvent() TokenBorrow {
	return b.Event
}
