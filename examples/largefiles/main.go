package main

import (
	"flag"
	"log"
	"os"
	"sync"

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

	salespool = &sync.Pool {
		New: func() interface{} {
			return new(Sales)
		},
	}
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
		var s = i.(*Sales)
		log.Printf("%s - %s - %s", s.Country, s.ItemType, s.Price)
		salespool.Put(s)
	}

}

func tosales(r []string) interface{} {
	var s = salespool.Get().(*Sales)
	s.Region =   r[0]
	s.Country =  r[1]
	s.ItemType = r [2]
	s.Price =    r[9]
	return s
}

func filtersales(r []string) bool {
	return r[1] == "Singapore"
}
