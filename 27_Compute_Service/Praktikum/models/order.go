package models

import "time"

type Order struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Product_id int     `json:"product_id"`
	Product    Product `gorm:"foreignKey:Product_id"`
	User_id    int     `json:"user_id"`
	User       User    `gorm:"foreignKey:User_id"`
	Quantity   uint    `json:"quantity"`
	Total      uint    `json:"total"`
	Created_at time.Time
	Updated_at time.Time
}
