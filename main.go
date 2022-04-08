package main

import (
	"github.com/gabrielsscti/ifood-backend-test/pkg/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	server.RunServer("localhost:8080")
}
