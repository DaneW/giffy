package main

import (
	"os"

	handler "github.com/danew/giffy/handler"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	server := handler.NewServer()
	server.Run(":" + port)
}
