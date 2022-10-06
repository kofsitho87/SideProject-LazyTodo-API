package todo

import (
	"gofiber-todo/config/database"
	"gofiber-todo/src/entity"

	"gorm.io/gorm"
)

type TodoService struct {
	repository *gorm.DB
}

func (s *TodoService) getTodos(dest interface{}) *gorm.DB {
	return s.repository.Find(dest, "")
}

func (s *TodoService) getTodo(dest interface{}, todoId int) *gorm.DB {
	return s.repository.First(dest, "id = ?", todoId)
}

func (s *TodoService) createTodo(todo *entity.Todo) *gorm.DB {
	return s.repository.Create(todo)
}

func (s *TodoService) deleteTodo(todoId int) *gorm.DB {
	return database.DB.Unscoped().Delete(&entity.Todo{}, todoId)
	// return s.repository.Unscoped().Delete("id = ?", todoId)
}
