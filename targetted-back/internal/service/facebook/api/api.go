package api

import (
	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/service/facebook/marketing"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/service/facebook/models"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/service/facebook/page"
)

type Facebook struct {
	PageManager page.API
	Marketing   marketing.API
}

func NewFacebook(cfg config.Facebook) Facebook {
	cr := models.Credentials{
		Token:       cfg.Token,
		ApiVersion:  cfg.APIVersion,
		BusinessID:  cfg.BusinessID,
		AdAccountID: cfg.AdAccountID,
	}
	return Facebook{
		PageManager: page.NewPageManager(cr),
		Marketing:   marketing.NewMarketingAPI(cr),
	}
}
