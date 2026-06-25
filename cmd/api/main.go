package main

import (
	"fmt"
	"go-starter-backend/internal/config"
	"go-starter-backend/internal/db"
	"go-starter-backend/internal/server"
	"go-starter-backend/pkg/storage"
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

	store, err := storage.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	serv := server.New(cfg, db, store)

	err = serv.Routes().Run(":" + cfg.AppPort)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Everything is fine")
}
