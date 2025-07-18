package domain

import "gorm.io/gorm"

// User represents a user in the system
type User struct {
	gorm.Model
	Username   string `json:"username" gorm:"unique" example:"john_doe"`
	Password   string `json:"password,omitempty" example:"password123"`
	Email      string `json:"email" example:"john@example.com"`
	Role       string `json:"role" example:"employee"`
	GoogleID   string `json:"google_id,omitempty" example:"123456789"`
	Avatar     string `json:"avatar,omitempty" example:"https://example.com/avatar.jpg"`
	Provider   string `json:"provider,omitempty" example:"google"`
	IsVerified bool   `json:"is_verified" example:"false"`
}
