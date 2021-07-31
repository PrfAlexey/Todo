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
		header, err := ctx.Cookie(authorizationHeader)

		if header == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		userId, err := h.services.Authorization.ParseToken(header.Value)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())

		}
		ctx.Set(userCtx, userId)

		return next(ctx)
	}
}
