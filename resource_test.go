package branchio

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_Resources_NewResourceAbstract(t *testing.T) {
	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)
	resource := NewResourceAbstract(transport)
	assert.NotEmpty(t, resource)
}

func Test_Resources_ResourceAbstract_UnmarshalResponseJson(t *testing.T) {
	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)
	resource := NewResourceAbstract(transport)
	expected, _ := ioutil.ReadFile("stubs/data/export/ontology/success.full.json")
	resp := BuildStubResponseFromFile(http.StatusOK, "stubs/data/export/ontology/success.full.json")
	resp.Header.Set("Content-Type", ResponseContentTypeJson)

	var data EventOntology
	err := resource.unmarshalResponse(resp, &data)
	assert.NoError(t, err)
	assert.Equal(t, "https://branch-exports-web.foo-bar.amazonaws.com/eo_click.csv", data.Click[0])
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(t, body)
	assert.Equal(t, expected, body)
}

func Test_Resources_ResourceAbstract_UnmarshalResponseCsv(t *testing.T) {
	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)
	resource := NewResourceAbstract(transport)
	//expected, _ := ioutil.ReadFile("stubs/data/export/events/eo-click-v2.csv")
	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/data/export/events/eo-click-v2.csv")
	resp.Header.Set("Content-Type", ResponseContentTypeOctetStream)

	events := []*Event{}
	err := resource.unmarshalResponse(resp, &events)
	assert.NoError(t, err)
	assert.Equal(t, 12345678900, events[0].Id.Value())
	assert.Equal(t, 1613320668570, events[0].Timestamp.Value())
	assert.Equal(t, 12345678900, events[0].LastAttributedTouchDataTildeId.Value())
	assert.Equal(t, false, events[0].DeepLinked.Value())
	assert.Equal(t, false, events[0].FirstEventForUser.Value())
	assert.Equal(t, 9876543210, events[0].DiMatchClickToken.Value())
	assert.Equal(t, float64(0), events[0].EventDataRevenueInUsd.Value())
	assert.Equal(t, float64(0), events[0].EventDataExchangeRate.Value())
	assert.Equal(t, 1613320668570, events[0].EventTimestamp.Value())
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(t, body)
	//assert.Equal(t, expected, body)
}
