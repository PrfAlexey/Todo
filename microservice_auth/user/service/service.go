package service

import (
	"Todo/microservice_auth/user"
	"Todo/pkg/service"
)

type UserService struct {
	urepo user.Repository
}

func NewUserService(urepo user.Repository) *UserService {
	return &UserService{urepo: urepo}
}

func (s *UserService) CheckUser(login, password string) (uint64, error) {
	user, err := s.urepo.GetUser(login, service.GeneratePasswordHash(password))
	if err != nil {
		return 0, err
	}

	return uint64(user.Id), nil
}
