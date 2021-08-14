package middleware

import (
	"Todo/microservice_auth/client"
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

type Auth struct {
	rpcAuth client.IAuthClient
}

func NewAuth(rpcAuth client.IAuthClient) Auth {
	return Auth{rpcAuth: rpcAuth}
}

func (h *Auth) UserIdentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		session, err := ctx.Cookie("Session_id")
		if session == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		userId, err := h.rpcAuth.Check(session.Value)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())

		}
		ctx.Set(userCtx, userId)
		return next(ctx)
	}
}

func GetUserID(c echo.Context) (uint64, error) {
	id := c.Get(userCtx)
	idInt, ok := id.(uint64)
	if !ok {
		return 0, errors.New("user id is not found")
	}
	return idInt, nil

}
