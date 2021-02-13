package branchio

import (
	"time"
)

type ExportsResource struct {
	*ResourceAbstract
}

/**
 * @unmarshal ExportData
 */
func (r *ExportsResource) Export(date time.Time) (*Response, error) {
	post := make(map[string]interface{})
	post["export_date"] = date.Format("2006-01-02")
	post["branch_key"] = r.cfg.Key
	post["branch_secret"] = r.cfg.Secret
	return r.Post("v3/export", post, nil)
}
