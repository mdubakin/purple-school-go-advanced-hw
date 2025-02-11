package usecase

import (
	"orderapi/internal/model"
	repository "orderapi/internal/repository"

	"gorm.io/gorm"
)

type ProductService struct {
	*repository.ProductRepository
}

func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{ProductRepository: productRepo}
}

func (p *ProductService) GetByID(id uint) (*model.Product, error) {
	return p.ProductRepository.GetByID(id)
}

func (p *ProductService) Create(name, category string, price int) (*model.Product, error) {
	product := model.Product{
		Name:     name,
		Category: category,
		Price:    price,
	}
	return p.ProductRepository.Create(&product)
}

func (p *ProductService) Update(id uint, name, category string, price int) (*model.Product, error) {
	product := model.Product{
		Model:    gorm.Model{ID: id},
		Name:     name,
		Category: category,
		Price:    price,
	}
	return p.ProductRepository.Update(&product)
}

func (p *ProductService) Delete(id uint) error {
	return p.ProductRepository.Delete(id)
}
