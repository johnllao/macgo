package main

import (
	"fmt"
	"time"

	"github.com/johnllao/macgo/examples/gobrpc/message"
	"github.com/johnllao/macgo/pkg/sock/client"
)

func main() {
	var err error
	var cli *client.Client
	cli, err = client.Connect("localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	sum(cli, 1, 2)
	diff(cli, 2, 1)
}

func sum(cli *client.Client, a, b int) {
	var err error

	var r []byte

	var start = time.Now()

	var o *message.Operation
	o, err  = message.NewSumOperation(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}

	var data []byte
	data, err  = o.ToBytes()
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err = cli.Send(data, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(time.Since(start), string(r))
}

func diff(cli *client.Client, a, b int) {
	var err error

	var r []byte

	var start = time.Now()

	var o *message.Operation
	o, err  = message.NewDiffOperation(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}

	var data []byte
	data, err  = o.ToBytes()
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err = cli.Send(data, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(time.Since(start), string(r))
}
