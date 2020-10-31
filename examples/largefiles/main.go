package main

import (
	"flag"
	"log"
	"os"

	"github.com/johnllao/macgo/pkg/csvrw"
)

type Sales struct {
	Region   string
	Country  string
	ItemType string
	Price    string
}

var (
	sourcepath  string
)

func init() {
	flag.StringVar(&sourcepath,  "s", "", "path of the source file")
	flag.Parse()
}

func main() {
	var err error

	log.Printf("opening source file. path: %s", sourcepath)

	var sourcefile *os.File
	sourcefile, err = os.Open(sourcepath)
	if err != nil {
		log.Fatalf("cannot open source file. path: %s. %s", sourcepath, err.Error())
		return
	}
	defer sourcefile.Close()

	var rows <-chan interface{}
	_, rows, err = csvrw.Read(sourcefile, tosales, filtersales, true)
	if err != nil {
		log.Fatalf("cannot read source file. path: %s. %s", sourcepath, err.Error())
		return
	}

	for i := range rows {
		var s = i.(Sales)
		log.Printf("%s - %s - %s", s.Country, s.ItemType, s.Price)
	}

}

func tosales(r []string) interface{} {
	return Sales {
		Region:   r[0],
		Country:  r[1],
		ItemType: r[2],
		Price:    r[9],
	}
}

func filtersales(r []string) bool {
	return r[1] == "Singapore"
}
