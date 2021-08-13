package user

type Service interface {
	CheckUser(login, password string) (uint64, error)
}
