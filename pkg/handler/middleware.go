package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		session, err := ctx.Cookie("Session_id")

		if session == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		userId, err := h.sessServices.Check(session.Value)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())

		}
		ctx.Set(userCtx, userId)

		return next(ctx)
	}
}
