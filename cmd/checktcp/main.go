package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Check TCP/port Utility")
	fmt.Println()

	for i, arg := range os.Args {
		if i == 0 {
			continue
		}

		conn, err := net.Dial("tcp", arg)
		if err != nil {
			fmt.Printf("ERR: Failed to dial. %v \n", err.Error())
		} else {
			fmt.Printf("SUCCESS %v \n", arg)
			defer conn.Close()
		}
	}
}