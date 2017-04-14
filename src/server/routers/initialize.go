package routers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"server/store"
	"server/utils"
)

func fetchInitialize(c echo.Context) error {
	config, e := store.Read()
	var inited = false
	if e == nil && config.Username != "" && config.HashedPassword != "" {
		inited = true
	}
	return c.JSON(http.StatusOK, inited)
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

	t, err := sign(u.Username)
	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
