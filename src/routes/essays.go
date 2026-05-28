package routes

import (
	"fmt"
	
	"main/util"
	
	"github.com/labstack/echo/v4"
)

// handlers

// POST /essays
func EssayCreateHandler(c echo.Context) error {
	// authentication

	cookie, err := c.Cookie("session_token")
	if err != nil {
		return c.JSON(401, "Missing session token.")
	}

	author := util.SessionTokenVerify(cookie.Value)
	if author == "" {
		return c.JSON(401, "Invalid author session token.")
	}
	
	// fetching inputs
	
	var title string = c.FormValue("title") // required, 4 to 200 chars
	if title == "" {
		return c.JSON(400, "Title is required.")
	} else if !util.InputLongEnough(title, 4, 200) {
		return c.JSON(400, "Title must be between 4-200 characters in length.")
	}

	var content string = c.FormValue("content") // required, 100 to 5000 characters
	if content == "" {
		return c.JSON(400, "Your essay cannot be empty.")
	} else if !util.InputLongEnough(content, 100, 5000) {
		return c.JSON(400, "Your essay must be between 100-5000 characters in length.")
	}

	// processing

	id := util.GenerateRandomID(8)
	language := "english"
	meta := ""

	if !essayCreate(id, language, author, title, content, meta) {
		return c.JSON(500, "Failed to create essay.")
	}
		
	return c.JSON(201, "Essay created successfully!")
}

// utility functions

func essayCreate(id string, language string, author string, title string, content string, meta string) bool {
	_, err := util.DB.Exec("INSERT INTO essays (id, language, author, title, content, meta) VALUES (?, ?, ?, ?, ?, ?)", id, language, author, title, content, meta)
	if err != nil { fmt.Println(err); return false }
	
	return true
}