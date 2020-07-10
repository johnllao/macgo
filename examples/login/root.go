package main

import (
	"html/template"
	"net/http"
	"path"
)

type RootModel struct {
	Title string
	Message string
}

func handleroot(w http.ResponseWriter, r *http.Request) {
	var err error

	var basepath = path.Base(r.URL.Path)
	if ext := path.Ext(basepath); ext == ".js" || ext == ".map" || ext == ".css" {
		fileserver.ServeHTTP(w, r)
		return
	}

	var tok string
	if r.Header.Get(AuthHeader) != "" {
		tok = r.Header.Get(AuthHeader)
	} else if r.URL.Query().Get("token") != "" {
		tok = r.URL.Query().Get("token")
	}
	
	if tok == "" {
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
		return
	}

	w.Header().Set(AuthHeader, r.URL.Query().Get("token"))

	var token *Token
	token, err = GetToken(tok)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ValidateSession(token.SessionID, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var t *template.Template
	t, err = template.New("root").Parse(roothtml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, RootModel{
		Title: "Simple Authentication",
		Message: "Welcome " + token.UserName + "!",
	})
}

var roothtml = `
<!doctype html>
<html>
<head>
	<title>{{ .Title }}</title>
	<link rel="stylesheet" type="text/css" href="/css" />
	<script type="text/javascript" src="js/vue.min.js"></script>
	<script type="text/javascript" src="js/axios.min.js"></script>
</head>
<body>
	{{ .Message }}
</body>
</html>
`
