//go:build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
)

func Run() error {
	fmt.Println("It is now runing ")
	return sh.Run("go", "run", "cmd/api/main.go")
}

func Tidy() error {
	fmt.Print("Tidy up ")
	if err := sh.Run("go", "mod", "tidy"); err != nil {
		return err
	}
	fmt.Println("and Cleaned")
	return sh.Run("go", "clean")
}

func Test() error {
	fmt.Println("Testing project")
	return sh.Run("go", "test", "./...")
}

func Build(tregger string) error {
	switch tregger {
	case "pm":
		fmt.Println("Building image")
		return sh.Run("podman", "build", "-t", "todolist-App", ".")

	case "run":

	case "run":
		fmt.Println("Building Go project inside a Podman container")
		return sh.Run(
			"podman", "run", "--rm",
			"-v", fmt.Sprintf("%s:/usr/src/myapp", os.Getenv("PWD")),
			"-w", "/usr/src/myapp",
			"golang:1.25",
			"go", "build", "-v", "./cmd/api",
		)

	case "go":
		fmt.Println("Building go project locally")
		return sh.Run("go", "build", "-o", "todo", "-v", "-a", "./cmd/api")

	case "print":
		fmt.Println("it is working")
		return sh.Run("echo", "it working")

	default:
		return fmt.Errorf("unknown fetcher: %s (expected: build, run, list)", tregger)
	}
}
