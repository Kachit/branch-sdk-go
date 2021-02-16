package branchio

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func NewDefaultHttpClient() *http.Client {
	tr := &http.Transport{}
	return &http.Client{Transport: tr}
}

type RequestBuilder struct {
	cfg *Config
}

func (rb *RequestBuilder) buildUri(path string, query map[string]interface{}) (uri *url.URL, err error) {
	u, err := url.Parse(rb.cfg.Uri)
	if err != nil {
		return nil, fmt.Errorf("RequestBuilder@buildUri parse: %v", err)
	}
	u.Path = "/" + path
	u.RawQuery = rb.buildQueryParams(query)
	return u, err
}

func (rb *RequestBuilder) buildQueryParams(query map[string]interface{}) string {
	q := url.Values{}
	if query != nil {
		for k, v := range query {
			q.Set(k, fmt.Sprintf("%v", v))
		}
	}
	return q.Encode()
}

func (rb *RequestBuilder) buildHeaders() http.Header {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	return headers
}

func (rb *RequestBuilder) buildBody(data map[string]interface{}) (io.Reader, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("RequestBuilder@buildBody json convert: %v", err)
	}
	return bytes.NewBuffer(b), nil
}

func NewHttpTransport(config *Config, h *http.Client) *Transport {
	if h == nil {
		h = NewDefaultHttpClient()
	}
	rb := &RequestBuilder{cfg: config}
	return &Transport{http: h, rb: rb}
}

type Transport struct {
	http *http.Client
	rb   *RequestBuilder
}

func (t *Transport) Request(method string, path string, query map[string]interface{}, body map[string]interface{}) (resp *http.Response, err error) {
	//build uri
	uri, err := t.rb.buildUri(path, query)
	if err != nil {
		return nil, fmt.Errorf("transport@request build uri: %v", err)
	}
	//build body
	bodyReader, err := t.rb.buildBody(body)
	if err != nil {
		return nil, fmt.Errorf("transport@request build request body: %v", err)
	}
	//build request
	req, err := http.NewRequest(strings.ToUpper(method), uri.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("transport@request new request error: %v", err)
	}
	//build headers
	req.Header = t.rb.buildHeaders()
	return t.http.Do(req)
}

func (t *Transport) Get(path string, query map[string]interface{}) (resp *http.Response, err error) {
	return t.Request("GET", path, query, nil)
}

func (t *Transport) Post(path string, body map[string]interface{}, query map[string]interface{}) (resp *http.Response, err error) {
	return t.Request("POST", path, query, body)
}

type Response struct {
	raw *http.Response
}

func (r *Response) IsSuccess() bool {
	return r.raw.StatusCode < http.StatusMultipleChoices
}

func (r *Response) GetRawResponse() *http.Response {
	return r.raw
}

func (r *Response) GetRawBody() string {
	body, _ := r.ReadBody()
	return string(body)
}

func (r *Response) Unmarshal(v interface{}) error {
	data, err := r.ReadBody()
	if err != nil {
		return fmt.Errorf("Response@Unmarshal read body: %v", err)
	}
	return json.Unmarshal(data, &v)
}

func (r *Response) UnmarshalGzip(v interface{}) error {
	data, err := r.ReadGzipBody()
	if err != nil {
		return fmt.Errorf("Response@Unmarshal read gzip body: %v", err)
	}
	return json.Unmarshal(data, &v)
}

func (r *Response) UnmarshalError(v interface{}) error {
	body, err := r.ReadBody()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

func (r *Response) ReadBody() ([]byte, error) {
	defer r.raw.Body.Close()
	return ioutil.ReadAll(r.raw.Body)
}

func (r *Response) ReadGzipBody() ([]byte, error) {
	defer r.raw.Body.Close()
	zr, _ := gzip.NewReader(r.raw.Body)
	defer zr.Close()
	return ioutil.ReadAll(zr)
}

func NewResponse(raw *http.Response) *Response {
	return &Response{raw: raw}
}
