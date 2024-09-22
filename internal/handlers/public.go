package handlers

import (
	"context"
	"net/http"

	"github.com/blaze-d83/blog-app/static/components"
	"github.com/labstack/echo/v4"
)

func Homepage(c echo.Context) error {
	w := c.Response().Writer

	homePage := components.HomePage()
	err := homePage.Render(context.Background(), w)
	if err != nil {
		http.Error(w, "failed to render homepage", http.StatusInternalServerError )
	}
	return nil
}
