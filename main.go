package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// logger logs all requests
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		if r.Method == "PUT" || r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatalf("Error reading request body: %s", err)
			}
			log.Println(string(body))
		}
		next.ServeHTTP(w, r)
	})
}

// fileParse handles file requests
func fileParse() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Write([]byte("hello"))
			//return http.FileServer(http.Dir("."))
		case "POST":
			w.Write([]byte("I don't want this no mo"))
		default:
			w.Write([]byte("Uuhhh..."))
		}
	})
}

func main() {
	// Defaults
	var err error
	listenString := ":8080"
	serveDir, _ := filepath.Abs(".")
	if len(os.Getenv("HTTP_ADDR")) > 0 {
		listenString = os.Getenv("HTTP_ADDR")
	}
	if len(os.Getenv("HTTP_DIR")) > 0 {
		serveDir, err = filepath.Abs(os.Getenv("HTTP_DIR"))
		if err != nil {
			log.Fatalf("Unable to parse directory: %s\n", err)
		}
	}
	if len(os.Args) > 1 {
		listenString = os.Args[1]
	}
	if len(os.Args) > 2 {
		serveDir, _ = filepath.Abs(os.Args[2])
	}

	log.Printf("Usage: %s [address:port] [directory]", filepath.Base(os.Args[0]))
	log.Printf("Listening on: %s", listenString)
	log.Printf("Serving from: %s", serveDir)

	//fp := http.HandlerFunc(fileParse)
	http.Handle("files/", fileParse)
	http.Handle("/", http.FileServer(http.Dir(serveDir)))
	err = http.ListenAndServe(listenString, logger(http.DefaultServeMux))
	if err != nil {
		log.Fatalln(err)
	}
}
