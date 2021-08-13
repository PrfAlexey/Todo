package session

type Service interface {
	CreateSession(userID uint64) (string, error)
	Check(session string) (uint64, error)
	Logout(session string) error
}
