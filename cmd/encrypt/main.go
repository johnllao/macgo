package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/johnllao/macgo/pkg/encryption"
)

var (
	action     string
	passphrase string
	inpath     string
	outpath    string
)

func init() {
	flag.Usage = func() {
		fmt.Println("Encrypts or decrypts files. Results are written as base64")
		flag.PrintDefaults()
	}
	flag.StringVar(&action, "a", "enc", "action to indicate encrypt (enc) or decrypt (dec)")
	flag.StringVar(&passphrase, "p", "", "passphrase")
	flag.StringVar(&inpath, "i", "", "path of the input file")
	flag.StringVar(&outpath, "o", "", "path of the output file")
	flag.Parse()
}

func main() {

	var err error
	var data []byte

	if inpath == "" || outpath == "" {
		fmt.Println("ERR: invalid input or output file path")
		os.Exit(99)
	}

	if inpath != "" {
		var f *os.File
		f, err = os.Open(inpath)
		if err != nil {
			fmt.Println("ERR:", err)
			os.Exit(99)
		}
		defer f.Close()

		data, err = ioutil.ReadAll(f)
		if err != nil {
			fmt.Println("ERR:", err)
			os.Exit(99)
		}
	}


	if action == "enc" {
		var encdata []byte
		encdata, err = encryption.EncryptUsingAES(passphrase, data)
		if err != nil {
			fmt.Println("ERR:", err)
			os.Exit(99)
		}

		var f *os.File
		f, err = os.Create(outpath)
		if err != nil {
			fmt.Println("ERR:", err)
			os.Exit(99)
		}
		defer f.Close()

		var encoder = base64.NewEncoder(base64.StdEncoding, f)
		encoder.Write(encdata)
		encoder.Close()
	}

	if action == "dec" {

		var reader = bytes.NewReader(data)
		var decoder = base64.NewDecoder(base64.StdEncoding, reader)

		var buf = &bytes.Buffer{}
		io.Copy(buf, decoder)

		var decdata []byte
		decdata, err = encryption.DecryptUsingAES(passphrase, buf.Bytes())
		if err != nil {
			fmt.Println("ERR:", err)
			os.Exit(99)
		}
		ioutil.WriteFile(outpath, decdata, 0664)
	}

}
