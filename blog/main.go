package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

var (
	port int
	staticpath string

	filehandler http.Handler
	indextmpl *template.Template
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
		Addr: "localhost:" + strconv.Itoa(port),
		Handler: mux,
	}
	log.Panicln(s.ListenAndServe())
}

func indexhandle(w http.ResponseWriter, r *http.Request) {

	var err error

	if path.Ext(r.URL.Path) == ".css" {
		if filehandler == nil {
			var rootdir = http.Dir(staticpath)
			filehandler = http.FileServer(rootdir)
		}
		filehandler.ServeHTTP(w, r)
		return
	}

	if indextmpl == nil {
		indextmpl, err = template.New("index").Parse(indexhtml)
		if err != nil {
			http.Error(w, err.Error(), http.StatusOK)
			return
		}
	}

	w.Header().Set("Content-Type", "text/html")
	indextmpl.Execute(w, nil)
}

var indexhtml = `<!doctype html>
<html>
<head>
	<title>Notebook</title>
	<link rel="stylesheet" type="text/css" href="/css/blog.css" />
</head>
<body>
	<div class="container">
		<h1 class="header">My Notes</h1>
		<div class="row">
			<button>Create New</button>
		</div>
		<div class="row">
			&nbsp;
		</div>
	</div>
</body>
</html>
`
