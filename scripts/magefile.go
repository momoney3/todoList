//go:build mage

package main

import "github.com/magefile/mage/sh"

func run() {
	return sh.Run("go", "run", "./...")
}
