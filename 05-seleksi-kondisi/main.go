package main

import "fmt"

func main() {

	ifBasic()
	ifInitialization()
	scopeBug()

	switchBasic()
	switchMultiple()
	switchCondition()
	switchFallthrough()

	typeSwitch()
	nestedCondition()

}

func ifBasic() {
	point := 8

	// If basic
	if point == 10 {
		fmt.Println("perfect score")
	} else if point > 5 {
		fmt.Println("pass")
	} else if point == 4 {
		fmt.Println("almost pass")
	} else {
		fmt.Println("fail")
	}
}

func ifInitialization() {
	point := 8840.0

	// percent hanya hidup di seleksi kondisi(if, else if, else)
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%% perfect!\n", percent)
	} else if percent >= 70 {
		fmt.Printf("%.1f%% good\n", percent)
	} else {
		fmt.Printf("%.1f%% not bad\n", percent)
	}
}

func scopeBug() {
	if x := 10; x > 5 {
		fmt.Println("inside:", x)
	}

	// fmt.Println(x) // Error
}

func switchBasic() {
	point := 6

	// switch basic
	switch point {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
}

func switchMultiple() {
	point := 6

	// multiple case
	switch point {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
}

func switchCondition() {
	point := 6

	// switch condition style
	// replacment untuk if chain panjang
	switch {
	case point == 8:
		fmt.Println("perfect")
	case point < 8 && point > 3:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
}

func switchFallthrough() {
	point := 6

	switch {
	case point > 5:
		fmt.Println("grater than 5")
		fallthrough // fallthrough tidak mengecek kondisi berikutnya
	case point > 3:
		fmt.Println("greater than 3")
	default:
		fmt.Println("small")

	}
}

func typeSwitch() {
	var data any = "hello"

	switch v := data.(type) {
	case int:
		fmt.Println("integer:", v)
	case string:
		fmt.Println("string:", v)
	case bool:
		fmt.Println("boolean:", v)
	default:
		fmt.Println("unknown type")
	}
}

func nestedCondition() {
	point := 10

	// Nested condition
	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else {
			fmt.Println("keep trying")
		}
	}
}
