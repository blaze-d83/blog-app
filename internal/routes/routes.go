package routes

import (
	"github.com/blaze-d83/blog-app/internal/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, adminHandler *handlers.AdminHandler, publicHandler *handlers.PublicHandler) {

	// Public Routes
	e.GET("/", publicHandler.GetListOfAllPostsHandler())
	e.GET("/post", publicHandler.ViewFullPostHandler())
	

	// Admin Routes
	e.GET("/admin/login", adminHandler.GetAdminLoginPage())
	e.POST("/admin/login", adminHandler.ProcessHandler())

	// Protected Admin Routes - JWT Middleware
	adminGroup := e.Group("/admin", adminHandler.AdminJWTMiddleware())
	adminGroup.GET("/posts", adminHandler.GetListOfPosts())
	adminGroup.POST("/post", adminHandler.CreatePost())
	adminGroup.GET("/post/:id", adminHandler.GetPostToPreview())
	adminGroup.PUT("/post/:id", adminHandler.UpdatePost())
	adminGroup.DELETE("/post/:id", adminHandler.DeletePost())

	adminGroup.GET("/categories", adminHandler.GetListOfCategories())
	adminGroup.POST("/category", adminHandler.CreateCategory())
	adminGroup.DELETE("/category/:id", adminHandler.DeleteCategory())

}
