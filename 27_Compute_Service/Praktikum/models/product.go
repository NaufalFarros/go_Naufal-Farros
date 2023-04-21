package models

import "time"

type Product struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Serial_number string `json:"serial_number"`
	Created_at    time.Time
	Updated_at    time.Time
}
