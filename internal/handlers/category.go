package handlers

import (
	"net/http"
	"strconv"

	"github.com/blaze-d83/blog-app/internal/db"
	"github.com/blaze-d83/blog-app/types"
	"github.com/labstack/echo/v4"
)

func GetCategories(dbInstance *db.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		var categories []types.Category
		if err := dbInstance.DB.Find(&categories).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, categories)
	}
}

func CreateCategory(dbInstance *db.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		category := new(types.Category)
		if err := c.Bind(category); err != nil {
			return err
		}
		if err := dbInstance.DB.Create(category).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, category)
	}
}

func UpdateCategory(dbInstance *db.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		category := new(types.Category)
		if err := dbInstance.DB.First(&category, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, "Category not found")
		}
		if err := c.Bind(category); err != nil {
			return err
		}
		dbInstance.DB.Save(&category)
		return c.JSON(http.StatusOK, category)
	}
}

func DeleteCategory(dbInstance *db.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		category := new(types.Category)
		if err := dbInstance.DB.First(&category, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, "Category not found")
		}
		dbInstance.DB.Delete(&category)
		return c.NoContent(http.StatusNoContent)
	}
}
