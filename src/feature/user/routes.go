// Package user
package user

import (
	"github.com/gorilla/mux"
)

type Handler struct {
	repo UserRepo
}

func NewHandler(repo UserRepo) *Handler {
	return &Handler{
		repo,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.HandleLogin).Methods("POST")
	router.HandleFunc("/register", h.HandleRegister).Methods("POST")
}
