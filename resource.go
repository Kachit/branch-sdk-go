package branchio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//ResourceAbstract base resource
type ResourceAbstract struct {
	tr  *Transport
	cfg *Config
}

//UnmarshalResponse method
func (ra *ResourceAbstract) unmarshalResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bodyBytes))
	if err != nil {
		return fmt.Errorf("ResourceAbstract.unmarshalResponse read body: %v", err)
	}
	//reset the response body to the original unread state
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return json.Unmarshal(bodyBytes, &v)
}

//NewResourceAbstract Create new resource abstract
func NewResourceAbstract(transport *Transport) *ResourceAbstract {
	return &ResourceAbstract{tr: transport, cfg: transport.rb.cfg}
}
