package main

import (
	"log"
	"test/handler"
	"test/server"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(server.Server)
	err := srv.Run("8080", handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
