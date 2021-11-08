package marketing

type Ad struct {
}

// curl -X POST \
//   -F 'name="My Ad"' \
//   -F 'adset_id="<AD_SET_ID>"' \
//   -F 'creative={
//        "creative_id": "<CREATIVE_ID>"
//      }' \
//   -F 'status="PAUSED"' \
//   -F 'access_token=<ACCESS_TOKEN>' \
//   https://graph.facebook.com/v12.0/act_<AD_ACCOUNT_ID>/ads
func (ad *Ad) Create() {

}
