package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"
)

const (
	blank = ""
	keysize = 4096
)

func main() {
	var err error
	var pubkey *rsa.PublicKey
	var pvtkey *rsa.PrivateKey
	pvtkey, err = rsa.GenerateKey(rand.Reader, keysize)
	if err != nil {
		log.Fatalln("[rsa.GenerateKey]", err)
		return
	}
	pubkey = &pvtkey.PublicKey

	var msg = []byte("Block 431 Choa Chu Kang Ave 4, #08-579 Singapore 680431")
	log.Println("Message length: ", len(msg))

	var encmsg []byte
	encmsg, err = rsa.EncryptOAEP(sha256.New(), rand.Reader, pubkey, msg, nil)
	if err != nil {
		log.Fatalln("[rsa.EncryptOAEP]", err)
		return
	}

	log.Println("Encrypted message length: ", len(encmsg))

	var decmsg []byte
	decmsg, err = rsa.DecryptOAEP(sha256.New(), rand.Reader, pvtkey, encmsg, nil)
	if err != nil {
		log.Fatalln("[rsa.EncryptOAEP]", err)
		return
	}

	log.Println("Decrypted message: ", string(decmsg))

}
