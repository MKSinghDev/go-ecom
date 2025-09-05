// Package product
package product

import (
	"context"
	"fmt"

	"github.com/MKSinghDev/go-ecom/src/interfaces"
	"github.com/MKSinghDev/go-ecom/src/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	dbpool *pgxpool.Pool
}

func NewRepo(dbpool *pgxpool.Pool) *Repo {
	return &Repo{dbpool}
}

func (r *Repo) GetProducts() ([]interfaces.Product, error) {
	rows, err := r.dbpool.Query(context.Background(), "SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]interfaces.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (r *Repo) GetProductsByIDs(ids []int) ([]interfaces.Product, error) {
	pgPlaceholders := utils.BuildPostgreSQLPlaceholders(ids)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (%s)", pgPlaceholders)

	args := make([]any, len(ids))
	for i, v := range ids {
		args[i] = v
	}

	rows, err := r.dbpool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}

	products := []interfaces.Product{}
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (r *Repo) CreateProduct(product interfaces.CreateProductPayload) error {
	_, err := r.dbpool.Exec(
		context.Background(),
		"INSERT INTO products (name, description, image, price, quantity) VALUES ($1,$2,$3,$4,$5)",
		product.Name, product.Description, product.Image, product.Price, product.Quantity,
	)

	return err
}

func (r *Repo) UpdateProduct(product interfaces.Product) error {
	_, err := r.dbpool.Exec(
		context.Background(),
		"UPDATE products SET name = $1, price = $2, image = $3, description = $4, quantity = $5 WHERE id = $6",
		product.Name, product.Price, product.Image, product.Description, product.Quantity, product.ID,
	)
	return err
}

func scanRowsIntoProduct(rows pgx.Rows) (*interfaces.Product, error) {
	product := new(interfaces.Product)

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
