package main

import "fmt"

func main() {
	// slice (array dinamis)
	nums := []int{10, 20, 30}

	// for range loop untuk membaca data dalam slice
	for i, v := range nums {
		// i = index (posisi data)
		// v = value (nilai pada index)
		fmt.Println("index:", i, "value:", v)
	}
}
