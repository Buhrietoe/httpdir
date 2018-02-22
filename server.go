package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// File represents a file
type File struct {
	Name string
	Size int
}

// indexModel of file listing
type indexModel struct {
	CurrentPath string
	Files       []File
}

const page string = `<html>
<head>
<style type="text/css">
	body { color: white; background: black; }
</style>
</head>
<body>
<div>{{ .Text }}</div>
</body>
</html>`

// logger logs all requests to stdout
//func logger(next http.Handler) http.Handler {
//return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
//next.ServeHTTP(w, r)
//})
//}

// getFile sends a file or file index
func getFile(c *gin.Context) {
	resp := make(map[string]interface{})

	//resp["status"] = "OK"
	resp["request"], _ = c.GetRawData()
	resp["debug"] = c.Params

	c.JSON(http.StatusOK, resp)
}

// putFile retrieves a file
func putFile(c *gin.Context) {
	// stuff
}

// postFile retrieves a file
func postFile(c *gin.Context) {
	// stuff
}

// deleteFile deletes a file
func deleteFile(c *gin.Context) {
	// stuff
}

// fileHandler triages downloading/uploading
//func fileHandler(root string) http.Handler {
//return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//switch r.Method {
//case "GET":
//w.WriteHeader(http.StatusOK)
//t, err := template.New("listing").Parse(page)
//if err != nil {
//log.Printf("Problem generating template: %s", err)
//}
//m := indexModel{CurrentPath: "my path"}
//err = t.Execute(w, &m)
//case "POST":
//w.WriteHeader(http.StatusAccepted)
//case "PUT":
//w.WriteHeader(http.StatusCreated)
//case "TEST":
//newfile := filepath.Base(r.RequestURI)
//out, err := os.Create(newfile)
//if err != nil {
//log.Println(err)
//}
//defer out.Close()

//if _, err := io.Copy(out, r.Body); err != nil {
//log.Println("Things happened during upload...")
//}
//out.Sync()
//w.WriteHeader(http.StatusCreated)
//default:
//w.WriteHeader(http.StatusBadRequest)
//}
//})
//}

// StartServer starts the webserver using specified config
func StartServer(config ServerConfig) {
	// Startup info
	log.Printf("Usage: %s [address:port] [directory]", filepath.Base(os.Args[0]))
	log.Printf("Listening on: %s", config.ListenString)
	log.Printf("Serving from: %s", config.ServeDir)

	r := gin.Default()
	r.GET("/", getFile)
	r.PUT("/:path", putFile)
	r.POST("/:path", postFile)

	//r.GET("/ping", func(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{
	//"message": "pong",
	//})
	//})

	r.Run(config.ListenString)

	// Map routes
	//mux := http.NewServeMux()
	//fs := http.FileServer(http.Dir(config.ServeDir))
	//mux.Handle("/", fileHandler(config.ServeDir))
	//mux.Handle("/fallback/", http.StripPrefix("/fallback/", fs))

	// Serve it up
	//err := http.ListenAndServe(config.ListenString, logger(mux))
	//if err != nil {
	//log.Fatalln(err)
	//}
}
