package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	servercertfile    = "/Users/johnllao/src/macgo/examples/mutualtls/servercert.pem"

	certfile    = "/Users/johnllao/src/macgo/examples/mutualtls/client1cert.pem"
	pvtkeyfile  = "/Users/johnllao/src/macgo/examples/mutualtls/client1pvtkey.pem"
)

func main() {
	// Load the client certificate
	cert, err := tls.LoadX509KeyPair(certfile, pvtkeyfile)
	if err != nil {
		fmt.Println("ERR (tls.LoadX509KeyPair) ", err)
		return
	}

	// Load the CA certificate
	cacert, err := ioutil.ReadFile(servercertfile)
	if err != nil {
		fmt.Println("ERR", err)
		return
	}
	cacertpool := x509.NewCertPool()
	cacertpool.AppendCertsFromPEM(cacert)

	// Create a HTTPS client and supply the CA cert pool and the client certificate
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: cacertpool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	// Request /hello via the created HTTPS client over port 8443 via GET
	r, err := client.Get("https://localhost:8443")
	if err != nil {
		fmt.Println("ERR (client.Get)", err)
		return
	}

	// Read the response body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ERR", err)
		return
	}

	// Print the response body to stdout
	fmt.Printf("%s\n", body)
}