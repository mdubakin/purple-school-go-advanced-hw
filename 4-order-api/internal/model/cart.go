package model

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model

	User     User
	Products []Product `gorm:"many2many:cart_products;"`
}
