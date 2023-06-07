package main

import (
	"Song_API/internal"
	"Song_API/internal/database"
)

func main() {
	database.Connect()
	server := internal.NewServer()
	server.Start()
}
