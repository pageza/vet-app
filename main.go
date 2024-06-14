package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Loading environment variable from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: ")
	}

	// Accessing environment variable
	projectName := os.Getenv("PROJECT_NAME")

	log.Println(projectName)
}
