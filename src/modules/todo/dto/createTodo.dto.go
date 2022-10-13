package dto

type CreateTodoDTO struct {
	Title     string `json:"title" validate:"required,lte=20"`
	Completed bool   `json:"completed" validate:"boolean"`
	Memo      string `json:"memo"`
	// EndedAt   time.Time `json:"ended_at" validate:"required,datetime"`
}
