package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"log"
	"time"
)

func main() {
	var err error
	var n uint64
	for i := 0; i < 10; i++ {
		var start = time.Now()
		n, err = randuint64()
		if err != nil {
			log.Println(err, time.Since(start))
			return
		}
		log.Println(n, time.Since(start))
	}
}

func uuid() string {
	var ts = make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(ts, time.Now().UnixNano())

	var r = make([]byte, 16)
	rand.Read(r)

	return hex.EncodeToString(append(ts, r...))
}

func randuint64() (uint64, error) {
	var err error

	var buf = make([]byte, binary.MaxVarintLen64)
	rand.Read(buf)

	var r = bytes.NewReader(buf)
	var n uint64
	n, err = binary.ReadUvarint(r)
	if err != nil {
		return 0, err
	}

	return n, nil
}

