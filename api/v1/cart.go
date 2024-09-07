package v1

import (
	"ecomerce/helpers"
	"ecomerce/services"
	"encoding/json"
	"net/http"
)

var cartService *services.CartService

func InitCartHandlers(cs *services.CartService) {
	cartService = cs
}

func GetCart(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	cart, err := cartService.GetCart(userID)
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, cart, http.StatusOK)
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
	var item struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid request payload"}, http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(uint)
	if err := cartService.AddToCart(userID, item.ProductID, item.Quantity); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, map[string]string{"message": "Item added to cart successfully"}, http.StatusOK)
}

func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	var item struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid request payload"}, http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(uint)
	if err := cartService.UpdateCartItemQuantity(userID, item.ProductID, item.Quantity); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, map[string]string{"message": "Cart item updated successfully"}, http.StatusOK)
}

func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	var item struct {
		ProductID uint `json:"product_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid request payload"}, http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(uint)
	if err := cartService.RemoveFromCart(userID, item.ProductID); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, map[string]string{"message": "Item removed from cart successfully"}, http.StatusOK)
}
