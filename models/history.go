package models

import "time"

type History struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserID uint `json:"user_id"`
	SongID uint `json:"song_id"`

	User User `gorm:"foreignKey:UserID" json:"-"`
	Song Song `gorm:"foreignKey:SongID" json:"song"`

	PlayedAt time.Time `json:"played_at"`
}