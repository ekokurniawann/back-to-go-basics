package main

import (
	"fmt"
)

// Helper function
func show(name string, value any) {
	fmt.Printf("%s => value: %v | type: %T\n", name, value, value)
}
func main() {
	// Basic data types
	var age int = 25
	var temperature float64 = 36.6
	var isStudent bool = true
	var name string = "roger"

	show("age", age)
	show("temperature", temperature)
	show("isStudent", isStudent)
	show("name", name)

	// Type inference
	// Menentukan tipe secara otomatis dari nilai
	score := 90
	price := 12.5
	message := "hello roger"

	show("score", score)
	show("price", price)
	show("message", message)

	// Type conversion
	// tipe berbeda tidak bisa langsung di operasikan (tidak implicit conversion)
	var a int = 10
	var b int64 = 20

	// error jika langsung dijumlahkan
	// result := a + b

	result := int64(a) + b
	show("result", result)

	// Interger types
	var small int8 = 100
	var medium int32 = 200000
	var big int64 = 9000000000

	show("small", small)
	show("medium", medium)
	show("big", big)

	// Overflow example
	// Overflow terjadi ketika nilai melewati batas tipe data.
	var max uint8 = 255
	show("max before", max)
	max++
	show("max after overflow", max)

	// Bytes vs Rune
	/*
		byte = uint8 -> untuk data mentah
		rune = int32 -> untuk karakter unicode
	*/

	var letter byte = 'A'
	var character rune = '界'

	show("letter", letter)
	show("character", character)

	fmt.Printf("letter as char: %c\n", letter)
	fmt.Printf("character as char: %c\n", character)

	// Zero value
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string

	show("zeroInt", zeroInt)
	show("zeroFloat", zeroFloat)
	show("zeroBool", zeroBool)
	show("zeroString", zeroString)

	// Nil value
	// nilai kosong untuk tipe reference seperti pointer,map,slice,dll.
	var pointer *int
	show("pointer", pointer)

	// Pointer basic
	number := 42
	ptr := &number

	show("number", number)
	show("pointer address", ptr)
	show("pointer value", *ptr)

	// bug
	// type mismatch
	var x int = 5
	var y int64 = 10

	// fmt.Println(x + y) // compile erro
	fmt.Println("solution:", int64(x)+y)

	// float precision problem
	// Constant calculation(compile time)
	// go menghitung angka mentah dengan presisi sangat tinggi, saat kode di compile sebelum dimasukan ke ata float64
	fmt.Println(0.1 + 0.2)

	// Varibel calculation (runtime)
	// Ketika angka dimasukan ke variabel, mereka dipaksa menjai tipe float64.
	// karena 0.1 & 0.2 tidak bisa diubah ke biner dengan sempurna, terjadi pembulatan kecil yang menumpuk saat dijumlahkan.
	g := 0.1
	h := 0.2
	fmt.Println(g + h)

	// String immutable
	text := "hello"
	fmt.Println(text)

	// text[0] = 'H' // error
	fmt.Println("strings in go cannot be modified directly")

	// Solution
	// Convert to []byte
	bytes := []byte(text)
	bytes[0] = 'H'
	text = string(bytes)

	fmt.Println("solution with []byte:", text)

	// Convert to []rune (safe for unicode)
	unicodeText := "🤖ello"
	runes := []rune(unicodeText)
	runes[0] = 'H'
	unicodeText = string(runes)

	fmt.Println("solution with []rune:", unicodeText)
}
