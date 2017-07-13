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

	if len(os.Args) > 1 {
		log.Printf("Serving from: %s", os.Args[1])
		panic(http.ListenAndServe(os.Args[1], logger(http.FileServer(http.Dir("./")))))
	} else {
		log.Fatalln("Please specify address:port to listen from")
	}
}
