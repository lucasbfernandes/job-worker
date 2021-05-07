package main

import (
	"flag"
	"log"
	"os"
	"server/internal/controllers"
)

func main() {
	appPort := flag.Int("port", 8080, "application port")
	flag.Parse()

	server, err := controllers.NewServer()
	if err != nil {
		log.Printf("failed to create server: %s\n", err)
		os.Exit(1)
	}

	err = server.Start(*appPort)
	if err != nil {
		log.Printf("failed to start server: %s\n", err)
		os.Exit(1)
	}
}
