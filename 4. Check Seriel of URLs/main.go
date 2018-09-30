package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://stackoverflow",
		"http://golang.org",
	}

	for _, link := range links {
		checkLink(link)
	}

}

func checkLink(url string) {
	_, err := http.Get(url) // Blocking call. There is delay
	if err != nil {
		fmt.Println(url, "might be down!")
		return
	}

	fmt.Println(url, "is up!")
}
