package routes

import (
	"authentication/controllers"
	"authentication/middleware"

	"github.com/labstack/echo"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	r := e.Group("/user")
	r.Use(middleware.JWTMiddleware())
	r.GET("/profile", controllers.Profile)
}
