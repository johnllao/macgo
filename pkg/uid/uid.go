package uid

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"time"
)

func Generate() ([]byte, error) {
	var err error

	var r = make([]byte, 16)
	_, err = rand.Read(r)
	if err != nil {
		return nil, err
	}

	var n = make([]byte, binary.MaxVarintLen64)
	var t = time.Now().UnixNano()
	binary.PutVarint(n, t)

	return append(r, n...), nil
}

func GenerateString() (string, error) {
	var err error

	var r []byte

	r, err = Generate()
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(r), nil
}
