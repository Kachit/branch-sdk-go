package branchio

import "net/http"

type Client struct {
	transport *Transport
}

func NewClientFromConfig(config *Config, cl *http.Client) *Client {
	if cl == nil {
		cl = &http.Client{}
	}
	transport := NewHttpTransport(config, cl)
	return &Client{transport}
}

func (c *Client) Export() *ExportResource {
	return &ExportResource{ResourceAbstract: NewResourceAbstract(c.transport)}
}
