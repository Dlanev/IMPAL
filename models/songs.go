package models

import "time"

type Song struct {
	ID uint `gorm:"primaryKey"`

	Title string `gorm:"size:255;not null"`
	Artist string `gorm:"size:255;not null"`
	Album string `gorm:"size:255"`

	Duration int

	CoverURL string `gorm:"type:text"`
	AudioURL string `gorm:"type:text"`

	PlayCount int `gorm:"default:0"`

	CreatedAt time.Time
}