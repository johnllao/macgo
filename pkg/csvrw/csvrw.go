package csvrw

import (
	"encoding/csv"
	"io"
	"sync"
)

var (
	rowpool = sync.Pool{
		New: func() interface{} {
			return []string{}
		},
	}
)
func Read(src io.Reader, factory func([]string) interface{}, filter func([]string) bool, hasheader bool) ([]string, <-chan interface{}, error) {

	var err error

	var i = make(chan interface{})

	var r = csv.NewReader(src)

	var header []string
	if hasheader {
		var header = rowpool.Get().([]string)
		header, err = r.Read()
		if err != nil {
			return nil, nil, err
		}
		rowpool.Put(header)
	}

	go func() {
		var rerr error
		for {
			var row = rowpool.Get().([]string)

			row, rerr = r.Read()
			if rerr == io.EOF {
				break
			}
			if rerr != nil {
				continue
			}

			if filter(row) {
				i <- factory(row)
			}

			rowpool.Put(row)
		}
		close(i)
	}()

	return header, i, nil
}
