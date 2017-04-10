package routers

import (
	"github.com/labstack/echo"
	"net/http"
)

// User user model
type User struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

// Setup setup router
func Setup(e *echo.Echo) {
	// get app initialization status
	e.GET("/api/initialize", fetchInitialize)
	// initialize app
	e.POST("/api/initialize", doInitialize)
	// get user me
	e.GET("/api/me", me)
	// login auth
	e.POST("/api/auth/login", func(c echo.Context) error {
		// username := c.FormValue("username")
		// password := c.FormValue("password")
		return c.String(http.StatusOK, "ok")
	})
	// e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey:  []byte(secret)
	// }))
	// api group
	//apiG := e.Group("/api")
	//apiG.Use(middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte(secret)}))
}
