package service

import (
	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/repository"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
)

type UserService struct {
	repo repository.User
}

var _ User = (*UserService)(nil)

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetById(userId uuid.UUID) (models.User, error) {
	return s.repo.GetById(userId)
}

func (s *UserService) AddMoney(userId uuid.UUID, amount float64) error {
	return s.repo.AddMoney(userId, amount)
}

func (s *UserService) SetBalance(id uuid.UUID, amount float64) error {
	return s.repo.SetBalance(id, amount)
}
