package cart

import (
	"context"

	"github.com/MKSinghDev/go-ecom/src/interfaces"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	dbpool *pgxpool.Pool
}

func NewRepo(dbpool *pgxpool.Pool) *Repo {
	return &Repo{dbpool}
}

func (r *Repo) CreateOrder(order interfaces.Order) (int, error) {
	var id int
	err := r.dbpool.QueryRow(
		context.Background(),
		"INSERT INTO orders (userId, total, status, address) VALUES ($1, $2, $3, $4)",
		order.UserID, order.Total, order.Status, order.Address,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *Repo) CreateOrderItem(orderItem interfaces.OrderItem) error {
	_, err := r.dbpool.Exec(
		context.Background(),
		"INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)",
		orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price,
	)
	return err
}
