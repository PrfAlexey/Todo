package session

type Repository interface {
	InsertSession(userId uint64, value string) error
	CheckSession(value string) (int, error)
	DeleteSession(value string) error
}
