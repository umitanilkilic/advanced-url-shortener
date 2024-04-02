package model

import (
	"time"

	"gorm.io/gorm"
)

type ShortURL struct {
	gorm.Model
	CreatedBy uint `json:"created_by"`
	Active    bool `gorm:"default:true" json:"active"`
	//TODO: Change UrlID to uint
	UrlID     uint32    `gorm:"uniqueIndex;not null;" json:"url_id"`
	Alias     string    `gorm:"uniqueIndex;size:255;" json:"alias"`
	Name      string    `gorm:"not null;" json:"name"`
	Long      string    `gorm:"not null;" json:"long"`
	ExpiresAt time.Time `json:"expires_at"`
}
