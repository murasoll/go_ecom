package v1

import (
	"ecomerce/helpers"
	"ecomerce/models"
	"ecomerce/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var categoryService *services.CategoryService

func InitCategoryHandlers(cs *services.CategoryService) {
	categoryService = cs
}

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := categoryService.GetAllCategories()
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, categories, http.StatusOK)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid category ID"}, http.StatusBadRequest)
		return
	}

	category, err := categoryService.GetCategoryByID(uint(id))
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusNotFound)
		return
	}

	helpers.JsonResponse(w, category, http.StatusOK)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid request payload"}, http.StatusBadRequest)
		return
	}

	if err := categoryService.CreateCategory(&category); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, category, http.StatusCreated)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid category ID"}, http.StatusBadRequest)
		return
	}

	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid request payload"}, http.StatusBadRequest)
		return
	}
	category.ID = uint(id)

	if err := categoryService.UpdateCategory(&category); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, category, http.StatusOK)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid category ID"}, http.StatusBadRequest)
		return
	}

	if err := categoryService.DeleteCategory(uint(id)); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, map[string]string{"message": "Category deleted successfully"}, http.StatusOK)
}
