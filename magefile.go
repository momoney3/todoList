//go:build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

func Run() error {
	fmt.Print("It is now runing ")
	return sh.Run("go", "run", "cmd/main.go")
}

func Test() error {
	fmt.Print("Testing project")
	return sh.Run("go", "test", "./...")
}

func Build() error {
	fmt.Print("Building project")
	return sh.Run("go", "build", "-o", "todo", "-v", "-a", "cmd/api/")
}
