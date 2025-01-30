package main

import (
	"log"
	"my-service/internal/server"
	"os"
)

func main() {
	// Get database DSN from environment
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable required")
	}

	// Create server
	srv, err := server.NewServer(dsn)
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	// Run server
	log.Println("Starting server on :8080")
	if err := srv.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
