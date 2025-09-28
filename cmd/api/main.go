package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type Database struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	COMPLETED bool      `json:"completed"`
	CREATEDAT time.Time `json:"created_at"`
}

func queryUsers(db *sql.DB) ([]Database, error) {
	// rows, err := db.Query("SELECT id, title FROM todolist")
	rows, err := db.Query("SELECT * FROM todolist")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var databases []Database

	for rows.Next() {
		var database Database
		if err := rows.Scan(&database.ID, &database.Title, &database.COMPLETED, &database.CREATEDAT); err != nil {
			return nil, err
		}
		fmt.Println("Row:", database)
		databases = append(databases, database)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return databases, nil
}

func main() {
	dbName := os.Getenv("DATABASE_URL")
	dbToken := os.Getenv("AUTH_TOKEN")

	if dbName == "" || dbToken == "" {
		fmt.Printf("db: %s and token: %s\n", dbName, dbToken)
		panic("Database or Token are missing")
	}

	url := fmt.Sprintf("%s?authToken=%s", dbName, dbToken)

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		databases, err := queryUsers(db)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to query database error code %s", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "applcatoin/json")
		json.NewEncoder(w).Encode(databases)
	})

	fmt.Println("Server running on http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
