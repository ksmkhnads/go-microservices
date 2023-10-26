package main

import (
	"github.com/ksmkhnads/go-microservices/go-microservices/internal/database"
	"github.com/ksmkhnads/go-microservices/go-microservices/internal/server"
	"log"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to create database client: %v", err)
	}
	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
