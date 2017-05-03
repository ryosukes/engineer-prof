package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ryosukes/engineer-prof/handler"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/", handler.IndexPage())
	e.GET("/users", handler.ShowUsersList())
	e.GET("/users/:id", handler.ShowUsersDetail())

	e.Logger.Fatal(e.Start(":1323"))
}
