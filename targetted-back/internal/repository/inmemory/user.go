package inmemory

import (
	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
)

type UserMemory struct {
}

func NewUserMemory() *UserMemory {
	return &UserMemory{}
}

func (r *UserMemory) SetBalance(id uuid.UUID, amount float64) error {
	db.Lock()
	defer db.Unlock()

	return nil
}

func (r *UserMemory) AddMoney(userId uuid.UUID, amount float64) error {
	db.Lock()
	defer db.Unlock()

	return nil
}

func (r *UserMemory) GetAll() ([]models.User, error) {
	db.Lock()
	defer db.Unlock()

	return db.Users, nil
}

func (r *UserMemory) GetById(userId uuid.UUID) (models.User, error) {
	db.Lock()
	defer db.Unlock()

	for _, u := range db.Users {
		if u.Id == userId {
			return u, nil
		}
	}

	return models.User{}, errUserNotfound
}
