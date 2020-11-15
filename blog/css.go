package main

import (
	"net/http"
)

func WriteCSS(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/css")
	w.Write([]byte(indexcss))
}

var indexcss = `
body {
	font-family: Arial;
	font-size: 12pt;
	margin: 0;
	padding: 0;
}

header {
	background-color: #1d3557;
	color: #f1faee
	margin: 0;
	padding: 0;
}

header h1 {
	font-weight: normal;
	margin: 0;
	padding: 0.5em;
}

header h1 a {
	color: #f1faee;
	text-decoration: none;
}

article {
	margin: 0;
	padding: 0.5em;
}

article h1 {
	color: #457b9d;
}

article a {
	color: #1d3557;
	text-decoration: none;
}
`
