package todo

import (
	"gofiber-todo/src/entity"
	"gofiber-todo/src/modules/todo/dto"
	"gofiber-todo/src/utils/validator"

	"github.com/gofiber/fiber/v2"
)

type TodoController struct {
	service *TodoService
}

// GetTodos godoc
// @Summary      GetTodos
// @Description  GetTodos
// @Tags         todos
// @Accept       json
// @Produce      json
// @Security 		 ApiKeyAuth
// @Success      200  	 {object}  	dto.ListTodoDTO
// @Failure      400  	 {object}		response.ErrorResponse
// @Failure      500  	 {object}		response.ErrorResponse
// @Router       /api/todos [get]
func (ctrl *TodoController) GetTodos(c *fiber.Ctx) error {
	todos := &[]dto.TodoDTO{}

	userId := c.Locals("USER").(uint)

	if err := ctrl.service.getTodos(todos, userId).Error; err != nil {
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

// CreateTodo godoc
// @Summary      CreateTodo
// @Description  CreateTodo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Security 		 ApiKeyAuth
// @Param        CreateTodoDTO   body    dto.CreateTodoDTO  true  "todo item"
// @Success      200  	 {object}  	dto.TodoDTO
// @Failure      400  	 {object}		response.ErrorResponse
// @Failure      500  	 {object}		response.ErrorResponse
// @Router       /api/todos [post]
func (ctrl *TodoController) CreateTodo(c *fiber.Ctx) error {
	createTodoDto := new(dto.CreateTodoDTO)
	if err := validator.ParseBodyAndValidate(c, createTodoDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	todoItem := &entity.Todo{}
	todoItem.FromDto(createTodoDto)
	todoItem.Creator = c.Locals("USER").(uint)

	if err := ctrl.service.createTodo(todoItem).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": todoItem.ToDto(),
	})
}

// DeleteTodo godoc
// @Summary      DeleteTodo
// @Description  DeleteTodo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Security 		 ApiKeyAuth
// @Param        id   	 path				int true "Todo ID"
// @Success      200  	 {object} 	bool
// @Failure      400  	 {object}		response.ErrorResponse
// @Failure      500  	 {object}		response.ErrorResponse
// @Router       /api/todos/{id} 		[delete]
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
