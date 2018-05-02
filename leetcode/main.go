package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	r := []int{}
	for x, val1 := range nums {
		for y, val2 := range nums {
			if x != y && val1+val2 == target {
				fmt.Println(val1, "", val2)
				r = append(r, x)
			}
		}
	}
	return r
}
