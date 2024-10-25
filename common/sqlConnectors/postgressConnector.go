package sqlConnectors

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"time"
)

type postgresConnector struct {
}

func CreateNewPostgresConnection(dsn string) *pgxpool.Pool {
	counts := 0
	for {
		conn, err := pgxpool.New(context.Background(), dsn)
		if err != nil {
			slog.Info("Couldn't connect to database", err)
			counts++
		} else {
			slog.Info("Successfully connected to database")
			return conn
		}
		if counts == 5 {
			slog.Error("unable to connect to postgres: %s", err)
			slog.Info("retying in 5 seconds...")
			time.Sleep(time.Second * 5)
			counts = 0
		}
	}
}
