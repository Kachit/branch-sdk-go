package branchio

import (
	"context"
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

func Test_HTTP_Transport_SendRequestSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)

	body, _ := LoadStubResponseData("stubs/data/export/success.empty.json")
	httpmock.RegisterResponder(http.MethodGet, cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	resp, _ := transport.SendRequest(ctx, http.MethodGet, "foo", nil, nil)
	assert.NotEmpty(t, resp)
}

func Test_HTTP_Transport_SendRequestGET(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)

	body, _ := LoadStubResponseData("stubs/data/export/success.empty.json")

	httpmock.RegisterResponder(http.MethodGet, cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	resp, _ := transport.Get(ctx, "foo", nil)
	assert.NotEmpty(t, resp)
}

func Test_HTTP_Transport_SendRequestPOST(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)

	body, _ := LoadStubResponseData("stubs/data/export/success.empty.json")

	httpmock.RegisterResponder(http.MethodPost, cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	resp, _ := transport.Post(ctx, "foo", nil, nil)
	assert.NotEmpty(t, resp)
}
