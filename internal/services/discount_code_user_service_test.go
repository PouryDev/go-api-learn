package services_test

import (
	"github.com/bxcodec/faker/v3"
	"testing"
	"video/internal/models"
	"video/internal/services"
)

var tdPhone string
var dcus *services.DiscountCodeUserService
var tCodeID uint

func init() {
	tdPhone = faker.Phonenumber()
	dcus = services.NewDiscountCodeUserService()
}

func TestDiscountCodeUserService_Create(t *testing.T) {
	svc := services.NewDiscountCodeService()
	code := faker.Timestamp()
	codeType := models.DiscountCodeType(models.GIFT_CODE)
	dc, err := svc.Create(code, 10, 1000000, codeType)
	if err != nil {
		t.Error(err)
	}

	tCodeID = dc.ID
	err = dcus.Create(tCodeID, tdPhone)
	if err != nil {
		t.Error(err)
	}
}

func TestDiscountCodeUserService_CheckCodeForUserOnTrue(t *testing.T) {
	check, err := dcus.CheckCodeForUser(0, tdPhone)
	if err != nil {
		t.Error(err)
	}

	if !check {
		t.Error("Somehow error is nil and check is false!")
	}
}

func TestDiscountCodeUserService_CheckCodeForUserOnFalse(t *testing.T) {
	check, err := dcus.CheckCodeForUser(tCodeID, tdPhone)
	if err == nil {
		t.Error("No error on failure check")
	}

	if check {
		t.Error("Invalid true check response")
	}
}
