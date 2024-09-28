package handlers

import (
	"context"
	"net/http"

	"github.com/blaze-d83/blog-app/internal/services"
	"github.com/blaze-d83/blog-app/static/components"
	"github.com/blaze-d83/blog-app/types"
	"github.com/blaze-d83/blog-app/utils"
	"github.com/labstack/echo/v4"
)

type AdminServiceHandler struct {
	services.AdminService
}

func RenderAdminDashboard(c echo.Context) error {
	w := c.Response().Writer
	adminPage := components.AdminDashboard()
	err := adminPage.Render(context.Background(), w)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to render admin dashboard"})
	}
	return nil
}

func (s *AdminServiceHandler) GetListOfPosts() echo.HandlerFunc {
	return func(c echo.Context) error {
		posts, err := s.AdminService.GetAllPostsForAdmin()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve posts"})
		}
		return c.JSON(http.StatusOK, posts)
	}
}

func (s *AdminServiceHandler) GetPostToPreview() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := utils.GetInt(c.Param("id"))
		post, err := s.AdminService.GetPostByID(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve post to preview"})
		}
		return c.JSON(http.StatusOK, post)
	}
}

func (s *AdminServiceHandler) CreatePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		post := types.Post{
			Title:       c.FormValue("title"),
			Citation:    c.FormValue("citation"),
			Summary:     c.FormValue("summary"),
			Content:     c.FormValue("content"),
			PhotoIcon:   c.FormValue("photo-link"),
			BannerImage: c.FormValue("banner-link"),
		}
		err := s.AdminService.CreatePost(post)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create post"})
		}
		return c.JSON(http.StatusCreated, post.Title)
	}
}

func (s *AdminServiceHandler) UpdatePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := utils.GetInt(c.Param("id"))
		post := types.Post{
			Title:       c.FormValue("title"),
			Citation:    c.FormValue("citation"),
			Summary:     c.FormValue("summary"),
			Content:     c.FormValue("content"),
			PhotoIcon:   c.FormValue("photo-link"),
			BannerImage: c.FormValue("banner-link"),
		}

		err := s.AdminService.UpdatePost(id, post)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve post to update"})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Post updated successfully"})
	}
}

func (s *AdminServiceHandler) DeletePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := utils.GetInt(c.Param("id"))
		err := s.AdminService.DeletePost(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve post to delete"})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Post deleted successfully"})
	}
}

func (s *AdminServiceHandler) GetListOfCategories() echo.HandlerFunc {
	return func(c echo.Context) error {
		categories, err := s.AdminService.AdminGetAllCategories()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve categories"})
		}
		return c.JSON(http.StatusOK, categories)
	}
}

func (s *AdminServiceHandler) CreateCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		newCategory := types.Category{
			Name: c.FormValue("name"),
		}
		err := s.AdminService.CreateCategory(newCategory)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create category"})
		}
		return c.JSON(http.StatusCreated, newCategory.Name)
	}
}

func (s *AdminServiceHandler) DeleteCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := utils.GetInt(c.Param("id"))
		err := s.AdminService.DeleteCategory(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete category"})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Category deleted successfully"})
	}
}
