//go:build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	BinaryName = "myapp"
)

// Default target executed when `mage` is run without specifying a target.
var Default = Build

func Build() error {
	fmt.Println("Building binary...")
	return sh.Run("go", "build", "cmd/main.go", "-o", BinaryName)
}

// Run executes the compiled Go program. It depends on Build completing first.
func Run() error {
	mg.Deps(Build) // Ensure the binary is built before running
	return sh.Run("./" + BinaryName)
}

func Clean() error {
	fmt.Printf("Cleaning up %s...\n", BinaryName)
	err := os.Remove(BinaryName)
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

func DockerBuild() error {
	fmt.Println("Building Docker image...")
	return sh.Run("docker", "build", "-t", "backend:latest", ".")
}
