package router

import (
	Controller "aryajipang/controller"
	"aryajipang/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
)

// Register home
func Register(web fiber.Router, session *session.Session, sessionLookup string, db *database.Database) {
	// Home
	web.Get("/", Controller.GetHome(session, db))

	// Auth
	web.Get("/login", Controller.GetLogin(session, db))
	web.Get("/register", Controller.GetRegister(session, db))
	web.Post("/login", Controller.PostLogin(session, db))
	web.Post("/logout", Controller.PostLogout(session, db))

}
