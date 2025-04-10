//go:build mage

package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var (
	binaryName = "build"
	backendDir = "backend"
	srcDir     = "cmd/main.go"
)

func Dev() error {
	cmd := exec.Command("go", "run", srcDir)
	cmd.Dir = backendDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func DockerDev() error {
	cmd := exec.Command("docker", "compose", "-f", "docker-compose.dev.yaml", "up", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func DockerProd() error {
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

func LoadPasswords() {
	downloadURL := "https://github.com/brannondorsey/naive-hashcat/releases/download/data/rockyou.txt"
	inputFile := "rockyou.txt"
	outputFile := "db/import.sql"

	fmt.Println("⬇️ Downloading rockyou.txt...")
	resp, err := http.Get(downloadURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	outFile, err := os.Create(inputFile)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ Download complete!")

	inFile, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	outFile, err = os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	writer.WriteString("COPY compromised_passwords (hashed_password) FROM stdin;\n")

	scanner := bufio.NewScanner(inFile)

	fmt.Println("Processing passwords...")

	for scanner.Scan() {
		password := strings.TrimSpace(scanner.Text())
		hash := sha1.Sum([]byte(password))
		hashedPassword := hex.EncodeToString(hash[:])

		writer.WriteString(hashedPassword + "\n")
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	writer.WriteString("\\.\n")

	fmt.Println("✅ Passwords processed and prepared for PostgreSQL COPY import!")
}

func Hash(password string) {
	hasher := sha1.New()

	hasher.Write([]byte(password))

	hash := hasher.Sum(nil)

	hashHex := hex.EncodeToString(hash)

	fmt.Println("First 5 characters of the hash:", hashHex[:5])
}

var Default = Dev
