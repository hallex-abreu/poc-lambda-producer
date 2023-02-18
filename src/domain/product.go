package domain

import "time"

type Product struct {
	Sku           string     `json:"sku"`
	Description   string     `json:"description"`
	Price         float64    `json:"price"`
	StatusProcess *string    `json:"status_process"`
	DateProcess   *time.Time `json:"date_process"`
}
