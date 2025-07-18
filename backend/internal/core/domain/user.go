package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `json:"username" gorm:"unique"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	GoogleID   string `json:"google_id,omitempty"`
	Avatar     string `json:"avatar,omitempty"`
	Provider   string `json:"provider,omitempty"`
	IsVerified bool   `json:"is_verified"`
}
