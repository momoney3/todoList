//go:build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

func Run() error {
	fmt.Print("It is now runing ")
	return sh.Run("go", "run", "main.go")
}

func Test() error {
	fmt.Print("Testing project")
	return sh.Run("go", "test", "./...")
}
