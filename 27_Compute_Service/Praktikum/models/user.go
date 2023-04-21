package models

import (
	"time"
)

func init() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
}

type User struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	First_name string    `json:"first_name"`
	Last_name  string    `json:"last_name"`
	RoleId     int       `json:"role_id"`
	Role       Role      `json:"role" gorm:"foreignKey:RoleId"`
	Created_at time.Time `json:"created_at" gorm:"autoCreateTime:false"`
	Updated_at time.Time `json:"updated_at" gorm:"autoCreateTime:false"`
}
