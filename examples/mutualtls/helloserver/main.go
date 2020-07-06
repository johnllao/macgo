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
	serverpvtkeyfile  = "/Users/johnllao/src/macgo/examples/mutualtls/serverpvtkey.pem"

	client1certfile = "/Users/johnllao/src/macgo/examples/mutualtls/client1cert.pem"
	client2certfile = "/Users/johnllao/src/macgo/examples/mutualtls/client2cert.pem"
)

func main() {
	var err error

	var mux = http.NewServeMux()
	mux.HandleFunc("/", handleindex)

	// Load the client 1 certificate
	client1cert, err := ioutil.ReadFile(client1certfile)
	if err != nil {
		fmt.Println("ERR", err)
		return
	}
	// Load the client 2 certificate
	client2cert, err := ioutil.ReadFile(client2certfile)
	if err != nil {
		fmt.Println("ERR", err)
		return
	}

	// Add the client 1, client2 and server certificate into the cert pool.
	// These are the only certificates to be accepted by the server
	certpool := x509.NewCertPool()
	certpool.AppendCertsFromPEM(client1cert)
	certpool.AppendCertsFromPEM(client2cert)

	// Create the TLS Config with the CA pool and enable Client certificate validation
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		ClientCAs: certpool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	var s = http.Server{
		Addr: "localhost:8443",
		Handler: mux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Starting service")
	err = s.ListenAndServeTLS(servercertfile, serverpvtkeyfile)
	if err != nil {
		fmt.Println("ERR", err)
	}
}

func handleindex(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Welcome!")
	fmt.Fprint(w, "Certs: ", r.TLS.PeerCertificates[0].Subject.String())
}
