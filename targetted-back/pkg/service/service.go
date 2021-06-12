package service

import (
	"github.com/google/uuid"
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
}

type AdCompany interface {
	Create(ac models.AdCompany) (uuid.UUID, error)
	Delete(ID string) error
	Update(ac models.AdCompany, ID string) (uuid.UUID, error)
	GetAll(userID uuid.UUID) ([]models.AdCompany, error)
	GetByID(userID uuid.UUID, companyID string) (models.AdCompany, error)
}

type Service struct {
	Authorization
	User
	AdCompany
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
		AdCompany:     NewAdCompanyService(repos.AdCompany),
	}
}
