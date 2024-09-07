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

var shippingService *services.ShippingService

func InitShippingHandlers(ss *services.ShippingService) {
	shippingService = ss
}

func GetShippingCost(w http.ResponseWriter, r *http.Request) {
	cityID, err := strconv.Atoi(r.URL.Query().Get("city_id"))
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid city ID"}, http.StatusBadRequest)
		return
	}

	cost, err := shippingService.GetShippingCost(uint(cityID))
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, map[string]float64{"shipping_cost": cost}, http.StatusOK)
}

func GetAllCities(w http.ResponseWriter, r *http.Request) {
	cities, err := shippingService.GetAllCities()
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, cities, http.StatusOK)
}

func CreateCity(w http.ResponseWriter, r *http.Request) {
	var city models.City
	if err := json.NewDecoder(r.Body).Decode(&city); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid request payload"}, http.StatusBadRequest)
		return
	}

	if err := shippingService.CreateCity(&city); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, city, http.StatusCreated)
}

func UpdateCity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid city ID"}, http.StatusBadRequest)
		return
	}

	var city models.City
	if err := json.NewDecoder(r.Body).Decode(&city); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid request payload"}, http.StatusBadRequest)
		return
	}
	city.ID = uint(id)

	if err := shippingService.UpdateCity(&city); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, city, http.StatusOK)
}

func DeleteCity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid city ID"}, http.StatusBadRequest)
		return
	}

	if err := shippingService.DeleteCity(uint(id)); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, map[string]string{"message": "City deleted successfully"}, http.StatusOK)
}
