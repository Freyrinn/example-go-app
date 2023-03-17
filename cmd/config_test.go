package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Set environment variables for the test.
	os.Setenv("API_ENDPOINT", "http://localhost:8080")
	os.Setenv("GRPC_ADDRESS", "localhost:50051")

	// Load the config.
	config := LoadConfig()

	// Verify that the config was loaded correctly.
	expected := &Config{
		APIEndpoint: "http://localhost:8080",
		GRPCAddress: "localhost:50051",
	}
	assert.Equal(t, expected, config, "loaded config does not match expected config")
}

func TestLoadConfig_Empty(t *testing.T) {
	// Set environment variables for the test.
	os.Setenv("API_ENDPOINT", "")
	os.Setenv("GRPC_ADDRESS", "")

	// Load the config.
	config := LoadConfig()

	// Verify that the config was loaded correctly.
	expected := &Config{
		APIEndpoint: "",
		GRPCAddress: "",
	}
	assert.Equal(t, expected, config, "loaded config does not match expected config")
}
