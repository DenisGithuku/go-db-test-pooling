package config

import (
	"context"
	"fmt"
	"time"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context) (*pgxpool.Pool, error) {
	user := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PASS")
	dbName := os.Getenv("PG_DB")
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, dbName)
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	// pool tuning
	config.MaxConns = 20
	config.MinConns = 5
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute
	config.HealthCheckPeriod = 1 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}
	return pool, nil
}