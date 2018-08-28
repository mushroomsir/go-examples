package main

import (
	"fmt"
	"io"
	"log"
	"reflect"
)

type testError struct {
	err string
}

func (t testError) Error() string {
	return t.err
}

func main() {

	var t *testError
	log.Println(t == nil)
	print(t)
	fmt.Println("")
	checkError(t)
}

func checkError(err interface{}) {
	var w io.Writer
	//fmt.Println(w == empty)

	fmt.Printf("(%v, %T)\n", w, w)
	print(err)
	fmt.Println("")
	if err == nil || (reflect.ValueOf(err).Kind() == reflect.Ptr && reflect.ValueOf(err).IsNil()) {
		return
	}
	panic(err)
}
