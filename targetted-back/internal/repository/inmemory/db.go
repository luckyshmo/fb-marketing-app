package inmemory

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
)

type dbScheme struct {
	sync.Mutex
	Users       []models.User
	AdCampaigns []models.AdCampaign
}

var db dbScheme

func Init() {

	usrID, _ := uuid.FromBytes([]byte("8db1eadc-ff22-4b87-bfef-049e913d6122"))
	campaignID, _ := uuid.FromBytes([]byte("95f9bd3a-95a8-460f-a737-f7a946df42f7"))

	db = dbScheme{
		Users: []models.User{
			{
				Id:          usrID,
				Name:        "Admin",
				Email:       "admin@admin.com",
				Password:    "686a7172686a7177313234363137616a6668616a73601f1889667efaebb33b8c12572835da3f027f78",
				PhoneNumber: "+79995600000",
				Balance:     0.0,
				TimeCreated: time.Now(),
			},
		},
		AdCampaigns: []models.AdCampaign{
			{
				Id:              campaignID,
				UserId:          usrID,
				FbPageId:        "321",
				InstagramID:     "123",
				BusinessAddress: "Ekb",
				Field:           "Some field",
				Name:            "Some name",
				Objective:       "Some objective",
				CreativeStatus:  "??",
				PostDescription: "",
				Budget:          100.0,
				DailyBudget:     10.0,
				Duration:        10,
				IsStarted:       false,
				TimeCreated:     time.Now(),
				TimeStarted:     time.Now(),
			},
		},
	}
}
