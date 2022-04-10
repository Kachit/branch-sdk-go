package branchio

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_HTTP_RequestBuilder_BuildUriWithoutQueryParams(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}
	uri, err := builder.buildUri("qwerty", nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, uri)
	assert.Equal(t, ProdAPIUrl+"/qwerty", uri.String())
}

func Test_HTTP_RequestBuilder_BuildUriWithQueryParams(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	data := make(map[string]interface{})
	data["foo"] = "bar"
	data["bar"] = "baz"

	uri, err := builder.buildUri("qwerty", data)
	assert.Nil(t, err)
	assert.NotEmpty(t, uri)
	assert.Equal(t, ProdAPIUrl+"/qwerty?bar=baz&foo=bar", uri.String())
}

func Test_HTTP_RequestBuilder_BuildHeaders(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	headers := builder.buildHeaders()
	assert.NotEmpty(t, headers)
	assert.Equal(t, "application/json", headers.Get("Content-Type"))
}

func Test_HTTP_RequestBuilder_BuildBody(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	data := make(map[string]interface{})
	data["foo"] = "bar"
	data["bar"] = "baz"

	body, _ := builder.buildBody(data)
	assert.NotEmpty(t, body)
}

func Test_HTTP_RequestBuilder_BuildRequestGET(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	ctx := context.Background()
	result, err := builder.BuildRequest(ctx, "get", "foo", map[string]interface{}{"foo": "bar"}, map[string]interface{}{"foo": "bar"})
	assert.NoError(t, err)
	assert.Equal(t, http.MethodGet, result.Method)
	assert.Equal(t, ProdAPIUrl+"/foo?foo=bar", result.URL.String())
	assert.Equal(t, "application/json", result.Header.Get("Content-Type"))
	assert.Nil(t, result.Body)
}

func Test_HTTP_RequestBuilder_BuildRequestPOST(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	ctx := context.Background()
	result, err := builder.BuildRequest(ctx, "post", "foo", map[string]interface{}{"foo": "bar"}, map[string]interface{}{"foo": "bar"})
	assert.NoError(t, err)
	assert.Equal(t, http.MethodPost, result.Method)
	assert.Equal(t, ProdAPIUrl+"/foo?foo=bar", result.URL.String())
	assert.Equal(t, "application/json", result.Header.Get("Content-Type"))
	assert.NotEmpty(t, result.Body)
}

func Test_HTTP_NewHttpTransport(t *testing.T) {
	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)
	assert.NotEmpty(t, transport)
}

func Test_HTTP_Transport_SendRequestSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()

	body, _ := LoadStubResponseData("stubs/data/export/ontology/success.empty.json")
	httpmock.RegisterResponder(http.MethodGet, cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	resp, err := transport.SendRequest(ctx, http.MethodGet, "foo", nil, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)

	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, body, bodyRsp)
}

func Test_HTTP_Transport_SendRequestGET(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()

	body, _ := LoadStubResponseData("stubs/data/export/ontology/success.empty.json")

	httpmock.RegisterResponder(http.MethodGet, cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	resp, err := transport.Get(ctx, "foo", nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)

	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, body, bodyRsp)
}

func Test_HTTP_Transport_SendRequestPOST(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()

	body, _ := LoadStubResponseData("stubs/data/export/ontology/success.empty.json")

	httpmock.RegisterResponder(http.MethodPost, cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	resp, err := transport.Post(ctx, "foo", nil, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)

	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, body, bodyRsp)
}

func Test_HTTP_ResponseBody_IsSuccess(t *testing.T) {
	rsp := &ResponseBody{status: http.StatusAccepted}
	assert.True(t, rsp.IsSuccess())
	rsp.status = http.StatusMultipleChoices
	assert.False(t, rsp.IsSuccess())
	rsp.status = http.StatusBadRequest
	assert.False(t, rsp.IsSuccess())
}

func Test_HTTP_ResponseBody_GetErrorByDefault(t *testing.T) {
	rsp := &ResponseBody{}
	assert.Equal(t, "Unknown error", rsp.GetError())
}

func Test_HTTP_ResponseBody_GetErrorEmptyMessage(t *testing.T) {
	rsp := &ResponseBody{Error: &ResponseBodyError{}}
	assert.Equal(t, "Unknown error", rsp.GetError())
}

