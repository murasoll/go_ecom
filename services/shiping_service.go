package services

import (
	"ecomerce/models"
	"ecomerce/repositories/city_repo"
)

type ShippingService struct {
	cityRepo city_repo.CityRepository
}

func NewShippingService(cr city_repo.CityRepository) *ShippingService {
	return &ShippingService{cityRepo: cr}
}

func (s *ShippingService) GetShippingCost(cityID uint) (float64, error) {
	city, err := s.cityRepo.GetByID(cityID)
	if err != nil {
		return 0, err
	}
	return city.ShippingCost, nil
}

func (s *ShippingService) GetAllCities() ([]models.City, error) {
	return s.cityRepo.GetAll()
}

func (s *ShippingService) CreateCity(city *models.City) error {
	return s.cityRepo.Create(city)
}

func (s *ShippingService) UpdateCity(city *models.City) error {
	return s.cityRepo.Update(city)
}

func (s *ShippingService) DeleteCity(id uint) error {
	return s.cityRepo.Delete(id)
}
