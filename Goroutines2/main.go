package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func request(wg *sync.WaitGroup) {
	response, _ := http.Get("https://vpn.colorbit.ru/")
	fmt.Println(response.StatusCode)
	wg.Done()
}

func main() {
	startTime := time.Now()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go request(&wg)
	}

	wg.Wait()

	fmt.Println(time.Since(startTime))
}

func printHi() {
	fmt.Println("Hi from go func")
}
