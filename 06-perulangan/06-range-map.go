package main

import "fmt"

func main() {
	// map
	user := map[string]int{
		"age":  20,
		"rank": 1,
	}

	// for range iterasi semua data di map
	for k, v := range user {
		// k = key dari map
		// v = value dari key
		fmt.Println(k, "=", v)
	}
}
