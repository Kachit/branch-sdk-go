package branchio

import (
	"bytes"
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"io"
)

//Csv unmarshaler
type CSV struct {
}

//Unmarshal raw data to Reports structs
func (c *CSV) Unmarshal(in []byte, out interface{}) error {
	r := c.NewCSVReader(bytes.NewReader(in))
	return gocsv.UnmarshalCSV(r, out)
}

//Create new CSV reader for unmarshaler
func (c *CSV) NewCSVReader(in io.Reader) gocsv.CSVReader {
	r := csv.NewReader(in)
	r.LazyQuotes = true
	return r
}
