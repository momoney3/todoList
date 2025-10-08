package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type TodoList struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	COMPLETED bool      `json:"completed"`
	CREATEDAT time.Time `json:"created_at"`
}

type TodoService struct {
	db *sql.DB
}

func (s *TodoService) QueryAllList(ctx context.Context) ([]TodoList, error) {
	queru := "SELECT * FROM todolist"

	rows, err := s.db.QueryContext(ctx, queru)
	if err != nil {
		return nil, fmt.Errorf("falied to query todos: %w", err)
	}
	defer rows.Close()

	var todos []TodoList
	for rows.Next() {
		var todo TodoList
		err := rows.Scan(&todo.ID, &todo.Title, &todo.COMPLETED, &todo.CREATEDAT)
		if err != nil {
			return nil, fmt.Errorf("failed to scan todo: %w", err)
		}
		todos = append(todos, todo)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration err: %w", err)
	}
	return todos, nil
}

func main() {
	dbName := os.Getenv("DATABASE_URL")
	dbToken := os.Getenv("AUTH_TOKEN")

	if dbName == "" || dbToken == "" {
		log.Printf("URL: %s", dbName)
		log.Printf("Token: %s", dbToken[:8])
		log.Fatal("Database URL and authToken are needed")
		// panic("Database or Token are missing")
	}

	url := fmt.Sprintf("%s?authToken=%s", dbName, dbToken)

	log.Println("Connecting to database....")
	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatalf("Failed to open database %v", err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/data", func(w http.ResponseWriter, r *http.Request) {
		databases, err := TodoList(db)
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
