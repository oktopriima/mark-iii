package model

import "time"

type Payment struct {
	ID         int       `json:"id"`
	Method     string    `json:"method"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
	ExpireDate time.Time `json:"expire_date"`
}

func (p *Payment) TableName() string {
	return "payments"
}
