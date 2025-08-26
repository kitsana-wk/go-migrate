package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func SeedUsers() {
	dsn := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), `
        INSERT INTO users (name, email, phone)
        VALUES
            ('Alice', 'alice@example.com', '0812345678'),
            ('Bob', 'bob@example.com', '0898765432')
        ON CONFLICT (email) DO NOTHING;
    `)
	if err != nil {
		log.Fatalf("failed to seed users: %v", err)
	}

	fmt.Println("✅ Users seeded successfully")
}
