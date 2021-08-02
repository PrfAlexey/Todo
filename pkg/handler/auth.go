package handler

import (
	"Todo/models"
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
	userId, err:=h.services.CheckUser(input.Username,input.Password)
	if err!=nil {
		log.Fatal(err)
	}

	id, err := h.sessServices.CreateSession(userId)
	if err != nil {
		log.Fatal(err)
	}

	cookie := h.services.Authorization.CreateCookieWithValue(id)
	c.SetCookie(cookie)

	err1 := c.JSON(http.StatusOK, map[string]interface{}{
		"session": id,
	})
	if err1!=nil {
		log.Fatal(err1)
	}

	return nil
}


func (h *Handler) Logout(c echo.Context) error {
	session, err := c.Cookie("Session_id")
	if err!=nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	err1 := h.sessServices.Logout(session.Value)
	if err1!= nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}