package handlers

import (
	"net/http"

	"github.com/blaze-d83/blog-app/internal/services"
	"github.com/blaze-d83/blog-app/utils"
	"github.com/labstack/echo/v4"
)

type PublicHandler struct {
	public services.PublicService
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
