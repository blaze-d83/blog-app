package routes

import (
	"github.com/blaze-d83/blog-app/internal/db"
	"github.com/blaze-d83/blog-app/internal/handlers"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
}

// RegisterRoutes sets up the routes for the application.
func RegisterRoutes(e *echo.Echo, dbInstance *db.Database) {
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	// Public routes
	e.GET("/", handlers.Homepage)

	// Admin login routes
	admin := e.Group("/admin")
	admin.GET("/login", handlers.GetLoginPageHandler())
	admin.POST("/login", handlers.ProcessAdminLoginHandler(dbInstance))
}
