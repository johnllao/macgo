package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"strings"
)

const (
	Endpoint = "http://www.singaporepools.com.sg/_layouts/15/TotoApplication/TotoCommonPage.aspx/CalculatePrizeForTOTO"
)

var (
	bet string
	draw string

)

func init() {
	flag.StringVar(&bet, "bet", "", "bet numbers delimited by comma")
	flag.StringVar(&draw, "draw", "", "draw number")
	flag.Parse()
}

func main() {
	var err error

	var cli http.Client

	var payloadb []byte
	var payload = NewPayload(bet, draw)
	payloadb, err = payload.ToJSON()
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}

	var reader = bytes.NewReader(payloadb)
	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, Endpoint, reader)
	if err != nil {
		fmt.Println("ERR: Failed to initialize request. ", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	var res *http.Response
	res, err = cli.Do(req)
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}

	var raw *Raw
	raw, err = ToRaw(res.Body)
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}

	var result *Result
	result, err = raw.Result()
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}

	if len(result.Prizes) == 0 {
		fmt.Println("SORRY! You did not won")
	}
	if len(result.Prizes) > 0 {
		fmt.Println("CONGRATULATIONS! You WIN")
		for _, p := range result.Prizes {
			fmt.Printf("  Prize Group: %d \n", p.GroupNumber)
			fmt.Printf("  Share Prize: %d \n", p.ShareAmount)
			fmt.Println()
		}
	}
	fmt.Printf("Winning numbers   : %s \n", strings.Join(result.WinningNumbers, ","))
	fmt.Printf("Additional numbers: %s \n", result.AdditionalNumber)
}
