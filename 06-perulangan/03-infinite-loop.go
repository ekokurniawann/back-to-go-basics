package main

import "fmt"

func main() {
	// inisialisasi variabel counter
	i := 0

	// infinite loop(loop tanpa kondisi)
	for {
		fmt.Println(i)
		i++

		// kondisi untuk menghentikan loop
		if i == 3 {
			// break digunakan unutk keluar dari loop
			break
		}
	}
}
