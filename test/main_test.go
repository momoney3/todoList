package main

import (
	"os"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CompletedAt *time.Time
	CreatedAt   time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) delete(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	t := *todos

	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) toggle(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	t := *todos
	todo := &t[index]

	if !todo.Completed {
		completedTime := time.Now()
		todo.CompletedAt = &completedTime
	} else {
		todo.CompletedAt = nil
	}

	todo.Completed = !todo.Completed
	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowlines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "x"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(
			strconv.IToa(index),
			t.Title,
			completed,
			t.CreatedAt.Format(time.RFC1123),
		)
	}
	table.Render()
}

// func main() {
// 	todos := Todos{}
// 	storage := NewStorage[Todos]("todos.json")
// 	err := storage.Load(&todos)
// 	if err != nil {
// 		fmt.Println("Warning: Could not load todos from storage. Starting fresh todos.")
// 	}
//
// 	cmdFlags := NewCmdFlags()
// 	cmdFlags.Execute(&todos)
//
// 	err = storage.Save(todos)
// 	if err != nil {
// 		fmt.Printf("Error saving todos in storage: %v\n", err)
// 	}
// }
