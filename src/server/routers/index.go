package routers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"server/config"
)

// User user model
type User struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

// Setup setup router
func Setup(e *echo.Echo) {
	// api group
	apiG := e.Group("/api")
	apiG.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(config.Secret),
		ContextKey: "user",
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/api/initialize" || c.Path() == "/api/auth/login" {
				return true
			}
			return false
		},
	}))
	// get app initialization status
	apiG.GET("/initialize", fetchInitialize)
	// initialize app
	apiG.POST("/initialize", doInitialize)
	// get user me
	apiG.GET("/users/me", me)
	// login auth
	apiG.POST("/auth/login", login)
}
