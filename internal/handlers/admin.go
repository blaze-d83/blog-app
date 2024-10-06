package handlers

import (
	"net/http"

	"github.com/blaze-d83/blog-app/pkg/auth"
	"github.com/blaze-d83/blog-app/pkg/services"
	"github.com/blaze-d83/blog-app/pkg/types"
	"github.com/blaze-d83/blog-app/pkg/utils"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	Repository *services.AdminRepository
}

func (h *AdminHandler) GetAdminLoginPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (h *AdminHandler) ProcessHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		admin, err := h.Repository.CheckAdminExists(username)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
		}

		err = auth.CompareHashPassword(admin, password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid credentials"})
		}

		token, err := auth.GenerateJWT(username)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not generate token"})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}
}

func (h *AdminHandler) AdminJWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing or invalid token"})
			}
			claims, err := auth.ValidateJWT(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or expired token"})
			}
			c.Set("username", claims.Username)

			return next(c)
		}
	}

}

func (h *AdminHandler) GetListOfPosts() echo.HandlerFunc {
	return func(c echo.Context) error {
		posts, err := h.Repository.GetAllPostsForAdmin()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve posts"})
		}
		return c.JSON(http.StatusOK, posts)
	}
}

func (h *AdminHandler) GetPostToPreview() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := utils.GetInt(c.Param("id"))
		if err != nil {
			return err
		}
		post, err := h.Repository.AdminGetPostByID(id)
		if err != nil {
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
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create post"})
		}
		return c.JSON(http.StatusCreated, post.Title)
	}
}

func (h *AdminHandler) UpdatePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := utils.GetInt(c.Param("id"))
		if err != nil {
			return err
		}
		post := parsePostFromRequest(c)
		err = h.Repository.UpdatePost(id, post)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve post to update"})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Post updated successfully"})
	}
}

func (h *AdminHandler) DeletePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := utils.GetInt(c.Param("id"))
		if err != nil {
			return err
		}
		err = h.Repository.DeletePost(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve post to delete"})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Post deleted successfully"})
	}
}

func (h *AdminHandler) GetListOfCategories() echo.HandlerFunc {
	return func(c echo.Context) error {
		categories, err := h.Repository.AdminGetAllCategories()
		if err != nil {
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
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create category"})
		}
		return c.JSON(http.StatusCreated, newCategory.Name)
	}
}

func (h *AdminHandler) DeleteCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := utils.GetInt(c.Param("id"))
		if err != nil {
			return err
		}
		err = h.Repository.DeleteCategory(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete category"})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Category deleted successfully"})
	}
}

func parsePostFromRequest(c echo.Context) types.Post {
	return types.Post{
		Title:       c.FormValue("title"),
		Citation:    c.FormValue("citation"),
		Summary:     c.FormValue("summary"),
		PhotoIcon:   c.FormValue("photo-icon"),
		BannerImage: c.FormValue("banner-image"),
	}
}
