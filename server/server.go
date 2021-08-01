package server

import (
	"ads/pkg/handler"
	"ads/pkg/repository"
	"ads/pkg/service"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
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
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)
	handler.InitHandler(e)
	server.e = e
	/*done*/
	return &server
}

func (s Server) ListenAndServe() {
	s.e.Start(":8000")
}
