package dto

type LoginDTO struct {
	Email    string `json:"email" validate:"required,lte=100,email"`
	Password string `json:"password" validate:"required,lte=100"`
}

type LoginResult struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// func (dto LoginDTO) ToEntity() *entity.Todo {
// 	return &entity.Todo{
// 		ID:        dto.ID,
// 		Title:     dto.Title,
// 		Completed: dto.Completed,
// 		Memo:      dto.Memo,
// 		CreatedAt: dto.CreatedAt,
// 	}
// }
