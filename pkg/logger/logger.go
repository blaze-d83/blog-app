package logger

import (
	"log/slog"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type CustomLogger struct {
	logger *slog.Logger
}

func NewCustomLogger() *CustomLogger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})
	logger := slog.New(handler)
	return &CustomLogger{
		logger: logger,
	}

}

func (cl *CustomLogger) LogRequest(c echo.Context) {
	requestID := c.Response().Header().Get(echo.HeaderXRequestID)
	cl.logger.Info("HTTP request",
		slog.String("method", c.Request().Method),
		slog.String("path", c.Path()),
		slog.String("remote_addr", c.RealIP()),
		slog.String("request_id", requestID),
	)
}

func (cl *CustomLogger) LogResponse(c echo.Context, status int, duration time.Duration) {
	requestID := c.Response().Header().Get(echo.HeaderXRequestID)
	cl.logger.Info("HTTP response",
		slog.String("method", c.Request().Method),
		slog.String("path", c.Path()),
		slog.Int("status", status),
		slog.Duration("duration", duration),
		slog.String("request_id", requestID),
	)
}

func (cl *CustomLogger) LogError(c echo.Context, err error) {
	requestID := c.Response().Header().Get(echo.HeaderXRequestID)
	cl.logger.Error("Error processing request",
		slog.String("method", c.Request().Method),
		slog.String("path", c.Path()),
		slog.String("request_id", requestID),
		slog.Any("error", err),
	)
}

func (cl *CustomLogger) LogEvent(message string, fields ...slog.Attr) {
	args := make([]any, len(fields))
	for i, field := range fields {
		args[i] = field
	}
	cl.logger.Info(message, args...)
}
