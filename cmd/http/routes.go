package main

import (
	"github.com/kilvish/user-signup/cmd/http/handlers"
	"github.com/labstack/echo"
)

func AddRoutes(e *echo.Echo) {
	e.POST("/user/signup", handlers.SignUpHandler{}.Create)
	e.GET("/user/signin", handlers.SignINHandler{}.Get)
	e.GET("/user/profile", handlers.ProfileHandler{}.Get)
	e.PUT("/user/profile/update", handlers.ProfileHandler{}.Put)
}
