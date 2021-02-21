package branchio

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math"
	"testing"
)

func Test_CSV_NewCSVReader(t *testing.T) {
	c := &CSV{}
	result := c.NewCSVReader(bytes.NewReader([]byte("")))
	assert.NotEmpty(t, result)
}

func Test_CSV_UnmarshalClicksEvents(t *testing.T) {
	reportData, _ := ioutil.ReadFile("stubs/data/export/events/eo_click-v2.csv")
	c := &CSV{}
	events := []*Event{}
	_ = c.Unmarshal(reportData, &events)
	assert.Equal(t, 12345678900, events[0].Id.Value())
	assert.Equal(t, 1613320668570, events[0].Timestamp.Value())
	assert.Equal(t, 12345678900, events[0].LastAttributedTouchDataTildeId.Value())
	assert.Equal(t, false, events[0].DeepLinked.Value())
	assert.Equal(t, false, events[0].FirstEventForUser.Value())
	assert.Equal(t, 9876543210, events[0].DiMatchClickToken.Value())
	assert.Equal(t, float64(0), events[0].EventDataRevenueInUsd.Value())
	assert.Equal(t, float64(0), events[0].EventDataExchangeRate.Value())
	assert.Equal(t, 1613320668570, events[0].EventTimestamp.Value())
	//assert.Equal(t, "2021-02-14 19:05:23+0000", events[0].TimestampISO.Value().Format(CustomTimestampFormatDefault))
}

func Test_CSV_UnmarshalEcommerceEvents(t *testing.T) {
	reportData, _ := ioutil.ReadFile("stubs/data/export/events/eo_commerce_event-v2.csv")
	c := &CSV{}
	events := []*Event{}
	_ = c.Unmarshal(reportData, &events)
	assert.Equal(t, 12345678900, events[0].Id.Value())
	assert.Equal(t, 1613264412028, events[0].Timestamp.Value())
	assert.Equal(t, 0, events[0].LastAttributedTouchDataTildeId.Value())
	assert.Equal(t, false, events[0].DeepLinked.Value())
	assert.Equal(t, false, events[0].FirstEventForUser.Value())
	assert.Equal(t, 0, events[0].DiMatchClickToken.Value())
	assert.Equal(t, 5.523858, math.Round(events[0].EventDataRevenueInUsd.Value()*1000000)/1000000)
	assert.Equal(t, 75.9071, math.Round(events[0].EventDataExchangeRate.Value()*10000)/10000)
	assert.Equal(t, 1613264412028, events[0].EventTimestamp.Value())
	//assert.Equal(t, "2021-02-14 19:05:23+0000", events[0].TimestampISO.Value().Format(CustomTimestampFormatDefault))
}
