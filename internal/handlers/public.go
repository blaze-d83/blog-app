package handlers

import (
	"context"
	"net/http"

	"github.com/blaze-d83/blog-app/internal/services"
	"github.com/blaze-d83/blog-app/static/components"
	"github.com/blaze-d83/blog-app/utils"
	"github.com/labstack/echo/v4"
)

type PublicHandler struct {
	public services.PublicService
}

func Homepage(c echo.Context) error {
	w := c.Response().Writer

	homePage := components.HomePage()
	err := homePage.Render(context.Background(), w)
	if err != nil {
		http.Error(w, "failed to render homepage", http.StatusInternalServerError)
	}
	return nil
}

func (h *PublicHandler) GetListOfAllPostsHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		posts, err := h.public.UsersGetAllPosts()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve posts"})
		}
		return c.JSON(http.StatusOK, posts)
	}
}

func (h *PublicHandler) ViewFullPostHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := utils.GetInt(c.Param("id"))
		post, err := h.public.UsersGetPostsByID(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve posts"})
		}
		return c.JSON(http.StatusOK, post)
	}
}
