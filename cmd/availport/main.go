package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 0, "port to check")
	flag.Parse()
}

func main() {
	var err error
	if port == 0 {
		var l net.Listener
		l, err = net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			fmt.Println("ERR", err)
			return
		}
		addr := l.Addr()
		addrstr := strings.Split(addr.String(), ":")
		if len(addrstr) <= 1 {
			fmt.Println("ERR Cannot retrieve available port")
			return
		}
		fmt.Println(addrstr[1])
	}

	if port > 0 {
		var conn net.Conn
		conn, err = net.Dial("tcp", "127.0.0.1:" + strconv.Itoa(port))
		if err != nil  {
			if operr, ok := err.(*net.OpError); ok && strings.Contains(operr.Err.Error(), "connect:") {
				fmt.Printf("Port %d is available \n", port)
			} else {
				fmt.Println("ERR", err)
			}
			return
		}
		defer conn.Close()
		fmt.Printf("Port %d is used \n", port)
	}
}
