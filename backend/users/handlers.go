package users

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func SignUpHandler(c echo.Context) error {
	var id string
	q1 := strings.ToLower(c.QueryParam("id"))
	if q1 == "" {
		return c.JSON(400, "an ID is required!")
	}
	id = q1

	var name string
	q2 := c.QueryParam("name")
	if q2 == "" {
		q2 = id // defaults to the ID if no name is given
	}
	name = q2

	var password string
	q3 := c.QueryParam("password")
	if q3 == "" {
		return c.JSON(400, "a password is required!")
	}
	password = q3

	if signUp(id, name, password) == false {
		return c.JSON(500, "failed to sign up user!")
	}

	return c.JSON(200, "user sign up successful!")
}
