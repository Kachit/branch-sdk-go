package branchio

import (
	"context"
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type ExportTestSuite struct {
	suite.Suite
}

func (suite *ExportTestSuite) TestEventOntologyResponseIsEmpty() {
	r := &EventOntologyResponse{}
	assert.True(suite.T(), r.IsEmpty())
	r.Data = &EventOntology{}
	assert.False(suite.T(), r.IsEmpty())
}

func (suite *ExportTestSuite) TestEventOntologyIsEmpty() {
	r := &EventOntology{}
	assert.True(suite.T(), r.IsEmpty())
	r.Click = []string{"foo"}
	assert.False(suite.T(), r.IsEmpty())
}

func (suite *ExportTestSuite) TestEventMarshalJSON() {
	csvData, _ := ioutil.ReadFile("stubs/data/export/events/eo-click-v2.csv")
	expectedJsonData, _ := ioutil.ReadFile("stubs/data/export/events/eo-click-v2.json")
	events := []*Event{}
	_ = UnmarshalCSV(csvData, &events)
	jsonData, err := json.Marshal(&events)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedJsonData, jsonData)
}

func TestExportTestSuite(t *testing.T) {
	suite.Run(t, new(ExportTestSuite))
}

type ExportResourceTestSuite struct {
	suite.Suite
	cfg      *Config
	ctx      context.Context
	testable *ExportResource
}

func (suite *ExportResourceTestSuite) SetupTest() {
	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()
	suite.cfg = cfg
	suite.ctx = context.Background()
	suite.testable = &ExportResource{NewResourceAbstract(transport)}
	httpmock.Activate()
}

func (suite *ExportResourceTestSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func (suite *ExportResourceTestSuite) TestBuildEventOntologyRequestParams() {
	dt := time.Date(2022, 4, 07, 0, 0, 0, 0, time.Local)
	result := suite.testable.buildEventOntologyRequestParams(dt)
	assert.Equal(suite.T(), suite.cfg.Key, result["branch_key"])
	assert.Equal(suite.T(), suite.cfg.Secret, result["branch_secret"])
	assert.Equal(suite.T(), "2022-04-07", result["export_date"])
}

func (suite *ExportResourceTestSuite) TestGetEventOntologySuccess() {
	rs := BuildStubResponseFromFile(http.StatusOK, "stubs/data/export/ontology/success.full.json")
	rs.Header.Set("Content-Type", ResponseContentTypeJson)
	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri+"/v3/export", httpmock.ResponderFromResponse(rs))

	dt := time.Date(2022, 4, 07, 0, 0, 0, 0, time.Local)
	result, resp, err := suite.testable.GetEventOntology(suite.ctx, dt)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), "https://branch-exports-web.foo-bar.amazonaws.com/eo_click.csv", result.Data.Click[0])

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(suite.T(), body)
}

func (suite *ExportResourceTestSuite) TestGetEventOntologyError() {
	rs := BuildStubResponseFromFile(http.StatusUnauthorized, "stubs/data/export/ontology/error.auth.json")
	rs.Header.Set("Content-Type", ResponseContentTypeJson)
	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri+"/v3/export", httpmock.ResponderFromResponse(rs))

	dt := time.Date(2022, 4, 07, 0, 0, 0, 0, time.Local)
	result, resp, err := suite.testable.GetEventOntology(suite.ctx, dt)
	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), "Invalid or missing app id, Branch key, or secret", err.Error())

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(suite.T(), body)
}

func (suite *ExportResourceTestSuite) TestGetEventDataSuccess() {
	rs := buildStubResponseFromGzip(http.StatusOK, "stubs/data/export/events/eo-click-v2.csv")
	rs.Header.Set("Content-Type", ResponseContentTypeOctetStream)
	httpmock.RegisterResponder(http.MethodGet, "https://branch-exports-web.foo-bar.amazonaws.com/eo_click.csv", httpmock.ResponderFromResponse(rs))

	result, resp, err := suite.testable.GetEventData(suite.ctx, "https://branch-exports-web.foo-bar.amazonaws.com/eo_click.csv")
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), "12345678900", result.Data[0].Id)
	assert.Equal(suite.T(), 1613320668570, result.Data[0].Timestamp.Value())
	assert.Equal(suite.T(), 12345678900, result.Data[0].LastAttributedTouchDataTildeId.Value())
	assert.Equal(suite.T(), false, result.Data[0].DeepLinked.Value())
	assert.Equal(suite.T(), false, result.Data[0].FirstEventForUser.Value())
	assert.Equal(suite.T(), 9876543210, result.Data[0].DiMatchClickToken.Value())
	assert.Equal(suite.T(), float64(0), result.Data[0].EventDataRevenueInUsd.Value())
	assert.Equal(suite.T(), float64(0), result.Data[0].EventDataExchangeRate.Value())
	assert.Equal(suite.T(), 1613320668570, result.Data[0].EventTimestamp.Value())

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(suite.T(), body)
}

func (suite *ExportResourceTestSuite) TestGetEventDataError() {
	rs := BuildStubResponseFromFile(http.StatusForbidden, "stubs/data/export/events/forbidden.xml")
	rs.Header.Set("Content-Type", ResponseContentTypeXml)
	httpmock.RegisterResponder(http.MethodGet, "https://branch-exports-web.foo-bar.amazonaws.com/eo_click.csv", httpmock.ResponderFromResponse(rs))

	result, resp, err := suite.testable.GetEventData(suite.ctx, "https://branch-exports-web.foo-bar.amazonaws.com/eo_click.csv")
	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), "Access Denied", err.Error())
}

func TestExportResourceTestSuite(t *testing.T) {
	suite.Run(t, new(ExportResourceTestSuite))
}
