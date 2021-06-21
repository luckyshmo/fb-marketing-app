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
	AddMoney(userId uuid.UUID, amount float64) error
}

type AdCompany interface {
	Create(ac models.AdCompany) (uuid.UUID, error)
	Delete(idStr string) error
	Update(ac models.AdCompany, idStr string) (uuid.UUID, error)
	GetAll(userID uuid.UUID) ([]models.AdCompany, error)
	GetByID(companyID string) (models.AdCompany, error)
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
