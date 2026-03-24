package main

import (
	"fmt"
)

func main() {
	// Basic constant
	// Konstanta tidak bisa diubah.
	const name = "roger"
	fmt.Println("name:", name)

	// Typed constant
	const age int = 25
	fmt.Printf("age: %d (type: %T)\n", age, age)

	// Untyped constant
	// Belum punya tipe tetap sampai digunakan
	const number = 10

	var a int = number
	var b int64 = number
	var c float64 = number

	fmt.Printf("a: %d (type: %T)\n", a, a)
	fmt.Printf("b: %d (type: %T)\n", b, b)
	fmt.Printf("c: %f (type: %T)\n", c, c)

	// Constant arithmetic
	// Perhitungan konstanta dilakukan saat compile-time bukan runtime
	const x = 10
	const y = 20
	const result = x + y

	fmt.Println("result:", result)

	// Iota (Enum style constant)
	// Iota adalah counter otomatis
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
	)

	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)

	// Iota with skip value
	// Digunakan untuk mengabaikan nilai tertentu
	const (
		A = iota
		_
		C
	)

	fmt.Println("A:", A)
	fmt.Println("C:", C)

	// Overflow constant
	const bigNumber = 300
	// var oveflow uint8 = bigNumber // error karena uint8 max 255
	var safe int = bigNumber
	fmt.Println("safe:", safe)

	// Constant with multiple declaration
	const (
		first  = "Go"
		second = "is"
		third  = "awesome"
	)

	fmt.Println(first, second, third)

	// Constant copy value
	// Value copyMsg otomatis menjadi "Hello"
	const (
		message = "Hello"
		copyMsg
	)

	fmt.Println("message:", message)
	fmt.Println("copyMsg:", copyMsg)

	// Compile time behavior
	// Area dihitung saat compile time bukan saat program berjalan
	const width = 10
	const height = 5
	const area = width * height

	fmt.Println("area:", area)

	// Bug

	// Const tidak bisa runtime
	/*
		const now = time.Now()
		const hanya boleh berisi compile time value.
		time.Now() adalah runtime value.
	*/

	// Powerfull
	const huge = 1 << 100
	fmt.Println("huge constant created")

	// Bit flag
	const (
		Read = 1 << iota
		Write
		Execute
	)

	permission := Read | Write
	fmt.Println("permission:", permission)

	if permission&Read != 0 {
		fmt.Println("can read")
	}

	if permission&Write != 0 {
		fmt.Println("can write")
	}

	if permission&Execute != 0 {
		fmt.Println("can execute")
	}
}
