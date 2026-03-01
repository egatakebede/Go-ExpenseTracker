package models

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserID     uint      `gorm:" index" json:"user_id"`
	Title      string    `json:"title"`
	Amount     float64   `json:"amount"`
	Categories string    `json:"categories"`
	Notes      string    `json:"notes"`
	SpentAt    time.Time `json:"spent_at"`
}
