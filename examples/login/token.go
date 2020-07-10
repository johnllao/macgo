package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"net/http"
)

type User struct {
	UserName string
	PasswordHash []byte
	Salt []byte
	Email string
	ContactNo string
}

type Token struct {
	UserName string
	SessionID string
}

func GenerateToken(u User, r *http.Request) (string, error) {
	var err error
	var c *http.Cookie
	c, err = r.Cookie(SessionKey)
	if err != nil {
		return "", err
	}

	var t = Token{
		UserName: u.UserName,
		SessionID: c.Value,
	}
	var buf = &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(t)

	return hex.EncodeToString(buf.Bytes()), nil
}

func GetToken(t string) (*Token, error) {
	var err error
	var tokenb []byte
	tokenb, err = hex.DecodeString(t)
	if err != nil {
		return nil, err
	}
	var tok Token
	var r = bytes.NewReader(tokenb)
	err = gob.NewDecoder(r).Decode(&tok)
	if err != nil {
		return nil, err
	}
	return &tok, nil
}
