package middleware

import (
	"net/http"
	"time"

	"github.com/blaze-d83/blog-app/pkg/auth"
	"github.com/blaze-d83/blog-app/pkg/logger"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
}

func (m Middleware) AdminJWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("auth_token")
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing or invalid token"})
			}
			claims, err := auth.ValidateJWT(cookie.Value)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
			}
			c.Set("username", claims.Username)
			return next(c)
		}
	}
}

func (m Middleware) LoggingMiddleware(logger *logger.CustomLogger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// Start the time of request
			start := time.Now()

			// Log the incoming request
			logger.LogRequest(c)

			// Call the next handler
			err := next(c)
			if err != nil {
				logger.LogError(c, err)
				return err
			}

			// Log the response with status codes
			duration := time.Since(start)
			status := c.Response().Status
			logger.LogResponse(c, status, duration)

			return nil
		}
	}
}
