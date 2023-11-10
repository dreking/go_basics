package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"http://google.com", "http://facebook.com", "http://stackoverflow.com", "http://golang.org", "http://amazon.com"}

	c := make(chan string)
	c <- "hello"

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	res, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		c <- link
	}

	if res.StatusCode != 200 {
		fmt.Println(link, "might be down")
		c <- link
	}

	fmt.Println(link, "is up")
	c <- link
}
