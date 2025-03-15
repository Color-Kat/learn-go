package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func ping(url string, respCh chan int, errCh chan error) {
	resp, err := http.Get(url)
	if err != nil {
		errCh <- err
		return
	}

	respCh <- resp.StatusCode
}

func main() {
	path := flag.String("file", "url.txt", "path to file with list of urls")
	flag.Parse()

	fmt.Println(*path)

	file, err := os.ReadFile(*path)
	if err != nil {
		panic(err.Error())
	}

	urlSlice := strings.Split(string(file), "\r\n")

	respCh := make(chan int)
	errCh := make(chan error)

	for _, url := range urlSlice {
		go ping(url, respCh, errCh)
	}

	for _, url := range urlSlice {
		fmt.Printf("%s: ", url)

		select {
		case res := <-respCh:
			fmt.Println(res)
		case errRes := <-errCh:
			fmt.Println(errRes)
		}

	}
}
