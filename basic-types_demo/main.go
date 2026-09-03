package main

import "fmt"

func main(){
	// 整数
	age := 25
	// 浮点数
	price := 19.99
	fmt.Printf("age:%T%v\n",age,age)
	fmt.Printf("price:%T%v\n",price,price)
	// 字符串
	message := "Hello, Golang!"
	// 多行字符串,使用的是反引号而不是单引号
	message2 := `
	hello
	golang
	`
	// 布尔值
	isActive := true
	fmt.Printf("Age:%d,Price:%.2f,Message:%s,Active:%t\n",age,price,message,isActive)
	
	fmt.Println(message2)
}