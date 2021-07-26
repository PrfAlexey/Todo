package server

import (
	"ads/pkg/handler"
	"ads/pkg/repository"
	"ads/pkg/service"
	"log"

	"github.com/labstack/echo"
)

type Server struct {
	e *echo.Echo
}

func NewServer() *Server {
	var server Server
	e := echo.New()
	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)
	handler.InitHandler(e)
	server.e = e
	return &server
}

func (s Server) ListenAndServe() {
	s.e.Start(":8000")
}
