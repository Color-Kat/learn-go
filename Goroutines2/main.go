package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func requestWaitGroup(wg *sync.WaitGroup) {
	response, _ := http.Get("https://vpn.colorbit.ru/")
	fmt.Println(response.StatusCode)
	wg.Done()
}

func waitGroup() {
	startTime := time.Now()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go requestWaitGroup(&wg)
	}

	wg.Wait()

	fmt.Println(time.Since(startTime))
}

func requestChannel(codeCh chan int) {
	response, _ := http.Get("https://vpn.colorbit.ru/")
	codeCh <- response.StatusCode
}

func channel() {
	codeCh := make(chan int)
	for i := 0; i < 10; i++ {
		go requestChannel(codeCh)
	}
	for res := range codeCh {
		fmt.Println(res)
	}
}

func main() {
	//waitGroup()

	channel()
}
