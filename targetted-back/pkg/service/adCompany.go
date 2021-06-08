package service

import (
	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/pkg/repository"
)

type AdCompanyService struct {
	repo repository.AdCompany
}

func NewAdCompanyService(repo repository.AdCompany) *AdCompanyService {
	return &AdCompanyService{repo: repo}
}

func (s *AdCompanyService) CreateCompany(ac models.AdCompany) (uuid.UUID, error) {
	return s.repo.CreateCompany(ac)
}

func (s *AdCompanyService) GetCompanyList(userID uuid.UUID) ([]models.AdCompany, error) {
	return s.repo.GetCompanyList(userID)
}

//TODO call image REPOS SERVICE
func (s *AdCompanyService) GetCompanyByID(userID uuid.UUID, companyID string) (models.AdCompany, error) {
	return s.repo.GetCompanyByID(companyID)
}
