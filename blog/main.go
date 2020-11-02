package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/johnllao/macgo/pkg/markdown"
)

var (
	port      int
	roottempl *template.Template
)

func init() {
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
	log.Fatalln(s.ListenAndServe())
}

func root(w http.ResponseWriter, r *http.Request) {

	var err error

	var c []byte
	c, err = markdown.MarkdownContent(md).Convert()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err = roottempl.Execute(w, template.HTML(c))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


var indexhtml = `<!doctype html>
<html>
<head>
	<title>Notebook</title>
</head>
<body>
	<div class="container">
		{{ . }}
	</div>
</body>
</html>
`

var md = `# The Go Programming Language

Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.

Our canonical Git repository is located at https://go.googlesource.com/go.
There is a mirror of the repository at https://github.com/golang/go.

Unless otherwise noted, the Go source files are distributed under the
BSD-style license found in the LICENSE file.

### Download and Install

#### Binary Distributions

Official binary distributions are available at https://golang.org/dl/.

After downloading a binary release, visit https://golang.org/doc/install
or load [doc/install.html](./doc/install.html) in your web browser for installation
instructions.

#### Install From Source

If a binary distribution is not available for your combination of
operating system and architecture, visit
https://golang.org/doc/install/source or load [doc/install-source.html](./doc/install-source.html)
in your web browser for source installation instructions.

### Contributing

Go is the work of thousands of contributors. We appreciate your help!

To contribute, please read the contribution guidelines:
	https://golang.org/doc/contribute.html

Note that the Go project uses the issue tracker for bug reports and
proposals only. See https://golang.org/wiki/Questions for a list of
places to ask questions about the Go language.`