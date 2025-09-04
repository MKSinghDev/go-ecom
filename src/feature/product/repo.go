// Package product
package product

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	dbpool *pgxpool.Pool
}

func NewRepo(dbpool *pgxpool.Pool) *Repo {
	return &Repo{dbpool}
}

func (r *Repo) GetProducts() ([]Product, error) {
	rows, err := r.dbpool.Query(context.Background(), "SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (r *Repo) CreateProduct(product CreateProductPayload) error {
	_, err := r.dbpool.Exec(
		context.Background(),
		"INSERT INTO products (name, description, image, price, quantity) VALUES ($1,$2,$3,$4,$5)",
		product.Name, product.Description, product.Image, product.Price, product.Quantity,
	)

	return err
}

func scanRowsIntoProduct(rows pgx.Rows) (*Product, error) {
	product := new(Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}
