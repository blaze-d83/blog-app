package handlers

import (
	"net/http"

	templates "github.com/blaze-d83/blog-app/internal/templates/pages"
	"github.com/blaze-d83/blog-app/pkg/logger"
	"github.com/blaze-d83/blog-app/pkg/services"
	"github.com/blaze-d83/blog-app/pkg/utils"
	"github.com/labstack/echo/v4"
)

type PublicHandler struct {
	Repository *services.UserRepository
	Logger     logger.CustomLogger
}

func NewPublicHandler(repo *services.UserRepository, customLogger logger.CustomLogger) *PublicHandler {
	return &PublicHandler{
		Repository: &services.UserRepository{},
		Logger:     customLogger,
	}
}

func (h *PublicHandler) Homepage() echo.HandlerFunc {
	return func(c echo.Context) error {
		homePage := templates.HomePage()
		if err := homePage.Render(c.Request().Context(), c.Response().Writer); err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to render home page"})
		}
		return nil
	}
}

func (h *PublicHandler) Shop() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (h *PublicHandler) About() echo.HandlerFunc {
	return func(c echo.Context) error {
		aboutPage := templates.AboutPage()
		if err := aboutPage.Render(c.Request().Context(), c.Response().Writer); err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to render about page"})
		}
		return nil
	}
}

func (h *PublicHandler) Support() echo.HandlerFunc {
	return func(c echo.Context) error {
		supportPage := templates.SupportPage()
		if err := supportPage.Render(c.Request().Context(), c.Response().Writer); err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to render support page"})
		}
		return nil
	}
}

func (h *PublicHandler) Events() echo.HandlerFunc {
	return func(c echo.Context) error {
		eventsPage := templates.EventsPage()
		if err := eventsPage.Render(c.Request().Context(), c.Response().Writer); err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to render events page"})
		}
		return nil
	}
}

func (h *PublicHandler) GetListOfAllPostsHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		posts, err := h.Repository.GetAllPosts()
		if err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve posts"})
		}
		return c.JSON(http.StatusOK, posts)
	}
}

func (h *PublicHandler) ViewFullPostHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := utils.GetInt(c.Param("id"))
		if err != nil {
			h.Logger.LogError(c, err)
			return err
		}
		post, err := h.Repository.GetPostByID(id)
		if err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve posts"})
		}
		return c.JSON(http.StatusOK, post)
	}
}
