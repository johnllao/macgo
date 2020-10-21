package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/johnllao/macgo/examples/gobrpc/message"
	"github.com/johnllao/macgo/pkg/sock/server"
)

const (
	port = 8080
)

var opmap = map[string]func(io.Writer, []byte) {
	"Sum":  sumhandler,
	"Diff": diffhandler,
}

func main() {
	var s = server.NewServer(8080, servermsghandler)
	s.Start()
	<-s.Ready

	fmt.Print("Press ENTER key to exit ")
	var stdinr = bufio.NewReader(os.Stdin)
	stdinr.ReadLine()
}

func servermsghandler(w io.Writer, d []byte) {
	var err error
	var o *message.Operation
	o, err = message.ToOperation(d)
	if err != nil {
		w.Write([]byte("ERR: " + err.Error()))
		return
	}

	var h func(io.Writer, []byte)
	var ok bool
	h, ok = opmap[o.Name]
	if !ok {
		w.Write([]byte("invalid operation"))
		return
	}
	h(w, o.Parameter)
}

func sumhandler(w io.Writer, d []byte) {
	var err error
	var o *message.Sum
	o, err = message.ToSum(d)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Fprintf(w, "Sum is %d", o.A + o.B)
}

func diffhandler(w io.Writer, d []byte) {
	var err error
	var o *message.Diff
	o, err = message.ToDiff(d)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Fprintf(w, "Diff is %d", o.A - o.B)
}
