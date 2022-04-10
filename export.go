package branchio

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

//EventOntologyResponse struct
type EventOntologyResponse struct {
	*ResponseBody
	Data *EventOntology `json:"data,omitempty"`
}

//IsEmpty method
func (r *EventOntologyResponse) IsEmpty() bool {
	return r.Data == nil
}

//EventOntology common struct
type EventOntology struct {
	BranchCtaView        []string `json:"branch_cta_view"`
	Click                []string `json:"eo_click"`
	CommerceEvent        []string `json:"eo_commerce_event"`
	ContentEvent         []string `json:"eo_content_event"`
	CustomEvent          []string `json:"eo_custom_event"`
	Impression           []string `json:"eo_impression"`
	Install              []string `json:"eo_install"`
	Open                 []string `json:"eo_open"`
	PageView             []string `json:"eo_pageview"`
	Reinstall            []string `json:"eo_reinstall"`
	SmsSent              []string `json:"eo_sms_sent"`
	UserLifecycleEvent   []string `json:"eo_user_lifecycle_event"`
	WebSessionStart      []string `json:"eo_web_session_start"`
	WebToAppAutoRedirect []string `json:"eo_web_to_app_auto_redirect"`
	Dismissal            []string `json:"eo_dismissal"`
}

//IsEmpty method
func (r *EventOntology) IsEmpty() bool {
	return len(r.Click) == 0
}

//EventResponse struct
type EventResponse struct {
	*ResponseBody
	Data []*Event `json:"data,omitempty"`
}

