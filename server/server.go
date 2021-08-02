package server

import (
	"Todo/pkg/handler"
	"Todo/pkg/repository"
	"Todo/pkg/service"
	seRep "Todo/session/repository"
	seServ "Todo/session/service"
	"flag"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)
var (
	redisAddr = flag.String("addr", "redis://user:@localhost:6379/0", "redis addr")
)
type Server struct {
	e *echo.Echo
}

func NewServer() *Server {
	var server Server
	e := echo.New()
	db, err := repository.NewPostgresDB()
	if err != nil {
		logrus.Fatalf("Failed to initialization db: %s", err.Error())
	}
	flag.Parse()
	redisConn, err := redis.DialURL(*redisAddr)
	if err != nil {
		logrus.Fatalf("Failed to initialization db redis : %s", err.Error())
	}
	sessionRepos := seRep.NewSessionRepository(redisConn)
	repos := repository.NewRepository(db)

	services := service.NewService(repos)
	sessionServices := seServ.NewSessionService(sessionRepos)

	handler := handler.NewHandler(services, sessionServices)
	handler.InitHandler(e)
	server.e = e
	return &server
}

func (s Server) ListenAndServe() {

	s.e.Start(":8000")

}
