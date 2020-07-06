package main

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"log"
	"net/http"
	"time"
)

var (
	rootpath = "."
)

func main() {
	var mux = http.NewServeMux()
	mux.HandleFunc("/", roothandler)

	var server = http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}
	log.Fatalln(server.ListenAndServe())
}

func roothandler(w http.ResponseWriter, r *http.Request) {
	var traceid = uuid()
	log.Printf("%s %s %s", traceid, r.Method, r.URL.Path)

	var root = http.Dir(rootpath)
	var h = http.FileServer(root)
	h.ServeHTTP(w, r)
}

func uuid() string {
	var ts = make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(ts, time.Now().UnixNano())

	var r = make([]byte, 16)
	rand.Read(r)
	return hex.EncodeToString(append(ts, r...))
}
