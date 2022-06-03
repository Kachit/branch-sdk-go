package branchio

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"math"
	"testing"
)

type CSVTestSuite struct {
	suite.Suite
}

func (suite *CSVTestSuite) TestNewCSVReader() {
	result := NewCSVReader(bytes.NewReader([]byte("")))
	assert.NotEmpty(suite.T(), result)
}

func (suite *CSVTestSuite) TestUnmarshalCSVClicksEvents() {
	reportData, _ := ioutil.ReadFile("stubs/data/export/events/eo-click-v2.csv")
	events := []*Event{}
	_ = UnmarshalCSV(reportData, &events)
	assert.Equal(suite.T(), "12345678900", events[0].Id)
	assert.Equal(suite.T(), 1613320668570, events[0].Timestamp.Value())
	assert.Equal(suite.T(), 12345678900, events[0].LastAttributedTouchDataTildeId.Value())
	assert.Equal(suite.T(), false, events[0].DeepLinked.Value())
	assert.Equal(suite.T(), false, events[0].FirstEventForUser.Value())
	assert.Equal(suite.T(), 9876543210, events[0].DiMatchClickToken.Value())
	assert.Equal(suite.T(), float64(0), events[0].EventDataRevenueInUsd.Value())
	assert.Equal(suite.T(), float64(0), events[0].EventDataExchangeRate.Value())
	assert.Equal(suite.T(), 1613320668570, events[0].EventTimestamp.Value())
	//assert.Equal(suite.T(), "2021-02-14 19:05:23+0000", events[0].TimestampISO.Value().Format(CustomTimestampFormatDefault))
}

func (suite *CSVTestSuite) TestUnmarshalCSVEcommerceEvents() {
	reportData, _ := ioutil.ReadFile("stubs/data/export/events/eo-commerce-event-v2.csv")
	events := []*Event{}
	_ = UnmarshalCSV(reportData, &events)
	assert.Equal(suite.T(), "12345678900", events[0].Id)
	assert.Equal(suite.T(), 1613264412028, events[0].Timestamp.Value())
	assert.Equal(suite.T(), 0, events[0].LastAttributedTouchDataTildeId.Value())
	assert.Equal(suite.T(), false, events[0].DeepLinked.Value())
	assert.Equal(suite.T(), false, events[0].FirstEventForUser.Value())
	assert.Equal(suite.T(), 0, events[0].DiMatchClickToken.Value())
	assert.Equal(suite.T(), 5.523858, math.Round(events[0].EventDataRevenueInUsd.Value()*1000000)/1000000)
	assert.Equal(suite.T(), 75.9071, math.Round(events[0].EventDataExchangeRate.Value()*10000)/10000)
	assert.Equal(suite.T(), 1613264412028, events[0].EventTimestamp.Value())
	//assert.Equal(suite.T(), "2021-02-14 19:05:23+0000", events[0].TimestampISO.Value().Format(CustomTimestampFormatDefault))
}

func TestCSVTestSuite(t *testing.T) {
	suite.Run(t, new(CSVTestSuite))
}
