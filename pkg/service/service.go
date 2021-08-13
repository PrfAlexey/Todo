package service

import (
	"Todo/models"
	"Todo/pkg"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

const (
	salt       = "asfsdgsg"
	signingKey = "fsddgsdgsgvgs"
)

type Service struct {
	repo pkg.Repository
}

func NewService(rep pkg.Repository) pkg.Service {
	return &Service{
		repo: rep,
	}
}

func (s *Service) CreateUser(user models.User) (int, error) {
	user.Password = GeneratePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (s *Service) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))

}

func (s *Service) ParseToken(accesstoken string) (int, error) {
	token, err := jwt.ParseWithClaims(accesstoken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *Service) CreateCookieWithValue(value string) *http.Cookie {
	newCookie := &http.Cookie{
		Name:     "Session_id",
		Value:    value,
		Expires:  time.Now().Add(72 * time.Hour),
		HttpOnly: true,
	}

	return newCookie
}

func (s *Service) CheckUser(username, password string) (int, error) {
	user, err := s.repo.GetUser(username, GeneratePasswordHash(password))
	if err != nil {
		return 0, err
	}

	return user.Id, nil
}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *Service) CreateList(userId int, list models.TodoList) (int, error) {
	return s.repo.CreateList(userId, list)
}

func (s *Service) GetAllLists(userID int) ([]models.TodoList, error) {
	return s.repo.GetAllLists(userID)
}

func (s *Service) GetListByID(userID int, listID int) (models.TodoList, error) {
	return s.repo.GetListByID(userID, listID)
}

func (s *Service) DeleteList(userID, listID int) error {
	return s.repo.DeleteList(userID, listID)
}
func (s *Service) UpdateList(userID, listID int, input models.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateList(userID, listID, input)
}

func (s *Service) CreateItem(userID, listID int, item models.TodoItem) (int, error) {
	_, err := s.repo.GetListByID(userID, listID)
	if err != nil {
		return 0, err
	}
	return s.repo.CreateItems(listID, item)
}

func (s *Service) GetAllItems(userID, listID int) ([]models.TodoItem, error) {
	return s.repo.GetAllItems(userID, listID)
}
