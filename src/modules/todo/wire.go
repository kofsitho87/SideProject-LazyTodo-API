package todo

import (
	"gofiber-todo/src/config/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

func InitializeTodoModule(app fiber.Router) *TodoModule {
	wire.Build(newTodoModule, newTodoController, newTodoService)
	return &TodoModule{}
}

func newTodoModule(c *TodoController) *TodoModule {
	return &TodoModule{
		ctrl: c,
	}
}

func newTodoController(s *TodoService) *TodoController {
	return &TodoController{
		service: s,
	}
}

func newTodoService() *TodoService {
	return &TodoService{
		repository: database.DB,
	}
}
