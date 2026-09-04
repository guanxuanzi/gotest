package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	// 方式1 字面量初始化 {最常用}
	p1 := Person{Name:"John",Age:20}
	fmt.Println(p1)
	
	// 方式2 先声明后赋值
	var p2 Person
	p2.Name = "Amy"
	p2.Age = 30
	fmt.Println(p2)

	// 方式3 用new创建指针
	p3 := new(Person)
	p3.Name = "Xiaoming"
	p3.Age = 35
	fmt.Println(p3) 
	fmt.Println(*p3)

	// 方式4 使用&直接获得指针
	p4 := &Person{Name:"Liuqiang",Age:32}
	fmt.Println(p4)
}