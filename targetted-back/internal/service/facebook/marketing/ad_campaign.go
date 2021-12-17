package marketing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// https://developers.facebook.com/docs/marketing-api/reference/ad-campaign-group

// curl -X POST \
//   -F 'name="My campaign"' \
//   -F 'objective="LINK_CLICKS"' \
//   -F 'status="PAUSED"' \
//   -F 'special_ad_categories=[]' \
//   -F 'access_token=<ACCESS_TOKEN>' \
//   https://graph.facebook.com/v12.0/act_<AD_ACCOUNT_ID>/campaigns
func (api *MarketingAPI) CreateCampaign(Name, Objective, Status, SpecialAdCategories string) (string, error) {
	type campaignRequest struct {
		Name                string `json:"name"`
		Objective           string `json:"objective"`
		Status              string `json:"status"`
		SpecialAdCategories string `json:"special_ad_categories"`
		AccessToken         string `json:"access_token"`
	}
	creationRequest, err := json.Marshal(campaignRequest{
		Name:                Name,
		Objective:           Objective,
		Status:              Status,
		SpecialAdCategories: "[]", // В первой итерации без них будем
		AccessToken:         api.credentials.Token,
	})
	if err != nil {
		return "", fmt.Errorf("marshall req: %w", err)
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"https://graph.facebook.com/%s/act_%s/campaigns",
			api.credentials.ApiVersion,
			api.credentials.AdAccountID,
		),
		bytes.NewBuffer(creationRequest),
	)
	if err != nil {
		return "", fmt.Errorf("new campaign creation request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := api.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("do create campaign request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var pResp facebookID
		if err := json.NewDecoder(resp.Body).Decode(&pResp); err != nil {
			return "", fmt.Errorf("decode body: %w", err)
		}
		return pResp.ID, nil
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err //!
	}
	bodyString := string(bodyBytes)
	logrus.Info(bodyString)

	//msg="{\"error\":{\"message\":\"(#100) The parameter special_ad_categories is required.\",\"type\":\"OAuthException\",\"code\":100,\"fbtrace_id\":\"AEk1ktrFhSnMu8f9G3UaD47\"}}"

	return "", errFacebook
}
