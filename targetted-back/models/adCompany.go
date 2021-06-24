package models

import (
	"time"

	"github.com/google/uuid"
)

type AdCompany struct {
	Id                     uuid.UUID `db:"id"`
	UserId                 uuid.UUID `db:"user_id"`
	FbPageId               string    `db:"fb_page_id"`
	BusinessAddress        string    `db:"business_address"` //TODO RENAME
	CompanyField           string    `db:"field"`
	CompanyName            string    `db:"name"`
	CompnayPurpose         string    `db:"purpose"`
	CreativeStatus         string    `db:"creative_status"`
	ImagesDescription      []string  `db:"images_description"`
	ImagesSmallDescription []string  `db:"images_small_description"`
	PostDescription        string    `db:"post_description"`
	CurrentAmount          int       `db:"current_amount"`
	DailyAmount            int       `db:"daily_amount"`
	Days                   int       `db:"days"`
	IsStarted              bool      `db:"is_started"`
	CreationDate           time.Time `db:"date_created"`
	StartDate              time.Time `db:"date_started"`
}
