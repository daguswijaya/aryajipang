package main

import (
	configuration "aryajipang/config"
	"aryajipang/database"
	"aryajipang/router"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
)

type App struct {
	*fiber.App

	DB      *database.Database
	Session *session.Session
	w       http.ResponseWriter
	r       *http.Response
}

func main() {
	config := configuration.New()
	app := App{
		App:     fiber.New(*config.GetFiberConfig()),
		Session: session.New(config.GetSessionConfig()),
	}

	app.Static("/", "./public")

	web := app.Group("")
	router.Register(web, app.Session, config.GetString("SESSION_LOOKUP"), app.DB)

	log.Fatal(app.Listen("localhost:3000"))
}
