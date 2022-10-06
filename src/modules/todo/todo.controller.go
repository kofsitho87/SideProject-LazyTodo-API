package todo

import (
	"gofiber-todo/src/modules/todo/dto"
	"gofiber-todo/utils/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TodoController struct {
	service *TodoService
}

func (ctrl *TodoController) GetTodos(c *fiber.Ctx) error {
	todos := &[]dto.TodoDTO{}

	if err := ctrl.service.getTodos(todos).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": &dto.ListTodoDTO{
			List: todos,
		},
	})
}

func (ctrl *TodoController) GetTodo(c *fiber.Ctx) error {
	// c.Params("name") string

	todoId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	todo := &dto.TodoDTO{}

	if err := ctrl.service.getTodo(todo, todoId).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": todo,
	})
}

func (ctrl *TodoController) CreateTodo(c *fiber.Ctx) error {
	createTodoDto := new(dto.CreateTodoDTO)
	if err := validator.ParseBodyAndValidate(c, createTodoDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	createTodoDto.ID = uint(uuid.New().ID())

	todo := createTodoDto.ToEntity()
	if err := ctrl.service.createTodo(todo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result := createTodoDto.FromEntity(todo)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": result,
	})
}

func (ctrl *TodoController) DeleteTodo(c *fiber.Ctx) error {
	todoId, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res := ctrl.service.deleteTodo(todoId)

	if err := res.Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if res.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Unable to delete todo",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": true,
	})
}
