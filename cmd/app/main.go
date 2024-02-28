package main

import "card2go_service/internal/api/server"

func main() {
	app := server.New()

	app.Listen(":8080")
}
