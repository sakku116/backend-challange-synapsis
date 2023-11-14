package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        string `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	SessionID string `json:"session_id"`
	Carts     []Cart `json:"carts"`
}
