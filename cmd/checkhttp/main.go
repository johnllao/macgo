package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"strings"
	"time"
)

const (
	EmptyString = ""
	MethodGet   = "GET"
	MethodPost  = "POST"
)

var (
	method   string
	certpath string
)

func init() {
	flag.StringVar(&method, "m", MethodGet, "http method (e.g. GET, POST)")
	flag.StringVar(&certpath, "c", "", "path of the certificate file")
	flag.Parse()
}

func main() {
	var err error
	var addr string

	var args = flag.Args()
	if len(args) > 0 {
		addr = args[0]
	}

	if addr == EmptyString {
		fmt.Printf("ERR: Invalid or empty URL parameter \n")
		return
	}

	_, err = url.Parse(addr)
	if err != nil {
		fmt.Printf("ERR: Invalid URL parameter. %s \n", err.Error())
		return
	}

	fmt.Printf("Checking HTTP Request - %s \n", addr)

	var events = NewEvents()
	var trace = &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			events.Add("GetConn", Property{
				Key:  "HostPort",
				Value: hostPort,
			})
		},

		GotConn: func(info httptrace.GotConnInfo) {
			var ra, la string
			if conn, ok := info.Conn.(*tls.Conn); ok {
				ra = conn.RemoteAddr().String()
				la = conn.LocalAddr().String()
			}
			events.Add("GetConn",
				Property{
					Key:   "RemoteAddr",
					Value: ra,
				},
				Property{
					Key:   "LocalAddr",
					Value: la,
				})
		},

		DNSStart: func(info httptrace.DNSStartInfo) {
			events.Add("DNSStart", Property{
				Key:   "Host",
				Value: info.Host,
			})
		},

		DNSDone: func(info httptrace.DNSDoneInfo) {
			var aa = make([]string, 0)
			if len(info.Addrs) > 0 {
				for _, a := range info.Addrs {
					aa = append(aa, a.String())
				}
			}
			events.Add("DNSDone", Property{
				Key:   "Addr",
				Value: strings.Join(aa, ","),
			})
		},

		TLSHandshakeStart: func() {
			events.Add("TLSHandshakeStart")
		},

		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			events.Add("TLSHandshakeDone", Property{
				Key:   "ServerName",
				Value: state.ServerName,
			})
		},

		ConnectStart: func(network, addr string) {
			events.Add("ConnectStart", Property{
				Key:   "NetworkAddr",
				Value: network + "://" + addr,
			})
		},

		ConnectDone: func(network, addr string, err error) {
			events.Add("ConnectDone", Property{
				Key:   "NetworkAddr",
				Value: network + "://" + addr,
			})
		},

		GotFirstResponseByte: func() {
			events.Add("GotFirstResponseByte")
		},
		WroteHeaders: func() {
			events.Add("WroteHeaders")
		},
		WroteRequest: func(info httptrace.WroteRequestInfo) {
			events.Add("WroteRequest")
		},
	}

	var req *http.Request
	req, err = http.NewRequest(method, addr, nil)
	if err != nil {
		fmt.Printf("ERR: Failed initializing request. %s \n", err.Error())
		return
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	var transport = &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	if certpath != "" {
		var cert []byte
		cert, err = ioutil.ReadFile(certpath)
		if err != nil {
			fmt.Printf("ERR: Failed reading certificate. %s \n", err.Error())
			return
		}
		var certpool = x509.NewCertPool()
		certpool.AppendCertsFromPEM(cert)

		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: false,
			RootCAs: certpool,
		}
	} else {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	var cli = &http.Client{
		Transport: transport,
	}

	var res *http.Response

	res, err = cli.Do(req)
	if err != nil {
		fmt.Printf("ERR: Failed initializing request. %s \n", err.Error())
		return
	}

	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("ERR: Failed initializing request. %s \n", err.Error())
		return
	}

	events.Print()
	fmt.Printf("%s (Content length: %d) \n", res.Status, len(body))
}
