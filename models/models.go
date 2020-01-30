package models

import "math/big"

// Account represents an account with debt.
type Account struct {
	Address string
	Borrows map[string]*big.Int
}
