package entity

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model

	ID        uint   `gorm:"primaryKey;autoIncrement;not null"`
	Title     string `gorm:"not null"`
	Completed bool   `gorm:"not null;default:false"`
	Memo      string
	CreatedAt time.Time
}
