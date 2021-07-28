package handler

import (
	"ads/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) SignUp(c echo.Context) error {
	var input models.User

	if err := c.Bind(&input); err != nil {
		log.Fatal(err)
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	return nil
}

func (h *Handler) SignIn(c echo.Context) error {
	return nil
}
