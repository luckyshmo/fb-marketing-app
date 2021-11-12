package marketing

// https://developers.facebook.com/docs/marketing-api/reference/ad-campaign

// type AdCompany interface {
// 	Create()
// }

type AdCompany struct {
}

// curl -X POST \
//   -F 'name="My campaign"' \
//   -F 'objective="LINK_CLICKS"' \
//   -F 'status="PAUSED"' \
//   -F 'special_ad_categories=[]' \
//   -F 'access_token=<ACCESS_TOKEN>' \
//   https://graph.facebook.com/v12.0/act_<AD_ACCOUNT_ID>/campaigns
func (ac *AdCompany) Create() {

}