//Event common struct
type Event struct {
	Id                                                 CustomInteger   `csv:"id" json:"id"`
	Name                                               string          `csv:"name" json:"name"`
	Timestamp                                          CustomInteger   `csv:"timestamp" json:"timestamp"`
	TimestampISO                                       CustomTimestamp `csv:"timestamp_iso" json:"timestamp_iso"`
	Origin                                             string          `csv:"origin" json:"origin"`
	LastAttributedTouchType                            string          `csv:"last_attributed_touch_type" json:"last_attributed_touch_type"`
	LastAttributedTouchTimestamp                       string          `csv:"last_attributed_touch_timestamp" json:"last_attributed_touch_timestamp"`
	LastAttributedTouchTimestampIso                    string          `csv:"last_attributed_touch_timestamp_iso" json:"last_attributed_touch_timestamp_iso"`
	LastAttributedTouchDataTildeId                     CustomInteger   `csv:"last_attributed_touch_data_tilde_id" json:"last_attributed_touch_data_tilde_id"`
	LastAttributedTouchDataTildeCampaign               string          `csv:"last_attributed_touch_data_tilde_campaign" json:"last_attributed_touch_data_tilde_campaign"`
	LastAttributedTouchDataTildeCampaignId             string          `csv:"last_attributed_touch_data_tilde_campaign_id" json:"last_attributed_touch_data_tilde_campaign_id"`
	LastAttributedTouchDataTildeChannel                string          `csv:"last_attributed_touch_data_tilde_channel" json:"last_attributed_touch_data_tilde_channel"`
	LastAttributedTouchDataTildeFeature                string          `csv:"last_attributed_touch_data_tilde_feature" json:"last_attributed_touch_data_tilde_feature"`
	LastAttributedTouchDataTildeStage                  string          `csv:"last_attributed_touch_data_tilde_stage" json:"last_attributed_touch_data_tilde_stage"`
	LastAttributedTouchDataTildeTags                   string          `csv:"last_attributed_touch_data_tilde_tags" json:"last_attributed_touch_data_tilde_tags"`
	LastAttributedTouchDataTildeAdvertisingPartnerName string          `csv:"last_attributed_touch_data_tilde_advertising_partner_name" json:"last_attributed_touch_data_tilde_advertising_partner_name"`
	LastAttributedTouchDataTildeSecondaryPublisher     string          `csv:"last_attributed_touch_data_tilde_secondary_publisher" json:"last_attributed_touch_data_tilde_secondary_publisher"`
	LastAttributedTouchDataTildeCreativeName           string          `csv:"last_attributed_touch_data_tilde_creative_name" json:"last_attributed_touch_data_tilde_creative_name"`
	LastAttributedTouchDataTildeCreativeId             string          `csv:"last_attributed_touch_data_tilde_creative_id" json:"last_attributed_touch_data_tilde_creative_id"`
	LastAttributedTouchDataTildeAdSetName              string          `csv:"last_attributed_touch_data_tilde_ad_set_name" json:"last_attributed_touch_data_tilde_ad_set_name"`
	LastAttributedTouchDataTildeAdSetId                string          `csv:"last_attributed_touch_data_tilde_ad_set_id" json:"last_attributed_touch_data_tilde_ad_set_id"`
	LastAttributedTouchDataTildeAdName                 string          `csv:"last_attributed_touch_data_tilde_ad_name" json:"last_attributed_touch_data_tilde_ad_name"`
	LastAttributedTouchDataTildeAdId                   string          `csv:"last_attributed_touch_data_tilde_ad_id" json:"last_attributed_touch_data_tilde_ad_id"`
	LastAttributedTouchDataTildeBranchAdFormat         string          `csv:"last_attributed_touch_data_tilde_branch_ad_format" json:"last_attributed_touch_data_tilde_branch_ad_format"`
	LastAttributedTouchDataTildeTechnologyPartner      string          `csv:"last_attributed_touch_data_tilde_technology_partner" json:"last_attributed_touch_data_tilde_technology_partner"`
	LastAttributedTouchDataTildeBannerDimensions       string          `csv:"last_attributed_touch_data_tilde_banner_dimensions" json:"last_attributed_touch_data_tilde_banner_dimensions"`
	LastAttributedTouchDataTildePlacement              string          `csv:"last_attributed_touch_data_tilde_placement" json:"last_attributed_touch_data_tilde_placement"`
	LastAttributedTouchDataTildeKeywordId              string          `csv:"last_attributed_touch_data_tilde_keyword_id" json:"last_attributed_touch_data_tilde_keyword_id"`
	LastAttributedTouchDataTildeAgency                 string          `csv:"last_attributed_touch_data_tilde_agency" json:"last_attributed_touch_data_tilde_agency"`
	LastAttributedTouchDataTildeOptimizationModel      string          `csv:"last_attributed_touch_data_tilde_optimization_model" json:"last_attributed_touch_data_tilde_optimization_model"`
	LastAttributedTouchDataTildeSecondaryAdFormat      string          `csv:"last_attributed_touch_data_tilde_secondary_ad_format" json:"last_attributed_touch_data_tilde_secondary_ad_format"`
	LastAttributedTouchDataTildeJourneyName            string          `csv:"last_attributed_touch_data_tilde_journey_name" json:"last_attributed_touch_data_tilde_journey_name"`
	LastAttributedTouchDataTildeJourneyId              string          `csv:"last_attributed_touch_data_tilde_journey_id" json:"last_attributed_touch_data_tilde_journey_id"`
	LastAttributedTouchDataTildeViewName               string          `csv:"last_attributed_touch_data_tilde_view_name" json:"last_attributed_touch_data_tilde_view_name"`
	LastAttributedTouchDataTildeViewId                 string          `csv:"last_attributed_touch_data_tilde_view_id" json:"last_attributed_touch_data_tilde_view_id"`
	LastAttributedTouchDataPlusCurrentFeature          string          `csv:"last_attributed_touch_data_plus_current_feature" json:"last_attributed_touch_data_plus_current_feature"`
	LastAttributedTouchDataPlusViaFeatures             string          `csv:"last_attributed_touch_data_plus_via_features" json:"last_attributed_touch_data_plus_via_features"`
	LastAttributedTouchDataDollar3P                    string          `csv:"last_attributed_touch_data_dollar_3p" json:"last_attributed_touch_data_dollar_3p"`
	LastAttributedTouchDataPlusWebFormat               string          `csv:"last_attributed_touch_data_plus_web_format" json:"last_attributed_touch_data_plus_web_format"`
	LastAttributedTouchDataCustomFields                string          `csv:"last_attributed_touch_data_custom_fields" json:"last_attributed_touch_data_custom_fields"`
	DaysFromLastAttributedTouchToEvent                 string          `csv:"days_from_last_attributed_touch_to_event" json:"days_from_last_attributed_touch_to_event"`
	HoursFromLastAttributedTouchToEvent                string          `csv:"hours_from_last_attributed_touch_to_event" json:"hours_from_last_attributed_touch_to_event"`
	MinutesFromLastAttributedTouchToEvent              string          `csv:"minutes_from_last_attributed_touch_to_event" json:"minutes_from_last_attributed_touch_to_event"`
	SecondsFromLastAttributedTouchToEvent              string          `csv:"seconds_from_last_attributed_touch_to_event" json:"seconds_from_last_attributed_touch_to_event"`
	LastCtaViewTimestamp                               string          `csv:"last_cta_view_timestamp" json:"last_cta_view_timestamp"`
	LastCtaViewTimestampIso                            string          `csv:"last_cta_view_timestamp_iso" json:"last_cta_view_timestamp_iso"`
	LastCtaViewDataTildeId                             string          `csv:"last_cta_view_data_tilde_id" json:"last_cta_view_data_tilde_id"`
	LastCtaViewDataTildeCampaign                       string          `csv:"last_cta_view_data_tilde_campaign" json:"last_cta_view_data_tilde_campaign"`
	LastCtaViewDataTildeCampaignId                     string          `csv:"last_cta_view_data_tilde_campaign_id" json:"last_cta_view_data_tilde_campaign_id"`
	LastCtaViewDataTildeChannel                        string          `csv:"last_cta_view_data_tilde_channel" json:"last_cta_view_data_tilde_channel"`
	LastCtaViewDataTildeFeature                        string          `csv:"last_cta_view_data_tilde_feature" json:"last_cta_view_data_tilde_feature"`
	LastCtaViewDataTildeStage                          string          `csv:"last_cta_view_data_tilde_stage" json:"last_cta_view_data_tilde_stage"`
	LastCtaViewDataTildeTags                           string          `csv:"last_cta_view_data_tilde_tags" json:"last_cta_view_data_tilde_tags"`
	LastCtaViewDataTildeAdvertisingPartnerName         string          `csv:"last_cta_view_data_tilde_advertising_partner_name" json:"last_cta_view_data_tilde_advertising_partner_name"`
	LastCtaViewDataTildeSecondaryPublisher             string          `csv:"last_cta_view_data_tilde_secondary_publisher" json:"last_cta_view_data_tilde_secondary_publisher"`
	LastCtaViewDataTildeCreativeName                   string          `csv:"last_cta_view_data_tilde_creative_name" json:"last_cta_view_data_tilde_creative_name"`
	LastCtaViewDataTildeCreativeId                     string          `csv:"last_cta_view_data_tilde_creative_id" json:"last_cta_view_data_tilde_creative_id"`
	LastCtaViewDataTildeAdSetName                      string          `csv:"last_cta_view_data_tilde_ad_set_name" json:"last_cta_view_data_tilde_ad_set_name"`
	LastCtaViewDataTildeAdSetId                        string          `csv:"last_cta_view_data_tilde_ad_set_id" json:"last_cta_view_data_tilde_ad_set_id"`
	LastCtaViewDataTildeAdName                         string          `csv:"last_cta_view_data_tilde_ad_name" json:"last_cta_view_data_tilde_ad_name"`
	LastCtaViewDataTildeAdId                           string          `csv:"last_cta_view_data_tilde_ad_id" json:"last_cta_view_data_tilde_ad_id"`
	LastCtaViewDataTildeBranchAdFormat                 string          `csv:"last_cta_view_data_tilde_branch_ad_format" json:"last_cta_view_data_tilde_branch_ad_format"`
	LastCtaViewDataTildeTechnologyPartner              string          `csv:"last_cta_view_data_tilde_technology_partner" json:"last_cta_view_data_tilde_technology_partner"`
	LastCtaViewDataTildeBannerDimensions               string          `csv:"last_cta_view_data_tilde_banner_dimensions" json:"last_cta_view_data_tilde_banner_dimensions"`
	LastCtaViewDataTildePlacement                      string          `csv:"last_cta_view_data_tilde_placement" json:"last_cta_view_data_tilde_placement"`
	LastCtaViewDataTildeKeywordId                      string          `csv:"last_cta_view_data_tilde_keyword_id" json:"last_cta_view_data_tilde_keyword_id"`
	LastCtaViewDataTildeAgency                         string          `csv:"last_cta_view_data_tilde_agency" json:"last_cta_view_data_tilde_agency"`
	LastCtaViewDataTildeOptimizationModel              string          `csv:"last_cta_view_data_tilde_optimization_model" json:"last_cta_view_data_tilde_optimization_model"`
	LastCtaViewDataTildeSecondaryAdFormat              string          `csv:"last_cta_view_data_tilde_secondary_ad_format" json:"last_cta_view_data_tilde_secondary_ad_format"`
	LastCtaViewDataPlusViaFeatures                     string          `csv:"last_cta_view_data_plus_via_features" json:"last_cta_view_data_plus_via_features"`
	LastCtaViewDataDollar3P                            string          `csv:"last_cta_view_data_dollar_3p" json:"last_cta_view_data_dollar_3p"`
	LastCtaViewDataPlusWebFormat                       string          `csv:"last_cta_view_data_plus_web_format" json:"last_cta_view_data_plus_web_format"`
	LastCtaViewDataCustomFields                        string          `csv:"last_cta_view_data_custom_fields" json:"last_cta_view_data_custom_fields"`
	DeepLinked                                         CustomBoolean   `csv:"deep_linked" json:"deep_linked"`
	FirstEventForUser                                  CustomBoolean   `csv:"first_event_for_user" json:"first_event_for_user"`
	UserDataOs                                         string          `csv:"user_data_os" json:"user_data_os"`
	UserDataOsVersion                                  string          `csv:"user_data_os_version" json:"user_data_os_version"`
	UserDataModel                                      string          `csv:"user_data_model" json:"user_data_model"`
	UserDataBrowser                                    string          `csv:"user_data_browser" json:"user_data_browser"`
	UserDataGeoCountryCode                             string          `csv:"user_data_geo_country_code" json:"user_data_geo_country_code"`
	UserDataAppVersion                                 string          `csv:"user_data_app_version" json:"user_data_app_version"`
	UserDataSdkVersion                                 string          `csv:"user_data_sdk_version" json:"user_data_sdk_version"`
	UserDataGeoDmaCode                                 string          `csv:"user_data_geo_dma_code" json:"user_data_geo_dma_code"`
	UserDataEnvironment                                string          `csv:"user_data_environment" json:"user_data_environment"`
	UserDataPlatform                                   string          `csv:"user_data_platform" json:"user_data_platform"`
	UserDataAaid                                       string          `csv:"user_data_aaid" json:"user_data_aaid"`
	UserDataIdfa                                       string          `csv:"user_data_idfa" json:"user_data_idfa"`
	UserDataIdfv                                       string          `csv:"user_data_idfv" json:"user_data_idfv"`
	UserDataAndroidId                                  string          `csv:"user_data_android_id" json:"user_data_android_id"`
	UserDataLimitAdTracking                            string          `csv:"user_data_limit_ad_tracking" json:"user_data_limit_ad_tracking"`
	UserDataUserAgent                                  string          `csv:"user_data_user_agent" json:"user_data_user_agent"`
	UserDataIp                                         string          `csv:"user_data_ip" json:"user_data_ip"`
	UserDataDeveloperIdentity                          string          `csv:"user_data_developer_identity" json:"user_data_developer_identity"`
	UserDataLanguage                                   string          `csv:"user_data_language" json:"user_data_language"`
	UserDataBrand                                      string          `csv:"user_data_brand" json:"user_data_brand"`
	DiMatchClickToken                                  CustomInteger   `csv:"di_match_click_token" json:"di_match_click_token"`
	EventDataRevenueInUsd                              CustomFloat64   `csv:"event_data_revenue_in_usd" json:"event_data_revenue_in_usd"`
	EventDataExchangeRate                              CustomFloat64   `csv:"event_data_exchange_rate" json:"event_data_exchange_rate"`
	EventDataTransactionId                             string          `csv:"event_data_transaction_id" json:"event_data_transaction_id"`
	EventDataRevenue                                   CustomFloat64   `csv:"event_data_revenue" json:"event_data_revenue"`
	EventDataCurrency                                  string          `csv:"event_data_currency" json:"event_data_currency"`
	EventDataShipping                                  string          `csv:"event_data_shipping" json:"event_data_shipping"`
	EventDataTax                                       string          `csv:"event_data_tax" json:"event_data_tax"`
	EventDataCoupon                                    string          `csv:"event_data_coupon" json:"event_data_coupon"`
	EventDataAffiliation                               string          `csv:"event_data_affiliation" json:"event_data_affiliation"`
	EventDataSearchQuery                               string          `csv:"event_data_search_query" json:"event_data_search_query"`
	EventDataDescription                               string          `csv:"event_data_description" json:"event_data_description"`
	CustomData                                         string          `csv:"custom_data" json:"custom_data"`
	LastAttributedTouchDataTildeKeyword                string          `csv:"last_attributed_touch_data_tilde_keyword" json:"last_attributed_touch_data_tilde_keyword"`
	UserDataCrossPlatformId                            string          `csv:"user_data_cross_platform_id" json:"user_data_cross_platform_id"`
	UserDataPastCrossPlatformIds                       string          `csv:"user_data_past_cross_platform_ids" json:"user_data_past_cross_platform_ids"`
	UserDataProbCrossPlatformIds                       string          `csv:"user_data_prob_cross_platform_ids" json:"user_data_prob_cross_platform_ids"`
	StoreInstallBeginTimestamp                         string          `csv:"store_install_begin_timestamp" json:"store_install_begin_timestamp"`
	ReferrerClickTimestamp                             string          `csv:"referrer_click_timestamp" json:"referrer_click_timestamp"`
	UserDataOsVersionAndroid                           string          `csv:"user_data_os_version_android" json:"user_data_os_version_android"`
	UserDataGeoCityCode                                string          `csv:"user_data_geo_city_code" json:"user_data_geo_city_code"`
	UserDataGeoCityEn                                  string          `csv:"user_data_geo_city_en" json:"user_data_geo_city_en"`
	UserDataHttpReferrer                               string          `csv:"user_data_http_referrer" json:"user_data_http_referrer"`
	EventTimestamp                                     CustomInteger   `csv:"event_timestamp" json:"event_timestamp"`
	CustomerEventAlias                                 string          `csv:"customer_event_alias" json:"customer_event_alias"`
	LastAttributedTouchDataTildeCustomerCampaign       string          `csv:"last_attributed_touch_data_tilde_customer_campaign" json:"last_attributed_touch_data_tilde_customer_campaign"`
	LastAttributedTouchDataTildeCampaignType           string          `csv:"last_attributed_touch_data_tilde_campaign_type" json:"last_attributed_touch_data_tilde_campaign_type"`
	LastCtaViewDataTildeCampaignType                   string          `csv:"last_cta_view_data_tilde_campaign_type" json:"last_cta_view_data_tilde_campaign_type"`
	LastAttributedTouchDataTildeAgencyId               string          `csv:"last_attributed_touch_data_tilde_agency_id" json:"last_attributed_touch_data_tilde_agency_id"`
	LastAttributedTouchDataPlusTouchId                 string          `csv:"last_attributed_touch_data_plus_touch_id" json:"last_attributed_touch_data_plus_touch_id"`
	LastCtaViewDataPlusTouchId                         string          `csv:"last_cta_view_data_plus_touch_id" json:"last_cta_view_data_plus_touch_id"`
	UserDataInstallerPackageName                       string          `csv:"user_data_installer_package_name" json:"user_data_installer_package_name"`
	UserDataCpuType                                    string          `csv:"user_data_cpu_type" json:"user_data_cpu_type"`
	UserDataScreenWidth                                string          `csv:"user_data_screen_width" json:"user_data_screen_width"`
	UserDataScreenHeight                               string          `csv:"user_data_screen_height" json:"user_data_screen_height"`
	UserDataBuild                                      string          `csv:"user_data_build" json:"user_data_build"`
	UserDataInternetConnectionType                     string          `csv:"user_data_internet_connection_type" json:"user_data_internet_connection_type"`
	HashVersion                                        string          `csv:"hash_version" json:"hash_version"`
}

