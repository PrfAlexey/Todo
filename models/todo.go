package models

import "errors"

type TodoList struct {
	Id          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UserList struct {
	Id     uint64
	UserId uint64
	ListId uint64
}

type TodoItem struct {
	Id          uint64    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id     uint64
	ListId uint64
	ItemId uint64
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (l *UpdateListInput) Validate() error {
	if l.Title == nil && l.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
