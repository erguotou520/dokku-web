package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"server/routers"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))
	e.Use(middleware.Gzip())
	routers.Setup(e)
	// Start as a web server
	e.Logger.Fatal(e.Start(":8000"))
}
