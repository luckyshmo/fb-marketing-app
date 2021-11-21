package inmemory

import (
	"errors"

	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
	"github.com/sirupsen/logrus"
)

var errUserNotfound = errors.New("user not found")

type AuthMemory struct {
}

func NewAuthMemory() *AuthMemory {
	return &AuthMemory{}
}

func (r *AuthMemory) CreateUser(user models.User) (uuid.UUID, error) {
	db.Lock()
	defer db.Unlock()

	user.Id = uuid.New()
	db.Users = append(db.Users, user)

	logrus.Infof("create user: %v", user)

	return user.Id, nil
}

func (r *AuthMemory) GetUser(email, passwordHash string) (models.User, error) {
	db.Lock()
	defer db.Unlock()

	for _, u := range db.Users {
		if u.Email == email && u.Password == passwordHash {
			return u, nil
		}
	}
	return models.User{}, errUserNotfound
}
