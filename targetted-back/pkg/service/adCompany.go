package service

import (
	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/pkg/repository"
)

type AdCompanyService struct {
	repo repository.AdCompany
}

var _ AdCompany = (*AdCompanyService)(nil)

func NewAdCompanyService(repo repository.AdCompany) *AdCompanyService {
	return &AdCompanyService{repo: repo}
}

func (s *AdCompanyService) Create(ac models.AdCompany) (uuid.UUID, error) {
	return s.repo.Create(ac)
}

func (s *AdCompanyService) GetAll(userID uuid.UUID) ([]models.AdCompany, error) {
	return s.repo.GetAll(userID)
}

func (s *AdCompanyService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *AdCompanyService) Update(ac models.AdCompany, id string) (uuid.UUID, error) {
	return s.repo.Update(ac, id)
}

func (s *AdCompanyService) Start(id uuid.UUID) error {
	return s.repo.Start(id)
}

func (s *AdCompanyService) Stop(id uuid.UUID) error {
	return s.repo.Stop(id)
}

//TODO call image REPOS SERVICE
func (s *AdCompanyService) GetByID(userID uuid.UUID, companyID string) (models.AdCompany, error) {
	return s.repo.GetByID(companyID)
}
