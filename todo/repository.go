package todo

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]TodoInput, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]TodoInput, error) {
	var todos []TodoInput

	err := r.db.Find(&todos).Error

	if err != nil {
		return todos, err
	}

	return todos, nil
}
