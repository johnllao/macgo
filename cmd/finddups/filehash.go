package main

import (
	"crypto/sha256"
	"io/ioutil"
)

func FileHash(path string) ([sha256.Size]byte, error) {
	var err error

	var data []byte
	data, err = ioutil.ReadFile(path)
	if err != nil {
		return [sha256.Size]byte{}, err
	}

	var cksum = sha256.Sum256(data)
	return cksum, nil
}