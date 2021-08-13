package server

import (
	"Todo/microservice_auth/client"
	"Todo/pkg/handler"
	"Todo/pkg/repository"
	"Todo/pkg/service"
	"Todo/server/middleware"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo"
	"log"
)

const (
	DBConnect = "user=postgres dbname=postgres password=4444 host=localhost port=5432 sslmode=disable pool_max_conns=50"
)

type Server struct {
	rpcAuth client.IAuthClient
	e       *echo.Echo
}

func NewServer() *Server {
	var server Server
	e := echo.New()
	pool, err := pgxpool.Connect(context.Background(), DBConnect)
	if err != nil {
		log.Fatal(err)
	}
	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	rpcAuth, err := client.NewAuthClient(":3001")
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(pool)

	services := service.NewService(repos)

	auth := middleware.NewAuth(rpcAuth)
	handler := handler.NewHandler(services, rpcAuth)

	handler.InitHandler(e, auth)
	server.e = e
	server.rpcAuth = rpcAuth
	return &server
}

func (s Server) ListenAndServe() {
	s.e.Start(":8000")
	defer s.rpcAuth.Close()
}
