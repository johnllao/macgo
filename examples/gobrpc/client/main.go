package main

import (
	"fmt"
	"time"

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

	var start time.Time
	var data []byte

	start = time.Now()
	data, err = cli.Send([]byte("Hello1"), true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(time.Since(start), string(data))

	start = time.Now()
	data, err = cli.Send([]byte("Hello2"), true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(time.Since(start), string(data))

	start = time.Now()
	data, err = cli.Send([]byte("Hello3"), true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(time.Since(start), string(data))
}
