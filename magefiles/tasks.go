//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Build() error {
	return sh.Run("go", "build", "-o", "backend", "cmd/backend/main.go")
}

func Watch() error {
	return sh.RunV("wgo", "mage", "build", "::", "mage", "run")
}

func Run() error {
	return sh.RunV("./backend")
}
