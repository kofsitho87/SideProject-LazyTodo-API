package dto

import (
	"gofiber-todo/src/entity"
	"time"
)

type CreateTodoDTO struct {
	ID        uint   `json:"id"`
	Creator   uint   `json:"creator"`
	Title     string `json:"title" validate:"required,lte=100"`
	Completed bool   `json:"completed" validate:"boolean"`
	Memo      string `json:"memo"`
	// EndedAt   time.Time `json:"ended_at" validate:"required,datetime"`
	CreatedAt time.Time `json:"created_at"`
}

func (dto CreateTodoDTO) ToEntity() *entity.Todo {
	return &entity.Todo{
		ID:        dto.ID,
		Creator:   dto.Creator,
		Title:     dto.Title,
		Completed: dto.Completed,
		Memo:      dto.Memo,
		// EndedAt:   dto.EndedAt,
		CreatedAt: dto.CreatedAt,
	}
}

func (dto CreateTodoDTO) FromEntity(todo *entity.Todo) *CreateTodoDTO {
	return &CreateTodoDTO{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		Memo:      todo.Memo,
		// EndedAt:   todo.EndedAt,
		CreatedAt: todo.CreatedAt,
	}
}
