package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ryosukes/engineer-prof/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/", index.Show())

	e.Logger.Fatal(e.Start(":1323"))
}
