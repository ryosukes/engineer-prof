package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/ryosukes/engineer-prof/models"
)

type (
	resultUserJson struct {
		User userModel.User
	}
	handler struct {
		userModel userModel.User
	}
)

var h handler

func IndexPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "hello", "Echo")
	}
}

func ShowUsersList() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "users list")
	}
}

func ShowUsersDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		user := h.userModel.Fetch(int64(id))
		return c.JSON(http.StatusOK, resultUserJson{User: user})
	}
}
