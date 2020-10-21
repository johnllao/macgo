package message

import (
	"testing"
)

func TestSumOperation(t *testing.T) {

	var err error

	var op *Operation
	op, err = NewSumOperation(1, 2)
	if err != nil {
		t.Error(err)
		return
	}

	var d []byte
	d, err = op.ToBytes()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("size: %d", len(d))

	op, err = ToOperation(d)
	if err != nil {
		t.Error(err)
		return
	}

	if op.Name != "Sum" {
		t.Errorf("Invalid operation. Actual: %s, Expected: Sum", op.Name)
	}
}

func TestDiffOperation(t *testing.T) {

	var err error

	var op *Operation
	op, err = NewDiffOperation(2, 1)
	if err != nil {
		t.Error(err)
		return
	}

	var d []byte
	d, err = op.ToBytes()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("size: %d", len(d))

	op, err = ToOperation(d)
	if err != nil {
		t.Error(err)
		return
	}

	if op.Name != "Diff" {
		t.Errorf("Invalid operation. Actual: %s, Expected: Sum", op.Name)
	}
}
