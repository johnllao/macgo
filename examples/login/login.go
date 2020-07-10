package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
)

type LoginModel struct {
	Title string
	UserName string
	Password string
}

func handlelogin(w http.ResponseWriter, r *http.Request) {
	var err error

	r.ParseForm()

	var model = LoginModel{
		Title: "Simple Authentication",
		UserName: r.Form.Get("username"),
		Password: r.Form.Get("password"),
	}

	w.Header().Set("Content-Type", "text/html")

	if model.UserName == "" {
		var t *template.Template
		t, err = template.New("login").Parse(loginhtml)
		if err != nil {
			http.Error(w, err.Error(), http.StatusOK)
			return
		}
		t.Execute(w, model)
	}

	if model.UserName != "" {
		var token string
		token, err = authenticate(model, r)
		if err == nil {
			http.Redirect(w, r, "/?token=" + token, http.StatusPermanentRedirect)
			return
		}
		if err != nil {
			log.Printf("%s", err.Error())

			var t *template.Template
			t, err = template.New("login").Parse(loginfailedhtml)
			if err != nil {
				http.Error(w, err.Error(), http.StatusOK)
				return
			}
			t.Execute(w, nil)
		}

	}
}

func authenticate(model LoginModel, req *http.Request) (string, error) {
	var err error
	var userb []byte
	err = db.View(func(tx *bolt.Tx) error {
		var b = tx.Bucket([]byte(UsersBucket))
		userb = b.Get([]byte(model.UserName))
		return nil
	})

	if userb == nil || len(userb) == 0 {
		return "", errors.New(model.UserName + " not found")
	}

	var u User
	var r = bytes.NewReader(userb)
	err = gob.NewDecoder(r).Decode(&u)
	if err != nil {
		return "", errors.New("Failed to decode " + model.UserName + ". " + err.Error())
	}

	var h = sha256.New()
	h.Write(append([]byte(model.Password), u.Salt...))

	if !bytes.Equal(u.PasswordHash, h.Sum(nil)) {
		return "", errors.New( model.UserName + " password mismatch")
	}

	var token string
	token, err = GenerateToken(u, req)
	if err != nil {
		return "", err
	}

	return token, nil
}

var loginhtml = `
<!doctype html>
<html>
<head>
	<title>{{ .Title }}</title>
	<link rel="stylesheet" type="text/css" href="/css" />
	<script type="text/javascript" src="js/vue.min.js"></script>
	<script type="text/javascript" src="js/axios.min.js"></script>
</head>
<body>
	<div>
		<form method="POST" action="/login">
			<div class="row">
				<label class="cell width-2">User name</label>
				<div class="cell width-5">
					<input type="text" name="username" />
				</div>
			</div>
			<div class="row">
				<label class="cell width-2">Password</label>
				<div class="cell width-5">
					<input type="password" name="password" />
				</div>
			</div>
			<div class="row">
				<div class="cell width-2">
					<button type="submit">Login</button>
				</div>
				<div class="cell width-2">
					<a href="/register">Register here</a>
				</div>
			</div>
		</form>
    </div>
</body>
</html>
`

var loginfailedhtml = `
<!doctype html>
<html>
<head>
	<title>Simple Authentication</title>
	<link rel="stylesheet" type="text/css" href="/css" />
	<script type="text/javascript" src="js/vue.min.js"></script>
	<script type="text/javascript" src="js/axios.min.js"></script>
</head>
<body>
	<div class="row">
		Invalid username and/or password
    </div>
	<div class="row">
		Please <a href="/login">try</a> again
    </div>
</body>
</html>
`