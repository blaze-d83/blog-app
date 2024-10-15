package handlers

import (
	"net/http"
	"time"

	templates "github.com/blaze-d83/blog-app/internal/templates/pages"
	"github.com/blaze-d83/blog-app/pkg/auth"
	"github.com/blaze-d83/blog-app/pkg/logger"
	"github.com/blaze-d83/blog-app/pkg/services"
	"github.com/blaze-d83/blog-app/pkg/types"
	"github.com/blaze-d83/blog-app/pkg/utils"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	Repository *services.AdminRepository
	Logger     logger.CustomLogger
}

func NewAdminHandler(repo *services.AdminRepository, customLogger logger.CustomLogger) *AdminHandler {
	return &AdminHandler{
		Repository: repo,
		Logger:     customLogger,
	}
}

func (h *AdminHandler) LoginPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginPage := templates.LoginPage()
		if err := loginPage.Render(c.Request().Context(), c.Response()); err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to render login page"})
		}
		return nil
	}
}

func (h *AdminHandler) AdminDashboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		dashboard := templates.AdminDashboard()
		if err := dashboard.Render(c.Request().Context(), c.Response()); err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to render admin dashboard"})
		}
		return nil
	}
}

func (h *AdminHandler) ProcessLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		admin, err := h.Repository.CheckAdminExists(username)
		if err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
		}

		err = auth.CompareHashPassword(admin, password)
		if err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid credentials"})
		}

		token, err := auth.GenerateJWT(username)
		if err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not generate token"})
		}

		cookie := &http.Cookie{
			Name:     "auth_token",
			Value:    token,
			Path:     "",
			Expires:  time.Now().Add(24 * time.Hour),
			Secure:   false,
			HttpOnly: false,
		}
		c.SetCookie(cookie)

		return c.Redirect(http.StatusSeeOther, "/admin/dashboard")
	}
}

func (h *AdminHandler) GetListOfPosts() echo.HandlerFunc {
	return func(c echo.Context) error {
		posts, err := h.Repository.GetAllPostsForAdmin()
		if err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve posts"})
		}
		return c.JSON(http.StatusOK, posts)
	}
}

func (h *AdminHandler) GetPostToPreview() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := utils.GetInt(c.Param("id"))
		if err != nil {
			h.Logger.LogError(c, err)
			return err
		}
		post, err := h.Repository.AdminGetPostByID(id)
		if err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve post to preview"})
		}
		return c.JSON(http.StatusOK, post)
	}
}

func (h *AdminHandler) CreatePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		post := parsePostFromRequest(c)
		err := h.Repository.CreatePost(post)
		if err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create post"})
		}
		h.Logger.LogEvent("Post created successfully")
		return c.JSON(http.StatusCreated, post.Title)
	}
}

func (h *AdminHandler) UpdatePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := utils.GetInt(c.Param("id"))
		if err != nil {
			h.Logger.LogError(c, err)
			return err
		}
		post := parsePostFromRequest(c)
		err = h.Repository.UpdatePost(id, post)
		if err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve post to update"})
		}
		h.Logger.LogEvent("Post updated successfully")
		return c.JSON(http.StatusOK, map[string]string{"message": "Post updated successfully"})
	}
}

func (h *AdminHandler) DeletePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := utils.GetInt(c.Param("id"))
		if err != nil {
			h.Logger.LogError(c, err)
			return err
		}
		err = h.Repository.DeletePost(id)
		if err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve post to delete"})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Post deleted successfully"})
	}
}

func (h *AdminHandler) GetListOfCategories() echo.HandlerFunc {
	return func(c echo.Context) error {
		categories, err := h.Repository.AdminGetAllCategories()
		if err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve categories"})
		}
		return c.JSON(http.StatusOK, categories)
	}
}

func (h *AdminHandler) CreateCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		newCategory := types.Category{
			Name: c.FormValue("name"),
		}
		err := h.Repository.CreateCategory(newCategory)
		if err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create category"})
		}
		return c.JSON(http.StatusCreated, newCategory.Name)
	}
}

func (h *AdminHandler) DeleteCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := utils.GetInt(c.Param("id"))
		if err != nil {
			h.Logger.LogError(c, err)
			return err
		}
		err = h.Repository.DeleteCategory(id)
		if err != nil {
			h.Logger.LogError(c, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete category"})
		}
		h.Logger.LogEvent("Category deleted:")
		return c.JSON(http.StatusOK, map[string]string{"message": "Category deleted successfully"})
	}
}

/*
------------------------------------------------
				Local Utilities
------------------------------------------------
*/

func parsePostFromRequest(c echo.Context) types.Post {
	return types.Post{
		Title:       c.FormValue("title"),
		Citation:    c.FormValue("citation"),
		Summary:     c.FormValue("summary"),
		PhotoIcon:   c.FormValue("photo-icon"),
		BannerImage: c.FormValue("banner-image"),
	}
}
