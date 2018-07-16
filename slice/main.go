package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// 空切片和 nil 切片的区别在于，nil 切片的指针指向 nil。空切片指向的地址不是nil，指向的是一个内存地址，但是它没有分配任何内存空间，即底层元素包含0个元素。
	var slice []int
	slice2 := make([]int, 0)
	slice3 := []int{}
	fmt.Println(slice == nil, slice2 == nil, slice3 == nil)
	// true false false
	b, _ := json.Marshal(slice)
	fmt.Println(string(b))

	b, _ = json.Marshal(slice2)
	fmt.Println(string(b))
}
