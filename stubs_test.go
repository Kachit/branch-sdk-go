package branchio

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"net/http"
	"strings"
)

func BuildStubConfig() *Config {
	return &Config{
		Uri:    ProdAPIUrl,
		Key:    "Key",
		Secret: "Secret",
	}
}

func BuildStubHttpTransport() *Transport {
	return NewHttpTransport(BuildStubConfig(), &http.Client{})
}

func LoadStubResponseData(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func BuildStubResponseFromString(statusCode int, json string) *http.Response {
	body := ioutil.NopCloser(strings.NewReader(json))
	return &http.Response{Body: body, StatusCode: statusCode, Header: http.Header{}}
}

func BuildStubResponseFromFile(statusCode int, path string) *http.Response {
	data, _ := LoadStubResponseData(path)
	body := ioutil.NopCloser(bytes.NewReader(data))
	return &http.Response{Body: body, StatusCode: statusCode, Header: http.Header{}}
}

func buildStubResponseFromGzip(statusCode int, path string) *http.Response {
	data, _ := loadStubResponseDataGzipped(path)
	body := ioutil.NopCloser(bytes.NewReader(data))
	return &http.Response{Body: body, StatusCode: statusCode, Header: http.Header{}}
}

func loadStubResponseDataGzipped(path string) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	_, err = zw.Write(data)
	err = zw.Close()
	return buf.Bytes(), err
}
