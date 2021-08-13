package pkg

import (
	"Todo/models"
	"net/http"
)

type Service interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	CreateCookieWithValue(value string) *http.Cookie
	CheckUser(username, password string) (int, error)

	CreateList(userId int, list models.TodoList) (int, error)
	GetAllLists(userID int) ([]models.TodoList, error)
	GetListByID(userID, listID int) (models.TodoList, error)
	DeleteList(userID, listID int) error
	UpdateList(userID, listID int, input models.UpdateListInput) error

	CreateItem(userID, listID int, item models.TodoItem) (int, error)
	GetAllItems(userID, listID int) ([]models.TodoItem, error)
}
