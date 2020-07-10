package main

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"net/http"
	"path"
	"time"
)
const(
	SessionKey = "GOSESSION"
)
func WithSession(h func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		var basepath = path.Base(r.URL.Path)
		if ext := path.Ext(basepath); ext == ".js" || ext == ".map" || ext == ".css" {
			h(w, r)
			return
		}

		var c *http.Cookie
		c, err = r.Cookie(SessionKey)
		if err == http.ErrNoCookie {
			c = &http.Cookie{
				Name: SessionKey,
				Value: uuid(),
				Expires: time.Now().Add(8*time.Hour),
			}
			http.SetCookie(w, c)
			h(w, r)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		h(w, r)
	}
}

func ValidateSession(s string, r *http.Request) error {
	var err error
	var c *http.Cookie
	c, err = r.Cookie(SessionKey)
	if err != nil {
		return err
	}
	if c.Value != s {
		return errors.New("Invalid session")
	}

	return nil
}

func uuid() string {
	var ts = make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(ts, time.Now().UnixNano())

	var r = make([]byte, 16)
	rand.Read(r)

	return hex.EncodeToString(append(ts, r...))
}
