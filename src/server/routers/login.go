package routers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"server/store"
	"server/utils"
)

func login(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusUnauthorized)
	}
	stored := store.Read()
	if u.Username != stored.Username {
		// fmt.Println("user %s not allowed", u.Username)
		return c.NoContent(http.StatusUnauthorized)
	}
	ePwd := utils.EncryptPassword(u.Password, stored.Salt)
	if stored.HashedPassword == ePwd {
		t, err := sign(u.Username)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
	return c.NoContent(http.StatusUnauthorized)
}
