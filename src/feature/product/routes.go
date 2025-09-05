package product

import (
	"net/http"

	"github.com/MKSinghDev/go-ecom/src/interfaces"
	"github.com/gorilla/mux"
)

type Handler struct {
	repo interfaces.ProductRepo
}

func NewHandler(repo interfaces.ProductRepo) *Handler {
	return &Handler{repo}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.GetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", h.CreateProduct).Methods(http.MethodPost)
}
