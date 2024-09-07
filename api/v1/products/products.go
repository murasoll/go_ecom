package products

import (
	"ecomerce/helpers"
	"ecomerce/models"
	"ecomerce/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products := services.GetProducts()
	helpers.JsonResponse(w, map[string][]models.Product{
		"products": products,
	}, http.StatusOK)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid request body"}, http.StatusBadRequest)
		return
	}

	product, err = services.CreateProduct(product)
	if err != nil {
		helpers.JsonResponse(w, map[string]any{
			"message": "error creating product",
		}, http.StatusInternalServerError)
	}
	helpers.JsonResponse(w, map[string]any{
		"message": "created successfully",
		"product": product,
	}, http.StatusOK)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	product, err := services.GetProductByID(id)
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid product ID"}, http.StatusBadRequest)
		return
	}
	helpers.JsonResponse(w, product, http.StatusOK)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid product ID"}, http.StatusBadRequest)
		return
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid request body"}, http.StatusBadRequest)
		return
	}

	product.ID = id // Assign the ID from the URL to the product

	err = services.UpdateProduct(product)
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Error updating product"}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, map[string]any{
		"message": "Product updated successfully",
		"product": product,
	}, http.StatusOK)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := services.DeleteProduct(id)
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Error Deleting product"}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, map[string]any{
		"message": "Product Deleted successfully",
	}, http.StatusOK)
}
