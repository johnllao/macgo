package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"github.com/boltdb/bolt"
	"html/template"
	"net/http"
)

type RegisterModel struct {
	Title string
	UserName string
	Password string
	Email string
	ContactNo string
}

func handleregister(w http.ResponseWriter, r *http.Request) {
	var err error

	r.ParseForm()

	var model = RegisterModel{
		Title: "Simple Authentication",
		UserName: r.Form.Get("username"),
		Password: r.Form.Get("password"),
		Email: r.Form.Get("email"),
		ContactNo: r.Form.Get("contactno"),
	}

	w.Header().Set("Content-Type", "text/html")

	if model.UserName == "" {
		var t *template.Template
		t, err = template.New("register").Parse(registerhtml)
		if err != nil {
			http.Error(w, err.Error(), http.StatusOK)
			return
		}
		t.Execute(w, model)
	}
	if model.UserName != "" {

		var tf *template.Template
		tf, err = template.New("registerfail").Parse(registerfailedhtml)
		if err != nil {
			http.Error(w, err.Error(), http.StatusOK)
			return
		}

		if err == nil {
			err = checkifuserexists(model.UserName)
		}
		if err == nil {
			err = registeruser(model)
		}
		if err != nil {
			tf.Execute(w, err.Error())
			return
		}

		var t *template.Template
		t, err = template.New("register").Parse(registersuccesshtml)
		if err != nil {
			http.Error(w, err.Error(), http.StatusOK)
			return
		}
		t.Execute(w, model)
	}
}

func checkifuserexists(username string) error {
	var err error
	err = db.View(func (tx *bolt.Tx) error {
		var b = tx.Bucket([]byte(UsersBucket))
		if k := b.Get([]byte(username)); k != nil || len(k) > 0 {
			return errors.New(username + " already exists")
		}
		return nil
	})
	return err
}

func registeruser(model RegisterModel) error {
	var salt = make([]byte, 16)
	rand.Read(salt)

	var pwdhash []byte

	var h = sha256.New()
	h.Write(append([]byte(model.Password), salt...))
	pwdhash = h.Sum(nil)

	var u = User{
		UserName: model.UserName,
		PasswordHash: pwdhash,
		Salt: salt,
		Email: model.Email,
		ContactNo: model.ContactNo,
	}

	var buff = &bytes.Buffer{}
	gob.NewEncoder(buff).Encode(u)

	var err error
	err = db.Update(func(tx *bolt.Tx) error {
		var b = tx.Bucket([]byte(UsersBucket))
		return b.Put([]byte(u.UserName), buff.Bytes())
	})

	return err
}

var registerhtml = `
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
		<form action="/register" method="post">
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
				<label class="cell width-2">Email</label>
				<div class="cell width-5">
					<input type="text" name="email" />
				</div>
			</div>
			<div class="row">
				<label class="cell width-2">Contact No</label>
				<div class="cell width-5">
					<input type="text" name="contactno" />
				</div>
			</div>
			<div class="row">
				<div class="cell width-2">
					<input type="submit" value="Register" />
				</div>
				<div class="cell width-2">
					<a href="/login">Login here</a>
				</div>
			</div>
		</form>
    </div>
</body>
</html>
`

var registersuccesshtml = `
<!doctype html>
<html>
<head>
	<title>{{ .Title }}</title>
	<link rel="stylesheet" type="text/css" href="/css" />
	<script type="text/javascript" src="js/vue.min.js"></script>
	<script type="text/javascript" src="js/axios.min.js"></script>
</head>
<body>
	<div class="row">
		User {{ .UserName }} successfully registered
    </div>
	<div class="row">
		Please proceed to <a href="/login">login</a>
    </div>
</body>
</html>
`

var registerfailedhtml = `
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
		Registration failed. {{ . }}
    </div>
	<div class="row">
		Please <a href="/register">register</a> again
    </div>
</body>
</html>
`
