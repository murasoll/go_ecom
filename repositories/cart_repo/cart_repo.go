// repositories/cart_repo/cart_repo.go

package cart_repo

import (
	"ecomerce/config"
	"ecomerce/models"
	"log"
)

type CartRepository interface {
	GetByUserID(userID uint) (models.Cart, error)
	AddItem(userID uint, productID uint, quantity int) error
	UpdateItemQuantity(userID uint, productID uint, quantity int) error
	RemoveItem(userID uint, productID uint) error
	Clear(userID uint) error
}

type cartRepo struct{}

func NewCartRepo() CartRepository {
	return &cartRepo{}
}

func (r *cartRepo) GetByUserID(userID uint) (models.Cart, error) {
	var cart models.Cart
	cart.UserID = userID

	query := `
		SELECT ci.product_id, ci.quantity, p.name, p.price 
		FROM cart_items ci
		JOIN products p ON ci.product_id = p.id
		WHERE ci.user_id = $1
	`
	rows, err := config.DB.Query(query, userID)
	if err != nil {
		log.Println("Error fetching cart items:", err)
		return cart, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.CartItem
		var productName string
		var productPrice float64
		err := rows.Scan(&item.ProductID, &item.Quantity, &productName, &productPrice)
		if err != nil {
			log.Println("Error scanning cart item:", err)
			return cart, err
		}
		item.Product = models.Product{ID: item.ProductID, Name: productName, Price: productPrice}
		cart.Items = append(cart.Items, item)
	}

	return cart, nil
}

func (r *cartRepo) AddItem(userID uint, productID uint, quantity int) error {
	query := `
		INSERT INTO cart_items (user_id, product_id, quantity)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, product_id) DO UPDATE SET quantity = cart_items.quantity + $3
	`
	_, err := config.DB.Exec(query, userID, productID, quantity)
	if err != nil {
		log.Println("Error adding item to cart:", err)
		return err
	}
	return nil
}

func (r *cartRepo) UpdateItemQuantity(userID uint, productID uint, quantity int) error {
	query := `UPDATE cart_items SET quantity = $1 WHERE user_id = $2 AND product_id = $3`
	_, err := config.DB.Exec(query, quantity, userID, productID)
	if err != nil {
		log.Println("Error updating cart item quantity:", err)
		return err
	}
	return nil
}

func (r *cartRepo) RemoveItem(userID uint, productID uint) error {
	query := `DELETE FROM cart_items WHERE user_id = $1 AND product_id = $2`
	_, err := config.DB.Exec(query, userID, productID)
	if err != nil {
		log.Println("Error removing item from cart:", err)
		return err
	}
	return nil
}

func (r *cartRepo) Clear(userID uint) error {
	query := `DELETE FROM cart_items WHERE user_id = $1`
	_, err := config.DB.Exec(query, userID)
	if err != nil {
		log.Println("Error clearing cart:", err)
		return err
	}
	return nil
}
