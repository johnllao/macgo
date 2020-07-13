package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"time"
)

var (
	countriespath string
)

func init() {
	flag.StringVar(&countriespath, "countries", "", "path of the file containing the list of countries")
	flag.Parse()
}

func main() {
	if countriespath == "" {
		fmt.Printf("ERR: missing country files path")
		return
	}

	var err error

	var f *os.File
	f, err = os.Open(countriespath)
	if err != nil {
		fmt.Printf("ERR: cannot open file '%s'. %s", countriespath, err.Error())
		return
	}
	defer f.Close()

	var countrylist = make([]string, 0)

	var r = csv.NewReader(f)
	r.Comma = '|'
	for {
		var ferr error
		var rec[]string
		rec, ferr = r.Read()
		if ferr == io.EOF {
			break
		}
		if ferr != nil || len(rec) != 2 {
			continue
		}
		countrylist = append(countrylist, rec[1])
	}

	sort.Strings(countrylist)

	var i = sort.SearchStrings(countrylist, "Singapore")
	fmt.Println(countrylist[i])
}

func RandomizeStrings(list []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
}
