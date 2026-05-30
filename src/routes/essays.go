package routes

import (
	"main/util"
	
	"github.com/labstack/echo/v4"
)

type Essay struct {
	ID string `json:"id"`
	Language string `json:"language"`
	Author string `json:"author"`
	Title string `json:"title"`
	Content string `json:"content"`
	Meta string `json:"meta"`
}

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

// GET /essays
func EssayGetHandler(c echo.Context) error {
	var essays []Essay
	rows, err := util.DB.Query("SELECT id, language, author, title, content, meta FROM essays")
	if err != nil { util.LogError(err); return c.JSON(500, "Failed to fetch essays.") }
	defer rows.Close()

	for rows.Next() {
		var essay Essay
		if err := rows.Scan(&essay.ID, &essay.Language, &essay.Author, &essay.Title, &essay.Content, &essay.Meta); err != nil {
			util.LogError(err)
		}
		essays = append(essays, essay)
	}

	return c.JSON(200, essays)
}

// GET /essays/user
func UserEssayGetHandler(c echo.Context) error {
	var user_id string = c.QueryParam("user_id")
	essays, err := essayGetByUser(user_id)
	if err != nil { util.LogError(err); return c.JSON(500, "Failed to fetch essays.") }
	return c.JSON(200, essays)
}

// utility functions

func essayCreate(id string, language string, author string, title string, content string, meta string) bool {
	_, err := util.DB.Exec("INSERT INTO essays (id, language, author, title, content, meta) VALUES (?, ?, ?, ?, ?, ?)", id, language, author, title, content, meta)
	if err != nil { util.LogError(err); return false }
	
	return true
}

func essayGetByUser(user_id string) ([]Essay, error) {
	var essays []Essay
	rows, err := util.DB.Query("SELECT id, language, author, title, content, meta FROM essays WHERE author = ?", user_id)
	if err != nil { return nil, err }
	defer rows.Close()

	for rows.Next() {
		var essay Essay
		if err := rows.Scan(&essay.ID, &essay.Language, &essay.Author, &essay.Title, &essay.Content, &essay.Meta); err != nil {
			return nil, err
		}
		
		essays = append(essays, essay)
	}

	return essays, nil
}