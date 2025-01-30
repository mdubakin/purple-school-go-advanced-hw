package cart

import (
	"orderapi/internal/product"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model

	UserID   uint              `gorm:"uniqueIndex"`
	Products []product.Product `gorm:"many2many:cart_products;"`
}
