package main

import (
	"fmt"
	"net/http"
)

func fetch(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("%s: error", url)
		return
	}
	c <- fmt.Sprintf("%s: %s", url, resp.Status)
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.example.com",
		"https://www.github.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go fetch(url, c) // start a goroutine for each URL
	}

	for range urls {
		fmt.Println(<-c) // receive from channel
	}
}
