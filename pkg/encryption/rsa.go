package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"io/ioutil"
	"os"
)

const keysize = 4096

// GenerateKeys generates RSA private and public keys
func GenerateKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	var err error
	var pvtkey *rsa.PrivateKey
	pvtkey, err = rsa.GenerateKey(rand.Reader, keysize)
	if err != nil {
		return nil, nil, err
	}
	return pvtkey, &pvtkey.PublicKey, nil
}

// GenerateKeysToFile generates RSA private and public keys to file
func GenerateKeysToFile(pvtkeyfile string, pubkeyfile string) error {
	var err error
	var pvtkey *rsa.PrivateKey
	var pubkey *rsa.PublicKey

	pvtkey, pubkey, err = GenerateKeys()
	if err != nil {
		return err
	}

	var pvtkeyb []byte
	pvtkeyb = PrivateKeyToBytes(pvtkey)

	var pubkeyb []byte
	pubkeyb = PublicKeyToBytes(pubkey)

	err = ioutil.WriteFile(pvtkeyfile, pvtkeyb, 0644)
	if err != nil {
		return err
	}

	err  = ioutil.WriteFile(pubkeyfile, pubkeyb, 0644)
	if err != nil {
		os.Remove(pvtkeyfile)
		return err
	}

	return nil
}

// PrivateKeyToBytes converts private key to bytes
func PrivateKeyToBytes(key *rsa.PrivateKey) []byte {
	return x509.MarshalPKCS1PrivateKey(key)
}

// PublicKeyToBytes converts public key to bytes
func PublicKeyToBytes(key *rsa.PublicKey) []byte {
	return x509.MarshalPKCS1PublicKey(key)
}

// BytesToPrivateKey converts bytes to private key
func BytesToPrivateKey(key []byte) (*rsa.PrivateKey, error) {
	return x509.ParsePKCS1PrivateKey(key)
}

// BytesToPublicKey converts bytes to public key
func BytesToPublicKey(key []byte) (*rsa.PublicKey, error) {
	return x509.ParsePKCS1PublicKey(key)
}

// FileToPrivateKey opens a private key file
func FileToPrivateKey(path string) (*rsa.PrivateKey, error) {
	var err error
	var file *os.File

	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var pvtkeyb []byte
	pvtkeyb, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var pvtkey *rsa.PrivateKey
	pvtkey, err = BytesToPrivateKey(pvtkeyb)
	if err != nil {
		return nil, err
	}
	return pvtkey, nil
}

// FileToPublicKey opens a private key file
func FileToPublicKey(path string) (*rsa.PublicKey, error) {
	var err error
	var file *os.File

	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var pubkeyb []byte
	pubkeyb, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var pubkey *rsa.PublicKey
	pubkey, err = BytesToPublicKey(pubkeyb)
	if err != nil {
		return nil, err
	}
	return pubkey, nil
}

// EncryptFromRSAPublicKey encrypts data using an RSA public key
func EncryptFromRSAPublicKey(key *rsa.PublicKey, data []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, key, data, nil)
}

// EncryptFromRSAPrivateKey encrypts data using an RSA private key
func EncryptFromRSAPrivateKey(key *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, &key.PublicKey, data, nil)
}

// DecryptFromRSAPrivateKey decrypts data using an RSA private key
func DecryptFromRSAPrivateKey(key *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, key, data, nil)
}