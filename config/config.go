package config

import (
	"log"
	"os"
	"path/filepath"
)

// Config contains config for the webserver
type ServerConfig struct {
	ListenString string
	ServeDir     string
}

// Load gets configuration details for the webserver
func Load() (config ServerConfig, err error) {
	// Defaults
	config.ListenString = ":8080"
	config.ServeDir, err = filepath.Abs(".")

	// Environment variables override defaults
	if len(os.Getenv("HTTP_ADDR")) > 0 {
		config.ListenString = os.Getenv("HTTP_ADDR")
	}
	if len(os.Getenv("HTTP_DIR")) > 0 {
		config.ServeDir, err = filepath.Abs(os.Getenv("HTTP_DIR"))
		if err != nil {
			log.Fatalf("Unable to parse directory: %s\n", err)
		}
	}

	// Command-line arguments take highest precedence
	if len(os.Args) > 1 {
		config.ListenString = os.Args[1]
	}
	if len(os.Args) > 2 {
		config.ServeDir, _ = filepath.Abs(os.Args[2])
	}

	return config, err
}
