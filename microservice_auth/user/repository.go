package user

import "Todo/models"

type Repository interface {
	GetUser(login, password string) (models.User, error)
}
