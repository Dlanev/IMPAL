package models

import "time"

type Favorite struct {
	ID uint `gorm:"primaryKey"`

	UserID uint
	User User `gorm:"foreignKey:UserID"`

	SongID uint
	Song Song `gorm:"foreignKey:SongID"`

	CreatedAt time.Time
}