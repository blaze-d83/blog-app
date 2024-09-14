package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminGetAllPostsHandler(c echo.Context) error {
	w := c.Request()
	if w.Method == http.MethodGet {
	}
	return nil
}
