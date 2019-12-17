package payment

import (
	"fmt"
	"github.com/oktopriima/mark-ii/conf"
	"github.com/oktopriima/mark-iii/services"
)

func UpdateExpired(cfg conf.Config) {
	db, err := conf.MysqlConnection(cfg)
	if err != nil {
		fmt.Printf("someting when wrong : %s", err.Error())
		return
	}

	paymentContract := services.NewPaymentServicesContract(db)

	data, err := paymentContract.FindExpiredPayment()
	if err != nil {
		fmt.Printf("someting when wrong : %s", err.Error())
		return
	}

	tx := db.Begin()
	defer tx.Rollback()

	for _, payment := range data {
		payment.Status = "EXPIRED"
		if err := paymentContract.UpdatePaymentStatus(payment, tx); err != nil {
			fmt.Printf("someting when wrong : %s", err.Error())
			return
		}
	}

	tx.Commit()

	return
}
