package utils

import (
	"encoding/binary"
)

type PayloadError struct {
	errmsg string
	InnerErr error
}

func (e *PayloadError) Error() string {
	return e.errmsg
}

func Payload(data []byte) ([]byte, error) {

	if len(data) == 0 {
		return nil, &PayloadError{ errmsg: "Payload() Payload is empty" }
	}

	var p []byte

	var size = UInt32ToBytes(uint32(len(data)))

	p = append(size, data...)
	p = append(p, FileSeparator)

	return p, nil
}

func PayloadData(payload []byte) (int, []byte, error) {
	var data []byte
	var size = -1

	var s = payload[:binary.MaxVarintLen32]
	size = int(BytesToUInt32(s))

	data = payload[binary.MaxVarintLen32:len(payload)-1]

	if size != len(data) {
		return -1, nil, &PayloadError{ errmsg: "PayloadData() Invalid payload size" }
	}

	if payload[len(payload) - 1] != FileSeparator {
		return -1, nil, &PayloadError{ errmsg: "PayloadData() Invalid payload delimiter" }
	}

	return size, data, nil
}
