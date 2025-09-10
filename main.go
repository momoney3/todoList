package main

import "time"

type Todo struct {
	Title     string
	Completed bool
	CreatedAt time.time
	Completed *time.time
}

type Todos []Todo

func (t Todos) add(title string) {
	todo := Todo{
		Tile:      title,
		Completed: false,
		CreatedAt: nil,
		Completed: time.Now(),
	}
	*todos = append(*todos, todo)
}

func (t Todos) delete(index int) error {
	t := *Todos
}

func main() {
}
