package model

import "gorm.io/gorm"

type ProductOrder struct {
	gorm.Model
	CartID   string    `json:"cart_id"`
	Quantity int       `json:"quantity"`
	Products []Product `json:"products"`
}
