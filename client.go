package branchio

import "net/http"

//Client common
type Client struct {
	transport *Transport
}

//NewClientFromConfig Create new client from config
func NewClientFromConfig(config *Config, cl *http.Client) (*Client, error) {
	err := config.IsValid()
	if err != nil {
		return nil, err
	}
	if cl == nil {
		cl = &http.Client{}
	}
	transport := NewHttpTransport(config, cl)
	return &Client{transport}, nil
}

//Export resource
func (c *Client) Export() *ExportResource {
	return &ExportResource{NewResourceAbstract(c.transport)}
}
