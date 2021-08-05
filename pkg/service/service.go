package service

import (
	"github.com/romon267/go-rest/internal/entities"
	"github.com/romon267/go-rest/pkg/repository"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Authorization: NewAuthService(repos)}
}
