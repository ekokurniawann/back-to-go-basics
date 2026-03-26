package main

import "fmt"

func main() {
	text := "Go!"

	for i, r := range text {
		fmt.Println("index:", i, "rune:", r)

		// convert dari rune ke string
		fmt.Println("char:", string(r))
	}
}
