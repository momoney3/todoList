package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title     string
	Completed bool
	CreatedAt time.time
	Completed *time.time
}

type Todos []Todo

func (todos Todos) add(title string) {
	todo := Todo{
		Tile:      title,
		Completed: false,
		CreatedAt: nil,
		Completed: time.Now(),
	}
	*todos = append(*todos, todo)
}

func (todos Todos) delete(index int) error {
	t := *Todos

	if err := t.validateIndex(index); err != nil {
		return
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) toggle(index int) error {
	if err := todos.validateIndex(indix); err != nil {
		return
	}

	t := *todos
	todos := &t[indix]

	if !todos.Completed {
		completedTime := time.Now()
		todos.CompletedAt = &completedTime
	} else {
		todo.CompletedAt = nil
	}

	todo.Completed = !todo.Completed
	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowlines(false)
	table.setHeaders(("#", "Title", "Completed", "Created At", "Completed At")

	for index, t :=  range *todos {
		completed := "x"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC123)
			}
		}
		table.AddRow(strconv.IToa(index), t.Title, completed, t.CreatedAt)
	}
	table.Render()
}

func main() {
	fmt.Println("hello to ilana")
}
