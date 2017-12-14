package main

import (
	"io"
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
		next.ServeHTTP(w, r)
	})
}

// thing does a thing
func thing() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("My url: " + r.RequestURI + "\n"))
		case "POST":
			w.WriteHeader(http.StatusAccepted)
			buf, _ := ioutil.ReadAll(r.Body)
			w.Write(buf)
		case "PUT":
			w.WriteHeader(http.StatusCreated)
			buf, _ := ioutil.ReadAll(r.Body)
			w.Write(buf)
		case "FILE":
			newfile := filepath.Base(r.RequestURI)
			out, err := os.Create(newfile)
			if err != nil {
				log.Println(err)
			}
			defer out.Close()

			if _, err := io.Copy(out, r.Body); err != nil {
				log.Println("Things happened during upload...")
			}
			out.Sync()
			w.WriteHeader(http.StatusCreated)
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Uuhhh...\n"))
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

	// Start info
	log.Printf("Usage: %s [address:port] [directory]", filepath.Base(os.Args[0]))
	log.Printf("Listening on: %s", listenString)
	log.Printf("Serving from: %s", serveDir)

	// Map routes
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(serveDir))
	mux.Handle("/files/", http.StripPrefix("/files/", fs))
	mux.Handle("/", thing())

	// Serve it up
	err = http.ListenAndServe(listenString, logger(mux))
	if err != nil {
		log.Fatalln(err)
	}
}
