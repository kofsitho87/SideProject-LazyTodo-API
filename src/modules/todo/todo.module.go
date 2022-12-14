package todo

import (
	"gofiber-todo/config/database"
	"gofiber-todo/src/entity"
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

	//TODO: database.DB.Model(&entity.Todo{}) 이방법은 작동이 안됨!
	m.ctrl.service = &TodoService{database.DB.Model(&entity.Todo{})}

	m.setRoutes()
	return &m
}

func (m *TodoModule) setRoutes() {
	route := m.app.Group("/todos").Use(middleware.Auth)

	route.Get("/", m.ctrl.GetTodos)
	route.Get("/:id", m.ctrl.GetTodo)
	route.Post("/", m.ctrl.CreateTodo)
	route.Delete("/:id", m.ctrl.DeleteTodo)
}
