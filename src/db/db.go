// Package db
package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPsqlStorage(connString string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse connection string: %v\n", err)
		os.Exit(1)
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	return dbpool
}

func InitStorage(db *pgxpool.Pool) {
	err := db.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to establish db connection: %v\n", err)
		os.Exit(1)
	}

	log.Println("🛢️ Database connected successfully")
}
