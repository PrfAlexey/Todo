package service

import (
	"Todo/microservice_auth/session"
	"math/rand"
	"time"
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

func (s *SessionService) CreateSession(userID uint64) (string, error) {

	var id string
	var err error
	for {
		id = RandStringRunes(sessKeyLen)
		err = s.repo.InsertSession(userID, id)

		if err == nil {
			break
		}
	}
	return id, nil
}

func (s *SessionService) Check(session string) (uint64, error) {
	id, err := s.repo.CheckSession(session)
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func (s *SessionService) Logout(session string) error {
	return s.repo.DeleteSession(session)
}
