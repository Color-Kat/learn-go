package main

import (
	"demo/http/configs"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("123")
}

func main() {
	_ := configs.LoadConfig()

	router := http.NewServeMux()
	router.HandleFunc("/hello", helloHandler)

	port := "8081"
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Println("Server is listening on port " + port)
	server.ListenAndServe()

}
