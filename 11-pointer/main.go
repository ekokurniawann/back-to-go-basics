package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	basicPointer()

	// value vs pointer
	x := 10
	updateValue(x)
	fmt.Println(x) // tetap 10

	updatePointer(&x)
	fmt.Println(x) // berubah 100

	// nil safety
	var p *int
	safePointer(p)

	// struct pointer
	user := User{Name: "roger", Age: 20}
	updateUser(&user)
	fmt.Println(user)

	// return pointer
	n := newNumber()
	fmt.Println(*n)
}

func basicPointer() {
	a := 10
	p := &a // ambil address a

	fmt.Println(a)  // value
	fmt.Println(p)  // address
	fmt.Println(*p) // dereference (value dari address)

	*p = 99        // ubah value asli
	fmt.Println(a) // ikut berubah
}

func updateValue(v int) {
	// v = copy dari nilai asli (bukan data asli)
	// perubahan di sini tidak mempengaruhi luar
	v = 100
}

func updatePointer(v *int) {
	// v = alamat dari variabel asli (pointer ke int)
	// *v = nilai asli di alamat tersebut (dereference)
	*v = 100
}

func safePointer(p *int) {
	// nil pointer safety
	if p == nil { // wajib check
		return
	}
	fmt.Println(*p)
}

func updateUser(u *User) {
	u.Age++ // langsung modify struct
	u.Name = "updated " + u.Name
}

func newNumber() *int {
	n := 50
	return &n // aman, Go handle escape analysis
}
