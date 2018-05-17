package main

import (
	"fmt"
)

func main() {
	s := []byte("")
	s1 := append(s, 'a')
	s2 := append(s, 'b')

	// // //println(s1, ",", s2)
	fmt.Println(s1, ",", s2)
	//fmt.Println(cap(s), len(s))
	// s := "a"
	// foo(s)
}

// func foo(s1 ...interface{}) {
// 	reflect.TypeOf(s1).Kind()
// }

// package main

// import "fmt"

// // const tmpStringBufSize = 32

// // type tmpBuf [tmpStringBufSize]byte

// func main() {

// 	// print((7 / 8))
// 	// var s string
// 	// var b []byte
// 	// buf := tmpBuf{}
// 	// b = buf[:len(s)]
// 	// var b []byte
// 	// println(b)
// 	// print(cap(b))
// 	//s := []byte(strings.Repeat("c", 33))
// 	s := []byte("a")
// 	print(cap(s), len(s))
// 	s1 := append(s, 'a')
// 	s2 := append(s, 'b')
// 	//fmt.Println(s1, ",", s2)
// 	fmt.Println(string(s1), ",", string(s2))
// 	// 	PS D:\go\src\github.com\mushroomsir\go-examples> C:\Go164\bin\go.exe  run -gcflags '-m -l' main.go
// 	// # command-line-arguments
// 	// .\main.go:4: main ([]byte)("") does not escape
// 	// .\main.go:9: main string(s1) does not escape
// 	// .\main.go:9: main string(s2) does not escape
// 	// 00a========b
// }
