package dto

import (
	"gofiber-todo/src/entity"
	"time"
)

type SignupDTO struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email" validate:"required,lte=100,email"`
	Password  string    `json:"password" validate:"required,lte=100"`
	CreatedAt time.Time `json:"created_at"`
}

func (dto SignupDTO) ToEntity() *entity.User {
	return &entity.User{
		ID:       dto.ID,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func (dto SignupDTO) FromEntity(user *entity.User) *SignupDTO {
	return &SignupDTO{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}
}
