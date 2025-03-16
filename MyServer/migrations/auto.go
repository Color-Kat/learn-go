package main

import (
	"demo/http/configs"
	"demo/http/internal/link"
	"demo/http/pkg/db"
)

func main() {
	config := configs.LoadConfig()
	database := database.NewDB(config)

	database.AutoMigrate(&link.Link{})
}
