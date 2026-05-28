package users

import (
	"main/util"
	"strings"

	"github.com/labstack/echo/v4"
)

func SignUpHandler(c echo.Context) error {
	var id string // required, between 4 and 64 characters
	q1 := strings.ToLower(c.QueryParam("id"))
	if q1 == "" {
		return c.JSON(400, "an ID is required!")
	} else if !util.InputLongEnough(q1, 4, 64) {
		return c.JSON(400, "ID must be between 4 and 64 characters long")
	}
	id = q1

	var name string // optional, defaults to the ID, and is between 2 and 96 characters
	q2 := c.QueryParam("name")
	if q2 == "" {
		q2 = id
	} else if !util.InputLongEnough(q2, 2, 96) {
		return c.JSON(400, "name must be between 2 and 96 characters long")
	}
	name = q2

	var password string // required, between 8 and 64 characters
	q3 := c.QueryParam("password")
	if q3 == "" {
		return c.JSON(400, "a password is required!")
	} else if !util.InputLongEnough(q3, 8, 64) {
		return c.JSON(400, "password must be between 8 and 64 characters long")
	}
	password = util.Hash(q3)

	if signUp(id, name, password) == false {
		return c.JSON(500, "failed to sign up user!")
	}

	return c.JSON(200, "user sign up successful!")
}
