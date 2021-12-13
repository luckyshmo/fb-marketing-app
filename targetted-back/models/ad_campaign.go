package models

import (
	"time"

	"github.com/google/uuid"
)

// AdCampaign represent user ad campaign in our service
type AdCampaign struct {
	Id          uuid.UUID `db:"id"`
	UserId      uuid.UUID `db:"user_id"`
	FbPageId    string    `db:"fb_page_id"`
	InstagramID string    `db:"instagram_id"`
	//! rename
	BusinessAddress string `db:"business_address"`
	//! rename
	Field           string    `db:"field"`
	Name            string    `db:"name"`
	AdName          string    `db:"ad_name"`
	Objective       string    `db:"objective"`
	CreativeStatus  string    `db:"creative_status"`
	PostDescription string    `db:"post_description"`
	DailyBudget     float64   `db:"daily_budget"`
	Duration        int       `db:"duration"`
	IsStarted       bool      `db:"is_started"`
	TimeCreated     time.Time `db:"time_created"`
	TimeStarted     time.Time `db:"time_started"`
}
