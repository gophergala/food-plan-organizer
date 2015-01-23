package etl

import (
	"encoding/csv"
	"io"
)

type Extract interface {
	Parse(io.Reader, chan<- interface{}) error
}

// SR27Reader wraps a CSV Reader, configurierung it to the CSV encoding standard
type SR27Reader struct {
	*csv.Reader
}

func newSR27Reader(r io.Reader) *SR27Reader {
	var sr27 = SR27Reader{csv.NewReader(r)}
	sr27.Comma = '^'
	sr27.LazyQuotes = true
	return &sr27
}

func (r *SR27Reader) Read() ([]string, error) {
	var s, e = r.Reader.Read()
	for i := range s {
		l := len(s[i])

		if l > 0 && s[i][0] == '~' && s[i][l-1] == '~' {
			s[i] = s[i][1 : l-1]
		}
	}
	return s, e
}
