package routers

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"server/store"
	"server/utils"
	"time"
)

func fetchInitialize(c echo.Context) error {
	return c.JSON(http.StatusOK, false)
}

func doInitialize(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusUnauthorized)
	}
	// encrypt password
	salt := utils.GenerateSalt()
	ePwd := utils.EncryptPassword(u.Password, salt)
	store.Update(store.Config{
		ID:             1,
		Username:       u.Username,
		Salt:           salt,
		HashedPassword: ePwd,
		Role:           "admin",
	})

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = u.Username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	var secret = os.Getenv("SECRET")
	if secret == "" {
		secret = "sdfgsdfg7867df6asdfas"
	}
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
