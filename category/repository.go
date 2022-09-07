package category

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(req CategoryPaginationRequest) ([]Category, error)
	FindByID(ID string) (Category, error)
	Save(category Category) (Category, error)
	Update(category Category) (Category, error)
	Delete(ID string) (Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll(req CategoryPaginationRequest) ([]Category, error) {
	var category []Category
	orm := r.db.Model(Category{})

	if req.Search != nil {
		orm = orm.Where("name like ?", "%"+*req.Search+"%")
	}
	orm = orm.Limit(req.Limit).Offset((req.Page - 1*req.Limit))
	err := orm.Find(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) FindByID(ID string) (Category, error) {
	var category Category

	err := r.db.Where("id = ?", ID).Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) Save(category Category) (Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return Category{}, err
	}

	err = r.db.Model(Category{}).Where("slug = ?", category.Slug).Take(&category).Error

	if err != nil {
		return Category{}, err
	}

	return category, nil
}

func (r *repository) Update(category Category) (Category, error) {
	err := r.db.Save(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) Delete(ID string) (Category, error) {
	var category Category

	err := r.db.Delete(&category, ID).Error

	if err != nil {
		return category, err
	}

	return category, nil
}
