package service

import (
	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/repository"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/service/facebook/api"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/service/facebook/img"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/service/payment"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
)

type Authorization interface {
	CreateUser(user models.User) (uuid.UUID, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (uuid.UUID, error)
}

type User interface {
	GetById(userId uuid.UUID) (models.User, error)
	GetAll() ([]models.User, error)
	AddMoney(userId uuid.UUID, amount float64) error
	SetBalance(id uuid.UUID, amount float64) error
}

type AdCampaign interface {
	Create(ac models.AdCampaign, imgWithInfo []ImageWithMessage) (uuid.UUID, error)
	Delete(ID string) error
	Start(id uuid.UUID) error
	Stop(id uuid.UUID) error
	Update(ac models.AdCampaign, ID string, imgWithInfo []ImageWithMessage) (uuid.UUID, error)
	GetAll(userID uuid.UUID) ([]models.AdCampaign, error)
	GetByID(userID uuid.UUID, campaignID string) (models.AdCampaign, error)
}

type Service struct {
	Authorization Authorization
	User          User
	AdCampaign    AdCampaign
	FacebookAPI   api.Facebook
	Payment       *payment.Youkassa
}

func NewService(repos *repository.Repository, storage img.ImageStorage, cfg *config.Config) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
		AdCampaign:    NewAdCampaignService(repos.AdCampaign, storage),
		FacebookAPI:   api.NewFacebook(cfg.Facebook),
		Payment:       payment.NewYoukassaService(cfg.Youkassa),
	}
}
