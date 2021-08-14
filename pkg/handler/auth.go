package handler

import (
	"Todo/models"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (h *Handler) SignUp(c echo.Context) error {
	var input models.User

	//session, err := c.Cookie("Session_id")
	//if err != nil {
	//	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	//}
	//
	//userID, _ := h.rpcAuth.Check(session.Value)
	//if userID != 0 {
	//	return echo.NewHTTPError(http.StatusBadRequest, "user is already logged in")
	//}

	if err := c.Bind(&input); err != nil {
		log.Fatal(err)
	}
	id, err := h.services.CreateUser(input)
	if err != nil {
		log.Fatal(err)
	}
	sessionID, err := h.rpcAuth.Login(input.Username, input.Password)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	cookie := h.services.CreateCookieWithValue(sessionID)
	c.SetCookie(cookie)

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
	session, err := h.rpcAuth.Login(input.Username, input.Password)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cookie := h.services.CreateCookieWithValue(session)
	c.SetCookie(cookie)

	err1 := c.JSON(http.StatusOK, map[string]interface{}{
		"session": session,
	})
	if err1 != nil {
		log.Fatal(err1)
	}
	return nil
}

func (h *Handler) Logout(c echo.Context) error {
	session, err := c.Cookie("Session_id")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = h.rpcAuth.Logout(session.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	c.SetCookie(session)
	return nil
}
