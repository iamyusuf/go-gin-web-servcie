package main

import (
	"fmt"
	"log"
	"my-service/internal/config"
	"my-service/internal/server"
)

func main() {
	port := config.EnvConfigs.AppPort
	fmt.Printf("Server running on port %v\n", port)
	dsn := config.EnvConfigs.GetDSN()
	fmt.Println("dsn:", dsn)
	srv, err := server.NewServer(dsn)
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	// Run server
	log.Printf("Starting server on :%d\n", port)
	if err := srv.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
