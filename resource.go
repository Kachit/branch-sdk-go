package branchio

import (
	"fmt"
	"net/http"
)

//ResourceAbstract base resource
type ResourceAbstract struct {
	tr  *Transport
	cfg *Config
}

//UnmarshalResponse method
func (ra *ResourceAbstract) unmarshalResponse(resp *http.Response, v interface{}) error {
	contentType := resp.Header.Get("Content-Type")
	responseHandler := NewResponseHandler(contentType)

	bodyBytes, err := responseHandler.ReadBody(resp)
	if err != nil {
		return fmt.Errorf("ResourceAbstract.unmarshalResponse read body: %v", err)
	}
	//reset the response body to the original unread state
	body, err := responseHandler.RestoreBody(bodyBytes)
	if err != nil {
		return fmt.Errorf("ResourceAbstract.unmarshalResponse read body: %v", err)
	}
	resp.Body = body
	return responseHandler.UnmarshalBody(bodyBytes, v)
}

//NewResourceAbstract Create new resource abstract
func NewResourceAbstract(transport *Transport) ResourceAbstract {
	return ResourceAbstract{tr: transport, cfg: transport.rb.cfg}
}
