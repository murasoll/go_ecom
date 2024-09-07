// services/cart_service.go

package services

import (
	"ecomerce/models"
	"ecomerce/repositories/cart_repo"
	"ecomerce/repositories/product_repo"
	"errors"
)

type CartService struct {
	cartRepo    cart_repo.CartRepository
	productRepo product_repo.ProductRepository
}

func NewCartService(cr cart_repo.CartRepository, pr product_repo.ProductRepository) *CartService {
	return &CartService{
		cartRepo:    cr,
		productRepo: pr,
	}
}

func (s *CartService) GetCart(userID uint) (models.Cart, error) {
	return s.cartRepo.GetByUserID(userID)
}

func (s *CartService) AddToCart(userID uint, productID uint, quantity int) error {
	// Check if product exists and has enough inventory
	product, err := s.productRepo.GetByID(productID)
	if err != nil {
		return err
	}
	if product.Inventory < int(quantity) {
		return errors.New("not enough inventory")
	}
	return s.cartRepo.AddItem(userID, productID, quantity)
}

func (s *CartService) UpdateCartItemQuantity(userID uint, productID uint, quantity int) error {
	// Check if product exists and has enough inventory
	product, err := s.productRepo.GetByID(productID)
	if err != nil {
		return err
	}
	if product.Inventory < int(quantity) {
		return errors.New("not enough inventory")
	}
	return s.cartRepo.UpdateItemQuantity(userID, productID, quantity)
}

func (s *CartService) RemoveFromCart(userID uint, productID uint) error {
	return s.cartRepo.RemoveItem(userID, productID)
}

func (s *CartService) ClearCart(userID uint) error {
	return s.cartRepo.Clear(userID)
}
