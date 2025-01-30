package user

import (
	"orderapi/internal/cart"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name  string    `gorm:"size:255;not null;index:idx_name_category"`
	Email string    `gorm:"uniqueIndex;size:255"`
	Cart  cart.Cart `gorm:"foreignKey:UserID;references:ID"`
}
