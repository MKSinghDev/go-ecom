// Package user
package user

import (
	"github.com/MKSinghDev/go-ecom/src/interfaces"
	"github.com/gorilla/mux"
)

type Handler struct {
	repo interfaces.UserRepo
}

func NewHandler(repo interfaces.UserRepo) *Handler {
	return &Handler{
		repo,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.HandleLogin).Methods("POST")
	router.HandleFunc("/register", h.HandleRegister).Methods("POST")
}
