package main

import (
	"fmt"
	"go-starter-backend/internal/config"
)

func main() {

	cfg := config.Load()

	fmt.Println("Db something ", cfg.DBSSLMode)
}
