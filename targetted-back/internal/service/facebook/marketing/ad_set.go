package marketing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// https://developers.facebook.com/docs/marketing-api/reference/ad-campaign

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
func (api *MarketingAPI) CreateAdSet(
	Name string,
	OptimizationGoal string,
	BillingEvent string,
	BidAmount string,
	DailyBudget string,
	CampaignID string,
	Targeting string,
	StartTime string,
	Status string,

	PageID string,
) (string, error) {

	type adSetReq struct {
		Name             string `json:"name"`
		OptimizationGoal string `json:"optimization_goal"`
		BillingEvent     string `json:"billing_event"`
		BidAmount        string `json:"bid_amount"`
		DailyBudget      string `json:"daily_budget"`
		CampaignID       string `json:"campaign_id"`
		Targeting        string `json:"targeting"`
		StartTime        string `json:"start_time"`
		Status           string `json:"status"`
		AccessToken      string `json:"access_token"`

		PromotedObject string `json:"promoted_object"`
	}
	createAdSetRequest, err := json.Marshal(adSetReq{
		Name:             Name,
		OptimizationGoal: OptimizationGoal,
		BillingEvent:     BillingEvent,
		BidAmount:        BidAmount,
		DailyBudget:      DailyBudget,
		CampaignID:       CampaignID,
		Targeting:        Targeting,
		StartTime:        StartTime,
		Status:           Status,
		PromotedObject: fmt.Sprintf(`{
			"page_id": "%s"
			}`, PageID),
		AccessToken: api.credentials.Token,
	})
	if err != nil {
		return "", fmt.Errorf("marshall req: %w", err)
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"https://graph.facebook.com/%s/act_%s/adsets?locale=en_EN",
			api.credentials.ApiVersion,
			api.credentials.AdAccountID,
		),
		bytes.NewBuffer(createAdSetRequest),
	)
	if err != nil {
		return "", fmt.Errorf("create ad set request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := api.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("do create ad set request: %w", err)
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

// curl -X GET -G \
//   -d 'fields="cost_per_store_visit_action,store_visit_actions"' \
//   -d 'access_token=<ACCESS_TOKEN>' \
//   https://graph.facebook.com/v12.0/<AD_SET_ID>/insights
// func (as *AdSet) Insight() {
// https://developers.facebook.com/docs/marketing-api/reference/ad-campaign/insights/
// }
