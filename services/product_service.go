package services

import (
	"ecomerce/models"
	"ecomerce/repositories/product_repo"
)

func GetProducts() []models.Product {
	return product_repo.GetAllProducts()
}

func CreateProduct(product models.Product) (models.Product, error) {
	return product_repo.CreateProduct(product)
}

func GetProductByID(id string) (models.Product, error) {
	return product_repo.GetProductByID(id)
}

func UpdateProduct(product models.Product) error {
	return product_repo.UpdateProduct(product)
}

func DeleteProduct(id string) error {
	return product_repo.DeleteProduct(id)
}
