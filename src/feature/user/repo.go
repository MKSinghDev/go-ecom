// Package user
package user

import (
	"context"
	"fmt"

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

func (ur *Store) GetUserByEmail(email string) (*User, error) {
	rows, err := ur.dbpool.Query(context.Background(), "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	u := new(User)
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

func (ur *Store) GetUserByID(id int) (*User, error) {
	rows, err := ur.dbpool.Query(context.Background(), "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	u := new(User)
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

func (ur *Store) CreateUser(user RegisterUserPayload) error {
	_, err := ur.dbpool.Exec(
		context.Background(),
		"INSERT INTO users (first_name, last_name, email, password) VALUES ($1,$2,$3,$4)",
		user.FirstName, user.LastName, user.Email, user.Password,
	)
	if err != nil {
		return err
	}

	return nil
}

func scanRowIntoUser(rows pgx.Rows) (*User, error) {
	user := new(User)
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
