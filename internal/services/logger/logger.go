package service 

import (
	"bytes"
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/color"
)

func GetCustomLoggerConfig(e *echo.Echo) *middleware.LoggerConfig {
	return &middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}",` +
			`"id":"${id}",` +
			`"status":"${status}",` +
			`"latency":"${latency_human}",` +
			`"method":"${method}",` +
			`"path":"${uri}"}` + "\n", // Removed trailing comma here

		CustomTagFunc: func(c echo.Context, buf *bytes.Buffer) (int, error) {
			// Optional Custom Fields if needed
			method := color.Green(c.Request().Method)
			path := color.Cyan(c.Request().URL.Path)

			// Adding colored method and path as custom log entries
			customLog := fmt.Sprintf(`,"custom_method":"%s","custom_path":"%s"`, method, path)
			buf.WriteString(customLog)

			return buf.Len(), nil
		},
		Output: io.MultiWriter(e.Logger.Output()),
	}
}

