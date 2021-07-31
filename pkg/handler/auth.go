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

type SignInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) SignIn(c echo.Context) error {
	var input SignInInput

	if err := c.Bind(&input); err != nil {
		log.Fatal(err)
	}
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		log.Fatal(err)
	}
	cookie := h.services.Authorization.CreateCookieWithValue(token)
	c.SetCookie(cookie)

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

	return nil
}
