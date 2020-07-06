package encryption_test

import (
	"crypto/rsa"
	"testing"

	"github.com/johnllao/macgo/pkg/encryption"
)

func TestRSAEncryptDecrypt(t *testing.T) {
	var err error
	var pvtkey *rsa.PrivateKey
	var pubkey *rsa.PublicKey

	pvtkey, pubkey, err = encryption.GenerateKeys()
	if err != nil {
		t.Fatal(err)
	}

	var msg = "Passw0rd"

	var encb []byte
	encb, err = encryption.EncryptFromRSAPublicKey(pubkey, []byte(msg))
	if err != nil {
		t.Error(err)
	}

	var decb []byte
	decb, err = encryption.DecryptFromRSAPrivateKey(pvtkey, encb)
	if err != nil {
		t.Error(err)
	}

	if msg != string(decb) {
		t.Errorf("Decrypt mismatch. Expected: %s, Actual: %s", msg, string(decb))
	}
}

