package middleware

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		ContextKey: "user", // default context key
	})
}
