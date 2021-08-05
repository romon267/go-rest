package service

import (
	"github.com/romon267/go-rest/internal/entities"
	"github.com/romon267/go-rest/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user entities.User) (int, error) {

	return s.repo.CreateUser(user)
}
