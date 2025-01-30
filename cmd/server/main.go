package main

import (
	"fmt"
	"log"
	"my-service/internal/config"
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
	port := config.EnvConfigs.AppPort
	log.Printf("Starting server on :%d\n", port)
	if err := srv.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
