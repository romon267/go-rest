package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/romon267/go-rest/internal/entities"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GetUser(username string) (entities.User, error)
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

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authorization: NewAuthSql(db), TodoList: NewTodoListSql(db)}
}
