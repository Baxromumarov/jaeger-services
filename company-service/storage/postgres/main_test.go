package postgres

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"jaeger-services/company-service/config"
)

var (
	db     *pgxpool.Pool
	dbPool *Pool
)

func TestMain(m *testing.M) {
	cfg := config.Load()
	conf, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		panic(err)
	}

	conf.MaxConns = cfg.PostgresMaxConnections

	db, err = pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		panic(err)
	}

	dbPool = &Pool{
		db: db,
	}

	os.Exit(m.Run())
}
