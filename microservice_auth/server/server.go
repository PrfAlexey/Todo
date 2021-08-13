package server

import (
	"Todo/microservice_auth/proto"
	srepo "Todo/microservice_auth/session/repository"
	session "Todo/microservice_auth/session/service"
	urepo "Todo/microservice_auth/user/repository"
	users "Todo/microservice_auth/user/service"
	"context"
	"flag"
	"github.com/gomodule/redigo/redis"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	DBConnect = "user=postgres dbname=postgres password=4444 host=localhost port=5432 sslmode=disable pool_max_conns=50"
)

var (
	redisAddr = flag.String("addr", "redis://user:@localhost:6379/0", "redis addr")
)

type Server struct {
	port string
	auth *AuthServer
}

func NewServer(port string) *Server {
	pool, err := pgxpool.Connect(context.Background(), DBConnect)
	if err != nil {
		log.Fatal(err)
	}
	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()
	redisConn, err := redis.DialURL(*redisAddr)
	if err != nil {
		logrus.Fatalf("Failed to initialization db redis : %s", err.Error())
	}

	sessionRepos := srepo.NewSessionRepository(redisConn)
	userRepos := urepo.NewUserDatabase(pool)
	uService := users.NewUserService(userRepos)

	s := session.NewSessionService(sessionRepos)

	return &Server{
		port: port,
		auth: NewAuthServer(s, uService),
	}
}

func (s *Server) ListenAndServe() error {

	gServer := grpc.NewServer()

	listener, err := net.Listen("tcp", s.port)
	defer listener.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	proto.RegisterAuthServer(gServer, s.auth)
	log.Println("starting server at " + s.port)
	err = gServer.Serve(listener)

	if err != nil {
		return nil
	}

	return nil
}
