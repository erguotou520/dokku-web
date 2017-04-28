package routers

import (
	"fmt"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os/exec"
)

// deploy app via tar file
func deployTarApp(c echo.Context) error {
	name := c.Param("name")
	file, ef := c.FormFile("file")
	if ef != nil || file == nil {
		return c.NoContent(http.StatusBadRequest)
	}
	tempfile, et := file.Open()
	if et != nil {
		return c.JSON(http.StatusInternalServerError, et)
	}
	cmd := exec.Command(DokkuPath, "tar:in", name)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	go func() {
		defer stdin.Close()
		io.Copy(stdin, tempfile)
	}()
	out, e := cmd.CombinedOutput()
	if e != nil {
		fmt.Println(string(out), e)
		return c.JSON(http.StatusInternalServerError, e)
	}
	return c.NoContent(http.StatusOK)
}
