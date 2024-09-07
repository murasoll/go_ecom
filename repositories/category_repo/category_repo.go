// repositories/category_repo/category_repo.go

package category_repo

import (
	"ecomerce/config"
	"ecomerce/models"
	"log"
)

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetByID(id uint) (models.Category, error)
	Create(category *models.Category) error
	Update(category *models.Category) error
	Delete(id uint) error
}

type categoryRepo struct{}

func NewCategoryRepo() CategoryRepository {
	return &categoryRepo{}
}

func (r *categoryRepo) GetAll() ([]models.Category, error) {
	var categories []models.Category
	query := `SELECT id, name FROM categories`
	rows, err := config.DB.Query(query)
	if err != nil {
		log.Println("Error fetching categories:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Category
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			log.Println("Error scanning category:", err)
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func (r *categoryRepo) GetByID(id uint) (models.Category, error) {
	var category models.Category
	query := `SELECT id, name FROM categories WHERE id = $1`
	err := config.DB.QueryRow(query, id).Scan(&category.ID, &category.Name)
	if err != nil {
		log.Println("Error fetching category by ID:", err)
		return category, err
	}
	return category, nil
}

func (r *categoryRepo) Create(category *models.Category) error {
	query := `INSERT INTO categories (name) VALUES ($1) RETURNING id`
	err := config.DB.QueryRow(query, category.Name).Scan(&category.ID)
	if err != nil {
		log.Println("Error creating category:", err)
		return err
	}
	return nil
}

func (r *categoryRepo) Update(category *models.Category) error {
	query := `UPDATE categories SET name = $1 WHERE id = $2`
	_, err := config.DB.Exec(query, category.Name, category.ID)
	if err != nil {
		log.Println("Error updating category:", err)
		return err
	}
	return nil
}

func (r *categoryRepo) Delete(id uint) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := config.DB.Exec(query, id)
	if err != nil {
		log.Println("Error deleting category:", err)
		return err
	}
	return nil
}
