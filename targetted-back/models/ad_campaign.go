package models

import (
	"time"

	"github.com/google/uuid"
)

// AdCampaign represent user ad campaign in our service
type AdCampaign struct {
	Id                     uuid.UUID `db:"id"`
	UserId                 uuid.UUID `db:"user_id"`
	FbPageId               string    `db:"fb_page_id"`
	BusinessAddress        string    `db:"business_address"` //TODO RENAME
	CampaignField          string    `db:"field"`
	CampaignName           string    `db:"name"`
	CompnayPurpose         string    `db:"purpose"`
	CreativeStatus         string    `db:"creative_status"`
	ImagesDescription      []string  `db:"images_description"`
	ImagesSmallDescription []string  `db:"images_small_description"`
	PostDescription        string    `db:"post_description"`
	CurrentAmount          float64   `db:"current_amount"`
	DailyAmount            int       `db:"daily_amount"`
	Days                   int       `db:"days"`
	IsStarted              bool      `db:"is_started"`
	CreationDate           time.Time `db:"date_created"`
	StartDate              time.Time `db:"date_started"`
}
