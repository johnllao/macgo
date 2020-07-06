package encryption_test

import (
	"encoding/base64"
	"testing"

	"github.com/johnllao/macgo/pkg/encryption"
)

func TestAESEncryptDecrypt(t *testing.T) {

	var err error

	var passphrase = "secret"
	var msg = "the quick brown fox"

	var encdata []byte
	encdata, err = encryption.EncryptUsingAES(passphrase, []byte(msg))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("encoded data: ", base64.StdEncoding.EncodeToString(encdata))

	var decdata []byte
	decdata, err = encryption.DecryptUsingAES(passphrase, encdata)
	if err != nil {
		t.Error(err)
		return
	}

	if msg != string(decdata) {
		t.Errorf("Decrypt mismatch. Expected: %s, Actual: %s", msg, string(decdata))
	}
}