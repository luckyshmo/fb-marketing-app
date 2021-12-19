package template

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
	"github.com/stretchr/testify/require"
)

func Test_TemplateMerge(t *testing.T) {
	ac, err := NewAdCampaign(models.AdCampaign{
		Id:              uuid.UUID{},
		UserId:          uuid.UUID{},
		FbPageId:        "321",
		InstagramID:     "123",
		BusinessAddress: "Ekb",
		Field:           "Some field",
		Name:            "Some name",
		Objective:       "Some objective",
		CreativeStatus:  "??",
		DailyBudget:     10.0,
		Duration:        10,
		IsStarted:       false,
		TimeCreated:     time.Now(),
		TimeStarted:     time.Now(),

		Images: models.Images{
			Img: []models.Image{
				{
					Hash:    "123123123",
					URL:     "some-url.com/img123",
					Message: "Some story",
					IsStory: true,
				},
				{
					Hash:    "456456456",
					URL:     "some-url.com/img456",
					Message: "Some post",
					IsStory: false,
				},
			},
			PostDescription: "I'M POST DESC WITH EMOJI 😀😃😄😁😆😅😂🤣",
		},
	})
	require.NoError(t, err)
	ac.print()
}
