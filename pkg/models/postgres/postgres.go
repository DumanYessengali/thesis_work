package postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type PgModel struct {
	Pool *pgxpool.Pool
}
