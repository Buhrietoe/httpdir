package main

import (
	"log"
	"os"

	"github.com/Buhrietoe/httpdir/config"
	"github.com/Buhrietoe/httpdir/server"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Printf("Error loading config: %s", err)
		os.Exit(1)
	}

	server.Start(config)
}
