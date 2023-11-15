package model

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	UserID        string         `json:"user_id"`
	ProductOrders []ProductOrder `json:"product_orders"`
}
