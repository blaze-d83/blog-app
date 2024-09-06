package handlers

import (
	"net/http"
	"strconv"

	"github.com/blaze-d83/blog-app/internal/db"
	"github.com/blaze-d83/blog-app/types"
	"github.com/labstack/echo/v4"
)

func CreatePost(dbInstance *db.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		post := new(types.Post)
		if err := c.Bind(post); err != nil {
			return err
		}

		if err := dbInstance.DB.Create(post).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, post)
	}
}

func UpdatePost(dbInstance *db.Database) echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))
		post := new(types.Post)

		if err := dbInstance.DB.First(&post, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, "Post not found")
		}	

		if err := c.Bind(post); err != nil {
			return err
		}

		dbInstance.DB.Save(&post)
		return c.JSON(http.StatusOK, post)
	}
}

func DeletePost(dbInstance *db.Database) echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))
		post := new(types.Post)

		if err := dbInstance.DB.First(&post, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, "Post not found")
		}

		dbInstance.DB.Delete(&post)
		return c.NoContent(http.StatusNoContent)
	}
}

func GetPosts(dbInstance *db.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		var posts []types.Post
		dbInstance.DB.Find(&posts)
		return c.JSON(http.StatusOK, posts)
	}
}

func GetPostByID(dbInstance *db.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		post := new(types.Post)
		id, _ := strconv.Atoi(c.Param("id"))

		if err := dbInstance.DB.First(&post, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, "Post not found")
		}
		return c.JSON(http.StatusOK, post)
	}
	
}
