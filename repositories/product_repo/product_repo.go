package product_repo

import (
	"ecomerce/config"
	"ecomerce/models"
	"fmt"
	"log"
)

type ProductRepository interface {
	GetAll(filter map[string]interface{}) ([]models.Product, error)
	GetByID(id uint) (models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id uint) error
	UpdateInventory(id uint, quantity int) error
}

type productRepo struct{}

func NewProductRepo() ProductRepository {
	return &productRepo{}
}

func (r *productRepo) GetAll(filter map[string]interface{}) ([]models.Product, error) {
	var products []models.Product
	query := `SELECT id, name, description, price, category_id, inventory FROM products`

	if categoryID, ok := filter["category_id"]; ok {
		query += fmt.Sprintf(" WHERE category_id = %v", categoryID)
	}

	if search, ok := filter["search"]; ok {
		if _, ok := filter["category_id"]; ok {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += fmt.Sprintf(" name ILIKE '%%%s%%'", search)
	}

	rows, err := config.DB.Query(query)
	if err != nil {
		log.Println("Error fetching products:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.CategoryID, &p.Inventory)
		if err != nil {
			log.Println("Error scanning product:", err)
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *productRepo) GetByID(id uint) (models.Product, error) {
	var product models.Product
	query := `SELECT id, name, description, price, category_id, inventory FROM products WHERE id = $1`
	err := config.DB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.Inventory)
	if err != nil {
		log.Println("Error fetching product by ID:", err)
		return product, err
	}
	return product, nil
}

func (r *productRepo) Create(product *models.Product) error {
	query := `INSERT INTO products (name, description, price, category_id, inventory) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := config.DB.QueryRow(query, product.Name, product.Description, product.Price, product.CategoryID, product.Inventory).Scan(&product.ID)
	if err != nil {
		log.Println("Error creating product:", err)
		return err
	}
	return nil
}

func (r *productRepo) Update(product *models.Product) error {
	query := `UPDATE products SET name = $1, description = $2, price = $3, category_id = $4, inventory = $5 WHERE id = $6`
	_, err := config.DB.Exec(query, product.Name, product.Description, product.Price, product.CategoryID, product.Inventory, product.ID)
	if err != nil {
		log.Println("Error updating product:", err)
		return err
	}
	return nil
}

func (r *productRepo) Delete(id uint) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := config.DB.Exec(query, id)
	if err != nil {
		log.Println("Error deleting product:", err)
		return err
	}
	return nil
}

func (r *productRepo) UpdateInventory(id uint, quantity int) error {
	query := `UPDATE products SET inventory = inventory + $1 WHERE id = $2`
	_, err := config.DB.Exec(query, quantity, id)
	if err != nil {
		log.Println("Error updating product inventory:", err)
		return err
	}
	return nil
}
