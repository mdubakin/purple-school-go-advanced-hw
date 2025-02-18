package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name     string `gorm:"size:255;not null;index:idx_name_category"`
	Category string `gorm:"size:255;not null;index:idx_name_category"`
	Price    int    `gorm:"check:price > 0"`
	Carts    []Cart `gorm:"many2many:cart_products;"`
}
