package service

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/romon267/go-rest/internal/entities"
	"github.com/romon267/go-rest/pkg/repository"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type AuthService struct {
	repo repository.Authorization
}

type customClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user entities.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	// Get user by username
	user, err := s.GetUser(username)
	if err != nil {
		return "", fmt.Errorf("no user with username %s: %s", username, err.Error())
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Println("passes", user.Password, password)
		return "", fmt.Errorf("authorization error, wrong password")
	}

	// Generate token and return
	claims := customClaims{
		Username: user.Username,
		Id:       user.Id,
		StandardClaims: jwt.StandardClaims{
			// Expires in 12 hours
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "go-rest",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", fmt.Errorf("error signing token: %s", err.Error())
	}

	return signedToken, nil
}

func (s *AuthService) GetUser(username string) (entities.User, error) {
	return s.repo.GetUser(username)
}

func (s *AuthService) ParseToken(token string) (int, error) {
	t, err := jwt.ParseWithClaims(token, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := t.Claims.(*customClaims)
	if !ok {
		return 0, fmt.Errorf("wrong token claims")
	}

	return claims.Id, nil
}
