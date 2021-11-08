package marketing

// https://developers.facebook.com/docs/marketing-api/reference/ad-creative

type Creative struct {
}

// curl -X POST \
//   -F 'name="Sample Creative"' \
//   -F 'object_story_spec={
//        "page_id": "<PAGE_ID>",
//        "link_data": {
//          "image_hash": "<IMAGE_HASH>",
//          "link": "https://facebook.com/<PAGE_ID>",
//          "message": "try it out"
//        }
//      }' \
//   -F 'access_token=<ACCESS_TOKEN>' \
//   https://graph.facebook.com/v12.0/act_<AD_ACCOUNT_ID>/adcreatives
func (c *Creative) Create() {

}
