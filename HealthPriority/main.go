package main

import (
	"healthpriority.com/connection"
	"healthpriority.com/models"
	"healthpriority.com/routes"
	"healthpriority.com/server"
)

func main() {

	conn := connection.GetConnection()

	models.Migrate(conn)

	router := routes.Init()
	server.Init(router)
}
