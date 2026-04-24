package main

import "fmt"

func main() {
	basicMap()
	nilMap()
	checkExist()
	iterateMap()
	deleteMap()
	referenceMap()
	mapOfSlice()
	sliceOfMap()
	mapCapacity()
}

func basicMap() {
	// deklarasi map kosong
	m := map[string]int{}

	// tambah data
	m["a"] = 1
	m["b"] = 2

	// akses data
	fmt.Println("a:", m["a"])
	// key tidak ada -> default value
	fmt.Println("c:", m["c"]) // 0
}

func nilMap() {
	var m map[string]int // nil

	// fmt.Println(m["a"]) // aman -> 0

	// m["a"] = 1 // error -> panic
	// panic : assignment to entry in nil map

	// solusi: init dulu
	m = make(map[string]int)
	m["a"] = 1

	fmt.Println("fixed:", m["a"])
}

func checkExist() {
	m := map[string]int{
		"a": 1,
	}

	v, ok := m["a"]
	fmt.Println(v, ok) // 1 true

	v, ok = m["b"]
	fmt.Println(v, ok) // 0 false
}

func iterateMap() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for k, v := range m {
		fmt.Println(k, v) // urutan output tidak dijamin
	}
}

func deleteMap() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	delete(m, "a")
	fmt.Println(m) // a hilang
}

func referenceMap() {
	m1 := map[string]int{"a": 1}

	// map bukan value type, tapi reference
	m2 := m1 // copy reference, bukan data
	m2["a"] = 100

	fmt.Println(m1["a"])
}

func mapOfSlice() {
	m := map[string][]int{}

	m["a"] = []int{1, 2}
	m["a"] = append(m["a"], 3)

	fmt.Println(m["a"])
}

func sliceOfMap() {
	data := []map[string]string{
		{"name": "roger"},
		{"name": "danuarta"},
	}

	for _, v := range data {
		fmt.Println(v["name"])
	}
}

func mapCapacity() {
	// 100 bukan size, tapi hint capacity
	m := make(map[string]int, 100)

	m["a"] = 1

	fmt.Println(len(m))
}
