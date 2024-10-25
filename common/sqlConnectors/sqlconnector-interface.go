package sqlConnectors

import "github.com/jackc/pgx/v5/pgxpool"

type SqlConnectors interface {
	CreateNewPostgresConnection(dsn string) *pgxpool.Pool
}
