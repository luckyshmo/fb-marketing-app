package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/repository/inmemory"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/repository/pg"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
)

type Authorization interface {
	CreateUser(user models.User) (uuid.UUID, error)
	GetUser(username, passwordHash string) (models.User, error)
}

type User interface {
	GetById(userId uuid.UUID) (models.User, error)
	GetAll() ([]models.User, error)
	AddMoney(userId uuid.UUID, amount float64) error
	SetBalance(userId uuid.UUID, amount float64) error
}

type AdCompany interface {
	Create(ac models.AdCompany) (uuid.UUID, error)
	Delete(idStr string) error
	Update(ac models.AdCompany, idStr string) (uuid.UUID, error)
	Start(id uuid.UUID) error
	Stop(id uuid.UUID) error
	GetAll(userID uuid.UUID) ([]models.AdCompany, error)
	GetByID(companyID string) (models.AdCompany, error) //TODO ID -> uuid
}

type Repository struct {
	Authorization Authorization
	User          User
	AdCompany     AdCompany
}

func NewSqlxRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: pg.NewAuthPostgres(db),
		User:          pg.NewUserPG(db),
		AdCompany:     pg.NewAdCompanyPg(db),
	}
}

func NewInMemoryRepository() *Repository {
	inmemory.Init()
	return &Repository{
		Authorization: inmemory.NewAuthMemory(),
		User:          inmemory.NewUserMemory(),
		AdCompany:     inmemory.NewAdCompanyMemory(),
	}
}
