package img

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/service/facebook/models"
	"github.com/sirupsen/logrus"
)

type ImgInfo struct {
	Hash string
	Url  string
}

type ImageStorage interface {
	Upload(img []byte) (hash, url string, err error)
	Delete(hash string) error
}

type FBImgStorage struct {
	credentials models.Credentials
	client      http.Client
}

func NewFBImgStorage(cr models.Credentials) *FBImgStorage {
	return &FBImgStorage{
		credentials: cr,
		client: http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

// https://developers.facebook.com/docs/marketing-api/reference/ad-image

// curl \
//   -F 'bytes=<img_bytes>'
//   -F 'access_token=<ACCESS_TOKEN>' \
// "https://graph.facebook.com/<API_VERSION>/act_<ACCOUNT_ID>/adimages"
// demoImg - []byte("iVBORw0KGgoAAAANSUhEUgAAABEAAAARCAMAAAAMs7fIAAAAOVBMVEX///87WZg7WZg7WZg7WZg7WZg7WZg7WZg7WZg7WZg7WZhMeMJEaa5Xi9tKdb0+Xp5Wi9tXjNxThNH+wk/7AAAACnRSTlMAsHIoaM7g/fx9Zr/g5QAAAGlJREFUeNplkFsOwCAIBPGJrtbX/Q/bqm1qwnxuJrBAE6OVD15pQy/WYePsDiIjp9FGyuC4DK7l6pOrVH4s41D6R4EzpJGXsa0MTQqp/yQo8hhHMuApoB1JQ5COnCN3yT6ys7xL3i7/cwMYsAveYa+MxAAAAABJRU5ErkJggg==")
func (storage *FBImgStorage) Upload(img []byte) (hash, url string, err error) {

	type imgResp struct {
		Images struct {
			Bytes struct {
				Hash string `json:"hash"`
				URL  string `json:"url"`
			} `json:"bytes"`
		} `json:"images"`
	}

	type imageReq struct {
		Bytes       string `json:"bytes"`
		AccessToken string `json:"access_token"`
	}

	addRequest, err := json.Marshal(imageReq{
		Bytes:       string(img), //! check
		AccessToken: storage.credentials.Token,
	})
	if err != nil {
		err = fmt.Errorf("marshall req: %w", err)
		return
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"https://graph.facebook.com/%s/act_%s/adimages",
			storage.credentials.ApiVersion,
			storage.credentials.AdAccountID,
		),
		bytes.NewBuffer(addRequest),
	)
	if err != nil {
		err = fmt.Errorf("add img request: %w", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := storage.client.Do(req)
	if err != nil {
		err = fmt.Errorf("do add img request: %w", err)
		return
	}
	defer resp.Body.Close()

	var imgR imgResp
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("read resp body: %w", err)
		return
	}
	json.Unmarshal(bodyBytes, &imgR)

	return imgR.Images.Bytes.Hash, imgR.Images.Bytes.URL, nil
}

func (storage *FBImgStorage) Delete(hash string) error {
	logrus.Error("Call not implemented method")
	return errors.New("call not implemented method")
}

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
