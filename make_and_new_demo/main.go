package main

import "fmt"

func main() {
	// make用于创建切片,map,channel
	slices := make([]int,3)// 长度为3的切片
	m := make(map[string]int)// 空的map
	ch := make(chan int,2)// 缓冲区为2的channel
	fmt.Println(slices,m,ch) // [0 0 0] map[] 0x1316a18be180
	// new返回的是指针
	ptr := new(int)
	fmt.Println(ptr)// 0xc00001a0a8（地址）
	fmt.Println(*ptr)// 0
	
	// new一个结构体,返回结构体指针
	type User struct {
		Name string
		Age int
	}
	u := new(User)
	u.Name = "张三"
	u.Age = 18
	fmt.Println(u)// &{张三 18}
	fmt.Println(*u)// {张三 18}

}