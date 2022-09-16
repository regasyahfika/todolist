package todo

import (
	"learning/todo/helper"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(req helper.PaginationRequest) ([]Todo, error)
	FindByID(ID string) (Todo, error)
	Save(todo Todo) (Todo, error)
	Update(todo Todo) (Todo, error)
	Delete(ID string) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll(req helper.PaginationRequest) ([]Todo, error) {
	var todos []Todo
	orm := r.db.Model(Todo{})

	if req.Search != nil {
		orm = orm.Where("name like ?", "%"+*req.Search+"%")
	}
	orm = orm.Limit(req.Limit).Offset((req.Page - 1*req.Limit))
	err := orm.Preload("Category").Find(&todos).Error

	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (r *repository) FindByID(ID string) (Todo, error) {
	var todos Todo

	err := r.db.Preload("Category").Where("id = ?", ID).Find(&todos).Error
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (r *repository) Save(todo Todo) (Todo, error) {
	err := r.db.Create(&todo).Error

	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func (r *repository) Update(todo Todo) (Todo, error) {
	err := r.db.Save(&todo).Error

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (r *repository) Delete(ID string) (bool, error) {
	var todo Todo

	err := r.db.Delete(&todo, "id = ?", ID).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
