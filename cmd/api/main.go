package main

import (
	"context"
	"fmt"
	"go-migrate/internal/db"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go + Postgres + Redis is running 🚀")
}

func main() {

	log.Println("🚀 Starting app...")

	// Connect to Postgres
	dsn := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer conn.Close(context.Background())

	database := db.Connect(dsn)
	defer database.Close()
	log.Println("Database connected ✅")

	db.RunMigrations() // รัน migration
	db.SeedUsers()     // รัน seed ข้อมูล

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	log.Println("🎉 App ready")

	// TODO: start HTTP server
}
