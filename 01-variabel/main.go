package main

import "fmt"

func main() {
	// Variable declaration

	// Short declaration (:=) -> type inference
	customerFirstName := "John"

	// var -> digunakan jika butuh zero value
	var customerLastName string
	customerLastName = "Wick"

	// idiomatic untuk nilai langsung
	customerNickName := "BabaYaga"

	fmt.Printf("Hello %s %s (%s)\n", customerFirstName, customerLastName, customerNickName)

	// Multiple variables

	// Grouping meningkatkan keterbacaan
	var (
		coffeeName  = "Arabica Gayo"
		coffeeStock = 25
		coffeePrice = 75000.0
	)

	fmt.Println(coffeeName, coffeeStock, coffeePrice)

	// Mixed type inference
	totalPrice, isDiscountApplied, discountRate, orderNote := 150000.0, true, 0.1, "Discount applied"
	fmt.Println(totalPrice, isDiscountApplied, discountRate, orderNote)

	// Zero value
	// Semua variable punya default value
	var customerNote string
	fmt.Println("Customer note:", customerNote) // ""

	// Constant
	// Nilai tetap, tidak bisa diubah (compile time)
	const appName = "Coffee ERP"
	fmt.Println("App:", appName)

	// Blank identifier
	// Digunakan untuk mengabaikan nilai yang tidak diperlukan
	customerName, _ := "John Wick", "ignored"
	fmt.Println(customerName)

	// Pointer Dasar
	// Referensi ke alamat memori
	customer := "Alice"
	customerPtr := &customer

	fmt.Println("Address:", customerPtr)
	fmt.Println("Value:", *customerPtr)

	// Keyword new()
	// Alokasi memori + return pointer
	newCustomerPtr := new(string)
	fmt.Println("New pointer value:", *newCustomerPtr)

	// Short declaration reuse
	x := 10
	x, y := 20, 30 // x di-update, y variable baru
	fmt.Println(x, y)

	/*
		// Shadowing
			a := 10

			if true {
				a, b := 20, 30
				fmt.Println(a, b)
			}

			fmt.Println(a)
	*/
}
