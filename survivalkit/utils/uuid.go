package utils

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"time"
)

func UUID() string {
	var ts = make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(ts, time.Now().UnixNano())

	var r = make([]byte, 16)
	rand.Read(r)

	return hex.EncodeToString(append(ts, r...))
}
