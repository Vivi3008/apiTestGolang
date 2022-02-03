package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

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

	migration, err := migrate.New(
		"file://../../migrations",
		databaseUrl)

	if err != nil {
		log.Fatalf("Could not read migration files: %s", err)
	}

	if err := migration.Up(); err != nil {
		log.Fatal(err)
	}

	// tearDown should be called to destroy container at the end of the test
	tearDown := func() {
		db.Close()
		dockerPool.Purge(resource)
	}

	return db, tearDown
}
