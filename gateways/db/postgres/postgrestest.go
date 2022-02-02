package postgres

import (
	"context"
	"embed"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

//go:embed migrations
var fstes embed.FS //nolint:gochecknoglobals

func GetTestPool() (*pgxpool.Pool, func()) {
	var db *pgxpool.Pool

	dockerPool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := dockerPool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13.2-alpine",
		Env: []string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_USER=postgres",
			"POSTGRES_DB=viviBank",
			"listen_addresses = '*'",
		},
	}, func(hc *docker.HostConfig) {
		hc.AutoRemove = true
		hc.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://postgres:secret@%s/viviBank?sslmode=disable", hostAndPort)

	log.Println("Connecting to database on url: ", databaseUrl)

	resource.Expire(60) // Tell docker to hard kill the container in 60 seconds

	dockerPool.MaxWait = 10 * time.Second
	// connects to db in container, with exponential backoff-retry,
	// because the application in the container might not be ready to accept connections yet
	if err = dockerPool.Retry(func() error {
		db, err = pgxpool.Connect(context.Background(), databaseUrl)

		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	err = RunMigrations(databaseUrl, fstes)
	if err != nil {
		log.Fatalf("error run migration: %s", err)
	}
	/* 	d, err := iofs.New(fs, "migrations")
	   	if err != nil {
	   		log.Fatalf("error in migration: %s", err)
	   	}
	   	m, err := migrate.NewWithSourceInstance("iofs", d, databaseUrl)
	   	if err != nil {
	   		log.Fatalf("error new source: %s", err)
	   	}

	   	if err := m.Up(); err != nil {
	   		if !errors.Is(err, migrate.ErrNoChange) {
	   			log.Fatalf("error in migration: %s", err)
	   		}
	   	} */

	// tearDown should be called to destroy container at the end of the test
	tearDown := func() {
		db.Close()
		dockerPool.Purge(resource)
	}

	return db, tearDown
}
