package video

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"os"
	"testing"
	"video/internal/api/services"
)

var phoneNumber string

func TestMain(m *testing.M) {
	phoneNumber = faker.Phonenumber()

	code := m.Run()
	os.Exit(code)
}

func TestCreateWallet(t *testing.T) {
	ws := services.NewWalletService()
	wallet, err := ws.Create(phoneNumber)
	if err != nil {
		t.Error(fmt.Sprintf("Failed to create wallet: %s", err))
	}

	wallet, err = ws.Find(wallet.PhoneNumber)
	if err != nil {
		t.Error(fmt.Sprintf("Wallet hasn't been created but no error has been generated: %s", err))
	}
}

func TestIncreaseBalance(t *testing.T) {
	ws := services.NewWalletService()
	wallet, err := ws.Find(phoneNumber)
	if err != nil {
		t.Error(fmt.Sprintf("Could not find wallet: %s", err))
	}
	balance := wallet.Balance

	err = ws.IncreaseBalance(*wallet, 5000)
	if err != nil {
		t.Error(fmt.Sprintf("Could not increase wallet balance: %s", err))
	}
	balance += 5000

	wallet, err = ws.Find(phoneNumber)
	if err != nil {
		t.Error(err)
	}
	if wallet.Balance != balance {
		t.Error(fmt.Sprintf("Increased wallet balance must be %v but got %v", balance, wallet.Balance))
	}
}

func TestGetUserBalance(t *testing.T) {
	ws := services.NewWalletService()
	_, err := ws.GetUserBalance(phoneNumber)
	if err != nil {
		t.Error(fmt.Sprintf("Could not get wallet balance from db: %s", err))
	}
}
