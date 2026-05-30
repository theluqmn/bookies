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

	// frontend URLs
	e.File("/", "./web/index.html")
	e.File("/signup", "./web/signup.html")

	// API URLs
	e.POST("/api/signup", routes.SignupHandler)
	e.POST("/api/login", routes.LoginHandler)
	e.POST("/api/essays", routes.EssayCreateHandler)
	e.GET("/api/essays", routes.EssayGetHandler)

	util.LogSuccess("now online")
	e.Logger.Fatal(e.Start(":6969"))
}
