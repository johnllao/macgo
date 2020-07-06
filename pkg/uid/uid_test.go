package uid

import (
	"testing"
)

func TestGenerateString(t *testing.T) {
	var err error
	var u string
	u, err = GenerateString()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(u)

	u, err = GenerateString()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(u)

	u, err = GenerateString()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(u)

	u, err = GenerateString()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(u)
}
