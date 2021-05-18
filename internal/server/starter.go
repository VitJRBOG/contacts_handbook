package server

import (
	"database/sql"
	"fmt"
	"log"
	"runtime/debug"

	"github.com/VitJRBOG/contacts_handbook/internal/config"
	"github.com/VitJRBOG/contacts_handbook/internal/db"
	"github.com/gofiber/fiber/v2"
)

func Start(serverCfg config.ServerCfg, dbConn config.DBConn) {
	var dbase *sql.DB
	if (dbConn != config.DBConn{}) {
		var err error
		dbase, err = db.Connect(dbConn)
		if err != nil {
			log.Printf("%s\n%s\n", err, debug.Stack())
		}
	}
	defer func(dbase *sql.DB) {
		err := dbase.Close()
		if err != nil {
			log.Printf("%s\n%s\n", err, debug.Stack())
		}
	}(dbase)

	app := fiber.New()

	handling(app, dbase)

	s := fmt.Sprintf(":%s", serverCfg.Port)
	err := app.Listen(s)
	if err != nil {
		log.Printf("%s\n%s\n", err, debug.Stack())
	}
	defer func(app *fiber.App) {
		err := app.Shutdown()
		if err != nil {
			log.Printf("%s\n%s\n", err, debug.Stack())
		}
	}(app)
}
