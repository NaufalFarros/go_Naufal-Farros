package models

import "time"

// set time indonesia
func init() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
}

type Role struct {
	ID         int       `json:"id"`
	Role       string    `json:"role"`
	Created_at time.Time `json:"created_at" gorm:"autoCreateTime:false"`
	Updated_at time.Time `json:"updated_at" gorm:"autoCreateTime:false"`
}
