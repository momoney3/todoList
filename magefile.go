//go:build mage

package main

import (
	"fmt"

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

func Build() error {
	fmt.Println("Building project")
	return sh.Run("go", "build", "-o", "todo", "-v", "-a", "cmd/api/")
}
