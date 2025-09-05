// Package user
package user

import (
	"context"
	"fmt"

	"github.com/MKSinghDev/go-ecom/src/interfaces"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	dbpool *pgxpool.Pool
}

func NewRepo(dbpool *pgxpool.Pool) *Store {
	return &Store{
		dbpool,
	}
}

func (ur *Store) GetUserByEmail(email string) (*interfaces.User, error) {
	rows, err := ur.dbpool.Query(context.Background(), "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	u := new(interfaces.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func (ur *Store) GetUserByID(id int) (*interfaces.User, error) {
	rows, err := ur.dbpool.Query(context.Background(), "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	u := new(interfaces.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func (ur *Store) CreateUser(user interfaces.RegisterUserPayload) error {
	_, err := ur.dbpool.Exec(
		context.Background(),
		"INSERT INTO users (first_name, last_name, email, password) VALUES ($1,$2,$3,$4)",
		user.FirstName, user.LastName, user.Email, user.Password,
	)
	return err
}

func scanRowIntoUser(rows pgx.Rows) (*interfaces.User, error) {
	user := new(interfaces.User)
	err := rows.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
