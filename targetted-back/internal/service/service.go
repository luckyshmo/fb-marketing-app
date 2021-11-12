package service

import (
	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/repository"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/service/facebook"
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

type AdCompany interface {
	Create(ac models.AdCompany) (uuid.UUID, error)
	Delete(ID string) error
	Start(id uuid.UUID) error
	Stop(id uuid.UUID) error
	Update(ac models.AdCompany, ID string) (uuid.UUID, error)
	GetAll(userID uuid.UUID) ([]models.AdCompany, error)
	GetByID(userID uuid.UUID, companyID string) (models.AdCompany, error)
}

type Service struct {
	Authorization Authorization
	User          User
	AdCompany     AdCompany
	Facebook      facebook.Service
	Payment       *payment.Youkassa
}

func NewService(repos *repository.Repository, cfg *config.Config) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
		AdCompany:     NewAdCompanyService(repos.AdCompany),
		Facebook:      facebook.NewFacebookService(cfg.Facebook),
		Payment:       payment.NewYoukassaService(cfg.Youkassa),
	}
}
