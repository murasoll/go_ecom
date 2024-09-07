// repositories/city_repo/city_repo.go

package city_repo

import (
	"ecomerce/config"
	"ecomerce/models"
	"log"
)

type CityRepository interface {
	GetAll() ([]models.City, error)
	GetByID(id uint) (models.City, error)
	Create(city *models.City) error
	Update(city *models.City) error
	Delete(id uint) error
}

type cityRepo struct{}

func NewCityRepo() CityRepository {
	return &cityRepo{}
}

func (r *cityRepo) GetAll() ([]models.City, error) {
	var cities []models.City
	query := `SELECT id, name, shipping_cost FROM cities`
	rows, err := config.DB.Query(query)
	if err != nil {
		log.Println("Error fetching cities:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.City
		err := rows.Scan(&c.ID, &c.Name, &c.ShippingCost)
		if err != nil {
			log.Println("Error scanning city:", err)
			return nil, err
		}
		cities = append(cities, c)
	}

	return cities, nil
}

func (r *cityRepo) GetByID(id uint) (models.City, error) {
	var city models.City
	query := `SELECT id, name, shipping_cost FROM cities WHERE id = $1`
	err := config.DB.QueryRow(query, id).Scan(&city.ID, &city.Name, &city.ShippingCost)
	if err != nil {
		log.Println("Error fetching city by ID:", err)
		return city, err
	}
	return city, nil
}

func (r *cityRepo) Create(city *models.City) error {
	query := `INSERT INTO cities (name, shipping_cost) VALUES ($1, $2) RETURNING id`
	err := config.DB.QueryRow(query, city.Name, city.ShippingCost).Scan(&city.ID)
	if err != nil {
		log.Println("Error creating city:", err)
		return err
	}
	return nil
}

func (r *cityRepo) Update(city *models.City) error {
	query := `UPDATE cities SET name = $1, shipping_cost = $2 WHERE id = $3`
	_, err := config.DB.Exec(query, city.Name, city.ShippingCost, city.ID)
	if err != nil {
		log.Println("Error updating city:", err)
		return err
	}
	return nil
}

func (r *cityRepo) Delete(id uint) error {
	query := `DELETE FROM cities WHERE id = $1`
	_, err := config.DB.Exec(query, id)
	if err != nil {
		log.Println("Error deleting city:", err)
		return err
	}
	return nil
}
