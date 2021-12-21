package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logger() echo.MiddlewareFunc {
	config := middleware.LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Format: `{"time":"${time_custom}" "method":"${method}", "uri":"${uri}", "status":${status}}, "error":"${error}"` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}
	return middleware.LoggerWithConfig(config)
}
