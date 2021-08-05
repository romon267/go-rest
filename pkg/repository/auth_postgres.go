package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/romon267/go-rest/internal/entities"
	"golang.org/x/crypto/bcrypt"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user entities.User) (int, error) {
	var id int

	// Check if Username or Name are unique
	var existingUser entities.User
	query := fmt.Sprintf("SELECT * from %s WHERE (username) = ($1)", usersTable)
	row := r.db.QueryRow(query, user.Username)
	if err := row.Scan(&existingUser.Id, &existingUser.Name, &existingUser.Username, &existingUser.Password); err != nil {
		fmt.Println("Error searching dupe: ", err.Error())
	}

	if existingUser.Id != 0 {
		return 0, fmt.Errorf("username is already in use: %s", user.Username)
	}
	
	// Insert new user
	query = fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	passwordHash, err := HashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	row = r.db.QueryRow(query, user.Name, user.Username, passwordHash)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func HashPassword(password string) (string, error) {
	hashedPBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// TODO Not sure if it should Fatalf here, research and remove maybe
	if err != nil {
		return "", fmt.Errorf("Error hashing a password: %s\n", err.Error())
	}

	return string(hashedPBytes), err
}
