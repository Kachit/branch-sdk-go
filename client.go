package branchio

import "net/http"

//Client common
type Client struct {
	transport *Transport
}

//NewClientFromConfig Create new client from config
func NewClientFromConfig(config *Config, cl *http.Client) *Client {
	if cl == nil {
		cl = &http.Client{}
	}
	transport := NewHttpTransport(config, cl)
	return &Client{transport}
}

//Export resource
func (c *Client) Export() *ExportResource {
	return &ExportResource{ResourceAbstract: NewResourceAbstract(c.transport)}
}
