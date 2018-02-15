package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const page string = `<html>
<head>
<style type="text/css">
	body { color: white; background: black; }
</style>
</head>
<body>
<div>Things!</div>
</body>
</html>`

// logger logs all requests to stdout
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// fileHandler triages downloading/uploading
func fileHandler(root string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusOK)
			t, err := template.New("listing").Parse(page)
			if err != nil {
				log.Printf("Problem generating template: %s", err)
			}
			err = t.Execute(w, "")
			//w.Write([]byte("GET: " + r.RequestURI + "\n"))
		case "POST":
			w.WriteHeader(http.StatusAccepted)
		case "PUT":
			w.WriteHeader(http.StatusCreated)
		case "TEST":
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
		}
	})
}

// StartServer starts the webserver using specified config
func StartServer(config ServerConfig) {
	// Startup info
	log.Printf("Usage: %s [address:port] [directory]", filepath.Base(os.Args[0]))
	log.Printf("Listening on: %s", config.ListenString)
	log.Printf("Serving from: %s", config.ServeDir)

	// Map routes
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(config.ServeDir))
	mux.Handle("/", fileHandler(config.ServeDir))
	mux.Handle("/fallback/", http.StripPrefix("/fallback/", fs))

	// Serve it up
	err := http.ListenAndServe(config.ListenString, logger(mux))
	if err != nil {
		log.Fatalln(err)
	}
}
