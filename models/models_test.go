package models

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/globalsign/mgo/bson"
	"github.com/l3a0/carbon/contracts"
)

func TestAccount_GetBSON(t *testing.T) {
	// Arrange
	type fields struct {
		ID        bson.ObjectId
		ShardKey  string
		Address   string
		Borrows   map[string]*big.Int
		Liquidity *big.Int
		Shortfall *big.Int
	}
	tests := []struct {
		name          string
		fields        fields
		StringBorrows map[string]string
		wantErr       bool
	}{
		{
			name: "Should get BSON",
			fields: fields{
				ID:        bson.NewObjectId(),
				Address:   "FakeAddress",
				ShardKey:  "FakeShardKey",
				Liquidity: big.NewInt(10),
				Shortfall: big.NewInt(0),
				Borrows: map[string]*big.Int{
					contracts.CUSDCSymbol: big.NewInt(1000000000000),
				},
			},
			StringBorrows: map[string]string{
				contracts.CUSDCSymbol: "1000000000000",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				ID:        tt.fields.ID,
				ShardKey:  tt.fields.ShardKey,
				Address:   tt.fields.Address,
				Borrows:   tt.fields.Borrows,
				Liquidity: tt.fields.Liquidity,
				Shortfall: tt.fields.Shortfall,
			}
			// Act
			result, err := a.GetBSON()
			got := result.(bson.M)
			// Assert
			if (err != nil) != tt.wantErr {
				t.Errorf("Account.GetBSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			want := bson.M{
				"_id":       tt.fields.ID,
				"shardkey":  tt.fields.ShardKey,
				"address":   tt.fields.Address,
				"borrows":   tt.StringBorrows,
				"liquidity": tt.fields.Liquidity.String(),
				"shortfall": tt.fields.Shortfall.String(),
			}
			// if !reflect.DeepEqual(got, want) {
			if got["_id"] != want["_id"] {
				t.Errorf(`got["_id"] = %v, want %v`, got["_id"], want["_id"])
			}
			if got["shardkey"] != want["shardkey"] {
				t.Errorf(`got["shardkey"] = %v, want %v`, got["shardkey"], want["shardkey"])
			}
			if got["address"] != want["address"] {
				t.Errorf(`got["address"] = %v, want %v`, got["address"], want["address"])
			}
			gotBorrows := got["borrows"].(map[string]string)
			wantBorrows := want["borrows"].(map[string]string)
			if gotBorrows[contracts.CUSDCSymbol] != wantBorrows[contracts.CUSDCSymbol] {
				t.Errorf(`gotBorrows[contracts.CUSDCSymbol] = %v, want %v`, gotBorrows[contracts.CUSDCSymbol], wantBorrows[contracts.CUSDCSymbol])
			}
			gotLiquidity := got["liquidity"].(string)
			wantLiquidity := want["liquidity"].(string)
			if gotLiquidity != wantLiquidity {
				t.Errorf(`gotLiquidity = %v, want %v`, gotLiquidity, wantLiquidity)
			}
			gotShortfall := got["shortfall"].(string)
			wantShortfall := want["shortfall"].(string)
			if gotShortfall != wantShortfall {
				t.Errorf(`gotShortfall = %v, want %v`, gotShortfall, wantShortfall)
			}
		})
	}
}

func TestAccount_SetBSON(t *testing.T) {
	// Arrange
	type fields struct {
		ID        bson.ObjectId
		ShardKey  string
		Address   string
		Liquidity *big.Int
		Shortfall *big.Int
		Borrows   map[string]*big.Int
	}
	type args struct {
		raw bson.Raw
	}
	tests := []struct {
		name          string
		fields        fields
		StringBorrows map[string]string
		wantErr       bool
	}{
		{
			name: "Should set BSON.",
			fields: fields{
				ID:        bson.NewObjectId(),
				ShardKey:  "FakeShardKey",
				Address:   "FakeAddress",
				Liquidity: big.NewInt(10),
				Shortfall: big.NewInt(0),
				Borrows: map[string]*big.Int{
					contracts.CUSDCSymbol: big.NewInt(1000000000000),
				},
			},
			StringBorrows: map[string]string{
				contracts.CUSDCSymbol: "1000000000000",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := bson.M{
				"_id":       tt.fields.ID,
				"shardkey":  tt.fields.ShardKey,
				"address":   tt.fields.Address,
				"liquidity": tt.fields.Liquidity.String(),
				"shortfall": tt.fields.Shortfall.String(),
				"borrows":   tt.StringBorrows,
			}
			data, err := bson.Marshal(want)
			if err != nil {
				t.Fatal(err)
			}
			var raw bson.Raw
			err = bson.Unmarshal(data, &raw)
			if err != nil {
				t.Fatal(err)
			}
			account := &Account{}
			// Act
			if err := account.SetBSON(raw); (err != nil) != tt.wantErr {
				t.Errorf("Account.SetBSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Assert
			got, err := account.GetBSON()
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Account.GetBSON() = %v, want %v", got, want)
			}
		})
	}
}
