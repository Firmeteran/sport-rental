package models

import "gorm.io/gorm"

type TopUp struct {
	gorm.Model
	UserID  uint    `json:"user_id"`
	Amount  float64 `json:"amount"`
	OrderID string  `gorm:"unique" json:"order_id"`
	Status  string  `gorm:"default:'pending'" json:"status"`
	SnapURL string  `json:"snap_url"`
}
