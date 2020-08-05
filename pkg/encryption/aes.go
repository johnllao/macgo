package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
)

func EncryptUsingAES(key string, data []byte) ([]byte, error) {

	var err error

	var hasher = sha256.New()
	hasher.Write([]byte(key))

	var keyb = hasher.Sum(nil)

	var block cipher.Block
	block, err = aes.NewCipher(keyb)
	if err != nil {
		return nil, err
	}

	var ciphertxt = make([]byte, aes.BlockSize + len(data))
	var iv = ciphertxt[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	var stream = cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertxt[aes.BlockSize:], data)

	return ciphertxt, nil
}

func DecryptUsingAES(key string, data []byte) ([]byte, error) {

	var err error

	var hasher = sha256.New()
	hasher.Write([]byte(key))

	var keyb = hasher.Sum(nil)

	var block cipher.Block
	block, err = aes.NewCipher(keyb)
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		return nil, errors.New("Decrypted text is too short")
	}

	var iv = data[:aes.BlockSize]
	var encdata = data[aes.BlockSize:]
	var decdata = make([]byte, len(encdata))

	var stream = cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decdata, encdata)

	return decdata, nil
}

func EncryptUsingAESGCM(key, data []byte) ([]byte, error) {

	var err error

	var b cipher.Block
	b, err = aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("Encrypt: Failed to create new cipher. " + err.Error())
	}

	var g cipher.AEAD
	g, err = cipher.NewGCM(b)
	if err != nil {
		return nil, errors.New("Encrypt: Failed to create GCM. " + err.Error())
	}

	var n = make([]byte, g.NonceSize())
	if _, err = io.ReadFull(rand.Reader, n); err != nil {
		return nil, errors.New("Encrypt: Failed to run random number generator. " + err.Error())
	}

	return g.Seal(n, n, data, nil), nil
}

func DecryptUsingAESGCM(key, data []byte) ([]byte, error) {

	var err error

	var b cipher.Block
	b, err = aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("Decrypt: Failed to create new cipher. " + err.Error())
	}

	var g cipher.AEAD
	g, err = cipher.NewGCM(b)
	if err != nil {
		return nil, errors.New("Decrypt: Failed to create GCM. " + err.Error())
	}

	if len(data) < g.NonceSize() {
		return nil, errors.New("Decrypt: Invalid data size")
	}

	var n = data[:g.NonceSize()]

	return g.Open(nil, n, data[g.NonceSize():], nil)
}
