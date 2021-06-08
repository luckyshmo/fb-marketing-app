package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/pkg/repository/pg"
)

type Authorization interface {
	CreateUser(user models.User) (uuid.UUID, error)
	GetUser(username, passwordHash string) (models.User, error)
}

type User interface {
	GetById(userId uuid.UUID) (models.User, error)
	GetAll() ([]models.User, error)
}

type AdCompany interface {
	CreateCompany(ac models.AdCompany) (uuid.UUID, error)
	GetCompanyList(userID uuid.UUID) ([]models.AdCompany, error)
	GetCompanyByID(companyID string) (models.AdCompany, error)
}

type Repository struct {
	Authorization
	User
	AdCompany
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: pg.NewAuthPostgres(db),
		User:          pg.NewUserPG(db),
		AdCompany:     pg.NewAdCompanyPg(db),
	}
}
