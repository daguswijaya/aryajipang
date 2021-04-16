package controller

import (
	"aryajipang/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
)

// GetHome
func GetHome(session *session.Session, db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Render("views/pages/home", fiber.Map{})
		if err != nil {
			return c.SendStatus(500)
		}
		return err
	}
}

// GetLogin
func GetLogin(session *session.Session, db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Render("views/pages/login", fiber.Map{})
		if err != nil {
			return c.SendStatus(500)
		}
		return err
	}
}

//PostLogin
func PostLogin(session *session.Session, db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Login")
	}
}

//GetRegister
func GetRegister(session *session.Session, db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Render("views/pages/register", fiber.Map{})
		if err != nil {
			return c.SendStatus(500)
		}
		return err
	}
}

//PostRegister
func PostRegister(session *session.Session, db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Register")
	}
}

//PostLogout
func PostLogout(session *session.Session, db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Logged out")
	}
}
