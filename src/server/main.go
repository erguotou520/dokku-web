package main

import (
	"github.com/labstack/echo"
	"server/routers"
)

func main() {
	e := echo.New()
	routers.Setup(e)
	// Start as a web server
	e.Logger.Fatal(e.Start(":8000"))
}
