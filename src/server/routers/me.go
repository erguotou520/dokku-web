package routers

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
)

func me(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	domain := claims["domain"].(string)
	return c.JSON(http.StatusOK, map[string]string{
		"domain": domain,
	})
}
