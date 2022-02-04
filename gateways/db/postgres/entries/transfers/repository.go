package transfers

import "github.com/jackc/pgx/v4/pgxpool"

type Repository struct {
	Db *pgxpool.Pool
}

func NewRepository(pgx *pgxpool.Pool) *Repository {
	return &Repository{
		Db: pgx,
	}
}
