package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `gorm:"size:20;not null" json:"name"`
	Email     string         `gorm:"size:20;unique;not null" json:"email"`
	Password  string         `gorm:"size:20;not null" json:"-"`
}
