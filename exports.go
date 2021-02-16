package branchio

import (
	"time"
)

type ExportsResource struct {
	*ResourceAbstract
}

/**
 * @unmarshal EventOntology
 */
func (r *ExportsResource) GetEventOntology(date time.Time) (*Response, error) {
	post := make(map[string]interface{})
	post["export_date"] = date.Format("2006-01-02")
	post["branch_key"] = r.cfg.Key
	post["branch_secret"] = r.cfg.Secret
	return r.Post("v3/export", post, nil)
}

/**
 * @unmarshal ExportData
 */
func (r *ExportsResource) GetEventData(link string) (*Response, error) {
	rsp, err := r.tr.http.Get(link)
	return NewResponse(rsp), err
}

type EventOntology struct {
	Click         []*string `json:"eo_click"`
	CommerceEvent []*string `json:"eo_commerce_event"`
	CustomEvent   []*string `json:"eo_custom_event"`
	Install       []*string `json:"eo_install"`
	Open          []*string `json:"eo_open"`
	Reinstall     []*string `json:"eo_reinstall"`
}
