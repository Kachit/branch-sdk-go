package branchio

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"testing"
)

type ResourceAbstractTestSuite struct {
	suite.Suite
	testable *ResourceAbstract
}

func (suite *ResourceAbstractTestSuite) SetupTest() {
	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)
	suite.testable = NewResourceAbstract(transport)
}

func (suite *ResourceAbstractTestSuite) TestUnmarshalResponseJson() {
	expected, _ := ioutil.ReadFile("stubs/data/export/ontology/success.full.json")
	resp := BuildStubResponseFromFile(http.StatusOK, "stubs/data/export/ontology/success.full.json")
	resp.Header.Set("Content-Type", ResponseContentTypeJson)

	var data EventOntology
	err := suite.testable.unmarshalResponse(resp, &data)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "https://branch-exports-web.foo-bar.amazonaws.com/eo_click.csv", data.Click[0])
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(suite.T(), body)
	assert.Equal(suite.T(), expected, body)
}

func (suite *ResourceAbstractTestSuite) TestUnmarshalResponseCsv() {
	//expected, _ := ioutil.ReadFile("stubs/data/export/events/eo-click-v2.csv")
	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/data/export/events/eo-click-v2.csv")
	resp.Header.Set("Content-Type", ResponseContentTypeOctetStream)

	events := []*Event{}
	err := suite.testable.unmarshalResponse(resp, &events)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "12345678900", events[0].Id)
	assert.Equal(suite.T(), 1613320668570, events[0].Timestamp.Value())
	assert.Equal(suite.T(), 12345678900, events[0].LastAttributedTouchDataTildeId.Value())
	assert.Equal(suite.T(), false, events[0].DeepLinked.Value())
	assert.Equal(suite.T(), false, events[0].FirstEventForUser.Value())
	assert.Equal(suite.T(), 9876543210, events[0].DiMatchClickToken.Value())
	assert.Equal(suite.T(), float64(0), events[0].EventDataRevenueInUsd.Value())
	assert.Equal(suite.T(), float64(0), events[0].EventDataExchangeRate.Value())
	assert.Equal(suite.T(), 1613320668570, events[0].EventTimestamp.Value())
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(suite.T(), body)
	//assert.Equal(t, expected, body)
}

func TestResourceAbstractTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceAbstractTestSuite))
}
