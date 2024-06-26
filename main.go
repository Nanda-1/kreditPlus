package main

import (
	"kreditPlus/app/database"
	"kreditPlus/app/routes"
	"log"
)

func main() {
	log.Println("Starting KreditPlus server")

	conn := database.ConnectToDB()
	if conn == nil {
		log.Panic("Can't connect to MySql!")
	}

	setup := routes.SetupRoutes()
	setup.Run(":8080")
}
