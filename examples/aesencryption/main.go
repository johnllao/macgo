package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
)

func main() {
	var err error

	var h = sha256.New()
	h.Write([]byte("secret"))
	var key = h.Sum(nil)

	var data = []byte("The quick brown fox jumps over the river.")

	fmt.Println("Encrypt")
	var atad []byte
	atad, err = Encrypt(key, data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}

	fmt.Println("Decrypt")
	var data2 []byte
	data2, err = Decrypt(key, atad)
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}

	fmt.Println(string(data2))

}

func Encrypt(key, data []byte) ([]byte, error) {

	var err error

	var b cipher.Block
	b, err = aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("Encrypt: Failed to create new cipher. " + err.Error())
	}

	var g cipher.AEAD
	g, err = cipher.NewGCM(b)
	if err != nil {
		return nil, errors.New("Encrypt: Failed to create GCM. " + err.Error())
	}

	var n = make([]byte, g.NonceSize())
	if _, err = io.ReadFull(rand.Reader, n); err != nil {
		return nil, errors.New("Encrypt: Failed to run random number generator. " + err.Error())
	}

	return g.Seal(n, n, data, nil), nil
}

func Decrypt(key, data []byte) ([]byte, error) {

	var err error

	var b cipher.Block
	b, err = aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("Decrypt: Failed to create new cipher. " + err.Error())
	}

	var g cipher.AEAD
	g, err = cipher.NewGCM(b)
	if err != nil {
		return nil, errors.New("Decrypt: Failed to create GCM. " + err.Error())
	}

	if len(data) < g.NonceSize() {
		return nil, errors.New("Decrypt: Invalid data size")
	}

	var n = data[:g.NonceSize()]

	return g.Open(nil, n, data[g.NonceSize():], nil)
}