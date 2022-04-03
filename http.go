package branchio

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//NewDefaultHttpClient create new http client
func NewDefaultHttpClient() *http.Client {
	tr := &http.Transport{}
	return &http.Client{Transport: tr}
}

//RequestBuilder handler
type RequestBuilder struct {
	cfg *Config
}

//buildUri method
func (rb *RequestBuilder) buildUri(path string, query map[string]interface{}) (uri *url.URL, err error) {
	u, err := url.Parse(rb.cfg.Uri)
	if err != nil {
		return nil, fmt.Errorf("RequestBuilder.buildUri parse: %v", err)
	}
	u.Path = "/" + path
	u.RawQuery = rb.buildQueryParams(query)
	return u, err
}

//buildQueryParams method
func (rb *RequestBuilder) buildQueryParams(query map[string]interface{}) string {
	q := url.Values{}
	if query != nil {
		for k, v := range query {
			q.Set(k, fmt.Sprintf("%v", v))
		}
	}
	return q.Encode()
}

//buildHeaders method
func (rb *RequestBuilder) buildHeaders() http.Header {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	return headers
}

//buildBody method
func (rb *RequestBuilder) buildBody(data map[string]interface{}) (io.Reader, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("RequestBuilder.buildBody json convert: %v", err)
	}
	return bytes.NewBuffer(b), nil
}

//BuildRequest method
func (rb *RequestBuilder) BuildRequest(ctx context.Context, method string, path string, query map[string]interface{}, body map[string]interface{}) (req *http.Request, err error) {
	method = strings.ToUpper(method)
	//build body
	var bodyReader io.Reader
	if method == http.MethodPost {
		bodyReader, err = rb.buildBody(body)
		if err != nil {
			return nil, fmt.Errorf("transport.request build request body: %v", err)
		}
	}
	//build uri
	uri, err := rb.buildUri(path, query)
	if err != nil {
		return nil, fmt.Errorf("transport.request build uri: %v", err)
	}
	//build request
	req, err = http.NewRequestWithContext(ctx, method, uri.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("transport.request new request error: %v", err)
	}
	//build headers
	req.Header = rb.buildHeaders()
	return req, nil
}

//NewHttpTransport create new http transport
func NewHttpTransport(config *Config, h *http.Client) *Transport {
	if h == nil {
		h = NewDefaultHttpClient()
	}
	rb := &RequestBuilder{cfg: config}
	return &Transport{http: h, rb: rb}
}

//Transport wrapper
type Transport struct {
	http *http.Client
	rb   *RequestBuilder
}

//SendRequest method
func (t *Transport) SendRequest(ctx context.Context, method string, path string, query map[string]interface{}, body map[string]interface{}) (resp *http.Response, err error) {
	req, err := t.rb.BuildRequest(ctx, method, path, query, body)
	if err != nil {
		return nil, fmt.Errorf("transport.SendRequest: %v", err)
	}
	return t.http.Do(req)
}

//Get method
func (t *Transport) Get(ctx context.Context, path string, query map[string]interface{}) (resp *http.Response, err error) {
	return t.SendRequest(ctx, http.MethodGet, path, query, nil)
}

//Post method
func (t *Transport) Post(ctx context.Context, path string, body map[string]interface{}, query map[string]interface{}) (resp *http.Response, err error) {
	return t.SendRequest(ctx, http.MethodPost, path, query, body)
}

//Response wrapper
type Response struct {
	raw *http.Response
	csv *CSV
}

//IsSuccess method
func (r *Response) IsSuccess() bool {
	return r.raw.StatusCode < http.StatusMultipleChoices
}

//GetRawResponse method
func (r *Response) GetRawResponse() *http.Response {
	return r.raw
}

//GetRawBody method
func (r *Response) GetRawBody() string {
	body, _ := r.ReadBody()
	return string(body)
}

//Unmarshal method
func (r *Response) Unmarshal(v interface{}) error {
	data, err := r.ReadBody()
	if err != nil {
		return fmt.Errorf("Response@Unmarshal read body: %v", err)
	}
	return json.Unmarshal(data, &v)
}

//UnmarshalCSV method
func (r *Response) UnmarshalCSV(v interface{}) error {
	data, err := r.ReadGzipBody()
	if err != nil {
		return fmt.Errorf("Response@Unmarshal read gzip body: %v", err)
	}
	return r.csv.Unmarshal(data, v)
}

//UnmarshalError method
func (r *Response) UnmarshalError(v interface{}) error {
	body, err := r.ReadBody()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

//ReadBody method
func (r *Response) ReadBody() ([]byte, error) {
	defer r.raw.Body.Close()
	return ioutil.ReadAll(r.raw.Body)
}

//ReadGzipBody method
func (r *Response) ReadGzipBody() ([]byte, error) {
	defer r.raw.Body.Close()
	zr, _ := gzip.NewReader(r.raw.Body)
	defer zr.Close()
	return ioutil.ReadAll(zr)
}

//NewResponse create new response
func NewResponse(raw *http.Response) *Response {
	return &Response{raw: raw, csv: &CSV{}}
}

//ResponseBody struct
type ResponseBody struct {
	status int
	Error  *ResponseBodyError `json:"error,omitempty"`
}

//ResponseBodyError struct
type ResponseBodyError struct {
	Message string `json:"message,omitempty"`
	Code    int64  `json:"code,omitempty"`
}

//IsSuccess method
func (r *ResponseBody) IsSuccess() bool {
	return r.status < http.StatusMultipleChoices
}

//GetError method
func (r *ResponseBody) GetError() string {
	err := "Unknown error"
	if r.Error != nil && r.Error.Message != "" {
		err = r.Error.Message
	}
	return err
}

//IsEmptyObjectResponseData func
func isEmptyObjectResponseData(data []byte) bool {
	return data[0] == 123 && data[1] == 125
}

//UnmarshalJson method
func UnmarshalJson(data []byte, v interface{}) error {
	return json.Unmarshal(data, &v)
}
