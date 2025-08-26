package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
)

func RunMigrations() {

	dsn := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	driver, err := postgres.WithInstance(Connect(dsn), &postgres.Config{})
	if err != nil {
		log.Fatalf("could not create migrate driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	fmt.Println("✅ Migrations applied successfully")

	// dsn := os.Getenv("DATABASE_URL")
	// db, err := sql.Open("postgres", dsn)
	// if err != nil {
	// 	log.Fatalf("failed to connect database: %v", err)
	// }
	// defer db.Close()

	// driver, err := postgres.WithInstance(db, &postgres.Config{})
	// if err != nil {
	// 	log.Fatalf("could not create migrate driver: %v", err)
	// }

	// m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	// if err != nil {
	// 	log.Fatalf("failed to create migrate instance: %v", err)
	// }

	// if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	// 	log.Fatalf("migration failed: %v", err)
	// }

	// fmt.Println("✅ Migrations applied successfully")
}
