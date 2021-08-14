package pkg

import "Todo/models"

type Repository interface {
	CreateUser(user models.User) (uint64, error)
	GetUser(username, password string) (models.User, error)

	CreateList(userId uint64, list models.TodoList) (uint64, error)
	GetAllLists(userID uint64) ([]models.TodoList, error)
	GetListByID(userID, listID uint64) (models.TodoList, error)
	DeleteList(userID, listID uint64) error
	UpdateList(userID, listID uint64, input models.UpdateListInput) error

	CreateItems(listID uint64, item models.TodoItem) (uint64, error)
	GetAllItems(userID, listID uint64) ([]models.TodoItem, error)
}
