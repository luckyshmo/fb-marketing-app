package marketing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// https://developers.facebook.com/docs/marketing-api/reference/ad-image

// curl \
//   -F 'filename=@<IMAGE_PATH>' \
//   -F 'access_token=<ACCESS_TOKEN>' \
//   https://graph.facebook.com/v2.11/act_<AD_ACCOUNT_ID>/adimages
//
// should return img hash
func (api *MarketingAPI) AddImage(filename string) (string, error) {

	type imageReq struct {
		FileName    string `json:"filename"`
		AccessToken string `json:"access_token"`
	}
	addRequest, err := json.Marshal(imageReq{
		FileName:    filename,
		AccessToken: api.credentials.Token,
	})
	if err != nil {
		return "", fmt.Errorf("marshall req: %w", err)
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"https://graph.facebook.com/%s/act_%s/adimages",
			api.credentials.ApiVersion,
			api.credentials.AdAccountID,
		),
		bytes.NewBuffer(addRequest),
	)
	if err != nil {
		return "", fmt.Errorf("add img request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := api.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("do add img request: %w", err)
	}
	defer resp.Body.Close()
	logrus.Warn(resp.StatusCode)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err //!
	}
	bodyString := string(bodyBytes)
	logrus.Info(bodyString)

	return "", nil //! should be hash

}

// curl \
//   -F 'bytes=iVBORw0KGgoAAAANSUhEUgAAABEAAAARCAMAAAAMs7fIAAAAOVBMVEX///87WZg7WZg7WZg7WZg7WZg7WZg7WZg7WZg7WZg7WZhMeMJEaa5Xi9tKdb0+Xp5Wi9tXjNxThNH+wk/7AAAACnRSTlMAsHIoaM7g/fx9Zr/g5QAAAGlJREFUeNplkFsOwCAIBPGJrtbX/Q/bqm1qwnxuJrBAE6OVD15pQy/WYePsDiIjp9FGyuC4DK7l6pOrVH4s41D6R4EzpJGXsa0MTQqp/yQo8hhHMuApoB1JQ5COnCN3yT6ys7xL3i7/cwMYsAveYa+MxAAAAABJRU5ErkJggg=='
//   -F 'access_token=<ACCESS_TOKEN>' \
// "https://graph.facebook.com/<API_VERSION>/act_<ACCOUNT_ID>/adimages"
// func (img *FacebookImage) AddBytes() {

// }

// curl -G \
//   -d 'access_token=<ACCESS_TOKEN>' \
//   https://graph.facebook.com/v2.11/act_<AD_ACCOUNT_ID>/adimages
// func (img *FacebookImage) GetByHash() {

// }

// curl -G \
//   -d 'hashes=["<IMAGE_1_HASH>","<IMAGE_2_HASH>"]' \
//   -d 'access_token=<ACCESS_TOKEN>' \
//   https://graph.facebook.com/v2.11/act_<AD_ACCOUNT_ID>/adimages
// func (img *FacebookImage) GetByAddAccount() {

// }

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
