package routers

import (
	jwt "github.com/dgrijalva/jwt-go"
	"server/config"
	"server/store"
	"time"
)

func sign(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	c, _ := store.Read()
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["admin"] = true
	claims["domain"] = c.Domain
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(config.Secret))
}
