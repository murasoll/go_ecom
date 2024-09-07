// api/v1/products.go

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

var productService *services.ProductService

func InitProductHandlers(ps *services.ProductService) {
	productService = ps
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	filter := make(map[string]interface{})
	if category := r.URL.Query().Get("category"); category != "" {
		filter["category_id"] = category
	}
	if search := r.URL.Query().Get("search"); search != "" {
		filter["search"] = search
	}

	products, err := productService.GetAllProducts(filter)
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, products, http.StatusOK)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid product ID"}, http.StatusBadRequest)
		return
	}

	product, err := productService.GetProductByID(uint(id))
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusNotFound)
		return
	}

	helpers.JsonResponse(w, product, http.StatusOK)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid request payload"}, http.StatusBadRequest)
		return
	}

	if err := productService.CreateProduct(&product); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, product, http.StatusCreated)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid product ID"}, http.StatusBadRequest)
		return
	}

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid request payload"}, http.StatusBadRequest)
		return
	}
	product.ID = uint(id)

	if err := productService.UpdateProduct(&product); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, product, http.StatusOK)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid product ID"}, http.StatusBadRequest)
		return
	}

	if err := productService.DeleteProduct(uint(id)); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, map[string]string{"message": "Product deleted successfully"}, http.StatusOK)
}
