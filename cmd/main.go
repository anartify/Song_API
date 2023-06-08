package main

import (
	"Song_API/internal"
	"Song_API/pkg/database"
)

func main() {
	database.Connect()
	server := internal.NewServer()
	server.Start()
}
