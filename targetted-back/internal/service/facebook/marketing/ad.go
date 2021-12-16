package marketing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// у нас ad и ad_set почти всегда будет 1 к 1

// https://developers.facebook.com/docs/marketing-api/reference/adgroup

// curl -X POST \
//   -F 'name="My Ad"' \
//   -F 'adset_id="<AD_SET_ID>"' \
//   -F 'creative={
//        "creative_id": "<CREATIVE_ID>"
//      }' \
//   -F 'status="PAUSED"' \
//   -F 'access_token=<ACCESS_TOKEN>' \
//   https://graph.facebook.com/v12.0/act_<AD_ACCOUNT_ID>/ads
func (api *MarketingAPI) CreateAd(Name, AdSetID, CreativeID, Status string) (string, error) {
	type createAd struct {
		Name        string `json:"name"`
		AdSetID     string `json:"adset_id"`
		Creative    string `json:"creative"`
		Status      string `json:"status"`
		AccessToken string `json:"access_token"`
	}
	createAdReq, err := json.Marshal(createAd{
		Name:    Name,
		AdSetID: AdSetID,
		Creative: fmt.Sprintf(`{
			"creative_id": "%s"
		  }`, CreativeID),
		Status:      Status,
		AccessToken: api.credentials.Token,
	})
	if err != nil {
		return "", fmt.Errorf("marshall req: %w", err)
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"https://graph.facebook.com/%s/act_%s/ads",
			api.credentials.ApiVersion,
			api.credentials.AdAccountID,
		),
		bytes.NewBuffer(createAdReq),
	)
	if err != nil {
		return "", fmt.Errorf("create ad request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := api.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("do create ad request: %w", err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err //!
	}
	bodyString := string(bodyBytes)
	logrus.Info(bodyString)

	return "", nil //! should be ?
}
