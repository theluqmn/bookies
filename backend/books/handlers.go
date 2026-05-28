// handles API routes for books + input validation

package books

import (
	"strings"

	"main/util"
	
	"github.com/labstack/echo/v4"
)

func AddBookHandler(c echo.Context) error {
	var id string // required, 4 to 96 characters
	q1 := strings.ToLower(c.QueryParam("id"))
	if q1 == "" {
		return c.JSON(400, "an ID is required!")
	} else if !util.InputLongEnough(q1, 4, 96) {
		return c.JSON(400, "ID must be between 4 and 96 characters long")
	} else if bookExists(id) {
		return c.JSON(400, "book already exists!")
	}
	id = q1

	var title string // required, 4 to 96 characters
	q2 := c.QueryParam("title")
	if q2 == "" {
		return c.JSON(400, "the book must have a title!")
	} else if !util.InputLongEnough(q2, 4, 96) {
		return c.JSON(400, "book title must be between 4 and 96 characters long")
	}
	title = q2

	var author string // required, 4 to 96 characters
	q3 := c.QueryParam("author")
	if q3 == "" {
		return c.JSON(400, "the book must have an author!")
	} else if !util.InputLongEnough(q3, 4, 96) {
		return c.JSON(400, "book author must be between 4 and 96 characters long")
	}
	author = q3

	var description string // optional, max 256 characters
	q4 := c.QueryParam("description")
	if q4 == "" {
		q4 = "No description provided"
	} else if !util.InputLongEnough(q4, 0, 256) {
		return c.JSON(400, "description can only be up to 256 characters long")
	}

	if !addBook(id, title, author, description) {
		return c.JSON(500, "failed to add book")
	}
	
	return c.JSON(200, "successfully added book")
}