package models

import "time"

type Playlist struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserID uint `json:"user_id"`

	Name string `gorm:"size:255;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`

	CreatedAt time.Time `json:"created_at"`
}