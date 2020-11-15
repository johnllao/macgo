package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var (
	mdpath    string
	port      int
	roottempl *template.Template
)

func init() {
	flag.StringVar(&mdpath, "mdpath", ".", "path of thd markdown files")
	flag.IntVar(&port, "port", 8080, "service port number")
	flag.Parse()
}

func main() {

	var err error

	roottempl, err = template.New("Root").Parse(indexhtml)
	if err != nil {
		log.Fatalln(err)
	}

	var s = http.Server{
		Addr: "localhost:" + strconv.Itoa(port),
		Handler: http.HandlerFunc(root),
	}
	log.Print("starting blog service")
	log.Printf("service port: %d",  port)
	log.Printf("markdown path: %s", mdpath)
	log.Fatalln(s.ListenAndServe())
}