package main

import "fmt"

func main() {
	// outer loop
	for i := 0; i < 3; i++ {
		// inner loop
		for j := 0; j < 3; j++ {
			fmt.Println("i:", i, "j:", j)
		}
	}
}
