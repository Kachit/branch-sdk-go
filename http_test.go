package branchio

import (
	"context"
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
