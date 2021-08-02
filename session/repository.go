package session

type Repository interface {
	InsertSession(userId int, value string) error
	CheckSession(value string) (int, error)
	DeleteSession(value string) error
}
