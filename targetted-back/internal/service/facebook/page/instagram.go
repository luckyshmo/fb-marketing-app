package page

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// curl -G \
// -d "access_token=<ACCESS_TOKEN>"\
// -d "fields=id,username,profile_pic" \
// "https://graph.facebook.com/<API_VERSION>/<PAGE_ID>/instagram_accounts"
//
// это способ получения инсты из интерфейса с токеном пользователя
func (page *PageManager) getInstagramIDByPageID(pageID string) (instagramID string, err error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"https://graph.facebook.com/%s/%s/instagram_accounts?access_token=%s&fields=%s",
			page.Credentials.ApiVersion,
			pageID,
			page.Credentials.Token,
			`id,username,profile_pic`,
		),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("create inst request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := page.client.Do(req)
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
