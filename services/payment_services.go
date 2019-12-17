package services

import (
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/mark-iii/model"
	"time"
)

type PaymentServiceContract interface {
	FindExpiredPayment() ([]*model.Payment, error)
	UpdatePaymentStatus(payment *model.Payment, tx *gorm.DB) error
}

type paymentServicesContract struct {
	db *gorm.DB
}

func NewPaymentServicesContract(db *gorm.DB) PaymentServiceContract {
	return &paymentServicesContract{db}
}

func (srv *paymentServicesContract) FindExpiredPayment() ([]*model.Payment, error) {
	var data []*model.Payment

	if err := srv.db.Where("status = ? AND expire_date < ?", "PENDING", time.Now()).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (srv *paymentServicesContract) UpdatePaymentStatus(payment *model.Payment, tx *gorm.DB) error {
	var err error
	err = tx.Save(&payment).Error
	return err
}
