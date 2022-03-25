package main

import (
	"goCRUD/connection"
	"goCRUD/handlers"
)

func main() {
	connection.Connect()

	handlers.HandleReq()
}
