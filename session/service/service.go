package service

import (
	"Todo/session"
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const sessKeyLen = 10

type SessionID struct {
	ID string
}

type SessionService struct {
	repo session.Repository
}

func NewSessionService(s session.Repository) session.Service {
	return &SessionService{repo: s}
}

func (s *SessionService) CreateSession(userId int) (string, error) {
	id := RandStringRunes(sessKeyLen)
	err := s.repo.InsertSession(userId, id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *SessionService) Check(session string) (int, error) {
	id,err:=s.repo.CheckSession(session)
	if err!=nil {
		return 0, err
	}
	return id, nil
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (s *SessionService) Logout(session string) error {
	return s.repo.DeleteSession(session)
}