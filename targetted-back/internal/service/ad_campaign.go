package service

import (
	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/repository"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
)

type AdCampaignService struct {
	repo repository.AdCampaign
}

var _ AdCampaign = (*AdCampaignService)(nil)

func NewAdCampaignService(repo repository.AdCampaign) *AdCampaignService {
	return &AdCampaignService{repo: repo}
}

func (s *AdCampaignService) Create(ac models.AdCampaign) (uuid.UUID, error) {
	return s.repo.Create(ac)
}

func (s *AdCampaignService) GetAll(userID uuid.UUID) ([]models.AdCampaign, error) {
	return s.repo.GetAll(userID)
}

func (s *AdCampaignService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *AdCampaignService) Update(ac models.AdCampaign, id string) (uuid.UUID, error) {
	return s.repo.Update(ac, id)
}

func (s *AdCampaignService) Start(id uuid.UUID) error {
	return s.repo.Start(id)
}

func (s *AdCampaignService) Stop(id uuid.UUID) error {
	return s.repo.Stop(id)
}

//TODO call image REPOS SERVICE
func (s *AdCampaignService) GetByID(userID uuid.UUID, campaignID string) (models.AdCampaign, error) {
	return s.repo.GetByID(campaignID)
}
