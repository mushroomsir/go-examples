package main

import "fmt"

func main() {
	height := 5
	switch {
	case height <= 5:
		fmt.Println("Short")
		fallthrough
	case height <= 5:
		fmt.Println("Normal")
	case height > 5:
		fmt.Println("Tall")
	}

	id := 10
	switch id {
	case 10, 12, 14:
		fmt.Println("Even")
	case id, 13, 15:
		fmt.Println("Odd")
	}

	do(21)
	do("hello")
	do(true)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
