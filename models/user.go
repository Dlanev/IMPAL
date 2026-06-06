package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"size:100;not null"`
	Email     string    `gorm:"size:150;unique;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time
}