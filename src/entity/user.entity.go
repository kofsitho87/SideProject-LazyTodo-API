package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID        uint   `gorm:"primaryKey;autoIncrement;not null"`
	Email     string `gorm:"size:256;not null"`
	Password  string `gorm:"not null;"`
	CreatedAt time.Time
}

// func (todo *Todo) GetTodos(dest interface{}) *gorm.DB {
// 	return database.DB.Model(todo).Find(dest, "")
// }
