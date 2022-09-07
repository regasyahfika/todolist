package todo

import "time"

type Todo struct {
	ID          string `gorm:"primaryKey"`
	CategoryID  string
	Name        int
	Description int
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
	DeletedAt   time.Time `gorm:"null"`
}
