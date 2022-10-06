package dto

import (
	"gofiber-todo/src/entity"
	"time"
)

type TodoDTO struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	Memo      string    `json:"memo"`
	CreatedAt time.Time `json:"created_at"`
}

func (dto TodoDTO) ToEntity() *entity.Todo {
	return &entity.Todo{
		ID:        dto.ID,
		Title:     dto.Title,
		Completed: dto.Completed,
		Memo:      dto.Memo,
		CreatedAt: dto.CreatedAt,
	}
}
