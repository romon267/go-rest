package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/romon267/go-rest/internal/entities"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authorization: NewAuthPostgres(db)}
}
