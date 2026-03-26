package main

import "fmt"

func main() {

outer: // label untuk loop luar
	for i := 0; i < 5; i++ {
		// loop dalam
		for j := 0; j < 5; j++ {
			// jika i sudah 3 maka keluar dari loop yang diberi label "outer"
			if i == 3 {
				break outer
			}
			fmt.Println(i, j)
		}
	}
}
