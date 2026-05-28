// handles routes for users, as well as input validation

package users

import (
	"crypto/rand"
	"encoding/base64"
	"main/util"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

var sessions = make(map[string]string)

func SignupHandler(c echo.Context) error { // POST /signup
	var id string // required, 4 to 64 characters
	q1 := strings.ToLower(c.FormValue("id"))
	if q1 == "" {
		return c.JSON(400, "an ID is required!")
	} else if !util.InputLongEnough(q1, 4, 64) {
		return c.JSON(400, "ID must be between 4 and 64 characters long")
	} else if userExists(q1) {
		return c.JSON(400, "the ID given already exists")
	}
	id = q1

	var name string // optional, defaults to the ID, 2 to 96 characters
	q2 := c.FormValue("name")
	if q2 == "" {
		q2 = id
	} else if !util.InputLongEnough(q2, 2, 96) {
		return c.JSON(400, "name must be between 2 and 96 characters long")
	}
	name = q2

	var password string // required, 8 to 64 characters
	q3 := c.FormValue("password")
	if q3 == "" {
		return c.JSON(400, "a password is required!")
	} else if !util.InputLongEnough(q3, 8, 64) {
		return c.JSON(400, "password must be between 8 and 64 characters long")
	}
	password = util.Hash(q3)

	if signUp(id, name, password) == false {
		return c.JSON(500, "failed to sign up user!")
	}

	return c.JSON(200, "<p>user sign up successful!</p>")
}

func LoginHandler(c echo.Context) error { // POST /login
	// authentication
	var id string
	q1 := strings.ToLower(c.FormValue("id"))
	if q1 == "" {
		return c.JSON(400, "an ID is required!")
	} else if !userExists(q1) {
		return c.JSON(404, "user not found!")
	}
	id = q1

	q2 := c.FormValue("password")
	if q2 == "" {
		return c.JSON(400, "a password is required!")
	}
	if !comparePassword(id, q2) {
		return c.JSON(401, "invalid password!")
	}

	// session token and cookie generation
	b := make([]byte, 32)
	rand.Read(b)
	sessionToken := base64.StdEncoding.EncodeToString(b)
	sessions[sessionToken] = id

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(48 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	c.SetCookie(cookie)

	return c.JSON(200, "user login successful")
}
