package marketing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// https://developers.facebook.com/docs/marketing-api/reference/ad-creative

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
func (api *MarketingAPI) CreateCreative(name, pageID, hash string) (string, error) {
	type campaignRequest struct {
		Name            string `json:"name"`
		ObjectStorySpec string `json:"object_story_spec"`
		AccessToken     string `json:"access_token"`
	}
	creationRequest, err := json.Marshal(campaignRequest{
		Name: name,
		ObjectStorySpec: fmt.Sprintf(`{
			"page_id": "%s",
			"link_data": {
			  "image_hash": "%s",
			  "link": "https://facebook.com/%s",
			  "message": "try it out"
			}
		  }`, pageID, hash, pageID), //! message ??
		AccessToken: api.credentials.Token,
	})
	if err != nil {
		return "", fmt.Errorf("marshall req: %w", err)
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"https://graph.facebook.com/%s/act_%s/adcreatives",
			api.credentials.ApiVersion,
			api.credentials.AdAccountID,
		),
		bytes.NewBuffer(creationRequest),
	)
	if err != nil {
		return "", fmt.Errorf("create creative request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := api.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("do create creative request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var pResp facebookID
		if err := json.NewDecoder(resp.Body).Decode(&pResp); err != nil {
			return "", fmt.Errorf("decode body: %w", err)
		}
		return pResp.ID, nil
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err //!
	}
	bodyString := string(bodyBytes)
	logrus.Info(bodyString)

	return "", errFacebook
}
