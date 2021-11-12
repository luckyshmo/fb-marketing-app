package marketing

// You can upload an image instead of using an image hash when you create an ad or ad creative.
// Add the image file to the multi-part MIME POST and specify the file name. For example:
//
// curl \
//   -F 'campaign_id=<AD_SET_ID>' \
//   -F 'creative={"title":"test title","body":"test","object_url":"http:\/\/www.test.com","image_file":"test.jpg"}' \
//   -F 'test.jpg=@test.jpg'
//   -F 'name=My ad' \
//   -F 'access_token=<ACCESS_TOKEN>' \
// "https://graph.facebook.com/<API_VERSION>/act_<ACCOUNT_ID>/ads"

type AdSet struct {
}

// curl \
//   -F 'name=My Ad Set' \
//   -F 'optimization_goal=REACH' \
//   -F 'billing_event=IMPRESSIONS' \
//   -F 'bid_amount=2' \
//   -F 'daily_budget=1000' \
//   -F 'campaign_id=<CAMPAIGN_ID>' \
//   -F 'targeting={"geo_locations": {"countries":["US"]}, "interests": [{id: 6003139266461, 'name': 'Movies'}]}' \
//   -F 'start_time=2020-10-06T04:45:17+0000' \
//   -F 'status=PAUSED' \
//   -F 'access_token=<ACCESS_TOKEN>' \
//   https://graph.facebook.com/v<API_VERSION>/act_<AD_ACCOUNT_ID>/adsets
func (as *AdSet) Create() {

}

// curl -X GET -G \
//   -d 'fields="cost_per_store_visit_action,store_visit_actions"' \
//   -d 'access_token=<ACCESS_TOKEN>' \
//   https://graph.facebook.com/v12.0/<AD_SET_ID>/insights
func (as *AdSet) Insight() {
	// https://developers.facebook.com/docs/marketing-api/reference/ad-campaign/insights/
}
