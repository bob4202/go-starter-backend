package main

import (
	"fmt"
	"go-starter-backend/internal/config"
	"go-starter-backend/internal/db"
	"go-starter-backend/internal/server"
	"log"
)

func main() {

	cfg := config.Load()

	db, err := db.Connect(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("DB connected successfully")

	serv := server.New(cfg, db)

	err = serv.Routes().Run(":" + cfg.AppPort)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Everything is fine")
}
