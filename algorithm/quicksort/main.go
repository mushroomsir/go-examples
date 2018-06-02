package main

import "fmt"

func main() {
	q := []int{2, 3, 1}
	quickSort(q)
	fmt.Println(q)
}

func quickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
		quickSort(data[:head])
		quickSort(data[head+1:])
	}
}
