package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PoolConfig struct {
	MaxConnections  int
	MinConnections  int
	MaxConnLifetime time.Duration
	MaxConnIdleTime time.Duration
}

func ConnectPool(ctx context.Context, url string, config PoolConfig) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = int32(config.MaxConnections)
	poolConfig.MinConns = int32(config.MinConnections)
	poolConfig.MaxConnLifetime = config.MaxConnLifetime
	poolConfig.MaxConnIdleTime = config.MaxConnIdleTime

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func Ping(ctx context.Context, pool *pgxpool.Pool) error {
	var one int
	return pool.QueryRow(ctx, "SELECT 1").Scan(&one)
}

func GetStats(pool *pgxpool.Pool) *pgxpool.Stat {
	return pool.Stat()
}
