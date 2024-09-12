package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Homepage(c echo.Context) error {
	 return c.String(http.StatusOK, "Welcome to the Blog app")
}
