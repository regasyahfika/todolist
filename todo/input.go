package todo

import (
	"learning/todo/category"
	"time"
)

type TodoInput struct {
	ID          string    `json:"id"`
	CategoryID  string    `json:"category_id" gorm:"index" validate:"required"`
	Name        int       `json:"name" validate:"required"`
	Description int       `json:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"not null"`
	// DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"null"`
	Category category.Category
}
