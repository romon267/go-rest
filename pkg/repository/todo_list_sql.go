package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/romon267/go-rest/internal/entities"
	"strings"
)

type TodoListSql struct {
	db *sqlx.DB
}

func NewTodoListSql(db *sqlx.DB) *TodoListSql {
	return &TodoListSql{db: db}
}

func (r *TodoListSql) Create(userId int, list entities.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			// Not sure about fatal here
			// logrus.Fatalf("update drivers: unable to rollback: %v", rollbackErr)
			return 0, err
		}
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	if _, err := tx.Exec(createUsersListQuery, userId, id); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			// Not sure about fatal here
			// logrus.Fatalf("update drivers: unable to rollback: %v", rollbackErr)
			return 0, err
		}
	}

	return id, tx.Commit()
}

func (r *TodoListSql) GetAll(userId int) ([]entities.TodoList, error) {
	var lists []entities.TodoList
	query := fmt.Sprintf("SELECT tl.* FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1", todoListsTable, usersListsTable)

	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *TodoListSql) GetById(userId, listId int) (entities.TodoList, error) {
	var list entities.TodoList
	query := fmt.Sprintf("SELECT tl.* FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", todoListsTable, usersListsTable)

	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *TodoListSql) UpdateById(userId, listId int, updateDto entities.UpdateListInput) (entities.TodoList, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	var updatedList entities.TodoList

	if err := updateDto.Validate(); err != nil {
		return entities.TodoList{}, err
	}

	if updateDto.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *updateDto.Title)
		argId++
	}

	if updateDto.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *updateDto.Description)
		argId++
	}

	args = append(args, listId, userId)

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id = $%d AND ul.user_id = $%d RETURNING tl.*",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)

	row := r.db.QueryRow(query, args...)
	err := row.Scan(&updatedList.Id, &updatedList.Title, &updatedList.Description)

	return updatedList, err
}

func (r *TodoListSql) DeleteById(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2", todoListsTable, usersListsTable)

	_, err := r.db.Exec(query, userId, listId)

	return err
}
