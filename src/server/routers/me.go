package routers

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
)

func me(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
