package postgres

import (
	"context"
	"embed"
	"errors"

	"github.com/Vivi3008/apiTestGolang/commom/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v4/pgxpool"
)

//go:embed migrations
var fs embed.FS //nolint:gochecknoglobals

func ConnectPool(ctx context.Context, cfg config.Config) (*pgxpool.Pool, error) {
	conn, err := pgxpool.Connect(ctx, cfg.URL())

	if err != nil {
		return nil, err
	}

	var greeting string
	err = conn.QueryRow(ctx, "select 'Connect to database sucessfully'").Scan(&greeting)
	if err != nil {
		return nil, err
	}

	err = RunMigrations(cfg.URL(), fs)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func RunMigrations(databaseUrl string, fs embed.FS) error {
	d, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}
	m, err := migrate.NewWithSourceInstance("iofs", d, databaseUrl)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		} else {
			return err
		}
	}
	return nil
}
