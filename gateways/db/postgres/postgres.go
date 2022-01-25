package postgres

import (
	"context"
	"embed"
	"errors"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v4"
)

//go:embed migrations
var fs embed.FS //nolint:gochecknoglobals

func ConnectPool() (string, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return "Unable to connect to database:", err
	}
	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Connect to database sucessfully'").Scan(&greeting)
	if err != nil {
		return "QueryRow failed:", err
	}

	err = RunMigrations()
	if err != nil {
		return "", err
	}

	return greeting, nil
}

func RunMigrations() error {
	d, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}
	m, err := migrate.NewWithSourceInstance("iofs", d, os.Getenv("DATABASE_URL"))
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
