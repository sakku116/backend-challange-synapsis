package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID        string         `json:"user_id"`
	ProductOrders []ProductOrder `json:"product_orders"`
}
