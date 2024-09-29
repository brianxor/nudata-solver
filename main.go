package main

import (
	"github.com/brianxor/nudata-solver/server"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	serverHost := os.Getenv("SERVER_HOST")

	if serverHost == "" {
		log.Fatal("SERVER_HOST environment variable not set")
	}

	serverPort := os.Getenv("SERVER_PORT")

	if serverPort == "" {
		log.Fatal("SERVER_PORT environment variable not set")
	}

	if err := server.Start(serverHost, serverPort); err != nil {
		log.Fatal(err)
	}
}
