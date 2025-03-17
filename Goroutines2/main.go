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
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			requestChannel(codeCh)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(codeCh)
	}()

	for res := range codeCh {
		fmt.Println(res)
	}
}

func sumArr(arr []int, sumCh chan int) {
	sum := 0
	for _, number := range arr {
		sum += number
	}
	sumCh <- sum
}

func sliceSumExample() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	numGoroutines := 3
	chunkSize := len(arr) / numGoroutines

	sumCh := make(chan int, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		fmt.Println(end)

		go sumArr(arr[start:end], sumCh)
	}

	totalSum := 0
	for i := 0; i < numGoroutines; i++ {
		totalSum += <-sumCh
	}
	fmt.Println(totalSum)

}

func main() {
	//waitGroup()
	//channel()
	sliceSumExample()
}
