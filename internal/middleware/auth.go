package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("Authorization") != "Bearer valid-token" {
			return c.Redirect(http.StatusSeeOther, "/admin/login")
		}
		return next(c)
	}
}
