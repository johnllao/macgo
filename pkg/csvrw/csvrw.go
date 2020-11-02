package csvrw

import (
	"encoding/csv"
	"io"
)

func Read(src io.Reader, factory func([]string) interface{}, filter func([]string) bool, hasheader bool) ([]string, <-chan interface{}, error) {

	var err error

	var i = make(chan interface{})

	var r = csv.NewReader(src)

	var header []string
	if hasheader {
		header, err = r.Read()
		if err != nil {
			return nil, nil, err
		}
	}

	go func() {
		var rerr error
		for {
			var row []string
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
		}
		close(i)
	}()

	return header, i, nil
}
