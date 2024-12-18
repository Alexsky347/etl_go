package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"go-ddd/pkg/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := sql.Open("postgres", fmt.Sprintf("%s?sslmode=disable", cfg.Dsn))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/database/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully")
}
