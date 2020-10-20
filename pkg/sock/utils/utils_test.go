package utils

import (
	"encoding/binary"
	"testing"
)

func TestPayloadSizeToBytes(t *testing.T) {
	var n0 = 20201031

	var b = UInt32ToBytes(uint32(n0))
	if len(b) != binary.MaxVarintLen32 {
		t.Errorf("length mismatch. actual %d, expected %d", len(b), binary.MaxVarintLen32)
		return
	}

	var n1 = int(BytesToUInt32(b))
	if n0 != n1 {
		t.Errorf("value mismatch. actual %d. expected %d", n1, n0)
	}
}
