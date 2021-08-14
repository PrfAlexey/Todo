package handler

import (
	"Todo/models"
	"Todo/server/middleware"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (h *Handler) CreateList(c echo.Context) error {
	userId, err := middleware.GetUserID(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	var input models.TodoList
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	listId, err1 := h.services.CreateList(userId, input)
	if err1 != nil {
		return err1
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"listId": listId,
	})
	return nil
}

type getAllListsResponse struct {
	Data []models.TodoList `json:"data"`
}

func (h *Handler) GetAllLists(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	lists, err1 := h.services.GetAllLists(userID)
	if err1 != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
	return nil
}

func (h *Handler) GetListById(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	list, err1 := h.services.GetListByID(userID, uint64(listID))
	if err1 != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err1.Error())
	}

	if err = c.JSON(http.StatusOK, list); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (h *Handler) UpdateList(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var input models.UpdateListInput
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = h.services.UpdateList(userID, uint64(listID), input)
	if err.Error() == "update structure has no values" {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (h *Handler) DeleteList(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.services.DeleteList(userID, uint64(listID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, "OK")
	return nil
}
