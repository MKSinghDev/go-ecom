package user

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

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user paload is invalid", func(t *testing.T) {
		payload := interfaces.RegisterUserPayload{
			FirstName: "user",
			LastName:  "123",
			Email:     "invalid",
			Password:  "asd983jl",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.HandleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Error(fmt.Printf("expected status code %d, got %d", http.StatusBadRequest, rr.Code))
		}
	})

	t.Run("should correctly register the user", func(t *testing.T) {
		payload := interfaces.RegisterUserPayload{
			FirstName: "user",
			LastName:  "123",
			Email:     "asd@gmail.com",
			Password:  "asd983jl",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.HandleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Error(fmt.Printf("expected status code %d, got %d", http.StatusCreated, rr.Code))
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*interfaces.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserByID(id int) (*interfaces.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user interfaces.RegisterUserPayload) error {
	return nil
}

