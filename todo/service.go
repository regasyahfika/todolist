package todo

import "gorm.io/gorm"

type Service interface {
	FindAll() ([]TodoInput, error)
}

type service struct {
	repository Repository
}

func NewService(db *gorm.DB) Service {
	repo := NewRepository(db)
	return &service{repository: repo}
}

func (s *service) FindAll() ([]TodoInput, error) {
	return s.repository.FindAll()
}
