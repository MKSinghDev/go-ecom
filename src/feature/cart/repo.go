package cart

import "github.com/jackc/pgx/v5/pgxpool"

type Repo struct {
	dbpool *pgxpool.Pool
}
