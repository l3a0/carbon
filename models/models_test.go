package models

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/l3a0/carbon/contracts"
	"gopkg.in/mgo.v2/bson"
)

func TestAccount_GetBSON(t *testing.T) {
	// Arrange
	type fields struct {
		ID            bson.ObjectId
		Address       string
		Borrows       map[string]*big.Int
		StringBorrows map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Should get BSON",
			fields: fields{
				ID:      bson.NewObjectId(),
				Address: "FakeAddress",
				Borrows: map[string]*big.Int{
					contracts.CUSDCSymbol: big.NewInt(1000000000000),
				},
				StringBorrows: map[string]string{
					contracts.CUSDCSymbol: "1000000000000",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				ID:      tt.fields.ID,
				Address: tt.fields.Address,
				Borrows: tt.fields.Borrows,
			}
			// Act
			got, err := a.GetBSON()
			// Assert
			if (err != nil) != tt.wantErr {
				t.Errorf("Account.GetBSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			want := bson.M{"_id": tt.fields.ID, "address": tt.fields.Address, "borrows": tt.fields.StringBorrows}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Account.GetBSON() = %v, want %v", got, want)
			}
		})
	}
}

func TestAccount_SetBSON(t *testing.T) {
	// Arrange
	type fields struct {
		ID      bson.ObjectId
		Address string
		Borrows map[string]*big.Int
		StringBorrows map[string]string
	}
	type args struct {
		raw bson.Raw
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Should set BSON.",
			fields: fields{
				ID: bson.NewObjectId(),
				Address: "FakeAddress",
				Borrows: map[string]*big.Int{
					contracts.CUSDCSymbol: big.NewInt(1000000000000),
				},
				StringBorrows: map[string]string{
					contracts.CUSDCSymbol: "1000000000000",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := bson.M{"_id": tt.fields.ID, "address": tt.fields.Address, "borrows": tt.fields.StringBorrows}
			data, err := bson.Marshal(want)
			if err != nil {
				t.Fatal(err)
			}
			var raw bson.Raw
			err = bson.Unmarshal(data, &raw)
			if err != nil {
				t.Fatal(err)
			}
			a := &Account{}
			// Act
			if err := a.SetBSON(raw); (err != nil) != tt.wantErr {
				t.Errorf("Account.SetBSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Assert
			got, err := a.GetBSON()
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Account.GetBSON() = %v, want %v", got, want)
			}
		})
	}
}
