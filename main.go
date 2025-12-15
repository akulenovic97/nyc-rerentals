package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("NYC re-rentals scraper")

	resp, err := http.Get("https://residenewyork.com/my-properties/?status=open-to-applications")
	if err != nil {
		fmt.Println("Error with fetch: ", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error with read: ", err)
		return
	}
	fmt.Printf("Got %d bytes", len(body))
}