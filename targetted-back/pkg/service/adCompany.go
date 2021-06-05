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

func (s *AdCompanyService) GetCompanyList(userId uuid.UUID) ([]models.AdCompany, error) {
	return s.repo.GetCompanyList(userId)
}
