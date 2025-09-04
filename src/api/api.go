// Package api
package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MKSinghDev/go-ecom/src/config"
	"github.com/MKSinghDev/go-ecom/src/feature/product"
	"github.com/MKSinghDev/go-ecom/src/feature/user"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type APIServer struct {
	addr   string
	dbpool *pgxpool.Pool
}

func NewAPIServer(addr string, db *pgxpool.Pool) *APIServer {
	return &APIServer{
		addr,
		db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userRepo := user.NewRepo(s.dbpool)
	user.NewHandler(userRepo).RegisterRoutes(subrouter)

	productRepo := product.NewRepo(s.dbpool)
	product.NewHandler(productRepo).RegisterRoutes(subrouter)

	log.Printf("🚀 Server listening at %s:%s", config.Envs.PublicHost, s.addr)
	return http.ListenAndServe(fmt.Sprintf(":%s", s.addr), router)
}
