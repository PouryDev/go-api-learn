package services_test

import (
	"errors"
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
	"testing"
	"video/internal/models"
	"video/internal/services"
)

var code string
var svc *services.DiscountCodeService

func init() {
	code = faker.Timestamp()
	svc = services.NewDiscountCodeService()
}

func TestDiscountCodeService_Create(t *testing.T) {
	codeType := models.DiscountCodeType(models.GIFT_CODE)
	dc, err := svc.Create(code, 10, 1000000, codeType)
	if err != nil {
		t.Error(err)
	}

	_, err = svc.Find(dc.ID)
	if err != nil {
		t.Error(err)
	}
}

func TestDiscountCodeService_Update(t *testing.T) {
	dc, err := svc.GetByCode(code)
	if err != nil {
		t.Error(err)
	}

	expected := int64(50000)
	dc.Value = expected
	err = svc.Update(*dc)
	if err != nil {
		t.Error(err)
	}

	dc, err = svc.GetByCode(code)
	if err != nil {
		t.Error(err)
	}

	if dc.Value != expected {
		t.Error("Discount code failed at updating without throwing error")
	}
}

func TestDiscountCodeService_IncreaseUsage(t *testing.T) {
	dc, err := svc.GetByCode(code)
	if err != nil {
		t.Error(err)
	}

	usage := dc.MaxUsers
	err = svc.IncreaseUsage(dc)
	if err != nil {
		t.Error(err)
	}

	dc, err = svc.GetByCode(code)
	if err != nil {
		t.Error(err)
	}

	if dc.MaxUsers != usage {
		t.Error("Failed to increase max users without throwing error")
	}
}

func TestDiscountCodeService_Delete(t *testing.T) {
	dc, err := svc.GetByCode(code)
	if err != nil {
		t.Error(err)
	}

	err = svc.Delete(*dc)
	if err != nil {
		t.Error(err)
	}

	dc, err = svc.GetByCode(code)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Error(err)
		}
	}
}
