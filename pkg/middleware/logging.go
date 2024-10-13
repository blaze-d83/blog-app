package middleware_service

import (
	"time"

	"github.com/blaze-d83/blog-app/pkg/logger"
	"github.com/labstack/echo/v4"
)

func loggingMiddleware(logger *logger.CustomLogger) echo.MiddlewareFunc {
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
