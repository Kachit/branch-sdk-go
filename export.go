package branchio

import (
	"time"
)

type ExportResource struct {
	*ResourceAbstract
}

/**
 * @unmarshal EventOntology
 */
func (r *ExportResource) GetEventOntology(date time.Time) (*Response, error) {
	post := make(map[string]interface{})
	post["export_date"] = date.Format("2006-01-02")
	post["branch_key"] = r.cfg.Key
	post["branch_secret"] = r.cfg.Secret
	return r.Post("v3/export", post, nil)
}

/**
 * @unmarshal ClickEvent
 */
func (r *ExportResource) GetEventData(link string) (*Response, error) {
	rsp, err := r.tr.http.Get(link)
	return NewResponse(rsp), err
}
