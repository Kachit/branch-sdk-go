package branchio

import (
	"context"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"testing"
)

type HttpRequestBuilderTestSuite struct {
	suite.Suite
	cfg      *Config
	ctx      context.Context
	testable *RequestBuilder
}

func (suite *HttpRequestBuilderTestSuite) SetupTest() {
	suite.cfg = BuildStubConfig()
	suite.ctx = context.Background()
	suite.testable = &RequestBuilder{cfg: suite.cfg}
}

func (suite *HttpRequestBuilderTestSuite) TestBuildUriWithoutQueryParams() {
	uri, err := suite.testable.buildUri("qwerty", nil)
	assert.Nil(suite.T(), err)
	assert.NotEmpty(suite.T(), uri)
	assert.Equal(suite.T(), ProdAPIUrl+"/qwerty", uri.String())
}

func (suite *HttpRequestBuilderTestSuite) TestBuildUriWithQueryParams() {
	data := make(map[string]interface{})
	data["foo"] = "bar"
	data["bar"] = "baz"

	uri, err := suite.testable.buildUri("qwerty", data)
	assert.Nil(suite.T(), err)
	assert.NotEmpty(suite.T(), uri)
	assert.Equal(suite.T(), ProdAPIUrl+"/qwerty?bar=baz&foo=bar", uri.String())
}

func (suite *HttpRequestBuilderTestSuite) TestBuildHeaders() {
	headers := suite.testable.buildHeaders()
	assert.NotEmpty(suite.T(), headers)
	assert.Equal(suite.T(), "application/json", headers.Get("Content-Type"))
}

func (suite *HttpRequestBuilderTestSuite) TestBuildBody() {
	data := make(map[string]interface{})
	data["foo"] = "bar"
	data["bar"] = "baz"

	body, _ := suite.testable.buildBody(data)
	assert.NotEmpty(suite.T(), body)
}

func (suite *HttpRequestBuilderTestSuite) TestBuildRequestGET() {
	result, err := suite.testable.BuildRequest(suite.ctx, "get", "foo", map[string]interface{}{"foo": "bar"}, map[string]interface{}{"foo": "bar"})
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.MethodGet, result.Method)
	assert.Equal(suite.T(), ProdAPIUrl+"/foo?foo=bar", result.URL.String())
	assert.Equal(suite.T(), "application/json", result.Header.Get("Content-Type"))
	assert.Nil(suite.T(), result.Body)
}

func (suite *HttpRequestBuilderTestSuite) TestBuildRequestPOST() {
	result, err := suite.testable.BuildRequest(suite.ctx, "post", "foo", map[string]interface{}{"foo": "bar"}, map[string]interface{}{"foo": "bar"})
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.MethodPost, result.Method)
	assert.Equal(suite.T(), ProdAPIUrl+"/foo?foo=bar", result.URL.String())
	assert.Equal(suite.T(), "application/json", result.Header.Get("Content-Type"))
	assert.NotEmpty(suite.T(), result.Body)
}

func TestHttpRequestBuilderTestSuite(t *testing.T) {
	suite.Run(t, new(HttpRequestBuilderTestSuite))
}

type HttpTransportTestSuite struct {
	suite.Suite
	cfg      *Config
	ctx      context.Context
	testable *Transport
}

func (suite *HttpTransportTestSuite) SetupTest() {
	suite.cfg = BuildStubConfig()
	suite.ctx = context.Background()
	suite.testable = BuildStubHttpTransport()
	httpmock.Activate()
}

func (suite *HttpTransportTestSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func (suite *HttpTransportTestSuite) TestSendRequestSuccess() {
	body, _ := LoadStubResponseData("stubs/data/export/ontology/success.empty.json")
	httpmock.RegisterResponder(http.MethodGet, suite.cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, err := suite.testable.SendRequest(suite.ctx, http.MethodGet, "foo", nil, nil)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)

	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
}

func (suite *HttpTransportTestSuite) TestSendRequestGET() {
	body, _ := LoadStubResponseData("stubs/data/export/ontology/success.empty.json")

	httpmock.RegisterResponder(http.MethodGet, suite.cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, err := suite.testable.Get(suite.ctx, "foo", nil)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)

	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
}

func (suite *HttpTransportTestSuite) TestSendRequestPOST() {
	body, _ := LoadStubResponseData("stubs/data/export/ontology/success.empty.json")

	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, err := suite.testable.Post(suite.ctx, "foo", nil, nil)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)

	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
}

