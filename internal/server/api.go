package server

import (
	"database/sql"
	"fmt"
	"log"
	"runtime/debug"
	"strconv"

	"github.com/VitJRBOG/contacts_handbook/internal/db"
	"github.com/gofiber/fiber/v2"
)

func index(c *fiber.Ctx) error {
	return c.SendString("/createContact - post-запрос, принимает json в формате " +
		"{\"name\": \"Some Name\", \"phonenumber\": \"+71234567890\"}\n" +
		"/getContacts - get-запрос, отправлять в формате /getContacts/1\n" +
		"/changeContact - post-запрос, отправлять в формате /changeContact/1, с передачей " +
		"json в формате {\"name\": \"Some Name\", \"phonenumber\": \"+71234567890\"}\n" +
		"/deleteContact - post-запрос, отправлять в формате /deleteContact/1\n")
}

func createContact(c *fiber.Ctx, dbase *sql.DB) error {
	var contact db.Contact

	if err := c.BodyParser(&contact); err != nil {
		log.Printf("%s\n%s\n", err, debug.Stack())
		return c.Status(400).SendString(err.Error())
	}

	_, _, err := contact.InsertInto(dbase)
	if err != nil {
		log.Printf("%s\n%s\n", err, debug.Stack())
		return err
	}

	contacts, err := contact.SelectFrom(dbase)
	if err != nil {
		return err
	}

	if len(contacts) == 0 {
		return c.SendString(fmt.Sprintf("createContact: contact '%s' was not created",
			contact.Name))
	}

	return c.SendString(fmt.Sprintf("CREATED\nID: %d\nName: %s\nPhoneNumber: %s\n\n",
		contacts[0].ID, contacts[0].Name, contacts[0].PhoneNumber))
}

func getContacts(c *fiber.Ctx, dbase *sql.DB) error {
	var contact db.Contact

	contacts, err := contact.SelectFrom(dbase)
	if err != nil {
		log.Printf("%s\n%s\n", err, debug.Stack())
		return err
	}

	if len(contacts) == 0 {
		return c.SendString("getContacts: no contacts found")
	}

	s := ""
	for _, item := range contacts {
		s += fmt.Sprintf("ID: %d\nName: %s\nPhoneNumber: %s\n\n",
			item.ID, item.Name, item.PhoneNumber)
	}

	return c.SendString(s)
}

func changeContact(c *fiber.Ctx, dbase *sql.DB) error {
	var contact db.Contact

	if c.Params("id") != "" {
		var err error
		contact.ID, err = strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Printf("%s\n%s\n", err, debug.Stack())
			return err
		}
	} else {
		return c.SendString("changeContact: No 'id' value specified")
	}

	contacts, err := contact.SelectFrom(dbase)
	if err != nil {
		return err
	}

	if len(contacts) == 0 {
		return c.SendString("changeContact: contact with the specified 'id' was not found")
	}

	contact = contacts[0]
	if err := c.BodyParser(&contact); err != nil {
		log.Printf("%s\n%s\n", err, debug.Stack())
		return c.Status(400).SendString(err.Error())
	}

	_, _, err = contact.Update(dbase)
	if err != nil {
		log.Printf("%s\n%s\n", err, debug.Stack())
		return err
	}

	return c.SendString(fmt.Sprintf("CHANGED\nID: %d\nName: %s\nPhoneNumber: %s\n\n",
		contact.ID, contact.Name, contact.PhoneNumber))
}

func deleteContact(c *fiber.Ctx, dbase *sql.DB) error {
	var contact db.Contact

	if c.Params("id") != "" {
		var err error
		contact.ID, err = strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Printf("%s\n%s\n", err, debug.Stack())
			return err
		}
	} else {
		return c.SendString("deleteContact: No 'id' value specified")
	}

	contacts, err := contact.SelectFrom(dbase)
	if err != nil {
		return err
	}

	if len(contacts) == 0 {
		return c.SendString("deleteContact: contact with the specified 'id' was not found")
	}

	_, _, err = contacts[0].Delete(dbase)
	if err != nil {
		log.Printf("%s\n%s\n", err, debug.Stack())
		return err
	}

	return c.SendString(fmt.Sprintf("DELETED\nID: %d\nName: %s\nPhoneNumber: %s\n\n",
		contacts[0].ID, contacts[0].Name, contacts[0].PhoneNumber))
}
