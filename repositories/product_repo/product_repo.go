package product_repo

import (
	"database/sql"
	"ecomerce/config"
	"ecomerce/models"
	"log"
)

func GetAllProducts() []models.Product {
	var products []models.Product
	query := `SELECT id, name, price FROM products`

	rows, err := config.DB.Query(query)
	if err != nil {
		log.Println("Error fetching products:", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			log.Println("Error scanning product:", err)
			continue
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		log.Println("Rows error:", err)
	}

	return products
}

func CreateProduct(product models.Product) (models.Product, error) {
	query := `INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id`
	err := config.DB.QueryRow(query, product.Name, product.Price).Scan(&product.ID)
	if err != nil {
		log.Println("Error creating product:", err)
		return product, err
	}

	log.Println("Product created with ID:", product.ID)
	return product, nil
}

func GetProductByID(id string) (models.Product, error) {
	var product models.Product
	query := `SELECT id, name, price FROM products WHERE id = $1`

	err := config.DB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No product found with the given ID")
			return product, err
		}
		log.Println("Error fetching product by ID:", err)
		return product, err
	}

	return product, nil
}

func UpdateProduct(product models.Product) error {
	query := `UPDATE products SET name = $1, price = $2 WHERE id = $3`
	_, err := config.DB.Exec(query, product.Name, product.Price, product.ID)
	if err != nil {
		log.Println("Error updating product:", err)
		return err
	}

	log.Println("Product updated with ID:", product.ID)
	return nil
}

func DeleteProduct(id string) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := config.DB.Exec(query, id)
	if err != nil {
		log.Println("Error updating product:", err)
		return err
	}

	log.Println("Product Deleted with ID:", id)
	return nil
}
