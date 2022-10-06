package auth

import (
	"gofiber-todo/src/entity"

	"gorm.io/gorm"
)

type AuthService struct {
	repository *gorm.DB
}

func (s *AuthService) findUser(dest interface{}, conds ...interface{}) *gorm.DB {
	// return database.DB.Model(&entity.User{}).Take(dest, conds...)
	return s.repository.Take(dest, conds...)
}

func (s *AuthService) findUserByEmail(dest interface{}, email string) *gorm.DB {
	return s.findUser(dest, "email = ?", email)
}

func (s *AuthService) createUser(user *entity.User) *gorm.DB {
	return s.repository.Create(user)
}
