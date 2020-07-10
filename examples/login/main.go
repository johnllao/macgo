package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/boltdb/bolt"
)

const (
	UsersBucket = "users"
	AuthHeader = "Authentication"
)

var (
	port int
	static string
	dbpath string

	fileserver http.Handler
	db *bolt.DB
)

func init() {
	flag.IntVar(&port, "port", 8080, "port number of the http")
	flag.StringVar(&static, "static", "", "path of the static folder containing the javascripts")
	flag.StringVar(&dbpath, "dbpath", "", "path of the db file")
	flag.Parse()
}

func main() {
	var err error
	err = initdb()
	if err != nil {
		log.Fatal(err)
		return
	}
	starthttp()
}

func initdb() error {
	var err error
	db, err = bolt.Open(dbpath, 0664, nil)
	if err != nil {
		return err
	}

	var sigc = make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigc
		db.Close()
		os.Exit(0)
	}()

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(UsersBucket))
		return nil
	})

	return nil

}

func starthttp() {

	var statpath = http.Dir(static)
	fileserver = http.FileServer(statpath)

	var mux = http.NewServeMux()
	mux.HandleFunc("/", WithSession(handleroot))
	mux.HandleFunc("/login", WithSession(handlelogin))
	mux.HandleFunc("/register", WithSession(handleregister))
	mux.HandleFunc("/css", handlecss)

	var server = http.Server{
		Addr: "localhost:" + strconv.Itoa(port),
		Handler: mux,
	}

	log.Printf("starting web server at port %d", port)
	log.Fatal(server.ListenAndServe())
}
