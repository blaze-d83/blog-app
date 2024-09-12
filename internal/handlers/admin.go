package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/blaze-d83/blog-app/internal/db"
	"github.com/blaze-d83/blog-app/internal/templates"
	"github.com/blaze-d83/blog-app/types"
	"github.com/labstack/echo/v4"
)

func GetPosts(dbInstance *db.Database) ([]types.Post, error) {
	var posts []types.Post
	if err := dbInstance.DB.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func GetCategories(dbInstance *db.Database) ([]types.Category, error) {
	var categories []types.Category
	if err := dbInstance.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil

}

func GetAdminDashboard(dbInstance *db.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		posts, err := GetPosts(dbInstance)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to fectch posts")
		}
		categories, err := GetCategories(dbInstance)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to fectch categories")
		}

		dashboard := templates.AdminDashboard(posts, categories)
		return dashboard.Render(context.Background(), c.Response().Writer)
	}
}

func CreatePost(dbInstance *db.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.FormValue("title")
		summary := c.FormValue("summary")
		content := c.FormValue("content")

		post := types.Post{
			Title:   title,
			Summary: summary,
			Content: types.Content{
				MainBody: content,
			},
			Date: time.Now(),
		}

		if err := dbInstance.DB.Create(&post).Error; err != nil {
			return c.String(http.StatusInternalServerError, "Failed to create post")
		}

		return GetAdminDashboard(dbInstance)(c)
	}
}

func UpdatePost(dbInstance *db.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid post ID")
		}

		var post types.Post
		if err := dbInstance.DB.First(&post, id).Error; err != nil {
			return c.String(http.StatusNotFound, "Post not found")
		}

		title := c.FormValue("title")
		summary := c.FormValue("summary")
		content := c.FormValue("content")

		post.Title = title
		post.Summary = summary
		post.Content.MainBody = content

		if err := dbInstance.DB.Save(&post).Error; err != nil {
			return c.String(http.StatusInternalServerError, "Failed to update post")
		}

		return GetAdminDashboard(dbInstance)(c)
	}
}

func DeletePost(dbInstance *db.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid post ID")
		}

		if err := dbInstance.DB.Delete(&types.Post{}, id).Error; err != nil {
			return c.String(http.StatusInternalServerError, "Failed to delete post")
		}

		return GetAdminDashboard(dbInstance)(c)
	}
}

func CreateCategory(dbInstance *db.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")

		category := types.Category{
			Name: name,
		}

		if err := dbInstance.DB.Create(&category).Error; err != nil {
			return c.String(http.StatusInternalServerError, "Failed to create category")
		}

		return GetAdminDashboard(dbInstance)(c)
	}
}

func UpdateCategory(dbInstance *db.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid category ID")
		}

		var category types.Category
		if err := dbInstance.DB.First(&category, id).Error; err != nil {
			return c.String(http.StatusNotFound, "Category not found")
		}

		name := c.FormValue("name")

		category.Name = name

		if err := dbInstance.DB.Save(&category).Error; err != nil {
			return c.String(http.StatusInternalServerError, "Failed to update category")
		}

		return GetAdminDashboard(dbInstance)(c)
	}
}

func DeleteCategory(dbInstance *db.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid category ID")
		}

		if err := dbInstance.DB.Delete(&types.Category{}, id).Error; err != nil {
			return c.String(http.StatusInternalServerError, "Failed to delete category")
		}

		return GetAdminDashboard(dbInstance)(c)
	}
}
