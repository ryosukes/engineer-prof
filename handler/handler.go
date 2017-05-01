package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func IndexPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!!")
	}
}

func ShowUsersList() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "users list")
	}
}

func ShowUsersDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, "users detail"+id)
	}
}
