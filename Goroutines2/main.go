package main

import (
	"fmt"
	"net/http"
	"time"
)

func request() {
	response, _ := http.Get("https://google.com")
	fmt.Println(response.StatusCode)
}

func main() {
	startTime := time.Now()

	for i := 0; i < 10; i++ {
		go request()
	}

	time.Sleep(3 * time.Second)

	fmt.Println(time.Since(startTime))
}

func printHi() {
	fmt.Println("Hi from go func")
}
