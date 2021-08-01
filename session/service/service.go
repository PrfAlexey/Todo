package service

import (
	"ads/session/repository"
	"asfdasf/session"
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const sessKeyLen = 10

type SessionID struct {
	ID string
}

type SessionService struct {
	repo repository.SessionRepository
}

func NewSessionService(s repository.SessionRepository) session.Service {
	return &SessionService{repo: s}
}

func (s *SessionService) CreateSession(userId int) (string, error) {
	id := RandStringRunes(sessKeyLen)
	err := s.repo.InsertSession(userId, id)
	if err != nil {

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
