package pkg

import "Todo/models"

type Repository interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)

	CreateList(userId int, list models.TodoList) (int, error)
	GetAllLists(userID int) ([]models.TodoList, error)
	GetListByID(userID, listID int) (models.TodoList, error)
	DeleteList(userID, listID int) error
	UpdateList(userID, listID int, input models.UpdateListInput) error

	CreateItems(listID int, item models.TodoItem) (int, error)
	GetAllItems(userID, listID int) ([]models.TodoItem, error)
}
