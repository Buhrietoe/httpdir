package main

import (
	"log"
	"net/http"
	"os"
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

	if len(os.Args) > 1 {
		listenString = os.Args[1]
	}

	log.Printf("Serving from: %s", listenString)
	err := http.ListenAndServe(listenString, logger(http.FileServer(http.Dir("./"))))
	if err != nil {
		log.Fatalln(err)
	}
}
