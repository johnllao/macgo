package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"net/http"
)

const (
	PassKey = "secret"
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

	var enctoken []byte
	enctoken, err =  Encrypt(PassKey, buf.Bytes())
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(enctoken), nil
}

func GetToken(t string) (*Token, error) {
	var err error
	var enctoken []byte
	enctoken, err = hex.DecodeString(t)
	if err != nil {
		return nil, err
	}

	var tokenb []byte
	tokenb, err = Decrypt(PassKey, enctoken)
	if err != nil {
		return nil, err
	}

	var token Token
	var r = bytes.NewReader(tokenb)
	err = gob.NewDecoder(r).Decode(&token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
