package product

import (
	"orderapi/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	*gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	tx := r.DB.First(&product, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &product, nil
}

func (r *ProductRepository) Create(product *model.Product) (*model.Product, error) {
	tx := r.DB.Clauses(clause.Returning{}).Create(product)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return product, nil
}

func (r *ProductRepository) Update(product *model.Product) (*model.Product, error) {
	tx := r.DB.Clauses(clause.Returning{}).Updates(product)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return product, nil
}

func (r *ProductRepository) Delete(id uint) error {
	var product model.Product
	tx := r.DB.Delete(&product, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
