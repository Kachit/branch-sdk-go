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
 * @unmarshal ClickEvent
 */
func (r *ExportsResource) GetEventData(link string) (*Response, error) {
	rsp, err := r.tr.http.Get(link)
	return NewResponse(rsp), err
}

type EventOntology struct {
	Click         []string `json:"eo_click"`
	CommerceEvent []string `json:"eo_commerce_event"`
	CustomEvent   []string `json:"eo_custom_event"`
	Install       []string `json:"eo_install"`
	Open          []string `json:"eo_open"`
	Reinstall     []string `json:"eo_reinstall"`
}

type ClickEvent struct {
	Id                                                   string `csv:"id" json:"id"`
	Name                                                 string `csv:"name" json:"name"`
	Timestamp                                            string `csv:"timestamp" json:"timestamp"`
	TimestampISO                                         string `csv:"timestamp_iso" json:"timestamp_iso"`
	Origin                                               string `csv:"origin" json:"origin"`
	LastAttributedTouchType                              string `csv:"last_attributed_touch_type" json:"last_attributed_touch_type"`
	LastAttributedTouchTimestamp                         string `csv:"last_attributed_touch_timestamp" json:"last_attributed_touch_timestamp"`
	LastAttributedTouchTimestampIso                      string `csv:"last_attributed_touch_timestamp_iso" json:"last_attributed_touch_timestamp_iso"`
	LastAttributedTouchDataTildeId                       string `csv:"last_attributed_touch_data_tilde_id" json:"last_attributed_touch_data_tilde_id"`
	LastAttributedTouchDataTildeCampaign                 string `csv:"last_attributed_touch_data_tilde_campaign" json:"last_attributed_touch_data_tilde_campaign"`
	LastAttributedTouchDataTildeCampaignId               string `csv:"last_attributed_touch_data_tilde_campaign_id" json:"last_attributed_touch_data_tilde_campaign_id"`
	LastAttributedTouchDataTildeChannel                  string `csv:"last_attributed_touch_data_tilde_channel" json:"last_attributed_touch_data_tilde_channel"`
	LastAttributedTouchDataTildeFeature                  string `csv:"last_attributed_touch_data_tilde_feature" json:"last_attributed_touch_data_tilde_feature"`
	LastAttributedTouchDataTildeStage                    string `csv:"last_attributed_touch_data_tilde_stage" json:"last_attributed_touch_data_tilde_stage"`
	LastAttributedTouchDataTildeTags                     string `csv:"last_attributed_touch_data_tilde_tags" json:"last_attributed_touch_data_tilde_tags"`
	LastAttributedTouchDataTildeAdvertisingPartnerName   string `csv:"last_attributed_touch_data_tilde_advertising_partner_name" json:"last_attributed_touch_data_tilde_advertising_partner_name"`
	LastAttributedTouchDataTildeSecondaryPublisher       string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_attributed_touch_data_tilde_creative_name       string `csv:"last_attributed_touch_data_tilde_creative_name" json:"last_attributed_touch_data_tilde_creative_name"`
	last_attributed_touch_data_tilde_creative_id         string `csv:"last_attributed_touch_data_tilde_creative_id" json:"last_attributed_touch_data_tilde_creative_id"`
	last_attributed_touch_data_tilde_ad_set_name         string `csv:"last_attributed_touch_data_tilde_ad_set_name" json:"last_attributed_touch_data_tilde_ad_set_name"`
	last_attributed_touch_data_tilde_ad_set_id           string `csv:"last_attributed_touch_data_tilde_ad_set_id" json:"last_attributed_touch_data_tilde_ad_set_id"`
	last_attributed_touch_data_tilde_ad_name             string `csv:"last_attributed_touch_data_tilde_ad_name" json:"last_attributed_touch_data_tilde_ad_name"`
	last_attributed_touch_data_tilde_ad_id               string `csv:"last_attributed_touch_data_tilde_ad_id" json:"last_attributed_touch_data_tilde_ad_id"`
	last_attributed_touch_data_tilde_branch_ad_format    string `csv:"last_attributed_touch_data_tilde_branch_ad_format" json:"last_attributed_touch_data_tilde_branch_ad_format"`
	last_attributed_touch_data_tilde_technology_partner  string `csv:"last_attributed_touch_data_tilde_technology_partner" json:"last_attributed_touch_data_tilde_technology_partner"`
	last_attributed_touch_data_tilde_banner_dimensions   string `csv:"last_attributed_touch_data_tilde_banner_dimensions" json:"last_attributed_touch_data_tilde_banner_dimensions"`
	last_attributed_touch_data_tilde_placement           string `csv:"last_attributed_touch_data_tilde_placement" json:"last_attributed_touch_data_tilde_placement"`
	last_attributed_touch_data_tilde_keyword_id          string `csv:"last_attributed_touch_data_tilde_keyword_id" json:"last_attributed_touch_data_tilde_keyword_id"`
	last_attributed_touch_data_tilde_agency              string `csv:"last_attributed_touch_data_tilde_agency" json:"last_attributed_touch_data_tilde_agency"`
	last_attributed_touch_data_tilde_optimization_model  string `csv:"last_attributed_touch_data_tilde_optimization_model" json:"last_attributed_touch_data_tilde_optimization_model"`
	last_attributed_touch_data_tilde_secondary_ad_format string `csv:"last_attributed_touch_data_tilde_secondary_ad_format" json:"last_attributed_touch_data_tilde_secondary_ad_format"`
	last_attributed_touch_data_tilde_journey_name        string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_attributed_touch_data_tilde_journey_id          string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_attributed_touch_data_tilde_view_name           string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_attributed_touch_data_tilde_view_id             string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_attributed_touch_data_plus_current_feature      string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_attributed_touch_data_plus_via_features         string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_attributed_touch_data_dollar_3p                 string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_attributed_touch_data_plus_web_format           string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_attributed_touch_data_custom_fields             string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	days_from_last_attributed_touch_to_event             string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	hours_from_last_attributed_touch_to_event            string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	minutes_from_last_attributed_touch_to_event          string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	seconds_from_last_attributed_touch_to_event          string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_timestamp                              string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_timestamp_iso                          string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_id                          string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_campaign                    string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_campaign_id                 string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_channel                     string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_feature                     string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_stage                       string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_tags                        string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_advertising_partner_name    string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_secondary_publisher         string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_creative_name               string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_creative_id                 string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_ad_set_name                 string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_ad_set_id                   string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_ad_name                     string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_ad_id                       string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_branch_ad_forma             string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_technology_partner          string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_banner_dimensions           string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_placement                   string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_keyword_id                  string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_agency                      string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_optimization_model          string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_secondary_ad_format         string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_plus_via_features                 string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_dollar_3p                         string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_plus_web_format                   string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_custom_fields                     string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	deep_linked                                          string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	first_event_for_user                                 string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_os                                         string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_os_version                                 string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_model                                      string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_browser                                    string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_geo_country_code                           string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_app_version                                string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_sdk_version                                string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_geo_dma_code                               string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_environment                                string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_platform                                   string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_aaid                                       string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_idfa                                       string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_idfv                                       string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_android_id                                 string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_limit_ad_tracking                          string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_user_agent                                 string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_ip                                         string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_developer_identity                         string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_language                                   string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_brand                                      string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	di_match_click_token                                 string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	event_data_revenue_in_usd                            string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	event_data_exchange_rate                             string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	event_data_transaction_id                            string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	event_data_revenue                                   string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	event_data_currency                                  string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	event_data_shipping                                  string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	event_data_tax                                       string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	event_data_coupon                                    string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	event_data_affiliation                               string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	event_data_search_query                              string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	event_data_description                               string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	custom_data                                          string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_attributed_touch_data_tilde_keyword             string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_cross_platform_id                          string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_past_cross_platform_ids                    string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_prob_cross_platform_ids                    string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	store_install_begin_timestamp                        string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	referrer_click_timestamp                             string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_os_version_android                         string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_geo_city_code                              string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_geo_city_en                                string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_http_referrer                              string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	event_timestamp                                      string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	customer_event_alias                                 string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_attributed_touch_data_tilde_customer_campaign   string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_attributed_touch_data_tilde_campaign_type       string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_tilde_campaign_type               string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_attributed_touch_data_tilde_agency_id           string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_attributed_touch_data_plus_touch_id             string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	last_cta_view_data_plus_touch_id                     string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_installer_package_name                     string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_cpu_type                                   string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_screen_width                               string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_screen_height                              string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_build                                      string `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	user_data_internet_connection_type                   string `csv:"user_data_internet_connection_type" json:"user_data_internet_connection_type"`
	hash_version                                         string `csv:"hash_version" json:"hash_version"`
}
