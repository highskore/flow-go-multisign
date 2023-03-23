package main

import "github.com/lukaracki/flow-go-multisign/backend/internal/app"

func main() {
	// Create a new instance of the Server struct
	server := app.NewServer()

	// Start the server
	server.Start()
}
