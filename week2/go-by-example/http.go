package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response Status", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
