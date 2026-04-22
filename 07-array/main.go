package main

import (
	"fmt"
)

func main() {
	basic()
	valueCopy()
	passByValue()
	pointerExample()
	typeDifference()
	multidimensional()
	looping()
	edgeCase()
}

func basic() {
	// deklarasi array (fixed size)
	var names [4]string

	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names)
	fmt.Println("length:", len(names))
}

func valueCopy() {
	// array di copy, bukan reference
	a := [3]int{1, 2, 3}
	b := a // copy

	b[0] = 99

	fmt.Println("a:", a) // tidak berubah
	fmt.Println("b:", b) // berubah
}

// pass by value
func passByValue() {
	arr := [3]int{1, 2, 3}

	change(arr)
	fmt.Println("after function:", arr) // tidak berubah
}

func change(a [3]int) {
	a[0] = 100
}

// pointer(no copy)
// gunakan pointer jika ingin ubah asli
func pointerExample() {
	arr := [3]int{1, 2, 3}

	changeWithPointer(&arr)
	fmt.Println("after pointer:", arr) // berubah

}

func changeWithPointer(a *[3]int) {
	a[0] = 999
}

// type difference
// [3]int tidak sama dengan [4]int
func typeDifference() {
	a := [3]int{1, 2, 3}
	fmt.Println(a)

	// Error
	// assign(a)
}

// func assign(a [4]int){}

// multidimensional
// array di dalam array
func multidimensional() {
	// array berisi 2 elemen,dimana setiap elemennya adalah array berisi 3 int
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(matrix)
}

func looping() {
	fruits := [4]string{"apple", "grape", "banana", "melon"}

	// for biasa (pakai index)
	for i := 0; i < len(fruits); i++ {
		fmt.Println(i, fruits[i])
	}

	// range
	for i, v := range fruits {
		fmt.Println(i, v)
	}

	// hanya value
	for _, v := range fruits {
		fmt.Println("value:", v)
	}
}

func edgeCase() {
	// auto size
	// ukuran dihitung otomatis
	arr := [...]int{1, 2, 3, 4}
	fmt.Println(arr)

	// zero value
	var nums [3]int
	fmt.Println(nums) // [0 0 0]

	// out of bonds -> panic
	// fmt.Println(arr[10])
}
