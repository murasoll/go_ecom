// services/order_service.go

package services

import (
	"ecomerce/models"
	"ecomerce/repositories/cart_repo"
	"ecomerce/repositories/order_repo"
	"ecomerce/repositories/product_repo"
	"errors"
)

type OrderService struct {
	orderRepo   order_repo.OrderRepository
	cartRepo    cart_repo.CartRepository
	productRepo product_repo.ProductRepository
}

func NewOrderService(or order_repo.OrderRepository, cr cart_repo.CartRepository, pr product_repo.ProductRepository) *OrderService {
	return &OrderService{
		orderRepo:   or,
		cartRepo:    cr,
		productRepo: pr,
	}
}

func (s *OrderService) CreateOrder(userID uint, shippingAddress models.ShippingAddress) (*models.Order, error) {
	cart, err := s.cartRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	if len(cart.Items) == 0 {
		return nil, errors.New("cart is empty")
	}

	var total float64
	var orderItems []models.OrderItem

	for _, item := range cart.Items {
		product, err := s.productRepo.GetByID(item.ProductID)
		if err != nil {
			return nil, err
		}
		if product.Inventory < int(item.Quantity) {
			return nil, errors.New("not enough inventory for product: " + product.Name)
		}
		total += product.Price * float64(item.Quantity)
		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})

		// Update inventory
		err = s.productRepo.UpdateInventory(item.ProductID, -item.Quantity)
		if err != nil {
			return nil, err
		}
	}

	order := &models.Order{
		UserID:          userID,
		Items:           orderItems,
		Status:          "pending",
		Total:           total,
		ShippingAddress: shippingAddress,
		// You might want to calculate shipping cost based on the address
		ShippingCost: 10.0, // Example fixed shipping cost
	}

	err = s.orderRepo.Create(order)
	if err != nil {
		return nil, err
	}

	// Clear the cart after successful order creation
	err = s.cartRepo.Clear(userID)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) GetOrdersByUser(userID uint) ([]models.Order, error) {
	return s.orderRepo.GetByUserID(userID)
}

func (s *OrderService) GetOrderByID(orderID uint, userID uint) (models.Order, error) {
	order, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return order, err
	}
	if order.UserID != userID {
		return order, errors.New("unauthorized access to order")
	}
	return order, nil
}

func (s *OrderService) UpdateOrderStatus(orderID uint, status string) error {
	return s.orderRepo.UpdateStatus(orderID, status)
}
