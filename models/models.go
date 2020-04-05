package models

import (
	"math/big"

	"github.com/globalsign/mgo/bson"
)

// Account represents an account with debt.
type Account struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	ShardKey  string
	Address   string
	Borrows   map[string]*big.Int
	Liquidity *big.Int
	Shortfall *big.Int
}

// GetBSON marshals the account to BSON.
func (account *Account) GetBSON() (interface{}, error) {
	borrows := make(map[string]string)
	for tokenSymbol, borrow := range account.Borrows {
		borrows[tokenSymbol] = borrow.String()
	}
	bson := bson.M{
		"_id":       account.ID,
		"shardkey":  account.ShardKey,
		"address":   account.Address,
		"borrows":   borrows,
		"liquidity": account.Liquidity.String(),
		"shortfall": account.Shortfall.String(),
	}
	return bson, nil
}

// SetBSON unmarshals the BSON to Account.
func (account *Account) SetBSON(raw bson.Raw) error {
	var value bson.M
	raw.Unmarshal(&value)
	account.ID = value["_id"].(bson.ObjectId)
	account.ShardKey = value["shardkey"].(string)
	account.Address = value["address"].(string)
	borrows := value["borrows"].(bson.M)
	account.Borrows = make(map[string]*big.Int)
	for tokenSymbol, borrow := range borrows {
		borrowValue := big.NewInt(0)
		borrowValue.SetString(borrow.(string), 10)
		account.Borrows[tokenSymbol] = borrowValue
	}
	liquidity, ok := value["liquidity"]
	if ok {
		account.Liquidity = big.NewInt(0)
		account.Liquidity.SetString(liquidity.(string), 10)
	}
	shortfall, ok := value["shortfall"]
	if ok {
		account.Shortfall = big.NewInt(0)
		account.Shortfall.SetString(shortfall.(string), 10)
	}
	return nil
}
