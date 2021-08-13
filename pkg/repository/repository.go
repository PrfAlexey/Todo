package repository

import (
	"Todo/models"
	"Todo/pkg"
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
)

const (
	usersTable      = "users"
	todoListTable   = "todo_lists"
	userListsTable  = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) pkg.Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.pool.QueryRow(context.Background(), query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) GetUser(username, password string) (models.User, error) {
	var user models.User
	err := r.pool.QueryRow(context.Background(), `SELECT * FROM users WHERE username=$1 AND password_hash=$2`, username, password).Scan(&user.Id, &user.Username, &user.Password, &user.Name)

	return user, err
}

func (r *Repository) CreateList(userId int, list models.TodoList) (int, error) {
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return 0, err
	}
	var id int
	row := tx.QueryRow(context.Background(), "INSERT INTO todo_lists (title, description) VALUES ($1, $2) RETURNING id", list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback(context.Background())
		return 0, err
	}

	_, err = tx.Exec(context.Background(), "INSERT INTO user_lists (user_id, list_id) VALUES ($1, $2)", userId, id)
	if err != nil {
		tx.Rollback(context.Background())
		return 0, err
	}

	return id, tx.Commit(context.Background())
}

func (r *Repository) GetAllLists(userID int) ([]models.TodoList, error) {
	var lists []models.TodoList

	err := pgxscan.Select(context.Background(), r.pool, &lists, `SELECT tl.id, tl.title, tl.description FROM todo_lists tl 
                                                                  INNER JOIN user_lists ul ON tl.id = ul.list_id WHERE ul.user_id = $1 `, userID)
	if err != nil {
		return nil, err
	}
	return lists, nil
}

func (r *Repository) GetListByID(userID int, listID int) (models.TodoList, error) {
	var list models.TodoList
	err := pgxscan.Get(context.Background(), r.pool, &list, `SELECT tl.id, tl.title, tl.description FROM todo_lists tl 
                                                                  INNER JOIN user_lists ul ON tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`, userID, listID)
	if err != nil {
		return list, err
	}
	return list, nil
}

func (r *Repository) DeleteList(userID, listID int) error {
	_, err := r.pool.Exec(context.Background(), `DELETE FROM todo_lists tl USING user_lists ul 
                                                        WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2`, userID, listID)
	return err
}

func (r *Repository) UpdateList(userID, listID int, input models.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE todo_lists tl SET %s FROM user_lists ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		setQuery, argId, argId+1)
	args = append(args, listID, userID)
	_, err := r.pool.Exec(context.Background(), query, args...)
	return err
}

func (r *Repository) CreateItems(listID int, item models.TodoItem) (int, error) {
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return 0, err
	}
	var itemID int
	row := tx.QueryRow(context.Background(), "INSERT INTO todo_items (title, description) VALUES ($1, $2) RETURNING id", item.Title, item.Description)
	err = row.Scan(&itemID)
	if err != nil {
		tx.Rollback(context.Background())
		return 0, err
	}

	_, err = tx.Exec(context.Background(), "INSERT INTO lists_items (list_id, item_id) values ($1, $2)", listID, itemID)
	if err != nil {
		tx.Rollback(context.Background())
		return 0, err
	}
	return itemID, tx.Commit(context.Background())
}

func (r *Repository) GetAllItems(userID, listID int) ([]models.TodoItem, error) {
	var items []models.TodoItem
	err := pgxscan.Select(context.Background(), r.pool, &items, `SELECT ti.id, ti.title, ti.description, ti.done FROM todo_items ti INNER JOIN lists_items li on li.item_id = ti.id
									INNER JOIN user_lists ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`, listID, userID)
	if err != nil {
		return nil, err
	}
	return items, err
}
