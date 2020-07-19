package queue

import (
	"testing"
)

func TestStringQueue(t *testing.T) {
	var q = NewStringQueue()
	q.Push("Five")
	q.Push("Eight")
	q.Push("Two")

	if q.Len() != 3 {
		t.Errorf("Incorrect queue length. Expecting: 3. Actual: %d", q.Len())
	}

	var v1 = q.Pop()
	if v1 != "Five" {
		t.Errorf("Ivnvalid Pop value. Expecting: Five. Actual: %s", v1)
	}

	var v2 = q.Pop()
	if v2 != "Eight" {
		t.Errorf("Ivnvalid Pop value. Expecting: Eight. Actual: %s", v2)
	}

	var v3 = q.Pop()
	if v3 != "Two" {
		t.Errorf("Ivnvalid Pop value. Expecting: Two. Actual: %s", v3)
	}
}
