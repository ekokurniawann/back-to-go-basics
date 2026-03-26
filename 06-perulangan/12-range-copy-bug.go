package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	bug(nums)
	solution(nums)
}

func bug(nums []int) {

	// range mengembalikan index dan copy dari value
	for _, v := range nums {
		// ini hanya mengubah copy dari value bukan elemen asli di slice
		v = v * 2
	}
	fmt.Println("bug:", nums)
}

func solution(nums []int) {
	// gunakan index agar bisa mengubah elemen asli
	for i := range nums {
		// akses elemen asli dengan nums[i]
		nums[i] = nums[i] * 2
	}

	fmt.Println("solution:", nums)
}
