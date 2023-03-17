package main

import (
	"log"
	"myapp/internal/api"
	"myapp/internal/service"
	"net/http"
	"os"
)

func main() {
	// Load configuration from environment variables
	cfg := &service.Config{
		APIEndpoint: os.Getenv("API_ENDPOINT"),
		GRPCAddress: os.Getenv("GRPC_ADDRESS"),
	}

	// Create API client
	apiClient := api.NewAPIClient(cfg.APIEndpoint)

	// Create gRPC service
	svc := service.NewService(apiClient)

	// Start HTTP server
	http.HandleFunc("/data", svc.HandleHTTP)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
