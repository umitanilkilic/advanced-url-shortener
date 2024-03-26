package model

import (
	"time"

	"gorm.io/gorm"
)

type ShortURL struct {
	gorm.Model
	CreatedBy uint `json:"created_by"`
	Active    bool `gorm:"default:true" json:"active"`
	//Todo: Change UrlID to uint
	UrlID     string    `gorm:"uniqueIndex;not null;size:255;" json:"url_id"`
	Name      string    `gorm:"not null;" json:"name"`
	Long      string    `gorm:"not null;" json:"long"`
	ExpiresAt time.Time `json:"expires_at"`
	UrlStats
}
