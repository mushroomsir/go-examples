package main

import "fmt"

type Base struct {
	A string
	B string
}

type Something struct {
	C string
	Base
}

func main() {
	test := Something{
		A: "letter a",
		C: "letter c",
	}
	fmt.Println(test.A)
}
