package main

import (
	"demo/http/configs"
	"demo/http/internal/auth"
	"demo/http/internal/link"
	"demo/http/pkg/database"
	"fmt"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	db := database.NewDB(config)
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(db)

	// --- handlers --- //
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: config,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})
	// --- handlers --- //

	port := "8081"
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Println("Server is listening on port " + port)
	server.ListenAndServe()

}
