package todo

import (
	"gofiber-todo/src/entity"

	"gorm.io/gorm"
)

type TodoService struct {
	repository *gorm.DB
}

func (s *TodoService) getTodos(dest interface{}, userId uint) *gorm.DB {
	return s.repository.Table("todos").Find(dest, "creator = ?", userId)
}

func (s *TodoService) getTodo(dest interface{}, todoId int) *gorm.DB {
	return s.repository.Table("todos").First(dest, "id = ?", todoId)
}

func (s *TodoService) createTodo(todo *entity.Todo) *gorm.DB {
	return s.repository.Create(todo)
}

func (s *TodoService) deleteTodo(todoId int) *gorm.DB {
	return s.repository.Unscoped().Delete(&entity.Todo{}, todoId)
	// return s.repository.Unscoped().Delete("id = ?", todoId)
}
