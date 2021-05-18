package server

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func handling(app *fiber.App, dbase *sql.DB) {
	app.Get("/", func(c *fiber.Ctx) error {
		return index(c)
	})

	app.Post("/createContact", func(c *fiber.Ctx) error {
		return createContact(c, dbase)
	})

	app.Get("/getContacts", func(c *fiber.Ctx) error {
		return getContacts(c, dbase)
	})

	app.Post("/changeContact/:id?", func(c *fiber.Ctx) error {
		return changeContact(c, dbase)
	})

	app.Post("/deleteContact/:id?", func(c *fiber.Ctx) error {
		return deleteContact(c, dbase)
	})
}
