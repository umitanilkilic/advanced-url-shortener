package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;not null;size:255;" validate:"required,email" json:"email"`
	Password string `gorm:"not null;" validate:"required,min=6,max=50" json:"password"`
	// UUID is a unique identifier for the user
	// This is used for the JWT token
	// UUID string `gorm:"uniqueIndex;not null;size:255;" json:"uuid"`
}
