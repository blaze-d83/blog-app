package handlers

import (
	"net/http"
	"strconv"

	"github.com/blaze-d83/blog-app/internal/services"
	"github.com/blaze-d83/blog-app/types"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	AdminService services.AdminService
}

func NewHandler(service services.AdminService) *AdminHandler{
	return &AdminHandler{
		AdminService: service,
	}
}

func (h *AdminHandler) AdminGetListOfPosts() echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		if r.Method == http.MethodGet {
			posts, err := h.AdminService.GetAllPostsForAdmin()
			if err != nil {
				return c.JSON(http.StatusNotFound, err)
			}
			return c.JSON(http.StatusOK, posts)
		}
		return c.JSON(http.StatusMethodNotAllowed, r.Method)
	}
}

func (h *AdminHandler) AdminGetPostToPreview() echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		if r.Method == http.MethodGet {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				return err
			}
			post, err := h.AdminService.GetPostByID(uint(id))
			if err != nil {
				return c.JSON(http.StatusNotFound, err)
			}
			return c.JSON(http.StatusOK, post)
		}
		return c.JSON(http.StatusMethodNotAllowed, r.Method)
	}
}

func (h *AdminHandler) AdminCreatePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		if r.Method == http.MethodPost {
			post := types.Post{
				Title:       c.FormValue("title"),
				Citation:    c.FormValue("citation"),
				Summary:     c.FormValue("summary"),
				Content:     c.FormValue("content"),
				PhotoIcon:   c.FormValue("photo-link"),
				BannerImage: c.FormValue("banner-link"),
			}
			err := h.AdminService.CreatePost(post)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, err)
			}
			return c.JSON(http.StatusCreated, post.Title)
		}
		return c.JSON(http.StatusMethodNotAllowed, r.Method)
	}
}

func (h *AdminHandler) AdminUpdatePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		if r.Method == http.MethodPut {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				return err
			}
			post := types.Post{
				Title:       c.FormValue("title"),
				Citation:    c.FormValue("citation"),
				Summary:     c.FormValue("summary"),
				Content:     c.FormValue("content"),
				PhotoIcon:   c.FormValue("photo-link"),
				BannerImage: c.FormValue("banner-link"),
			}

			err = h.AdminService.UpdatePost(uint(id), post)
			if err != nil {
				return c.JSON(http.StatusNotFound, err)
			}
		}
		return c.JSON(http.StatusMethodNotAllowed, r.Method)
	}
}

func (h *AdminHandler) AdminDeletePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		if r.Method == http.MethodDelete {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				return err
			}
			err = h.AdminService.DeletePost(uint(id))
			if err != nil {
				return c.JSON(http.StatusNotFound, err)
			}
		}
		return c.JSON(http.StatusMethodNotAllowed, r.Method)
	}
}

func (h *AdminHandler) AdminGetListOfCategories() echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		if r.Method == http.MethodGet {
			categories, err := h.AdminService.AdminGetAllCategories()
			if err != nil {
				return c.JSON(http.StatusNotFound, err)
			}
			return c.JSON(http.StatusOK, categories)
		}
		return c.JSON(http.StatusMethodNotAllowed, r.Method)
	}
}

func (h *AdminHandler) AdminCreateCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		if r.Method == http.MethodPost {
			newCategory := types.Category{
				Name: c.FormValue("name"),
			}
			err := h.AdminService.CreateCategory(newCategory)
			if err != nil {
				return c.JSON(http.StatusNotFound, err)
			}
			return c.JSON(http.StatusCreated, newCategory.Name)
		}
		return c.JSON(http.StatusMethodNotAllowed, r.Method)
	}
}

func (h *AdminHandler) AdminDeleteCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		if r.Method == http.MethodDelete {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				return err
			}
			err = h.AdminService.DeleteCategory(uint(id))
			if err != nil {
				return c.JSON(http.StatusNotFound, err)
			}
		}
		return c.JSON(http.StatusMethodNotAllowed, r.Method)
	}
}
