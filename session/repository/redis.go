package repository

import (
	"Todo/session"
	"fmt"
	"log"

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

func (r *SessionRepository) CheckSession(value string) ( int, error) {
	mkey := "sessions:" + value
	data, err := redis.Int(r.Conn.Do("GET", mkey))
	if err != nil {
		log.Println("cant get data:", err)
		return 0,err
	}

	return data, nil
}

func (r *SessionRepository) DeleteSession(value string) error{
	mkey := "sessions:" + value
	_, err := redis.Int(r.Conn.Do("DEL", mkey))
	if err != nil {
		log.Println("redis error:", err)
		return err
	}
	return nil
}