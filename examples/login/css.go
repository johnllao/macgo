package main

import (
	"net/http"
)

func handlecss(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/css/app.css", http.StatusPermanentRedirect)
}

var appcss = `
body {
	background-color: #333;
	color: #eee;
	font-family: Arial;
	font-size: 12pt;
}
`


