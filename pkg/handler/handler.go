package handler

import (
	"ads/pkg/service"

	"github.com/labstack/echo"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitHandler(e *echo.Echo) {
	var handler Handler

	e.POST("/auth/sign-up", handler.SignUp)
	e.POST("/auth/sign-in", handler.SignIn)

	e.POST("/api/lists/", handler.CreateList)
	e.GET("/api/lists/", handler.GetAllLists)
	e.GET("/api/lists/:id", handler.GetListById)
	e.PUT("/api/lists/:id", handler.UpdateList)
	e.DELETE("/api/lists/:id", handler.DeleteList)

	e.POST("/api/lists/:id/items", handler.CreateItem)
	e.GET("/api/lists/:id/items", handler.GetAllItems)
	e.GET("/api/lists/:id/items/:item_id", handler.GetItemById)
	e.PUT("/api/lists/:id/items/:item_id", handler.UpdateItem)
	e.DELETE("/api/lists/:id/items/:item_id", handler.DeleteItem)
}
