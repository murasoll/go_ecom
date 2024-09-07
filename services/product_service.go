// services/product_service.go

package services

import (
	"ecomerce/models"
	"ecomerce/repositories/category_repo"
	"ecomerce/repositories/product_repo"
	"errors"
)

type ProductService struct {
	productRepo  product_repo.ProductRepository
	categoryRepo category_repo.CategoryRepository
}

func NewProductService(pr product_repo.ProductRepository, cr category_repo.CategoryRepository) *ProductService {
	return &ProductService{
		productRepo:  pr,
		categoryRepo: cr,
	}
}

func (s *ProductService) GetAllProducts(filter map[string]interface{}) ([]models.Product, error) {
	return s.productRepo.GetAll(filter)
}

func (s *ProductService) GetProductByID(id uint) (models.Product, error) {
	return s.productRepo.GetByID(id)
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	// Validate category exists
	_, err := s.categoryRepo.GetByID(product.CategoryID)
	if err != nil {
		return errors.New("invalid category")
	}
	return s.productRepo.Create(product)
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	// Validate category exists
	_, err := s.categoryRepo.GetByID(product.CategoryID)
	if err != nil {
		return errors.New("invalid category")
	}
	return s.productRepo.Update(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.productRepo.Delete(id)
}

func (s *ProductService) UpdateInventory(id uint, quantity int) error {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return err
	}
	if int(product.Inventory)+quantity < 0 {
		return errors.New("insufficient inventory")
	}
	return s.productRepo.UpdateInventory(id, quantity)
}
