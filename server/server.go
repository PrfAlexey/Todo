package server

import (
	"ads/pkg/handler"
	"ads/pkg/repository"
	"ads/pkg/service"

	"github.com/labstack/echo"
)

type Server struct {
	e *echo.Echo
}

func NewServer() *Server {
	var server Server
	e := echo.New()
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server.e = e
	return &server
}

func (s Server) ListenAndServe() {
	s.e.Start(":8000")
}
