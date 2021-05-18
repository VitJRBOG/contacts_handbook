package main

import (
	"log"
	"runtime/debug"

	"github.com/VitJRBOG/contacts_handbook/internal/config"
	"github.com/VitJRBOG/contacts_handbook/internal/server"
)

func main() {
	serverCfg, err := config.GetServerConfigs()
	if err != nil {
		log.Printf("%s\n%s\n", err, debug.Stack())
	}
	dbConn, err := config.GetDBConnectionData()
	if err != nil {
		log.Printf("%s\n%s\n", err, debug.Stack())
	}
	if (serverCfg != config.ServerCfg{}) && (dbConn != config.DBConn{}) {
		server.Start(serverCfg, dbConn)
	}
}
