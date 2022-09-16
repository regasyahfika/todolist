package category

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        string          `json:"id" gorm:"primaryKey"`
	Name      string          `json:"name" binding:"required"`
	Slug      string          `json:"slug" gorm:"unique"`
	CreatedAt time.Time       `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"not null"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"null"`
}

type CategoryInput struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" binding:"required"`
	Slug      string         `json:"slug" gorm:"unique"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"null"`
}

type GetCategoryID struct {
	ID string `uri:"id" binding:"required"`
}
