package routes

import (
	"github.com/blaze-d83/blog-app/internal/db"
	"github.com/blaze-d83/blog-app/internal/handlers"
	"github.com/blaze-d83/blog-app/internal/middleware"
	"github.com/labstack/echo/v4"
)

// RegisterRoutes sets up the routes for the application.
func RegisterRoutes(e *echo.Echo, dbInstance *db.Database) {
	// Public routes
	e.GET("/", handlers.Homepage)
	e.GET("/posts", handlers.GetPosts(dbInstance))
	e.GET("/posts/:id", handlers.GetPostByID(dbInstance))

	// Admin routes
	admin := e.Group("/admin")
	admin.GET("/login", handlers.AdminLoginHandler(dbInstance)) // Route for the login page
	admin.POST("/login", handlers.AdminLoginHandler(dbInstance)) // Route for login POST requests

	// Apply authentication middleware to all admin routes except login
	admin.Use(middleware.AuthMiddleware)

	// Admin CRUD routes
	admin.POST("/posts", handlers.CreatePost(dbInstance))
	admin.PUT("/posts/:id", handlers.UpdatePost(dbInstance))
	admin.DELETE("/posts/:id", handlers.DeletePost(dbInstance))

	// Category routes
	admin.GET("/categories", handlers.GetCategories(dbInstance))
	admin.POST("/categories", handlers.CreateCategory(dbInstance))
	admin.PUT("/categories/:id", handlers.UpdateCategory(dbInstance))
	admin.DELETE("/categories/:id", handlers.DeleteCategory(dbInstance))
}

