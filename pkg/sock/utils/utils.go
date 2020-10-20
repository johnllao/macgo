package utils

import (
	"encoding/binary"
)

const (
	FileSeparator = 0x1C
)

func UInt32ToBytes(n uint32) []byte {
	var b = make([]byte, binary.MaxVarintLen32)
	binary.LittleEndian.PutUint32(b, n)
	return b
}

func BytesToUInt32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}


