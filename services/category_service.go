package services

import (
	"ecomerce/models"
	"ecomerce/repositories/category_repo"
)

type CategoryService struct {
	categoryRepo category_repo.CategoryRepository
}

func NewCategoryService(cr category_repo.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepo: cr}
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.categoryRepo.GetAll()
}

func (s *CategoryService) GetCategoryByID(id uint) (models.Category, error) {
	return s.categoryRepo.GetByID(id)
}

func (s *CategoryService) CreateCategory(category *models.Category) error {
	return s.categoryRepo.Create(category)
}

func (s *CategoryService) UpdateCategory(category *models.Category) error {
	return s.categoryRepo.Update(category)
}

func (s *CategoryService) DeleteCategory(id uint) error {
	return s.categoryRepo.Delete(id)
}
