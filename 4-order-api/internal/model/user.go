package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name   string `gorm:"size:255;not null;index:idx_name_category"`
	Email  string `gorm:"uniqueIndex;size:255"`
	CartID uint
}
