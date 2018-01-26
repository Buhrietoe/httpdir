package main

import (
	"log"
	"os"
)

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Printf("Error loading config: %s", err)
		os.Exit(1)
	}

	startServer(config)
}