func Test_HTTP_ResponseBody_GetErrorNotEmptyMessage(t *testing.T) {
	rsp := &ResponseBody{Error: &ResponseBodyError{Message: "Foo error"}}
	assert.Equal(t, "Foo error", rsp.GetError())
}

func Test_HTTP_NewResponseHandler_ResponseHandlerStream(t *testing.T) {
	result := NewResponseHandler(ResponseContentTypeOctetStream)
	assert.Implements(t, (*ResponseHandlerInterface)(nil), result)
	assert.IsType(t, (*ResponseHandlerStream)(nil), result)
}

func Test_HTTP_NewResponseHandler_ResponseHandlerJson(t *testing.T) {
	result := NewResponseHandler(ResponseContentTypeJson)
	assert.Implements(t, (*ResponseHandlerInterface)(nil), result)
	assert.IsType(t, (*ResponseHandlerJson)(nil), result)
}

func Test_HTTP_NewResponseHandler_ByDefault(t *testing.T) {
	result := NewResponseHandler("foo")
	assert.Implements(t, (*ResponseHandlerInterface)(nil), result)
	assert.IsType(t, (*ResponseHandlerJson)(nil), result)
}

func Test_HTTP_ResponseHandlerStream_UnmarshalBody(t *testing.T) {
	handler := &ResponseHandlerStream{}
	reportData, _ := ioutil.ReadFile("stubs/data/export/events/eo-click-v2.csv")
	events := []*Event{}
	err := handler.UnmarshalBody(reportData, &events)
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
}

func Test_HTTP_ResponseHandlerStream_ReadBody(t *testing.T) {
	handler := &ResponseHandlerStream{}
	expected, _ := ioutil.ReadFile("stubs/data/export/events/eo-click-v2.csv")
	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/data/export/events/eo-click-v2.csv")
	data, err := handler.ReadBody(resp)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
	assert.Equal(t, expected, data)
	assert.Empty(t, body)
}

func Test_HTTP_ResponseHandlerStream_RestoreBody(t *testing.T) {
	handler := &ResponseHandlerStream{}
	expectedRaw, _ := ioutil.ReadFile("stubs/data/export/events/eo-click-v2.csv")
	expectedGzipped, _ := loadStubResponseDataGzipped("stubs/data/export/events/eo-click-v2.csv")
	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/data/export/events/eo-click-v2.csv")
	data, err := handler.ReadBody(resp)
	closer, err := handler.RestoreBody(data)
	resp.Body = closer

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, expectedRaw, data)
	assert.NotEmpty(t, body)
	fmt.Println(expectedGzipped)
	fmt.Println(body)
	//assert.Equal(t, expectedGzipped, body)
}

func Test_HTTP_ResponseHandlerJson_UnmarshalBody(t *testing.T) {
	handler := &ResponseHandlerJson{}
	reportData, _ := ioutil.ReadFile("stubs/data/export/ontology/success.full.json")
	var data EventOntology
	err := handler.UnmarshalBody(reportData, &data)
	assert.NoError(t, err)
	assert.Equal(t, "https://branch-exports-web.foo-bar.amazonaws.com/eo_click.csv", data.Click[0])
}

func Test_HTTP_ResponseHandlerJson_ReadBody(t *testing.T) {
	handler := &ResponseHandlerJson{}
	expected, _ := ioutil.ReadFile("stubs/data/export/ontology/success.full.json")
	resp := BuildStubResponseFromFile(http.StatusOK, "stubs/data/export/ontology/success.full.json")
	data, err := handler.ReadBody(resp)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
	assert.Equal(t, expected, data)
	assert.Empty(t, body)
}

func Test_HTTP_ResponseHandlerJson_RestoreBody(t *testing.T) {
	handler := &ResponseHandlerJson{}
	expected, _ := ioutil.ReadFile("stubs/data/export/ontology/success.full.json")
	resp := BuildStubResponseFromFile(http.StatusOK, "stubs/data/export/ontology/success.full.json")
	data, _ := handler.ReadBody(resp)
	closer, err := handler.RestoreBody(data)
	resp.Body = closer

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
	assert.NotEmpty(t, body)
	assert.Equal(t, expected, body)
}
