package todo

import (
	"gofiber-todo/src/config/database"
	"gofiber-todo/src/middleware"

	"github.com/gofiber/fiber/v2"
)

type TodoModule struct {
	app  fiber.Router
	ctrl *TodoController
}

func NewModule(app fiber.Router) *TodoModule {
	m := TodoModule{}
	m.app = app
	m.ctrl = &TodoController{}

	m.ctrl.service = &TodoService{database.DB}
	m.setRoutes()

	return &m
}

func (m *TodoModule) setRoutes() {
	route := m.app.Group("/todos").Use(middleware.Auth)

	route.Get("/", m.ctrl.GetTodos)
	route.Get("/:id", m.ctrl.GetTodo)
	route.Post("/", m.ctrl.CreateTodo)
	route.Delete("/:id", m.ctrl.DeleteTodo)
	route.Put("/:id/complete", m.ctrl.CompleteTodo)
}
