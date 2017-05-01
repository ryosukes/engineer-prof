package index

import (
	"net/http"

	"github.com/labstack/echo"
)

func Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!!")
	}
}
