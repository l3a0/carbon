package contracts

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
)

func TestNewTokenContracts(t *testing.T) {
	// Arrange
	// mockEthClient, _ := ethclient.Dial("https://mainnet.infura.io")
	mockEthClient, _ := ethclient.Dial("/home/l3a0/.ethereum/geth.ipc")
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
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			got, err := NewTokenContracts(tt.args.ethClient, MockNewToken)
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

func TestTokenContracts_GetTokens(t *testing.T) {
	// Arrange
	// mockEthClient, _ := ethclient.Dial("https://mainnet.infura.io")
	mockEthClient, _ := ethclient.Dial("/home/l3a0/.ethereum/geth.ipc")
	type fields struct {
		ethClient *ethclient.Client
	}
	mockToken := &MockToken{}
	tokenFactory := func(tokenSymbol string, ethClient *ethclient.Client) (Token, error) {
		var token Token = mockToken
		var err error = nil

		return token, err
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]Token
	}{
		{
			name: "Should get token contracts.",
			fields: fields{
				ethClient: mockEthClient,
			},
			want: map[string]Token{
				CBATSymbol: mockToken,
				CDAISymbol: mockToken,
				CETHSymbol: mockToken,
				CREPSymbol: mockToken,
				CSAISymbol: mockToken,
				CUSDCSymbol: mockToken,
				CWBTCSymbol: mockToken,
				CZRXSymbol: mockToken,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenContracts, err := NewTokenContracts(tt.fields.ethClient, tokenFactory)
			if err != nil {
				t.Errorf("NewTokenContracts() error = %v", err)
				return
			}
			// Act
			got := tokenContracts.GetTokens()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TokenContracts.GetTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewToken(t *testing.T) {
	// mockEthClient, _ := ethclient.Dial("https://mainnet.infura.io")
	mockEthClient, _ := ethclient.Dial("/home/l3a0/.ethereum/geth.ipc")
	type args struct {
		tokenSymbol string
		ethClient   *ethclient.Client
	}
	type want struct {
		tokenType string
		tokenName string
	}
	tests := []struct {
		name    string
		args    args
		wants   want
	}{
		{
			name: "Should create CBAT.",
			args: args{
				tokenSymbol: CBATSymbol,
				ethClient: mockEthClient,
			},
			wants: want{
				tokenType: "*contracts.CBAT",
				tokenName: "Compound Basic Attention Token",
			},
		},
		{
			name: "Should create CDAI.",
			args: args{
				tokenSymbol: CDAISymbol,
				ethClient: mockEthClient,
			},
			wants: want{
				tokenType: "*contracts.CDAI",
				tokenName: "Compound Dai",
			},
		},
		{
			name: "Should create CETH.",
			args: args{
				tokenSymbol: CETHSymbol,
				ethClient: mockEthClient,
			},
			wants: want{
				tokenType: "*contracts.CETH",
				tokenName: "Compound Ether",
			},
		},
		{
			name: "Should create CREP.",
			args: args{
				tokenSymbol: CREPSymbol,
				ethClient: mockEthClient,
			},
			wants: want{
				tokenType: "*contracts.CREP",
				tokenName: "Compound Augur",
			},
		},
		{
			name: "Should create CSAI.",
			args: args{
				tokenSymbol: CSAISymbol,
				ethClient: mockEthClient,
			},
			wants: want{
				tokenType: "*contracts.CSAI",
				tokenName: "Compound Dai",
			},
		},
		{
			name: "Should create CUSDC.",
			args: args{
				tokenSymbol: CUSDCSymbol,
				ethClient: mockEthClient,
			},
			wants: want{
				tokenType: "*contracts.CUSDC",
				tokenName: "Compound USD Coin",
			},
		},
		{
			name: "Should create CWBTC.",
			args: args{
				tokenSymbol: CWBTCSymbol,
				ethClient: mockEthClient,
			},
			wants: want{
				tokenType: "*contracts.CWBTC",
				tokenName: "Compound Wrapped BTC",
			},
		},
		{
			name: "Should create CZRX.",
			args: args{
				tokenSymbol: CZRXSymbol,
				ethClient: mockEthClient,
			},
			wants: want{
				tokenType: "*contracts.CZRX",
				tokenName: "Compound 0x",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewToken(tt.args.tokenSymbol, tt.args.ethClient)
			if (err != nil) {
				t.Fatalf("NewToken() error = %v", err)
			}
			gotType := reflect.TypeOf(got)
			if gotType.String() != tt.wants.tokenType {
				t.Errorf("gotType.String() = %v, want %v", gotType.String(), tt.wants.tokenType)
			}
			gotName, err := got.Name(nil)
			if (err != nil) {
				t.Fatalf("got.Name(nil) error = %v", err)
			}
			if gotName != tt.wants.tokenName {
				t.Errorf("got.Name(nil) = %v, want %v", gotName, tt.wants.tokenName)
			}
		})
	}
}
