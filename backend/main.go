package main

import (
	"fmt"

	"main/users"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("starting server")

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// root URL
	e.GET("/", func(c echo.Context) error {
		fmt.Println("Root URL called")
		return c.JSON(200, "Hello from the Bookies server!")
	})

	e.POST("/signup", users.SignUpHandler)

	fmt.Println("server is now online")
	e.Logger.Fatal(e.Start(":8080"))
}
