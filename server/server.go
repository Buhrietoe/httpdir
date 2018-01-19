package server

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Buhrietoe/httpdir/config"
)

func Start() {
	config := config.Load()

	// Startup info
	log.Printf("Usage: %s [address:port] [directory]", filepath.Base(os.Args[0]))
	log.Printf("Listening on: %s", config.ListenString)
	log.Printf("Serving from: %s", config.ServeDir)

	// Map routes
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(config.ServeDir))
	mux.Handle("/files/", http.StripPrefix("/files/", fs))
	mux.Handle("/", thing())

	// Serve it up
	err := http.ListenAndServe(config.ListenString, logger(mux))
	if err != nil {
		log.Fatalln(err)
	}
}
