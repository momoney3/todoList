package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var todos = []Todo{
	{ID: 1, Title: "go to school", Completed: true, CreatedAt: time.Now()},
	{ID: 2, Title: "pick up food", Completed: true, CreatedAt: time.Now()},
}

// func getTodo(t *chi.Todo)

func main() {
	url := "libsql://[DATABASE_URL].turso.io?authToken=[TOKEN]"

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":3000", r)
}
