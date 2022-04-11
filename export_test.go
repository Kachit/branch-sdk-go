package branchio

import (
	"context"
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func Test_Export_EventOntologyResponse_IsEmpty(t *testing.T) {
	r := &EventOntologyResponse{}
	assert.True(t, r.IsEmpty())
	r.Data = &EventOntology{}
	assert.False(t, r.IsEmpty())
}

func Test_Export_EventOntology_IsEmpty(t *testing.T) {
	r := &EventOntology{}
	assert.True(t, r.IsEmpty())
	r.Click = []string{"foo"}
	assert.False(t, r.IsEmpty())
}

func Test_Export_Event_MarshalJSON(t *testing.T) {
	csvData, _ := ioutil.ReadFile("stubs/data/export/events/eo-click-v2.csv")
	expectedJsonData, _ := ioutil.ReadFile("stubs/data/export/events/eo-click-v2.json")
	events := []*Event{}
	_ = UnmarshalCSV(csvData, &events)
	jsonData, err := json.Marshal(&events)
	assert.NoError(t, err)
	assert.Equal(t, expectedJsonData, jsonData)
}

func Test_Export_ExportResource_BuildEventOntologyRequestParams(t *testing.T) {
	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)
	ar := NewResourceAbstract(transport)
	resource := &ExportResource{ar}

	dt := time.Date(2022, 4, 07, 0, 0, 0, 0, time.Local)
	result := resource.buildEventOntologyRequestParams(dt)
	assert.Equal(t, cfg.Key, result["branch_key"])
	assert.Equal(t, cfg.Secret, result["branch_secret"])
	assert.Equal(t, "2022-04-07", result["export_date"])
}

func Test_Export_ExportResource_GetEventOntologySuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()
	ar := NewResourceAbstract(transport)
	resource := &ExportResource{ar}

	rs := BuildStubResponseFromFile(http.StatusOK, "stubs/data/export/ontology/success.full.json")
	rs.Header.Set("Content-Type", ResponseContentTypeJson)
	httpmock.RegisterResponder(http.MethodPost, cfg.Uri+"/v3/export", httpmock.ResponderFromResponse(rs))

	ctx := context.Background()
	dt := time.Date(2022, 4, 07, 0, 0, 0, 0, time.Local)
	result, resp, err := resource.GetEventOntology(ctx, dt)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)
	assert.Equal(t, "https://branch-exports-web.foo-bar.amazonaws.com/eo_click.csv", result.Data.Click[0])

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(t, body)
}

func Test_Export_ExportResource_GetEventOntologyError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()
	ar := NewResourceAbstract(transport)
	resource := &ExportResource{ar}

	rs := BuildStubResponseFromFile(http.StatusUnauthorized, "stubs/data/export/ontology/error.auth.json")
	rs.Header.Set("Content-Type", ResponseContentTypeJson)
	httpmock.RegisterResponder(http.MethodPost, cfg.Uri+"/v3/export", httpmock.ResponderFromResponse(rs))

	ctx := context.Background()
	dt := time.Date(2022, 4, 07, 0, 0, 0, 0, time.Local)
	result, resp, err := resource.GetEventOntology(ctx, dt)
	assert.Error(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)
	assert.Equal(t, "Invalid or missing app id, Branch key, or secret", err.Error())

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(t, body)
}

func Test_Export_ExportResource_GetEventDataSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	transport := BuildStubHttpTransport()
	ar := NewResourceAbstract(transport)
	resource := &ExportResource{ar}

	rs := buildStubResponseFromGzip(http.StatusOK, "stubs/data/export/events/eo-click-v2.csv")
	rs.Header.Set("Content-Type", ResponseContentTypeOctetStream)
	httpmock.RegisterResponder(http.MethodGet, "https://branch-exports-web.foo-bar.amazonaws.com/eo_click.csv", httpmock.ResponderFromResponse(rs))

	ctx := context.Background()
	result, resp, err := resource.GetEventData(ctx, "https://branch-exports-web.foo-bar.amazonaws.com/eo_click.csv")
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)
	assert.Equal(t, "12345678900", result.Data[0].Id)
	assert.Equal(t, 1613320668570, result.Data[0].Timestamp.Value())
	assert.Equal(t, 12345678900, result.Data[0].LastAttributedTouchDataTildeId.Value())
	assert.Equal(t, false, result.Data[0].DeepLinked.Value())
	assert.Equal(t, false, result.Data[0].FirstEventForUser.Value())
	assert.Equal(t, 9876543210, result.Data[0].DiMatchClickToken.Value())
	assert.Equal(t, float64(0), result.Data[0].EventDataRevenueInUsd.Value())
	assert.Equal(t, float64(0), result.Data[0].EventDataExchangeRate.Value())
	assert.Equal(t, 1613320668570, result.Data[0].EventTimestamp.Value())

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(t, body)
}

func Test_Export_ExportResource_GetEventDataError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	transport := BuildStubHttpTransport()
	ar := NewResourceAbstract(transport)
	resource := &ExportResource{ar}

	rs := BuildStubResponseFromFile(http.StatusForbidden, "stubs/data/export/events/forbidden.xml")
	rs.Header.Set("Content-Type", ResponseContentTypeOctetStream)
	httpmock.RegisterResponder(http.MethodGet, "https://branch-exports-web.foo-bar.amazonaws.com/eo_click.csv", httpmock.ResponderFromResponse(rs))

	ctx := context.Background()
	result, resp, err := resource.GetEventData(ctx, "https://branch-exports-web.foo-bar.amazonaws.com/eo_click.csv")
	assert.Error(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)
	assert.Equal(t, "Unknown error", err.Error())
}
