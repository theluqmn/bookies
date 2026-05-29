package main

import(
	"main/routes"
	"main/util"

	"github.com/labstack/echo/v4"
)

func main() {
	util.Clear()
	util.LogInfo("bookies is starting...")

	// initialisation
	err := util.Init("./database.sqlite")
	if err != nil { util.LogError(err) }

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// API URLs
	e.POST("/signup", routes.SignupHandler)
	e.POST("/login", routes.LoginHandler)
	// e.POST("/books", routes.AddBookHandler)
	e.POST("/essays", routes.EssayCreateHandler)
	e.GET("/essays", routes.EssayGetHandler)

	// frontend URLs
	e.File("/", "../web/index.html")
	e.File("/signup", "../web/signup.html")

	util.LogSuccess("now online")
	e.Logger.Fatal(e.Start(":6969"))
}
