package inmemory

import (
	"errors"

	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
	"github.com/sirupsen/logrus"
)

var errCampaignNotFound = errors.New("campaign not found")

type AdCampaignMemory struct {
}

func NewAdCampaignMemory() *AdCampaignMemory {
	return &AdCampaignMemory{}
}

func (r *AdCampaignMemory) Create(ac models.AdCampaign) (uuid.UUID, error) {
	db.Lock()
	defer db.Unlock()

	id := uuid.New()
	ac.Id = id
	db.AdCampaigns = append(db.AdCampaigns, ac)

	return id, nil
}

func (r *AdCampaignMemory) Delete(id string) error {
	db.Lock()
	defer db.Unlock()

	return errors.New("not implemented")
}

func (r *AdCampaignMemory) Update(ac models.AdCampaign, idStr string) (uuid.UUID, error) {
	db.Lock()
	defer db.Unlock()

	return uuid.Nil, errors.New("not implemented")
}

func (r *AdCampaignMemory) changeStatus(id uuid.UUID, status bool) error {
	db.Lock()
	defer db.Unlock()

	return errors.New("not implemented")
}

func (r *AdCampaignMemory) Start(id uuid.UUID) error {
	return r.changeStatus(id, true)
}

func (r *AdCampaignMemory) Stop(id uuid.UUID) error {
	return r.changeStatus(id, false)
}

func (r *AdCampaignMemory) GetAll(userId uuid.UUID) ([]models.AdCampaign, error) {
	db.Lock()
	defer db.Unlock()

	var userComp []models.AdCampaign

	logrus.Warn(userId)
	logrus.Warn(db.Users)
	logrus.Warn(db.AdCampaigns)

	for _, ac := range db.AdCampaigns {
		logrus.Errorf("%s \n %s", ac.UserId.String(), userId.String())
		if ac.UserId.String() == userId.String() {
			userComp = append(userComp, ac)
		}
	}
	if len(userComp) == 0 {
		return nil, errCampaignNotFound
	}
	return userComp, nil
}

func (r *AdCampaignMemory) GetByID(campaignID string) (models.AdCampaign, error) {
	db.Lock()
	defer db.Unlock()
	for _, ac := range db.AdCampaigns {
		if ac.Id.String() == campaignID {
			return ac, nil
		}
	}

	return models.AdCampaign{}, errCampaignNotFound
}
