package routes

import (
	"fmt"
	"strings"

	"main/util"
	
	"github.com/labstack/echo/v4"
)

// route handlers

// POST /books
func AddBookHandler(c echo.Context) error {
	id := strings.ToLower(c.FormValue("id")) // required, 4 to 96 characters
	if id == "" {
		return c.JSON(400, "an ID is required!")
	} else if !util.InputLongEnough(id, 4, 96) {
		return c.JSON(400, "ID must be between 4 and 96 characters long")
	} else if bookExists(id) {
		return c.JSON(400, "book already exists!")
	}

	title := c.FormValue("title") // required, 4 to 96 characters
	if title == "" {
		return c.JSON(400, "the book must have a title!")
	} else if !util.InputLongEnough(title, 4, 96) {
		return c.JSON(400, "book title must be between 4 and 96 characters long")
	}

	author := c.FormValue("author") // required, 4 to 96 characters
	if author == "" {
		return c.JSON(400, "the book must have an author!")
	} else if !util.InputLongEnough(author, 4, 96) {
		return c.JSON(400, "book author must be between 4 and 96 characters long")
	}

	description := c.FormValue("description") // optional, max 256 characters
	if description == "" {
		description = "No description provided"
	} else if !util.InputLongEnough(description, 0, 256) {
		return c.JSON(400, "description can only be up to 256 characters long")
	}

	if !addBook(id, title, author, description) {
		return c.JSON(500, "failed to add book")
	}
	
	return c.JSON(200, "successfully added book")
}

// utility functions

func addBook(id string, title string, author string, description string) bool {
	_, err := util.DB.Exec("INSERT INTO books (id, title, author, description) VALUES (?, ?, ?, ?)", id, title, author, description)
	if err != nil { fmt.Println(err); return false }
	
	return true
}

func bookExists(id string) bool {
	var count int
	err := util.DB.QueryRow("SELECT COUNT(*) FROM books WHERE id = ?", id).Scan(&count)
	if err != nil { return false }

	return count > 0
}