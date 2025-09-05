// Package cart
package cart

import (
	"net/http"

	"github.com/MKSinghDev/go-ecom/src/interfaces"
	"github.com/MKSinghDev/go-ecom/src/service/auth"
	"github.com/gorilla/mux"
)

type Handler struct {
	orderRepo   interfaces.OrderRepo
	productRepo interfaces.ProductRepo
	userRepo    interfaces.UserRepo
}

func NewHandler(orderRepo interfaces.OrderRepo, productRepo interfaces.ProductRepo, userRepo interfaces.UserRepo) *Handler {
	return &Handler{orderRepo, productRepo, userRepo}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", auth.WithJWTAuth(h.Checkout, h.userRepo)).Methods(http.MethodPost)
}
