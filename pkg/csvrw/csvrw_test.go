package csvrw

import (
	"os"
	"sync"
	"testing"
)

type Sales struct {
	Region   string
	Country  string
	ItemType string
	Price    string
}

var (
	salespool = &sync.Pool {
		New: func() interface{} {
			return new(Sales)
		},
	}
)

func BenchmarkRead(b *testing.B) {

	var err error

	for n := 0; n < b.N; n++ {
		const sourcepath = "/Users/johnllao/src/macgo/data/sales-records.csv"

		var sourcefile *os.File
		sourcefile, err = os.Open(sourcepath)
		if err != nil {
			b.Fatalf("cannot open source file. path: %s. %s", sourcepath, err.Error())
			return
		}
		defer sourcefile.Close()

		var rows <-chan interface{}
		_, rows, err = Read(sourcefile, tosales, filtersales, true)
		if err != nil {
			b.Fatalf("cannot read source file. path: %s. %s", sourcepath, err.Error())
			return
		}

		for i := range rows {
			var s = i.(*Sales)
			_, _, _ = s.Country, s.ItemType, s.Price
			salespool.Put(s)
		}
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
