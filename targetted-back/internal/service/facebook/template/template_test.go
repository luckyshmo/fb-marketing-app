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
		Id:                     uuid.UUID{},
		UserId:                 uuid.UUID{},
		FbPageId:               "321",
		InstagramID:            "123",
		BusinessAddress:        "Ekb",
		Field:                  "Some field",
		Name:                   "Some name",
		Objective:              "Some purpose",
		CreativeStatus:         "??",
		ImagesDescription:      []string{"1", "2", "3"},
		ImagesSmallDescription: []string{"1", "2", "3"},
		PostDescription:        "",
		Budget:                 100.0,
		DailyBudget:            10.0,
		Days:                   10,
		IsStarted:              false,
		CreationDate:           time.Now(),
		StartDate:              time.Now(),
	})
	require.NoError(t, err)
	ac.print()
}
