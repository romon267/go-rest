package service

import (
	"github.com/romon267/go-rest/internal/entities"
	"github.com/romon267/go-rest/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo}
}

func (s *TodoListService) Create(userId int, list entities.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]entities.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (entities.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) UpdateById(userId, listId int, updateDto entities.UpdateListInput) (entities.TodoList, error) {
	return s.repo.UpdateById(userId, listId, updateDto)
}

func (s *TodoListService) DeleteById(userId, listId int) error {
	return s.repo.DeleteById(userId, listId)
}
