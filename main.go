package main

import (
	"log"
	"os"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Printf("Error loading config: %s", err)
		os.Exit(1)
	}

	StartServer(config)
}
