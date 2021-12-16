package marketing

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// https://developers.facebook.com/docs/marketing-api/insights

// curl -G \
// -d "fields=actions,cpc,ctr,spend,cost_per_action_type,action_values,location" \
// -d "date_preset=this_year" \
// -d "access_token=EAALCaNdmu2oBADTIXWMgsLwzY2q5nm7lbbCWNHipxW2kC2eIntIGbHG7UNAs7p0QjADhI78gLf8EQXaGZA9KahXQsMdMLUVpMP97cLRDpOWH5vjMVJNtmccjF1EkjRlNzXcdCQFvJd9bMdvwppsNeZAAFzmeHfHMHOcIB9JsFnW0gqtiur" \
// "https://graph.facebook.com/v12.0/23846909377100457/insights"
//
//
func (api *MarketingAPI) adObjectInsight(objectID string) (string, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"https://graph.facebook.com/%s/%s/insights?fields=%s&access_token=%s&date_preset=%s",
			api.credentials.ApiVersion,
			objectID,
			"actions,cpc,ctr,spend,cost_per_action_type,action_values,location",
			api.credentials.Token,
			"this_year",
		),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("create inst request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := api.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("do inst request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err //!
	}
	bodyString := string(bodyBytes)
	logrus.Info(bodyString)

	return "", nil //!
}
