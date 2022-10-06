package auth

import (
	"gofiber-todo/config/database"
	"gofiber-todo/src/entity"

	"github.com/gofiber/fiber/v2"
)

type AuthModule struct {
	app  fiber.Router
	ctrl *AuthController
}

func NewModule(app fiber.Router) *AuthModule {
	m := AuthModule{}
	m.app = app
	m.ctrl = &AuthController{}
	m.ctrl.service = &AuthService{database.DB.Model(&entity.User{})}
	m.setRoutes(app)
	return &m
}

func (m *AuthModule) setRoutes(app fiber.Router) {
	route := app.Group("/auth")

	route.Post("/login", m.ctrl.Login)
	route.Post("/signup", m.ctrl.SignUp)
}
