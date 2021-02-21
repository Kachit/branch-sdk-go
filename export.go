package branchio

import (
	"fmt"
	"time"
)

//ExportResource resource wrapper
type ExportResource struct {
	*ResourceAbstract
}

//GetEventOntology Get events ontology data
//@unmarshal EventOntology
func (r *ExportResource) GetEventOntology(date time.Time) (*Response, error) {
	post := make(map[string]interface{})
	post["export_date"] = date.Format("2006-01-02")
	post["branch_key"] = r.cfg.Key
	post["branch_secret"] = r.cfg.Secret
	return r.Post("v3/export", post, nil)
}

//GetEventData Get event data by link
//@unmarshal EventOntology
func (r *ExportResource) GetEventData(link string) (*Response, error) {
	rsp, err := r.tr.http.Get(link)
	if err != nil {
		return nil, fmt.Errorf("ExportResource@GetEventData error: %v", err)
	}
	return NewResponse(rsp), err
}
