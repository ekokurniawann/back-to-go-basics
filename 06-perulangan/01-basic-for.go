package main

import "fmt"

func main() {
	// for loop klasik: inisialisasi; kondisi; increment
	for i := 0; i < 5; i++ {
		// mencetak nomor iterasi
		fmt.Println("iteration:", i)
	}

	// for loop go 1.22+ range pada integer
	for a := range 5 {
		fmt.Println("iteration:", a)
	}
}
