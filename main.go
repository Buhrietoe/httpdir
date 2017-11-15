package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// function logger logs all requests
func logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	listenString := ":8080"
	serveDir, _ := filepath.Abs(".")

	if len(os.Args) > 1 {
		listenString = os.Args[1]
	}
	if len(os.Args) > 2 {
		serveDir, _ = filepath.Abs(os.Args[2])
	}

	log.Printf("Serving from: %s", serveDir)
	log.Printf("Listening on: %s", listenString)
	err := http.ListenAndServe(listenString, logger(http.FileServer(http.Dir(serveDir))))
	if err != nil {
		log.Fatalln(err)
	}
}
