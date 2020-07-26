package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

var (
	port int
	staticpath string
)

func init() {
	flag.IntVar(&port, "port", 8080, "service port number")
	flag.StringVar(&staticpath, "static", "", "path of the static files")
	flag.Parse()
}

func main() {
	starthttp()
}

func starthttp() {

	var mux = http.NewServeMux()
	mux.HandleFunc("/", indexhandle)

	var s = http.Server{
		Addr: ":" + strconv.Itoa(port),
		Handler: mux,
	}
	log.Panicln(s.ListenAndServe())
}

func indexhandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Welcome!"))
}
