package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Common HTML Examples")

	var mux = http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var fs = http.FileServer(http.Dir("."))
		fs.ServeHTTP(w, r)
	})

	var s = http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	fmt.Println(s.ListenAndServe())
}
