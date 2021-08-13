package handler

import (
	"Todo/models"
	"Todo/server/middleware"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func (h *Handler) CreateItem(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var input models.TodoItem
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, err := h.services.CreateItem(userID, listID, input)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	return nil
}

func (h *Handler) GetAllItems(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var items []models.TodoItem
	items, err = h.services.GetAllItems(userID, listID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, items)
	return nil
}

func (h *Handler) GetItemById(c echo.Context) error {
	return nil
}

func (h *Handler) UpdateItem(c echo.Context) error {
	return nil
}

func (h *Handler) DeleteItem(c echo.Context) error {
	return nil
}
