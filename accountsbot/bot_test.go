package accountsbot_test

import (
	"fmt"
	"testing"

	ab "github.com/l3a0/carbon/accountsbot"
	c "github.com/l3a0/carbon/contracts"

	"gopkg.in/mgo.v2"
)

// TestNewAccountsBot validates a new AccountsBot can be created.
func TestNewAccountsBot(t *testing.T) {
	// Arrange
	t.Log("we have a connection\n")
	botsCollection := &mgo.Collection{
		Database: nil,
		Name:     "bots",
		FullName: "db.bots",
	}
	accountsCollection := &mgo.Collection{
		Database: nil,
		Name:     "accounts",
		FullName: "db.accounts",
	}
	tokensProvider := &c.MockTokenContracts{}
	// Act
	accountsBot := ab.NewAccountsBot(botsCollection, accountsCollection, tokensProvider)

	// Assert
	got := accountsBot.String()
	want := "AccountsBot{<nil>}"
	description := fmt.Sprintf("accountsBot.String()")
	if got != want {
		t.Errorf("%s = %q, want %q", description, got, want)
	}
}
