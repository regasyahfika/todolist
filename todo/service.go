package todo

import (
	"errors"
	"learning/todo/helper"

	"github.com/calasteo/uuid"
	"gorm.io/gorm"
)

type Service interface {
	FindAll(req helper.PaginationRequest) ([]Todo, error)
	FindByID(inputID GetTodoID) (Todo, error)
	Save(todo Todo) (Todo, error)
	Update(inputID GetTodoID, inputTodo Todo) (Todo, error)
	Delete(inputID GetTodoID) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(db *gorm.DB) Service {
	repo := NewRepository(db)
	return &service{repository: repo}
}

func (s *service) FindAll(req helper.PaginationRequest) ([]Todo, error) {
	return s.repository.FindAll(req)
}

func (s *service) FindByID(inputID GetTodoID) (Todo, error) {
	return s.repository.FindByID(inputID.ID)
}

func (s *service) Save(todo Todo) (Todo, error) {
	todo.ID = uuid.GenerateOrderedUUID()

	return s.repository.Save(todo)
}

func (s *service) Update(inputID GetTodoID, inputTodo Todo) (Todo, error) {
	todo, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return todo, err
	}

	if todo.ID != inputID.ID {
		return todo, errors.New("Salah ID todo")
	}

	todo.Name = inputTodo.Name
	todo.Description = inputTodo.Description
	todo.CategoryID = inputTodo.CategoryID

	return s.repository.Update(todo)
}

func (s *service) Delete(inputID GetTodoID) (bool, error) {
	todo, err := s.repository.FindByID(inputID.ID)

	if err != nil {
		return false, err
	}

	if todo.ID != inputID.ID {
		return false, errors.New("Salah ID")
	}

	return s.repository.Delete(inputID.ID)
}
