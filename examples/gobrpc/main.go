package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/johnllao/macgo/pkg/sock/server"
)

const (
	port = 8080
)

func main() {
	var s = server.Server{
		Port: port,
		MsgHandler: servermsghandler,
	}
	s.Start()

	<-s.Ready

	fmt.Print("Press ENTER key to exit ")
	var stdinr = bufio.NewReader(os.Stdin)
	stdinr.ReadLine()
}

func servermsghandler(w io.Writer, d []byte) {
	w.Write(append([]byte("ACK: "), d...))
}
