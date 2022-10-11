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

func (s *TodoService) getTodoByModel(model *entity.Todo) *gorm.DB {
	return s.repository.First(model)
}

func (s *TodoService) createTodo(todo *entity.Todo) *gorm.DB {
	return s.repository.Create(todo)
}

func (s *TodoService) deleteTodo(todoId int) *gorm.DB {
	return s.repository.Unscoped().Delete(&entity.Todo{}, todoId)
}

func (s *TodoService) completeTodoById(todoId int) *gorm.DB {
	return s.repository.Table("todos").Where("id", todoId).Update("completed", true)
}

func (s *TodoService) completeTodo(todo *entity.Todo) *gorm.DB {
	return s.repository.Model(todo).Update("completed", true)
}
