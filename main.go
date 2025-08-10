package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var now string
		err := db.QueryRow("SELECT NOW()").Scan(&now)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprintf(w, "Hello from Go! Time in DB: %s", now)
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
