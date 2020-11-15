package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/johnllao/macgo/pkg/markdown"
)

func root(w http.ResponseWriter, r *http.Request) {

	var err error

	var p = r.URL.Path
	log.Printf("%s %s", r.Method, p)

	var parts = strings.SplitN(p, "/", 2)

	if parts == nil || len(parts) == 1 || parts[1] == "" {
		var indexmd []byte
		indexmd, err = ioutil.ReadFile(filepath.Join(mdpath, "index.md"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		WriteMarkdown([]byte(indexmd), w)
		return
	}

	if parts[1] == "css" {
		WriteCSS(w)
		return
	}

	var contentmd []byte
	contentmd, err = ioutil.ReadFile(filepath.Join(mdpath, parts[1] + ".md"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteMarkdown(contentmd, w)
}

func WriteMarkdown(md []byte, w http.ResponseWriter) {
	var err error

	var c []byte
	c, err = markdown.MarkdownContent(md).Convert()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = roottempl.Execute(w, template.HTML(c))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html")
}


var indexhtml = `<!doctype html>
<html>
<head>
	<title>Notes</title>
	<link rel="stylesheet" type="text/css" href="/css" />
</head>
<body>
	<div class="container">
		<header>
			<h1><a href="/">Notes</a></h1>
		</header>
		<article>
			{{ . }}
		</article>
	</div>
</body>
</html>
`
