package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		log.Printf("Serving from: %s", os.Args[1])
		panic(http.ListenAndServe(os.Args[1], http.FileServer(http.Dir("./"))))
	} else {
		log.Fatalln("Please specify address:port to listen from")
	}
}
