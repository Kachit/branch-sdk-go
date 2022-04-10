package branchio

import (
	"bytes"
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"io"
)

//UnmarshalCSV raw data to Events structures
func UnmarshalCSV(in []byte, out interface{}) error {
	r := NewCSVReader(bytes.NewReader(in))
	return gocsv.UnmarshalCSV(r, out)
}

//NewCSVReader Create new CSV reader for unmarshaler
func NewCSVReader(in io.Reader) gocsv.CSVReader {
	r := csv.NewReader(in)
	r.LazyQuotes = true
	return r
}
