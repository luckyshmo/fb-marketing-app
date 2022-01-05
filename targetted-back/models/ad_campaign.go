package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
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
	Field          string    `db:"field"`
	Name           string    `db:"name"`
	AdName         string    `db:"ad_name"`
	Objective      string    `db:"objective"`
	CreativeStatus string    `db:"creative_status"`
	DailyBudget    float64   `db:"daily_budget"`
	Duration       int       `db:"duration"`
	IsStarted      bool      `db:"is_started"`
	TimeCreated    time.Time `db:"time_created"`
	TimeStarted    time.Time `db:"time_started"`

	Images Images `db:"images"`
}

type Images struct {
	Img             []Image `json:"images"`
	PostDescription string  `json:"post_description"`
}

type Image struct {
	Hash    string `json:"hash"`
	URL     string `json:"url"`
	Message string `json:"message"`
	IsStory bool   `json:"is_story"`
}

// Make the Image struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (img Images) Value() (driver.Value, error) {
	return json.Marshal(img)
}

// Make the Image struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (img *Images) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed") //!
	}

	return json.Unmarshal(b, &img)
}
