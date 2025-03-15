package main

import (
	"demo/http/configs"
	"demo/http/internal/auth"
	"fmt"
	"net/http"
)

func main() {
	config := configs.LoadConfig()

	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: config,
	})

	port := "8081"
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Println("Server is listening on port " + port)
	server.ListenAndServe()

}
