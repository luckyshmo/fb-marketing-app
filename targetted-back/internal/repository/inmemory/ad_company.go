package inmemory

import (
	"errors"

	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
	"github.com/sirupsen/logrus"
)

var errCompanyNotFound = errors.New("company not found")

type AdCompanyMemory struct {
}

func NewAdCompanyMemory() *AdCompanyMemory {
	return &AdCompanyMemory{}
}

func (r *AdCompanyMemory) Create(ac models.AdCompany) (uuid.UUID, error) {
	db.Lock()
	defer db.Unlock()

	id := uuid.New()
	ac.Id = id
	db.AdCompanies = append(db.AdCompanies, ac)

	return id, nil
}

func (r *AdCompanyMemory) Delete(id string) error {
	db.Lock()
	defer db.Unlock()

	return errors.New("not implemented")
}

func (r *AdCompanyMemory) Update(ac models.AdCompany, idStr string) (uuid.UUID, error) {
	db.Lock()
	defer db.Unlock()

	return uuid.Nil, errors.New("not implemented")
}

func (r *AdCompanyMemory) changeStatus(id uuid.UUID, status bool) error {
	db.Lock()
	defer db.Unlock()

	return errors.New("not implemented")
}

func (r *AdCompanyMemory) Start(id uuid.UUID) error {
	return r.changeStatus(id, true)
}

func (r *AdCompanyMemory) Stop(id uuid.UUID) error {
	return r.changeStatus(id, false)
}

func (r *AdCompanyMemory) GetAll(userId uuid.UUID) ([]models.AdCompany, error) {
	db.Lock()
	defer db.Unlock()

	var userComp []models.AdCompany

	logrus.Warn(userId)
	logrus.Warn(db.Users)
	logrus.Warn(db.AdCompanies)

	for _, ac := range db.AdCompanies {
		logrus.Errorf("%s \n %s", ac.UserId.String(), userId.String())
		if ac.UserId.String() == userId.String() {
			userComp = append(userComp, ac)
		}
	}
	if len(userComp) == 0 {
		return nil, errCompanyNotFound
	}
	return userComp, nil
}

func (r *AdCompanyMemory) GetByID(companyID string) (models.AdCompany, error) {
	db.Lock()
	defer db.Unlock()
	for _, ac := range db.AdCompanies {
		if ac.Id.String() == companyID {
			return ac, nil
		}
	}

	return models.AdCompany{}, errCompanyNotFound
}
