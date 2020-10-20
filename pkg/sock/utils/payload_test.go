package utils

import (
	"bytes"
	"testing"
)

func TestPayload(t *testing.T) {
	var data = "The quick brown fox jumps over the river"

	var err error
	var p []byte

	p, err = Payload([]byte(data))
	if err != nil {
		t.Error(err)
		return
	}

	var s int
	var d []byte
	s, d, err = PayloadData(p)
	if err != nil {
		t.Error(err)
		return
	}

	if s != len(data) {
		t.Errorf("mismatch in payload size. actual %d, expected %d", s, len(data))
	}

	if !bytes.Equal(d, []byte(data)) {
		t.Errorf("mismatch payload data")
	}
}