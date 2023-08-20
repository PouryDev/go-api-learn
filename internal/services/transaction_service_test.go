package services_test

import (
	"testing"
	"video/internal/models"
	"video/internal/services"
)

var ts *services.TransactionService

func init() {
	ts = services.NewTransactionService()
}

func TestTransactionService_Create(t *testing.T) {
	phone := "+989350539376"
	title := "Unit test title"
	transaction, err := ts.Create(50000, models.GATEWAY_TYPE, phone, models.OPEN_STATUS, &title)
	if err != nil {
		t.Error(err)
	}
	id := transaction.ID

	transaction, err = ts.Find(id)
	if err != nil {
		t.Error(err)
	}
}

func TestTransactionService_Pay(t *testing.T) {
	phone := "+989350539376"
	title := "Unit test title"
	transaction, err := ts.Create(50000, models.GATEWAY_TYPE, phone, models.OPEN_STATUS, &title)
	if err != nil {
		t.Error(err)
	}

	err = ts.Pay(*transaction)
	if err != nil {
		t.Error(err)
	}
	id := transaction.ID

	transaction, err = ts.Find(id)
	if err != nil {
		t.Error(err)
	}

	expected := models.PAID_STATUS
	if string(transaction.Status) != expected {
		t.Error("Transaction did not go to paid state")
	}
}
