package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
)

func Encrypt(key string, data []byte) ([]byte, error) {

	var err error

	var hasher = sha256.New()
	hasher.Write([]byte(key))

	var keyb = hasher.Sum(nil)

	var block cipher.Block
	block, err = aes.NewCipher(keyb)
	if err != nil {
		return nil, err
	}

	var ciphertxt = make([]byte, aes.BlockSize + len(data))
	var iv = ciphertxt[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	var stream = cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertxt[aes.BlockSize:], data)

	return ciphertxt, nil
}

func Decrypt(key string, data []byte) ([]byte, error) {

	var err error

	var hasher = sha256.New()
	hasher.Write([]byte(key))

	var keyb = hasher.Sum(nil)

	var block cipher.Block
	block, err = aes.NewCipher(keyb)
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		return nil, errors.New("Decrypted text is too short")
	}

	var iv = data[:aes.BlockSize]
	var encdata = data[aes.BlockSize:]
	var decdata = make([]byte, len(encdata))

	var stream = cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decdata, encdata)

	return decdata, nil
}

