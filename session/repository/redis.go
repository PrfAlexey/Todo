package repository

import (
	"Todo/session"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type SessionRepository struct {
	Conn redis.Conn
}

func NewSessionRepository(c redis.Conn) session.Repository {
	return &SessionRepository{Conn: c}
}

func (r *SessionRepository) InsertSession(userId int, value string) error {
	mkey := "sessions:" + value
	result, err := redis.String(r.Conn.Do("SET", mkey, userId, "EX", 86400))
	if err != nil {
		return err
	}
	if result != "OK" {
		return fmt.Errorf("result not OK")
	}
	return nil
}
