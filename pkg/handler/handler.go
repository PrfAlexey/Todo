package handler

import (
	"Todo/pkg/service"

	"github.com/labstack/echo"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitHandler(e *echo.Echo) {

	g := e.Group("/auth")
	g.POST("/sign-up", h.SignUp)
	g.POST("/sign-in", h.SignIn)

	a := e.Group("/api", h.userIdentity)

	a.POST("/lists/", h.CreateList)
	a.GET("/lists/", h.GetAllLists)
	a.GET("/lists/:id", h.GetListById)
	a.PUT("/lists/:id", h.UpdateList)
	a.DELETE("/lists/:id", h.DeleteList)

	a.POST("/lists/:id/items", h.CreateItem)
	a.GET("/lists/:id/items", h.GetAllItems)
	a.GET("/lists/:id/items/:item_id", h.GetItemById)
	a.PUT("/lists/:id/items/:item_id", h.UpdateItem)
	a.DELETE("/lists/:id/items/:item_id", h.DeleteItem)
}
