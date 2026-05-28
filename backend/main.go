package main

import (
	"fmt"

	"main/books"
	"main/users"
	"main/util"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("starting server")

	// initialisation
	_ = util.Init("./database.sqlite")

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// root URL
	e.GET("/", func(c echo.Context) error {
		fmt.Println("Root URL called")
		return c.JSON(200, "Hello from the Bookies server!")
	})

	// API URLs
	e.POST("/signup", users.SignupHandler)
	e.POST("/login", users.LoginHandler)
	e.POST("/books", books.AddBookHandler)

	// HTMX URLs
	e.File("/", "../frontend/index.html")
	e.File("/signup", "../frontend/signup.html")

	fmt.Println("server is now online")
	e.Logger.Fatal(e.Start(":6969"))
}
