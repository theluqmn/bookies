package routes

import (
	"net/http"
	"strings"
	"time"
	
	"main/util"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// POST /signup
func SignupHandler(c echo.Context) error {
	id := strings.ToLower(c.FormValue("id")) // required, 4 to 64 characters
	if id == "" {
		return c.JSON(400, "an ID is required!")
	} else if !util.InputLongEnough(id, 4, 64) {
		return c.JSON(400, "ID must be between 4 and 64 characters long")
	} else if userExists(id) {
		return c.JSON(400, "the ID given already exists")
	}
	
	name := c.FormValue("name") // optional, defaults to the ID, 2 to 96 characters
	if name == "" {
		name = id
	} else if !util.InputLongEnough(name, 2, 96) {
		return c.JSON(400, "name must be between 2 and 96 characters long")
	}

	password := c.FormValue("password") // required, 8 to 64 characters
	if password == "" {
		return c.JSON(400, "a password is required!")
	} else if !util.InputLongEnough(password, 8, 64) {
		return c.JSON(400, "password must be between 8 and 64 characters long")
	}
	password = util.Hash(password)

	if signUp(id, name, password) == false {
		return c.JSON(500, "failed to sign up user!")
	}

	return c.JSON(200, "<p>user sign up successful!</p>")
}

// POST /login
func LoginHandler(c echo.Context) error {
	// authentication
	id := strings.ToLower(c.FormValue("id"))
	if id == "" {
		return c.JSON(400, "an ID is required!")
	} else if !userExists(id) {
		return c.JSON(404, "user not found!")
	}

	password := c.FormValue("password")
	if password == "" {
		return c.JSON(400, "a password is required!")
	}
	if !comparePassword(id, password) {
		return c.JSON(401, "invalid password!")
	}

	// cookie generation
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    util.SessionTokenCreate(id),
		Expires:  time.Now().Add(48 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	c.SetCookie(cookie)

	return c.JSON(200, "user login successful")
}

// utility functions

func signUp(id string, name string, password string) bool {
	_, err := util.DB.Exec("INSERT INTO users (id, name, password) VALUES (?, ? ,?);", id, name, password)
	if err != nil { util.LogError(err); return false }

	util.Log("new user signed up: " + id + " " + name)
	
	return true
}

func userExists(id string) bool {
	var count int
	err := util.DB.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", id).Scan(&count)
	if err != nil { util.LogError(err); return false }

	return count > 0
}

func comparePassword(id string, password string) bool {
	var hashed string
	err := util.DB.QueryRow("SELECT password FROM users WHERE id = ?", id).Scan(&hashed)
	if err != nil { util.LogError(err); return false }

	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil { util.LogError(err); return false }

	return true
}