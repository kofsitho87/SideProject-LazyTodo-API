package entity

import (
	"gofiber-todo/src/modules/todo/dto"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model

	ID        uint   `gorm:"primaryKey;autoIncrement;not null"`
	Creator   uint   `gorm:"not null"`
	Title     string `gorm:"not nulll;size:10"`
	Completed bool   `gorm:"not null;default:false"`
	Memo      string
	EndedAt   time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}

func (todo *Todo) FromDto(dto *dto.CreateTodoDTO) *Todo {
	todo.Title = dto.Title
	todo.Completed = dto.Completed
	todo.Memo = dto.Memo
	// todo.EndedAt = dto.EndedAt
	todo.EndedAt = time.Now()
	return todo
}

func (todo *Todo) ToDto() *dto.TodoDTO {
	return &dto.TodoDTO{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		Memo:      todo.Memo,
		EndedAt:   todo.EndedAt,
		CreatedAt: todo.CreatedAt,
	}
}
