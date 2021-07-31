package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) CreateList(c echo.Context) error {
	id := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	return nil
}

func (h *Handler) GetAllLists(c echo.Context) error {
	return nil
}

func (h *Handler) GetListById(c echo.Context) error {
	return nil
}

func (h *Handler) UpdateList(c echo.Context) error {
	return nil
}

func (h *Handler) DeleteList(c echo.Context) error {
	return nil
}
