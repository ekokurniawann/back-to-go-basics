package main

import "fmt"

func main() {
	basic()
	slicing()
	sharedMemory()
	lenCap()
	appendBehavior()
	copyExample()
	threeIndex()
}

func basic() {
	// deklarasi slice
	s := []int{1, 2, 3} // slice

	fmt.Println(s)
	fmt.Println("len:", len(s)) // jumlah data
	fmt.Println("cap:", cap(s)) // kapasitas
}

func slicing() {
	a := []string{"a", "b", "c", "d"}

	// slicing(ambil sebagian data)
	b := a[1:3] // ambil index 1 sampai sebelum 3 (end - start)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

func sharedMemory() {
	a := []string{"a", "b", "c"}

	// slicing share memory
	b := a[:2] // slice[low:high]
	b[0] = "X" // ubah data

	fmt.Println("a:", a) // ikut berubah
	fmt.Println("b:", b)
}

func lenCap() {
	a := []int{1, 2, 3, 4}

	b := a[1:3] // cap = cap(a) - start index
	fmt.Println("a len/cap:", len(a), cap(a))
	fmt.Println("b len/cap:", len(b), cap(b))
}

func appendBehavior() {
	a := []int{1, 2, 3}

	b := a[:2]
	// append data ke slice
	b = append(b, 99) // tambah data

	fmt.Println("a:", a) // bisa berubah
	fmt.Println("b:", b)
}

func copyExample() {
	a := []int{1, 2, 3}

	// copy via append (shortcut)
	b1 := append([]int{}, a...)

	// copy via make + copy (explicit)
	b2 := make([]int, len(a))
	copy(b2, a)

	b1[0] = 10
	b2[0] = 20

	fmt.Println("a:", a)
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
}

func threeIndex() {
	a := []int{1, 2, 3, 4}

	b := a[0:2:2]     // cap dibatasi
	b = append(b, 99) // tidak ganggu a

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
