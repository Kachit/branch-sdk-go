package branchio

import (
	"fmt"
	"net/http"
	"time"
)

//ExportResource resource wrapper
type ExportResource struct {
	*ResourceAbstract
}

//GetEventOntology Get events ontology data
func (r *ExportResource) GetEventOntology(date time.Time) (*EventOntology, *http.Response, error) {
	post := make(map[string]interface{})
	post["export_date"] = date.Format("2006-01-02")
	post["branch_key"] = r.cfg.Key
	post["branch_secret"] = r.cfg.Secret
	rsp, err := r.tr.Post("v3/export", post, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("ExportResource.GetEventOntology error: %v", err)
	}
	var result EventOntology
	err = r.unmarshalResponse(rsp, &result)
	if err != nil {
		return nil, rsp, fmt.Errorf("ExportResource.GetEventOntology error: %v", err)
	}
	return &result, rsp, nil
}

//GetEventData Get event data by link
func (r *ExportResource) GetEventData(link string) (*http.Response, error) {
	rsp, err := r.tr.http.Get(link)
	if err != nil {
		return nil, fmt.Errorf("ExportResource.GetEventData error: %v", err)
	}
	return rsp, err
}
