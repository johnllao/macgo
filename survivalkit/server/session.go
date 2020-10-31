package server

import (
	"log"
	"net/http"
	"time"

	"github.com/johnllao/macgo/survivalkit/utils"
)

const (
	SessionKey =      "SESSIONID"
	SessionLifetime = 8 * time.Hour
)

func Session(w http.ResponseWriter, r *http.Request) string {
	var err error
	var c *http.Cookie
	c, err = r.Cookie(SessionKey)
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name: SessionKey,
			Value: utils.UUID(),
			Expires: time.Now().Add(SessionLifetime),
		}
		http.SetCookie(w, c)
	}
	if err != nil {
		log.Print(err)
		return ""
	}
	return c.Value
}
