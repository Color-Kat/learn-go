package main

import (
	"demo/http/configs"
	"demo/http/internal/link"
	"demo/http/internal/user"
	"demo/http/pkg/database"
)

func main() {
	config := configs.LoadConfig()
	database := database.NewDB(config)

	database.AutoMigrate(
		&link.Link{},
		&user.User{},
	)
}
