package services_test

import (
	"errors"
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
	"testing"
	"video/internal/models"
	"video/internal/services"
)

var dCode string
var dcs *services.DiscountCodeService

func init() {
	dCode = faker.Timestamp()
	dcs = services.NewDiscountCodeService()
}

func TestDiscountCodeService_Create(t *testing.T) {
	codeType := models.DiscountCodeType(models.GIFT_CODE)
	dc, err := dcs.Create(dCode, 10, 1000000, codeType)
	if err != nil {
		t.Error(err)
	}

	_, err = dcs.Find(dc.ID)
	if err != nil {
		t.Error(err)
	}
}

func TestDiscountCodeService_Update(t *testing.T) {
	dc, err := dcs.GetByCode(dCode)
	if err != nil {
		t.Error(err)
	}

	expected := int64(50000)
	dc.Value = expected
	err = dcs.Update(*dc)
	if err != nil {
		t.Error(err)
	}

	dc, err = dcs.GetByCode(dCode)
	if err != nil {
		t.Error(err)
	}

	if dc.Value != expected {
		t.Error("Discount code failed at updating without throwing error")
	}
}

func TestDiscountCodeService_IncreaseUsage(t *testing.T) {
	dc, err := dcs.GetByCode(dCode)
	if err != nil {
		t.Error(err)
	}

	usage := dc.MaxUsers
	err = dcs.IncreaseUsage(dc)
	if err != nil {
		t.Error(err)
	}

	dc, err = dcs.GetByCode(dCode)
	if err != nil {
		t.Error(err)
	}

	if dc.MaxUsers != usage {
		t.Error("Failed to increase max users without throwing error")
	}
}

func TestDiscountCodeService_Delete(t *testing.T) {
	dc, err := dcs.GetByCode(dCode)
	if err != nil {
		t.Error(err)
	}

	err = dcs.Delete(*dc)
	if err != nil {
		t.Error(err)
	}

	dc, err = dcs.GetByCode(dCode)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Error(err)
		}
	}
}
