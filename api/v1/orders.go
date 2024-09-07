// api/v1/orders.go

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

var orderService *services.OrderService

func InitOrderHandlers(os *services.OrderService) {
	orderService = os
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderRequest struct {
		ShippingAddress models.ShippingAddress `json:"shipping_address"`
	}
	if err := json.NewDecoder(r.Body).Decode(&orderRequest); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid request payload"}, http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(uint)
	order, err := orderService.CreateOrder(userID, orderRequest.ShippingAddress)
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, order, http.StatusCreated)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	orders, err := orderService.GetOrdersByUser(userID)
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, orders, http.StatusOK)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid order ID"}, http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(uint)
	order, err := orderService.GetOrderByID(uint(orderID), userID)
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusNotFound)
		return
	}

	helpers.JsonResponse(w, order, http.StatusOK)
}

func UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid order ID"}, http.StatusBadRequest)
		return
	}

	var statusUpdate struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&statusUpdate); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid request payload"}, http.StatusBadRequest)
		return
	}

	if err := orderService.UpdateOrderStatus(uint(orderID), statusUpdate.Status); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, map[string]string{"message": "Order status updated successfully"}, http.StatusOK)
}
