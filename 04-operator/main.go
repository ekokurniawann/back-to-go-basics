package main

import "fmt"

func main() {
	// Arithmetic operators
	a := 10
	b := 3

	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	/*
		10 / 3 = 3
		Bukan 3.33 karena int / int = int
	*/

	// Float divison
	result := float64(a) / float64(b)
	fmt.Println("float division = ", result)

	// Comparasion operators
	// Hasil operator perbandingan adalah boolean.
	fmt.Println("a == b", a == b)
	fmt.Println("a != b", a != b)
	fmt.Println("a > b", a > b)
	fmt.Println("a < b", a < b)
	fmt.Println("a >= b", a >= b)
	fmt.Println("a <= b", a <= b)

	// Logical operators
	// Operator logika digunakan untuk kombinasi boolean
	x := true
	y := false

	fmt.Println("x && y =", x && y)
	fmt.Println("x || y =", x || y)
	fmt.Println("!x =", !x)

	// Short circuit behavior
	/*
		expensiveFunction tidak di panggil.
		ini disebut short circuit evaluation.
	*/
	if false && expensiveFunction() {
		fmt.Println("will never execute")
	}

	// Assignment operators
	/*
		Operator ini mempersingkat penulisan:
		n = n + 5
	*/
	n := 10

	n += 5
	fmt.Println("n += 5:", n)
	n -= 3
	fmt.Println("n -= 3:", n)
	n *= 2
	fmt.Println("n *= 2:", n)

	// Increment & Decrement
	/*
		x := counter++ // Error
		++ bukan expression hanya statement.
	*/
	counter := 0
	counter++
	counter++
	fmt.Println("counter:", counter)
	counter--
	fmt.Println("counter after decrement:", counter)

	// Bitwise operators
	// Bitwise operato bekerja pada level bit. digunakan di programming sistem.
	p := 5
	q := 3

	fmt.Println("p & q =", p&q)
	fmt.Println("p | q=", p|q)
	fmt.Println("p ^ q =", p^q)

	// Shift operators
	fmt.Println("1 << 3 =", 1<<3) // Left shift
	fmt.Println("8 >> 2 =", 8>>2) // Right shift

	// String comparasion
	/*
		Hasilny false
		Karena string dibandingkan secara lexicographic bukan numeric.
	*/
	fmt.Println("10 > 2 ?", "10" > "2")

}

func expensiveFunction() bool {
	fmt.Println("expensive function executed")
	return true
}
