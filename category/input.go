package category

import "time"

type Category struct {
	ID        string     `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" binding:"required"`
	Slug      string     `json:"slug" gorm:"unique"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"not null"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"null"`
}

type CategoryInput struct {
	ID        string     `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" binding:"required"`
	Slug      string     `json:"slug" gorm:"unique"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"not null"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"null"`
}

type GetCategoryID struct {
	ID string `uri:"id" binding:"required"`
}

type CategoryPaginationRequest struct {
	Search *string `form:"search"`
	Page   int     `form:"page"`
	Limit  int     `form:"limit"`
}
