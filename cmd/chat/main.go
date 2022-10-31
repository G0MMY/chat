package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/G0MMY/chat/persistence"
	"github.com/G0MMY/chat/server"
)

func main() {
	// change that
	time.Sleep(5 * time.Second)
	config := persistence.DbConfig{
		Name: os.Getenv("DB_NAME"),
		Host: os.Getenv("DB_HOST"),
		User: os.Getenv("DB_USER"),
		Pwd:  os.Getenv("DB_PWD"),
		Port: os.Getenv("DB_PORT"),
	}

	conn, err := persistence.NewConnection(&config)
	if err != nil {
		log.Println(err)
		return
	}

	routes := server.CreateRoutes(conn)

	log.Println("running")
	log.Fatal(http.ListenAndServe(":8080", routes))
}
