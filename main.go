package main

import (
	"Song_API/database"
	"Song_API/internal"
)

func main() {
	database.Connect()
	server := internal.NewServer()
	server.Start()
}
