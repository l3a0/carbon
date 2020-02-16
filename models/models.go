package models

import (
	"math/big"

	"github.com/globalsign/mgo/bson"
)

// Account represents an account with debt.
type Account struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	ShardKey string
	Address  string
	Borrows  map[string]*big.Int
}

// GetBSON marshals the account to BSON.
func (a *Account) GetBSON() (interface{}, error) {
	borrows := make(map[string]string)
	for tokenSymbol, borrow := range a.Borrows {
		borrows[tokenSymbol] = borrow.String()
	}
	bson := bson.M{"_id": a.ID, "shardkey": a.ShardKey, "address": a.Address, "borrows": borrows}
	return bson, nil
}

// SetBSON unmarshals the BSON to Account.
func (a *Account) SetBSON(raw bson.Raw) error {
	var v bson.M
	raw.Unmarshal(&v)
	a.ID = v["_id"].(bson.ObjectId)
	a.ShardKey = v["shardkey"].(string)
	a.Address = v["address"].(string)
	borrows := v["borrows"].(bson.M)
	a.Borrows = make(map[string]*big.Int)
	for tokenSymbol, borrow := range borrows {
		bi := big.NewInt(0)
		bi.SetString(borrow.(string), 10)
		a.Borrows[tokenSymbol] = bi
	}
	return nil
}
