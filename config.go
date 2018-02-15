package main

import (
	"errors"
	"os"
	"path/filepath"
)

// ServerConfig contains config for the webserver
type ServerConfig struct {
	ListenString string
	ServeDir     string
}

// LoadConfig gets configuration details for the webserver
func LoadConfig() (config ServerConfig, err error) {
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
			return config, err
		}
	}

	// Command-line arguments take highest precedence
	if len(os.Args) > 1 {
		config.ListenString = os.Args[1]
	}
	if len(os.Args) > 2 {
		config.ServeDir, _ = filepath.Abs(os.Args[2])
	}

	// Validate path given
	if pathExists(config.ServeDir) {
		return config, err
	}

	return config, errors.New("Path does not exist: " + config.ServeDir)
}

func pathExists(path string) bool {
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		return true
	}
	return false
}
