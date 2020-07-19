package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

var (
	port int
	blogpath string
	logpath string
)

func init() {
	flag.IntVar(&port, "port", 8080, "service port number")
	flag.StringVar(&blogpath, "blog", "", "path of the blog article files")
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

}
