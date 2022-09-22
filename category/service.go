package category

import (
	"errors"
	"learning/todo/helper"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Service interface {
	FindAll(req helper.PaginationRequest) ([]Category, error)
	FindByID(inputID GetCategoryID) (Category, error)
	Save(category Category) (Category, error)
	Update(inputID GetCategoryID, inputCategory Category) (Category, error)
	Delete(inputID GetCategoryID) (error, error)
}

type service struct {
	repo Repository
}

func NewCategoryService(db *gorm.DB) Service {
	repo := NewCategoryRepository(db)
	return &service{repo: repo}
}

func (s *service) FindAll(req helper.PaginationRequest) ([]Category, error) {
	return s.repo.FindAll(req)
}

func (s *service) FindByID(inputID GetCategoryID) (Category, error) {
	return s.repo.FindByID(inputID.ID)
}

func (s *service) Update(inputID GetCategoryID, inputCategory Category) (Category, error) {
	category, err := s.repo.FindByID(inputID.ID)

	if err != nil {
		return category, err
	}

	if category.ID != inputID.ID {
		return category, errors.New("Salah ID")
	}

	category.Name = inputCategory.Name
	category.Slug = slug.Make(inputCategory.Name)

	return s.repo.Update(category)
}

func (s *service) Save(category Category) (Category, error) {
	category.Slug = slug.Make(category.Name)
	return s.repo.Save(category)
}

func (s *service) Delete(inputID GetCategoryID) (error, error) {
	category, err := s.repo.FindByID(inputID.ID)

	if err != nil {
		return err, err
	}

	if category.ID != inputID.ID {
		return err, errors.New("Salah ID")
	}

	return s.repo.Delete(category.ID)
}
