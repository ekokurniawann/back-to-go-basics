package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(add(2, 3))
	sayHello("roger")

	// multiple return
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Divide:", result)
	}

	// error case
	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Expected error:", err)
	}

	// named return
	a, p := rectangle(5, 3)
	fmt.Println("Area:", a, "Perimeter:", p)

	// variadic
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5))

	// function as value
	f := multiply
	fmt.Println("Multiply:", f(3, 4))

	// function as parameter
	fmt.Println("Operate:", operate(4, 5, multiply))

	// closure
	c := counter()
	fmt.Println("Counter:", c())
	fmt.Println("Counter:", c())
	fmt.Println("Counter:", c())

	// random
	fmt.Println("Random:", randomInRange(1, 10))

	// early return validation
	err = validateAge(15)
	if err != nil {
		fmt.Println("Validation:", err)
	}

	// defer
	process()

	// panic & recover
	safeExecute()

}

// function sederhana dengan parameter & return
func add(a, b int) int {
	return a + b
}

// function tanpa return
func sayHello(name string) {
	fmt.Println("Hello,", name)
}

// multiple return
// fungsi dengan error handling
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

// named return
func rectangle(width, height int) (area int, perimeter int) {
	area = width * width
	perimeter = 2 * (width + height)
	return // implicit return
}

// variadic function
func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// function as value
func multiply(a, b int) int {
	return a * b
}

// function as parameter(high order function)
func operate(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

// closure (stateful function)
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// random function
var randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomInRange(min, max int) int {
	return randomizer.Intn(max-min+1) + min
}

// early return
func validateAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}
	if age < 18 {
		return errors.New("underage")
	}
	return nil
}

// defer
func process() {
	defer fmt.Println("Process finished")
	fmt.Println("Processing...")
}

// panic & recover
func safeExecute() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("something went wrong")
}
