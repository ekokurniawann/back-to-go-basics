package main

import (
	"fmt"
)

// struct
type Student struct {
	Name  string
	Grade int
}

// embedded struct
type Person struct {
	Name string
	Age  int
}

type CollegeStudent struct {
	Person // embedded
	Grade  int
}

// field tabrakan (ambigu)
type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	A
	B
}

// nested struct
type Nested struct {
	Info struct {
		Name string
		Age  int
	}
}

// struct tag (metadata)
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// type alias vs type baru
type MyInt = int
type NewInt int

// memory (biaya copy)
type Big struct {
	Data [1000]int
}

func main() {
	basicStruct()
	initStruct()
	valueSemantics()
	pointerSemantics()
	functionStruct()
	methodStruct()
	embeddedStruct()
	fieldCollision()
	anonymousStruct()
	nestedStruct()
	sliceStruct()
	structTag()
	typeAlias()
	memoryNote()
}

func basicStruct() {
	var s Student // zero value
	fmt.Printf("Name: %q, Grade: %d\n", s.Name, s.Grade)

	s.Name = "roger"
	s.Grade = 2
	fmt.Printf("Name: %q, Grade: %d\n", s.Name, s.Grade)
}

func initStruct() {
	// menggunakan nama field
	s1 := Student{Name: "okta", Grade: 3}

	// tergantung urutan field
	s2 := Student{"Budi", 4}
	fmt.Println(s1, s2)
}

func valueSemantics() {
	// value semantics
	s1 := Student{Name: "roger", Grade: 2}
	s2 := s1 // copy, bukan reference

	s2.Name = "Changed"

	fmt.Println(s1.Name) // tidak berubah
	fmt.Println(s2.Name)
}

func pointerSemantics() {
	s1 := Student{Name: "roger", Grade: 2}
	s2 := &s1 // reference ke s1

	s2.Name = "Changed"

	fmt.Println(s1.Name) // ikut berubah
	fmt.Println(s2.Name)

	// dereference (go otomatis bantu)
	(*s2).Grade = 10
	fmt.Println(s1.Grade)
}

func changeValue(s Student) {
	// hanya copy, tidak mengubah aslii
	s.Name = "ValueChanged"
}

func changePointer(s *Student) {
	// mengubah data asli
	s.Name = "PointerChanged"
}

func functionStruct() {
	s := Student{Name: "roger", Grade: 2}

	changeValue(s)
	fmt.Println(s.Name) // tidak berubah

	changePointer(&s)
	fmt.Println(s.Name) // berubah
}

// value receiver (copy)
func (s Student) Print() {
	fmt.Println("Name:", s.Name)
}

// pointer receiver(bisa ubah data)
func (s *Student) setName(name string) {
	s.Name = name
}

func methodStruct() {
	s := Student{Name: "laras"}

	s.Print()
	s.setName("Updated")
	fmt.Println(s.Name)
}

func embeddedStruct() {
	s := CollegeStudent{
		Person: Person{Name: "wahyu", Age: 26},
		Grade:  3,
	}

	// field dari person bisa langsung diakses
	fmt.Println(s.Name)

	// akses eksplisit
	fmt.Println(s.Person.Name)
}

func fieldCollision() {
	c := C{
		A: A{Name: "A"},
		B: B{Name: "B"},
	}

	// fmt.Println(c.Name) // error: ambigu

	fmt.Println(c.A.Name)
	fmt.Println(c.B.Name)
}

func anonymousStruct() {
	s := struct {
		Name string
		Age  int
	}{
		Name: "ricep",
		Age:  21,
	}

	fmt.Println(s)
}

func nestedStruct() {
	var n Nested
	n.Info.Name = "muslim"
	n.Info.Age = 21

	fmt.Println(n)
}

func sliceStruct() {
	students := []Student{
		{Name: "muslim", Grade: 2},
		{Name: "daffa", Grade: 3},
	}

	for _, s := range students {
		fmt.Println(s.Name)
	}
}

func structTag() {
	u := User{Name: "daffa", Age: 21}
	fmt.Println(u) // tag tidak berpengaruh tanpa library
}

func typeAlias() {
	var a MyInt = 10
	var b int = a // valid

	var c NewInt = 10
	// var d int = c // tidak valid

	fmt.Println(a, b, c)
}

func memoryNote() {
	b1 := Big{}
	b2 := b1 // copy besar (mahal)

	_ = b2
}
