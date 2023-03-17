package main

import "os"

type Config struct {
	APIEndpoint string
	GRPCAddress string
}

func LoadConfig() *Config {
	return &Config{
		APIEndpoint: os.Getenv("API_ENDPOINT"),
		GRPCAddress: os.Getenv("GRPC_ADDRESS"),
	}
}