func TestHttpTransportTestSuite(t *testing.T) {
	suite.Run(t, new(HttpTransportTestSuite))
}

type HttpResponseBodyTestSuite struct {
	suite.Suite
}

func (suite *HttpResponseBodyTestSuite) TestIsSuccess() {
	rsp := &ResponseBody{status: http.StatusAccepted}
	assert.True(suite.T(), rsp.IsSuccess())
	rsp.status = http.StatusMultipleChoices
	assert.False(suite.T(), rsp.IsSuccess())
	rsp.status = http.StatusBadRequest
	assert.False(suite.T(), rsp.IsSuccess())
}

func (suite *HttpResponseBodyTestSuite) TestGetErrorByDefault() {
	rsp := &ResponseBody{}
	assert.Equal(suite.T(), "Unknown error", rsp.GetError())
}

func (suite *HttpResponseBodyTestSuite) TestGetErrorEmptyMessage() {
	rsp := &ResponseBody{Error: &ResponseBodyError{}}
	assert.Equal(suite.T(), "Unknown error", rsp.GetError())
}

func (suite *HttpResponseBodyTestSuite) TestGetErrorNotEmptyMessage() {
	rsp := &ResponseBody{Error: &ResponseBodyError{Message: "Foo error"}}
	assert.Equal(suite.T(), "Foo error", rsp.GetError())
}

func TestHttpResponseBodyTestSuite(t *testing.T) {
	suite.Run(t, new(HttpResponseBodyTestSuite))
}

type HttpNewResponseHandlerTestSuite struct {
	suite.Suite
}

func (suite *HttpNewResponseHandlerTestSuite) TestNewResponseHandlerStream() {
	result := NewResponseHandler(ResponseContentTypeOctetStream)
	assert.Implements(suite.T(), (*ResponseHandlerInterface)(nil), result)
	assert.IsType(suite.T(), (*ResponseHandlerStream)(nil), result)
}

func (suite *HttpNewResponseHandlerTestSuite) TestNewResponseHandlerJson() {
	result := NewResponseHandler(ResponseContentTypeJson)
	assert.Implements(suite.T(), (*ResponseHandlerInterface)(nil), result)
	assert.IsType(suite.T(), (*ResponseHandlerJson)(nil), result)
}

func (suite *HttpNewResponseHandlerTestSuite) TestNewResponseHandlerXml() {
	result := NewResponseHandler(ResponseContentTypeXml)
	assert.Implements(suite.T(), (*ResponseHandlerInterface)(nil), result)
	assert.IsType(suite.T(), (*ResponseHandlerXml)(nil), result)
}

func (suite *HttpNewResponseHandlerTestSuite) TestNewHandlerByDefault() {
	result := NewResponseHandler("foo")
	assert.Implements(suite.T(), (*ResponseHandlerInterface)(nil), result)
	assert.IsType(suite.T(), (*ResponseHandlerJson)(nil), result)
}

func TestHttpNewResponseHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HttpNewResponseHandlerTestSuite))
}

type HttpResponseHandlerStreamTestSuite struct {
	suite.Suite
	testable *ResponseHandlerStream
}

func (suite *HttpResponseHandlerStreamTestSuite) SetupTest() {
	suite.testable = &ResponseHandlerStream{}
}

func (suite *HttpResponseHandlerStreamTestSuite) TestUnmarshalBody() {
	reportData, _ := ioutil.ReadFile("stubs/data/export/events/eo-click-v2.csv")
	events := []*Event{}
	err := suite.testable.UnmarshalBody(reportData, &events)
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
}

func (suite *HttpResponseHandlerStreamTestSuite) TestReadBody() {
	expected, _ := ioutil.ReadFile("stubs/data/export/events/eo-click-v2.csv")
	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/data/export/events/eo-click-v2.csv")
	data, err := suite.testable.ReadBody(resp)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), data)
	assert.Equal(suite.T(), expected, data)
	assert.Empty(suite.T(), body)
}

