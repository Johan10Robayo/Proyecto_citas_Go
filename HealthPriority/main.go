package main

import (
	"healthpriority.com/routes"
	"healthpriority.com/server"
)

func main() {
	/*conn, err := connection.GetConnection()

	if err != nil {
		log.Panic(err)
	}

	_ = conn

	conn.AutoMigrate(&models.Curso{}, &models.Student{},
	&models.Student_Curso{})*/
	routes.Init()
	server.Init()
}
