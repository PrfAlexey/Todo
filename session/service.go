package session

type Service interface {
	CreateSession(userId int) (string, error)
	// Login(login string, password string) (string, bool, error)
	Check(session string) (int, error)
	Logout(session string) error
}
