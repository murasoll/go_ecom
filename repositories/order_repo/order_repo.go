// repositories/order_repo/order_repo.go

package order_repo

import (
	"ecomerce/config"
	"ecomerce/models"
	"log"
)

type OrderRepository interface {
	Create(order *models.Order) error
	GetByID(id uint) (models.Order, error)
	GetByUserID(userID uint) ([]models.Order, error)
	UpdateStatus(id uint, status string) error
}

type orderRepo struct{}

func NewOrderRepo() OrderRepository {
	return &orderRepo{}
}

func (r *orderRepo) Create(order *models.Order) error {
	tx, err := config.DB.Begin()
	if err != nil {
		return err
	}

	// Insert order
	query := `
		INSERT INTO orders (user_id, status, total, shipping_address, shipping_cost)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err = tx.QueryRow(query, order.UserID, order.Status, order.Total, order.ShippingAddress, order.ShippingCost).Scan(&order.ID)
	if err != nil {
		tx.Rollback()
		log.Println("Error creating order:", err)
		return err
	}

	// Insert order items
	for _, item := range order.Items {
		query := `
			INSERT INTO order_items (order_id, product_id, quantity, price)
			VALUES ($1, $2, $3, $4)
		`
		_, err = tx.Exec(query, order.ID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			tx.Rollback()
			log.Println("Error creating order item:", err)
			return err
		}
	}

	return tx.Commit()
}

func (r *orderRepo) GetByID(id uint) (models.Order, error) {
	var order models.Order
	query := `
		SELECT id, user_id, status, total, shipping_address, shipping_cost
		FROM orders
		WHERE id = $1
	`
	err := config.DB.QueryRow(query, id).Scan(&order.ID, &order.UserID, &order.Status, &order.Total, &order.ShippingAddress, &order.ShippingCost)
	if err != nil {
		log.Println("Error fetching order by ID:", err)
		return order, err
	}

	// Fetch order items
	itemsQuery := `
		SELECT oi.product_id, oi.quantity, oi.price, p.name
		FROM order_items oi
		JOIN products p ON oi.product_id = p.id
		WHERE oi.order_id = $1
	`
	rows, err := config.DB.Query(itemsQuery, id)
	if err != nil {
		log.Println("Error fetching order items:", err)
		return order, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.OrderItem
		var productName string
		err := rows.Scan(&item.ProductID, &item.Quantity, &item.Price, &productName)
		if err != nil {
			log.Println("Error scanning order item:", err)
			return order, err
		}
		item.Product = models.Product{ID: item.ProductID, Name: productName, Price: item.Price}
		order.Items = append(order.Items, item)
	}

	return order, nil
}

func (r *orderRepo) GetByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	query := `
		SELECT id, user_id, status, total, shipping_address, shipping_cost
		FROM orders
		WHERE user_id = $1
		ORDER BY id DESC
	`
	rows, err := config.DB.Query(query, userID)
	if err != nil {
		log.Println("Error fetching orders by user ID:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.UserID, &order.Status, &order.Total, &order.ShippingAddress, &order.ShippingCost)
		if err != nil {
			log.Println("Error scanning order:", err)
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *orderRepo) UpdateStatus(id uint, status string) error {
	query := `UPDATE orders SET status = $1 WHERE id = $2`
	_, err := config.DB.Exec(query, status, id)
	if err != nil {
		log.Println("Error updating order status:", err)
		return err
	}
	return nil
}
