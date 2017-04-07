package routers

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"
	"time"
)

func Setup(e *echo.Echo) {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "sdfgsdfg7867df6asdfas"
	}
	// e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey:  []byte(secret)
	// }))
	// api group
	apiG := e.Group("/api")
	apiG.Use(middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte(secret)}))

	// login auth
	e.Post("/auth/local", func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
	})
}
