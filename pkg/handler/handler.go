package handler

import (
	"Todo/microservice_auth/client"
	"Todo/pkg"
	"Todo/server/middleware"
	"github.com/labstack/echo"
)

type Handler struct {
	services pkg.Service
	rpcAuth  client.IAuthClient
}

func NewHandler(services pkg.Service, auth client.IAuthClient) *Handler {
	return &Handler{
		services: services,
		rpcAuth:  auth,
	}
}

func (h *Handler) InitHandler(e *echo.Echo, auth middleware.Auth) {

	g := e.Group("/auth")
	g.POST("/sign-up", h.SignUp)
	g.POST("/sign-in", h.SignIn)
	g.DELETE("/logout", h.Logout)

	a := e.Group("/api", auth.UserIdentity)

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
