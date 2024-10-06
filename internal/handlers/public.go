package handlers

import (
	"net/http"

	"github.com/blaze-d83/blog-app/pkg/services"
	"github.com/blaze-d83/blog-app/pkg/utils"
	"github.com/labstack/echo/v4"
)

type PublicHandler struct {
	service services.PublicService
}

func (h *PublicHandler) GetListOfAllPostsHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		posts, err := h.service.GetAllPosts()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve posts"})
		}
		return c.JSON(http.StatusOK, posts)
	}
}

func (h *PublicHandler) ViewFullPostHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := utils.GetInt(c.Param("id"))
		if err != nil {
			return err
		}
		post, err := h.service.GetPostsByID(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve posts"})
		}
		return c.JSON(http.StatusOK, post)
	}
}
