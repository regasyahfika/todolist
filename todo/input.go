package todo

import (
	"learning/todo/category"
	"time"

	"github.com/calasteo/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID          string             `json:"id" gorm:"primaryKey"`
	CategoryID  string             `json:"category_id" gorm:"index" binding:"required"`
	Name        string             `json:"name" binding:"required"`
	Description string             `json:"description" binding:"required"`
	CreatedAt   time.Time          `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time          `json:"updated_at" gorm:"not null"`
	DeletedAt   *gorm.DeletedAt    `json:"deleted_at,omitempty" gorm:"null"`
	Category    *category.Category `json:"category,omitempty" gorm:"null"`
}

type GetTodoID struct {
	ID string `uri:"id" binding:"required"`
}

func (c *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.GenerateOrderedUUID()
	return
}
