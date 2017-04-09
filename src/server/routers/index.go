package routers

import (
	"github.com/labstack/echo"
	"net/http"
	"os"
)

type Token struct {
	Token string `json:"token"`
}

func Setup(e *echo.Echo) {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "sdfgsdfg7867df6asdfas"
	}
	// get app initialization status
	e.GET("/api/initialize", func(c echo.Context) error {
		return c.JSON(http.StatusOK, false)
	})
	// initialize app
	e.POST("/api/initialize", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &Token{
			Token: "aa",
		})
	})
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
