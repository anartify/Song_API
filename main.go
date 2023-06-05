package main

import (
	"Song_API/api"
	"Song_API/database"
)

func main() {
	database.Connect()
	server := api.NewServer()
	server.Start()
}
