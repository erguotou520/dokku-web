package routers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os/exec"
	"server/config"
)

// User user model
type User struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

var (
	// DokkuPath dokku bin path
	DokkuPath = ""
)

// Setup setup router
func Setup(e *echo.Echo) {
	// api group
	apiG := e.Group("/api")
	path, err := exec.LookPath("dokku")
	if err != nil {
		panic(err)
	}
	DokkuPath = path
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

	// app apis
	// list apps
	apiG.GET("/apps", appList)
	// create app
	apiG.POST("/apps", createApp)
	// rename app
	apiG.PUT("/apps/:name", renameApp)
	// destroy app
	apiG.DELETE("/apps/:name", destroyApp)

	// plugin apis
	// add a plugin
	apiG.POST("/plugins", addPlugin)
	// disable|enable|update a plugin
	apiG.PUT("/plugins/:name", updatePlugin)
	// remove a plugin
	apiG.DELETE("/plugins/:name", removePlugin)
}
