package service

import (
	"github.com/romon267/go-rest/internal/entities"
	"github.com/romon267/go-rest/pkg/repository"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GetUser(username string) (entities.User, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list entities.TodoList) (int, error)
	GetAll(userId int) ([]entities.TodoList, error)
	GetById(userId, listId int) (entities.TodoList, error)
	UpdateById(userId, listId int, updateDto entities.UpdateListInput) (entities.TodoList, error)
	DeleteById(userId, listId int) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Authorization: NewAuthService(repos), TodoList: NewTodoListService(repos)}
}
