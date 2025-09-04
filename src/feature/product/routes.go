package product

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	repo ProductRepo
}

func NewHandler(repo ProductRepo) *Handler {
	return &Handler{repo}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.GetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", h.CreateProduct).Methods(http.MethodPost)
}
