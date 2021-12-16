package marketing

import (
	"errors"
	"net/http"
	"time"

	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/service/facebook/models"
	"github.com/sirupsen/logrus"
)

var errFacebook = errors.New("facebook error")

type API interface {
	//! all needed methods
}

type MarketingAPI struct {
	credentials models.Credentials
	client      http.Client
}

func NewMarketingAPI(cr models.Credentials) *MarketingAPI {
	return &MarketingAPI{
		credentials: cr,
		client: http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

type facebookID struct {
	ID string `json="id"`
}

func (api *MarketingAPI) StartCampaign(pageID string) {
	// steps from https://developers.facebook.com/docs/marketing-apis/get-started

	campaignID, err := api.CreateCampaign("TestCamp", "PAGE_LIKES", "PAUSED", "[]")
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info("campaignID=", campaignID)

	adSetID, err := api.CreateAdSet(
		"Test ad set",
		"REACH",
		"IMPRESSIONS",
		"3000",
		"30000",
		campaignID,
		`{"geo_locations": {"countries":["US"]}, "interests": [{id: 6003139266461, 'name': 'Movies'}]}`,
		"2020-10-06T04:45:17+0000",
		"PAUSED",
		pageID,
	)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info("adSetID=", adSetID)

	imgHash := "c4981a4aa76f8e5e8828b14d358fd96a" //!

	creativeID, err := api.CreateCreative(
		"Test creative",
		pageID,
		imgHash,
	)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info("creativeID=", creativeID)

	adID, err := api.CreateAd(
		"Test ad",
		adSetID,
		creativeID,
		"PAUSED",
	)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info(adID)
}
