// handles API routes for books + input validation

package books

import (
	"strings"
	
	"github.com/labstack/echo/v4"
)

func AddBookHandler(c echo.Context) error {
	var id string
	q1 := strings.ToLower(c.QueryParam("id"))
	if q1 == "" {
		return c.JSON(400, "id is required")
	}

	var title string

	var author string

	var description string
	
	return c.JSON(200, "new book added!")
}