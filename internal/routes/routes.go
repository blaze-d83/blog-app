package routes

import (
	"github.com/blaze-d83/blog-app/internal/db"
	"github.com/blaze-d83/blog-app/internal/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, dbInstance *db.Database) {

	// Public
	e.GET("/", handlers.Homepage)
	e.GET("/posts", handlers.GetPosts(dbInstance))
	e.GET("/posts/:id", handlers.GetPostByID(dbInstance))

	// Admin routes (Authenticated routes)
	admin := e.Group("/admin")
	admin.POST("/posts", handlers.CreatePost(dbInstance))
	admin.PUT("/posts/:id", handlers.UpdatePost(dbInstance))
	admin.DELETE("/posts/:id", handlers.DeletePost(dbInstance))

	//Category routes
	admin.GET("/categories", handlers.GetCategories(dbInstance))
	admin.POST("/categories", handlers.CreateCategory(dbInstance))
	admin.PUT("/categories/:id", handlers.UpdateCategory(dbInstance))
	admin.DELETE("/categories/:id", handlers.DeleteCategory(dbInstance))

}
