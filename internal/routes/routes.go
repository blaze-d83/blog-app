package routes

import (

	"github.com/blaze-d83/blog-app/internal/handlers"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo,
//	adminHandler *handlers.AdminHandler,
//	publicHandler *handlers.PublicHandler,
	loginHandler *handlers.LoginRepository) {

	store := sessions.NewCookieStore([]byte("secret"))
	e.Use(session.Middleware(store))

	// Public routes
//	e.GET("/posts", publicHandler.GetListOfAllPostsHandler())
//	e.GET("/posts/:id", publicHandler.ViewFullPostHandler())

	// Login routes
	// admin := e.Group("/admin")
	// admin.GET("/login", loginHandler.GetLoginPageHandler())
	// admin.POST("/login", loginHandler.ProcessAdminLoginHandler())

	// Admin post routes
//	admin.GET("/posts", adminHandler.GetListOfPosts())
//	admin.GET("/posts/:id", adminHandler.GetPostToPreview())
//	admin.POST("/posts", adminHandler.CreatePost())
//	admin.PUT("/posts/:id", adminHandler.UpdatePost())
//	admin.DELETE("/posts/:id", adminHandler.DeletePost())

	// Admin categories routes
//	admin.GET("/categories", adminHandler.GetListOfCategories())
//	admin.POST("/categories", adminHandler.CreatePost())
//	admin.DELETE("/categories/:id", adminHandler.DeleteCategory())

}
