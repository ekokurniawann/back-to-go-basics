package main

import "fmt"

func main() {
	nums := []int{10, 20, 30}

	for _, v := range nums {
		fmt.Println("value:", v)
	}
}
