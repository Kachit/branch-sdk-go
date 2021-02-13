package branchio

import "fmt"

type ResourceAbstract struct {
	tr  *Transport
	cfg *Config
}

func (r *ResourceAbstract) Get(path string, query map[string]interface{}) (*Response, error) {
	rsp, err := r.tr.Get(path, query)
	if err != nil {
		return nil, fmt.Errorf("ResourceAbstract@get request: %v", err)
	}
	return NewResponse(rsp), nil
}

func (r *ResourceAbstract) Post(path string, body map[string]interface{}, query map[string]interface{}) (*Response, error) {
	rsp, err := r.tr.Post(path, body, query)
	if err != nil {
		return nil, fmt.Errorf("ResourceAbstract@post request: %v", err)
	}
	return NewResponse(rsp), nil
}

func NewResourceAbstract(transport *Transport) *ResourceAbstract {
	return &ResourceAbstract{tr: transport, cfg: transport.rb.cfg}
}
