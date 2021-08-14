package pkg

import (
	"Todo/models"
	"net/http"
)

type Service interface {
	CreateUser(user models.User) (uint64, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (uint64, error)
	CreateCookieWithValue(value string) *http.Cookie
	CheckUser(username, password string) (uint64, error)

	CreateList(userId uint64, list models.TodoList) (uint64, error)
	GetAllLists(userID uint64) ([]models.TodoList, error)
	GetListByID(userID, listID uint64) (models.TodoList, error)
	DeleteList(userID, listID uint64) error
	UpdateList(userID, listID uint64, input models.UpdateListInput) error

	CreateItem(userID, listID uint64, item models.TodoItem) (uint64, error)
	GetAllItems(userID, listID uint64) ([]models.TodoItem, error)
}
