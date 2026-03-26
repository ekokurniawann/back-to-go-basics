package main

import "fmt"

func main() {
	nums := []int{5, 6, 7}

	for i := range nums {
		fmt.Println("index:", i)
	}
}
