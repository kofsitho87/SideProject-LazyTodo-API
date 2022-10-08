package entity

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model

	ID uint `gorm:"primaryKey;autoIncrement;not null"`
	// Creator   User   `gorm:"embedded"`
	Creator   uint   `gorm:"not null"`
	Title     string `gorm:"not nulll;size:20"`
	Completed bool   `gorm:"not null;default:false"`
	Memo      string
	EndedAt   time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}
