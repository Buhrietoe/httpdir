package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// logger logs all requests
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func prettyData(data io.ReadCloser, contentType string) string {
	var output string

	buf, _ := ioutil.ReadAll(data)
	if len(buf) > 0 {
		if strings.Contains(strings.ToLower(contentType), "application/json") {
			jsonoutput := new(bytes.Buffer)
			json.Indent(jsonoutput, buf, "", "  ")
			output = jsonoutput.String()
		} else {
			output = string(buf)
		}
	}

	return output
}

// thing does a thing
func thing() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			if ct, ok := r.Header["Content-Type"]; ok {
				log.Println(prettyData(r.Body, ct[0]))
			} else {
				log.Println(prettyData(r.Body, "text/plain"))
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("My url: " + r.RequestURI + "\n"))
		case "POST":
			log.Println(prettyData(r.Body, r.Header["Content-Type"][0]))
			w.WriteHeader(http.StatusAccepted)
		case "PUT":
			log.Println(prettyData(r.Body, r.Header["Content-Type"][0]))
			w.WriteHeader(http.StatusCreated)
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
	// Environment variables override defaults
	if len(os.Getenv("HTTP_ADDR")) > 0 {
		listenString = os.Getenv("HTTP_ADDR")
	}
	if len(os.Getenv("HTTP_DIR")) > 0 {
		serveDir, err = filepath.Abs(os.Getenv("HTTP_DIR"))
		if err != nil {
			log.Fatalf("Unable to parse directory: %s\n", err)
		}
	}
	// Command-line arguments take precedence
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
