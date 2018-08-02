package main

import "fmt"

func main() {
	test2()
	//test1()
}

type Student struct {
	Name string
	Age  int
}

func test2() {
	m := make(map[string]*Student)
	stus := []Student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus { // 内存地址相同，每次循环覆盖旧值
		m[stu.Name] = &stu // 同一块内存，所以内存地址相同,
	}
	// 显示m中的内容
	for k, v := range m {
		fmt.Println(k, v) // 都显示  {Name: "wang", Age: 22},
	}
}

// fmt.Printf("%p\n", &stu)
// 		fmt.Println(&stu)

// a := TestType1{Name: "x"}
// b := a
// a.Name = "y"
// fmt.Println(b)

// TestType1 ...
type TestType1 struct {
	Name string
}

func test1() {
	var array [3]TestType1
	for _, e := range array {
		e.Name = "foo"
	}
	for _, e := range array {
		fmt.Println("result: ", e.Name)
	}
}

// 	array := []*TestType1{&TestType1{Name: ""}}
