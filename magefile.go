//go:build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	binaryName = "build"
	backendDir = "backend"
	srcDir     = "cmd/main.go"
)

func Build() error {
	cmd := exec.Command("go", "build", "-o", binaryName, srcDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Run() error {
	cmd := exec.Command("go", "run", srcDir)
	cmd.Dir = backendDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Clean() error {
	fmt.Println("Cleaning...")
	return os.Remove(binaryName)
}

func Test() error {
	cmd := exec.Command("go", "test", "./...")
	cmd.Dir = backendDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func ComposeUp() error {
	cmd := exec.Command("docker", "compose", "up", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func BuildImages() error {
	cmd := exec.Command("docker", "compose", "build", "--no-cache")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func DevComposeUp() error {
	cmd := exec.Command("docker", "compose", "-f", "docker-compose.dev.yaml", "up", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

var Default = Build
