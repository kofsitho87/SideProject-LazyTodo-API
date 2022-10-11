package dto

type LoginDTO struct {
	Email    string `json:"email" validate:"required,lte=100,email" swaggertype:"string" example:"s1@s.com"`
	Password string `json:"password" validate:"required,lte=100" swaggertype:"string" example:"123456`
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