func (suite *HttpResponseHandlerStreamTestSuite) TestRestoreBody() {
	expectedRaw, _ := ioutil.ReadFile("stubs/data/export/events/eo-click-v2.csv")
	//expectedGzipped, _ := loadStubResponseDataGzipped("stubs/data/export/events/eo-click-v2.csv")
	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/data/export/events/eo-click-v2.csv")
	data, err := suite.testable.ReadBody(resp)
	closer, err := suite.testable.RestoreBody(data)
	resp.Body = closer

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRaw, data)
	assert.NotEmpty(suite.T(), body)
	//fmt.Println(expectedGzipped)
	//fmt.Println(body)
	//assert.Equal(suite.T(), expectedGzipped, body)
}

func TestHttpResponseHandlerStreamTestSuite(t *testing.T) {
	suite.Run(t, new(HttpResponseHandlerStreamTestSuite))
}

type HttpResponseHandlerJsonTestSuite struct {
	suite.Suite
	testable *ResponseHandlerJson
}

func (suite *HttpResponseHandlerJsonTestSuite) SetupTest() {
	suite.testable = &ResponseHandlerJson{}
}

func (suite *HttpResponseHandlerJsonTestSuite) TestUnmarshalBody() {
	reportData, _ := ioutil.ReadFile("stubs/data/export/ontology/success.full.json")
	var data EventOntology
	err := suite.testable.UnmarshalBody(reportData, &data)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "https://branch-exports-web.foo-bar.amazonaws.com/eo_click.csv", data.Click[0])
}

func (suite *HttpResponseHandlerJsonTestSuite) TestReadBody() {
	expected, _ := ioutil.ReadFile("stubs/data/export/ontology/success.full.json")
	resp := BuildStubResponseFromFile(http.StatusOK, "stubs/data/export/ontology/success.full.json")
	data, err := suite.testable.ReadBody(resp)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), data)
	assert.Equal(suite.T(), expected, data)
	assert.Empty(suite.T(), body)
}

func (suite *HttpResponseHandlerJsonTestSuite) TestRestoreBody() {
	expected, _ := ioutil.ReadFile("stubs/data/export/ontology/success.full.json")
	resp := BuildStubResponseFromFile(http.StatusOK, "stubs/data/export/ontology/success.full.json")
	data, _ := suite.testable.ReadBody(resp)
	closer, err := suite.testable.RestoreBody(data)
	resp.Body = closer

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, data)
	assert.NotEmpty(suite.T(), body)
	assert.Equal(suite.T(), expected, body)
}

func TestHttpResponseHandlerJsonTestSuite(t *testing.T) {
	suite.Run(t, new(HttpResponseHandlerJsonTestSuite))
}

type HttpResponseHandlerXmlTestSuite struct {
	suite.Suite
	testable *ResponseHandlerXml
}

func (suite *HttpResponseHandlerXmlTestSuite) SetupTest() {
	suite.testable = &ResponseHandlerXml{}
}

func (suite *HttpResponseHandlerXmlTestSuite) TestReadBody() {
	expected, _ := ioutil.ReadFile("stubs/data/export/events/forbidden.xml")
	resp := BuildStubResponseFromFile(http.StatusOK, "stubs/data/export/events/forbidden.xml")
	data, err := suite.testable.ReadBody(resp)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), data)
	assert.Equal(suite.T(), expected, data)
	assert.Empty(suite.T(), body)
}

func (suite *HttpResponseHandlerXmlTestSuite) TestUnmarshalBody() {
	reportData, _ := ioutil.ReadFile("stubs/data/export/events/forbidden.xml")
	var result EventError
	err := suite.testable.UnmarshalBody(reportData, &result)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "AccessDenied", result.Code)
	assert.Equal(suite.T(), "Access Denied", result.Message)
	assert.Equal(suite.T(), "QWERTY", result.RequestId)
	assert.Equal(suite.T(), "qwerty=", result.HostId)
}

func (suite *HttpResponseHandlerXmlTestSuite) TestRestoreBody() {
	expected, _ := ioutil.ReadFile("stubs/data/export/events/forbidden.xml")
	resp := BuildStubResponseFromFile(http.StatusOK, "stubs/data/export/events/forbidden.xml")
	data, _ := suite.testable.ReadBody(resp)
	closer, err := suite.testable.RestoreBody(data)
	resp.Body = closer

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, data)
	assert.NotEmpty(suite.T(), body)
	assert.Equal(suite.T(), expected, body)
}

func TestHttpResponseHandlerXmlTestSuite(t *testing.T) {
	suite.Run(t, new(HttpResponseHandlerXmlTestSuite))
}
