package main

import (
	"kreditPlus/app/database"
	"log"
)

func main() {
	log.Println("Starting KreditPlus server")

	conn := database.ConnectToDB()
	if conn == nil {
		log.Panic("Can't connect to MySql!")
	}

	
}