//ExportResource resource wrapper
type ExportResource struct {
	*ResourceAbstract
}

//GetEventOntology Get events ontology data
func (r *ExportResource) GetEventOntology(ctx context.Context, date time.Time) (*EventOntologyResponse, *http.Response, error) {
	post := make(map[string]interface{})
	post["export_date"] = date.Format("2006-01-02")
	post["branch_key"] = r.cfg.Key
	post["branch_secret"] = r.cfg.Secret
	rsp, err := r.tr.Post(ctx, "v3/export", post, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("ExportResource.GetEventOntology error: %v", err)
	}
	result := EventOntologyResponse{ResponseBody: &ResponseBody{}}
	result.status = rsp.StatusCode
	if result.IsSuccess() {
		var data EventOntology
		err = r.unmarshalResponse(rsp, &data)
		if err != nil {
			return &result, rsp, fmt.Errorf("ExportResource.GetEventOntology error: %v", err)
		}
		if !data.IsEmpty() {
			result.Data = &data
		}
	} else {
		err = r.unmarshalResponse(rsp, &result)
		if err != nil {
			return &result, rsp, fmt.Errorf("ExportResource.GetEventOntology error: %v", err)
		}
		return &result, rsp, fmt.Errorf(result.GetError())
	}
	return &result, rsp, nil
}

//GetEventData Get event data by link
func (r *ExportResource) GetEventData(ctx context.Context, link string) (*EventResponse, *http.Response, error) {
	rsp, err := r.tr.http.Get(link)
	if err != nil {
		return nil, nil, fmt.Errorf("ExportResource.GetEventData error: %v", err)
	}
	result := EventResponse{ResponseBody: &ResponseBody{}}
	result.status = rsp.StatusCode
	if result.IsSuccess() {
		events := []*Event{}
		err = r.unmarshalResponse(rsp, &events)
		if err != nil {
			return &result, rsp, fmt.Errorf("ExportResource.GetEventData error: %v", err)
		}
		result.Data = events
	} else {
		return &result, rsp, fmt.Errorf(result.GetError())
	}
	return &result, rsp, err
}
