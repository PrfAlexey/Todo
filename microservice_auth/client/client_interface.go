package client

type IAuthClient interface {
	Login(username, password string) (string, error)
	Logout(session string) error
	Check(session string) (uint64, error)
	Close()
}
