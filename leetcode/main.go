package main

import (
	"fmt"
)

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}
func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}
func main() {
	s := NewSlice()
	defer s.Add(1).Add(2)
	s.Add(3)
}

// func main() {
// 	fmt.Println(fun(10))
// 	c := 1
// 	var a *int
// 	a = &c

// 	//fmt.Println(twoSum([]int{3, 2, 4}, 6))
// }
// func fun(x int) int {
// 	if x == 1 {
// 		return 1
// 	}
// 	return x + fun(x-1)
// }
// func twoSum(nums []int, target int) []int {
// 	r := []int{}
// 	for x, val1 := range nums {
// 		for y, val2 := range nums {
// 			if x != y && val1+val2 == target {
// 				fmt.Println(val1, "", val2)
// 				r = append(r, x)
// 			}
// 		}
// 	}
// 	return r
// }
