package interfaces

import "time"

type ProductRepo interface {
	GetProducts() ([]Product, error)
	GetProductsByIDs(ids []int) ([]Product, error)
	CreateProduct(product CreateProductPayload) error
	UpdateProduct(product Product) error
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateProductPayload struct {
	Name        string  `json:"name" validate:"required,min=3,max=255"`
	Description string  `json:"description" validate:"required,min=5,max=512"`
	Image       string  `json:"image" validate:"url"`
	Price       float64 `json:"price" validate:"required,min=0,max=99999"`
	Quantity    int     `json:"quantity" validate:"required,min=0,max=9999"`
}
