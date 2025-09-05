package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MKSinghDev/go-ecom/src/interfaces"
	"github.com/gorilla/mux"
)

type mockProductRepo struct{}

func TestProductHandler(t *testing.T) {
	productRepo := &mockProductRepo{}
	handler := NewHandler(productRepo)

	t.Run("should fail if the product payload is invalid", func(t *testing.T) {
		payload := interfaces.CreateProductPayload{
			Name:        "ip",
			Description: "des",
			Image:       "img",
			Price:       -84.3,
			Quantity:    -1,
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.CreateProduct)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Error(fmt.Printf("expected status code %d, got %d", http.StatusBadRequest, rr.Code))
		}
	})

	t.Run("should correctly create the product", func(t *testing.T) {
		payload := interfaces.CreateProductPayload{
			Name:        "MacBook Air M4",
			Description: "This is the best macbook for software development",
			Image:       "http://localhost.com/images/image.png",
			Price:       1065,
			Quantity:    15,
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.CreateProduct)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Error(fmt.Printf("expected status code %d, got %d", http.StatusCreated, rr.Code))
		}
	})
}

func (r *mockProductRepo) GetProducts() ([]interfaces.Product, error) {
	return nil, nil
}

func (r *mockProductRepo) GetProductsByIDs(ids []int) ([]interfaces.Product, error) {
	return nil, nil
}

func (r *mockProductRepo) CreateProduct(product interfaces.CreateProductPayload) error {
	return nil
}

func (r *mockProductRepo) UpdateProduct(product interfaces.Product) error {
	return nil
}
