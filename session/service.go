package session

type Service interface {
	CreateSession(userId int) (string, error)
	// Login(login string, password string) (string, bool, error)
	// Check(value string) (bool, uint64, error)
}
