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
		return c.HTML(http.StatusBadRequest, "<p>An ID is required!</p>")
	} else if !util.InputLongEnough(id, 4, 64) {
		return c.HTML(http.StatusBadRequest, "<p>ID must be between 4 and 64 characters long.</p>")
	} else if userExists(id) {
		return c.HTML(http.StatusBadRequest, "<p>The ID given already exists.</p>")
	}
	
	name := c.FormValue("name") // optional, defaults to the ID, 2 to 96 characters
	if name == "" {
		name = id
	} else if !util.InputLongEnough(name, 2, 96) {
		return c.HTML(http.StatusBadRequest, "<p>The name must be between 2 and 96 characters long.</p>")
	}

	password := c.FormValue("password") // required, 8 to 64 characters
	if password == "" {
		return c.HTML(http.StatusBadRequest, "<p>A password is required!</p>")
	} else if !util.InputLongEnough(password, 8, 64) {
		return c.HTML(http.StatusBadRequest, "<p>Password must be between 8 and 64 characters long.</p>")
	}
	password = util.Hash(password)

	if signUp(id, name, password) == false {
		return c.HTML(http.StatusInternalServerError, "<p>The server failed to sign you up.</p>")
	}

	c.Response().Header().Set("HX-Trigger", "formSuccess")
	return c.HTML(http.StatusCreated, "<p>Sign up successful! Redirecting...</p>")
}

// POST /login
func LoginHandler(c echo.Context) error {
	// authentication
	id := strings.ToLower(c.FormValue("id"))
	if id == "" {
		return c.HTML(http.StatusBadRequest, "<p>An ID is required.</p>")
	} else if !userExists(id) {
		return c.HTML(http.StatusNotFound, "<p>The provided ID returns no existing user.</p>")
	}

	password := c.FormValue("password")
	if password == "" {
		return c.HTML(http.StatusBadRequest, "<p>A password is required.</p>")
	}
	if !comparePassword(id, password) {
		return c.HTML(http.StatusUnauthorized, "<p>Invalid password.</p>")
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

	c.Response().Header().Set("HX-Trigger", "formSuccess")
	return c.HTML(http.StatusOK, "<p>Login successful! Redirecting...</p>")
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