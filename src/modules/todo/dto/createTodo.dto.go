package dto

import (
	"gofiber-todo/src/entity"
	"time"
)

type CreateTodoDTO struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title" validate:"required,lte=100"`
	Completed bool      `json:"completed" validate:"boolean"`
	Memo      string    `json:"memo"`
	CreatedAt time.Time `json:"created_at"`
}

func (dto CreateTodoDTO) ToEntity() *entity.Todo {
	return &entity.Todo{
		ID:        dto.ID,
		Title:     dto.Title,
		Completed: dto.Completed,
		Memo:      dto.Memo,
		CreatedAt: dto.CreatedAt,
	}
}

func (dto CreateTodoDTO) FromEntity(todo *entity.Todo) *CreateTodoDTO {
	return &CreateTodoDTO{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		Memo:      todo.Memo,
		CreatedAt: todo.CreatedAt,
	}
}
