package branchio

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_HTTP_RequestBuilder_BuildUriWithoutQueryParams(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}
	uri, err := builder.buildUri("qwerty", nil)
	assert.NotEmpty(t, uri)
	assert.Equal(t, "https://github.com/qwerty", uri.String())
	assert.Nil(t, err)
}

func Test_HTTP_RequestBuilder_BuildUriWithQueryParams(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	data := make(map[string]interface{})
	data["foo"] = "bar"
	data["bar"] = "baz"

	uri, err := builder.buildUri("qwerty", data)
	assert.NotEmpty(t, uri)
	assert.Equal(t, "https://github.com/qwerty?bar=baz&foo=bar", uri.String())
	assert.Nil(t, err)
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

func Test_HTTP_NewHttpTransport(t *testing.T) {
	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)
	assert.NotEmpty(t, transport)
}

func Test_HTTP_Transport_RequestSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)

	body, _ := LoadStubResponseData("stubs/data/auth/token.success.json")

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := transport.Request("GET", "foo", nil, nil)
	assert.NotEmpty(t, resp)
}

func Test_HTTP_Transport_RequestGET(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)

	body, _ := LoadStubResponseData("stubs/data/exports/export.success.json")

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := transport.Get("foo", nil)
	assert.NotEmpty(t, resp)
}

func Test_HTTP_Transport_RequestPOST(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)

	body, _ := LoadStubResponseData("stubs/data/exports/export.success.json")

	httpmock.RegisterResponder("POST", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := transport.Post("foo", nil, nil)
	assert.NotEmpty(t, resp)
}

func Test_HTTP_Response_IsSuccessTrue(t *testing.T) {
	response := &Response{raw: BuildStubResponseFromFile(http.StatusOK, "stubs/data/exports/export.success.json")}
	assert.True(t, response.IsSuccess())
}

func Test_HTTP_Response_IsSuccessFalse(t *testing.T) {
	response := &Response{raw: BuildStubResponseFromFile(http.StatusBadRequest, "stubs/data/exports/export.success.json")}
	assert.False(t, response.IsSuccess())
}

func Test_HTTP_Response_GetRawResponse(t *testing.T) {
	rsp := BuildStubResponseFromFile(http.StatusOK, "stubs/data/exports/export.success.json")
	response := &Response{raw: rsp}
	raw := response.GetRawResponse()
	assert.NotEmpty(t, raw)
	assert.Equal(t, http.StatusOK, raw.StatusCode)
}

func Test_HTTP_Response_GetRawBody(t *testing.T) {
	data, _ := LoadStubResponseData("stubs/data/exports/export.success.json")
	rsp := BuildStubResponseFromFile(http.StatusBadRequest, "stubs/data/exports/export.success.json")
	response := &Response{raw: rsp}
	assert.Equal(t, string(data), response.GetRawBody())
}

func Test_HTTP_NewResponse(t *testing.T) {
	rsp := BuildStubResponseFromFile(http.StatusOK, "stubs/data/exports/export.success.json")
	response := NewResponse(rsp)
	assert.NotEmpty(t, response)
}
