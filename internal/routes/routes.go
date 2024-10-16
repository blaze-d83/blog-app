package routes

import (
	"net/http"

	"github.com/blaze-d83/blog-app/internal/handlers"
	"github.com/blaze-d83/blog-app/pkg/middleware"
	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo,
	adminHandler *handlers.AdminHandler,
	publicHandler *handlers.PublicHandler,
	m middleware.Middleware) {

	// Catch and log unregistered routes
	e.Any("/*", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Not found"})
	})

	// Static files
	e.Static("/static", "./static/")

	// Public Routes
	e.GET("/", publicHandler.Homepage())
	e.GET("/about", publicHandler.About())
	e.GET("/support", publicHandler.Support())
	e.GET("/events", publicHandler.Events())
	e.GET("/posts", publicHandler.GetListOfAllPostsHandler())
	e.GET("/post", publicHandler.ViewFullPostHandler())

	// Admin Routes
	e.GET("/admin/login", adminHandler.LoginPage())
	e.POST("/admin/login", adminHandler.ProcessLogin())

	// Protected Admin Routes - JWT Middleware
	adminGroup := e.Group("/admin", m.AdminJWTMiddleware())
	adminGroup.GET("/dashboard", adminHandler.AdminDashboard())
	adminGroup.GET("/posts", adminHandler.GetListOfPosts())
	adminGroup.POST("/post", adminHandler.CreatePost())
	adminGroup.GET("/post/:id", adminHandler.GetPostToPreview())
	adminGroup.PUT("/post/:id", adminHandler.UpdatePost())
	adminGroup.DELETE("/post/:id", adminHandler.DeletePost())

	adminGroup.GET("/categories", adminHandler.GetListOfCategories())
	adminGroup.POST("/category", adminHandler.CreateCategory())
	adminGroup.DELETE("/category/:id", adminHandler.DeleteCategory())

}
