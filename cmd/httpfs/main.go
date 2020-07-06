package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

var (
	port int
	root string
)

func init() {
	flag.StringVar(&root, "root", ".", "path of the root folder")
	flag.IntVar(&port, "port", 8080, "http port number")
	flag.Parse()
}

func main() {
	var mux = http.NewServeMux()
	mux.HandleFunc("/", handleindex)

	var server = http.Server{
		Addr: "localhost:" + strconv.Itoa(port),
		Handler: mux,
	}

	log.Printf("starting service...")
	log.Fatal(server.ListenAndServe())
}

func handleindex(w http.ResponseWriter, r *http.Request) {
	var p = http.Dir(root)
	var f = http.FileServer(p)

	f.ServeHTTP(w, r)
}
