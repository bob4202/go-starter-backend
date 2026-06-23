package main

import (
	"fmt"
	"go-starter-backend/internal/config"
	"go-starter-backend/internal/db"
)

func main() {

	cfg := config.Load()

	db, err := db.Connect(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("DB connected successfully")
}
