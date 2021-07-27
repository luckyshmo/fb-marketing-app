package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/pkg/repository"
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

type Facebook interface {
	PageClaim(ID string) (string, error)
	GetOwnedPages() ([]models.FacebookPage, error)
	GetPendingPagesID() ([]string, error)
	DeletePageByID(ID string) error
	CheckPageLimitTicker(ctx context.Context)
	IsPageOwnedByID(ID string) (bool, error)
}

type Service struct {
	Authorization Authorization
	User          User
	AdCompany     AdCompany
	Facebook      Facebook
}

func NewService(repos *repository.Repository, cfg config.Facebook) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
		AdCompany:     NewAdCompanyService(repos.AdCompany),
		Facebook:      NewFacebookService(cfg),
	}
}
