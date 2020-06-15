package main

import (
	"log"
	"quotes-rest-api/server"
)

func main() {

	log.Fatalln(server.InitializeServer())

}