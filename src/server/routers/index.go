package routers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os/exec"
	"server/config"
)

// User user model
type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Domain   string `json:"domain" form:"domain"`
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
	// operate app
	apiG.PUT("/apps/ps/:name", actionApp)
	// get app environments
	apiG.GET("/apps/:name/envs", getEnvironment)
	// add app environment
	apiG.PUT("/apps/:name/addEnv", addEnvironment)
	// unset app environment
	apiG.PUT("/apps/:name/removeEnv", removeEnvironment)
	// get app logs
	apiG.GET("/apps/:name/logs", getAppLogs)
	// get app git histories
	apiG.GET("/apps/:name/git-logs", getAppActivities)
	// deploy via tar upload
	apiG.POST("/apps/:name/tar-deploy", deployTarApp)

	// plugin apis
	// get plugin list
	apiG.GET("/plugins", getPlugins)
	// get installed plugin list
	apiG.GET("/installed-plugins", getInstalledPlugins)
	// add a plugin
	apiG.POST("/plugins", addPlugin)
	// disable|enable|update a plugin
	apiG.PUT("/plugins/:name", updatePlugin)
	// toggle plugin status
	apiG.PUT("/plugins/:name/status", togglePluginStatus)
	// remove a plugin
	apiG.DELETE("/plugins/:name", removePlugin)
}
