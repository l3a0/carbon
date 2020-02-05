package contracts

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
)

// func TestTokenContracts_GetTokens(t *testing.T) {
// 	type fields struct {
// 		ethClient *ethclient.Client
// 		tokens    map[string]c.Token
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   map[string]c.Token
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := &c.TokenContracts{
// 				ethClient: tt.fields.ethClient,
// 				tokens:    tt.fields.tokens,
// 			}
// 			if got := c.GetTokens(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("TokenContracts.GetTokens() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestNewTokenContracts(t *testing.T) {
	// Arrange
	mockEthClient, _ := ethclient.Dial("https://mainnet.infura.io")
	type args struct {
		ethClient *ethclient.Client
	}
	tests := []struct {
		name    string
		args    args
		want    TokensProvider
		wantErr bool
	}{
		{
			name: "Should create TokensProvider.",
			args: args{
				ethClient: mockEthClient,
			},
			want: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			got, err := NewTokenContracts(tt.args.ethClient)
			// Assert
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTokenContracts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenContracts() = %v, want not %v", got, tt.want)
			}
		})
	}
}
