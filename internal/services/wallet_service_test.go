package services_test

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"testing"
	"video/internal/services"
)

var tPhone string
var ws *services.WalletService

func init() {
	tPhone = faker.Phonenumber()
	ws = services.NewWalletService()
}

func TestCreateWallet(t *testing.T) {
	wallet, err := ws.Create(tPhone)
	if err != nil {
		t.Error(fmt.Sprintf("Failed to create wallet: %s", err))
	}

	wallet, err = ws.GetByPhoneNumber(wallet.PhoneNumber)
	if err != nil {
		t.Error(fmt.Sprintf("Wallet hasn't been created but no error has been generated: %s", err))
	}
}

func TestIncreaseBalance(t *testing.T) {
	wallet, err := ws.GetByPhoneNumber(tPhone)
	if err != nil {
		t.Error(fmt.Sprintf("Could not find wallet: %s", err))
	}
	balance := wallet.Balance

	err = ws.IncreaseBalance(*wallet, 5000)
	if err != nil {
		t.Error(fmt.Sprintf("Could not increase wallet balance: %s", err))
	}
	balance += 5000

	wallet, err = ws.GetByPhoneNumber(tPhone)
	if err != nil {
		t.Error(err)
	}
	if wallet.Balance != balance {
		t.Error(fmt.Sprintf("Increased wallet balance must be %v but got %v", balance, wallet.Balance))
	}
}

func TestGetUserBalance(t *testing.T) {
	_, err := ws.GetUserBalance(tPhone)
	if err != nil {
		t.Error(fmt.Sprintf("Could not get wallet balance from db: %s", err))
	}
}
