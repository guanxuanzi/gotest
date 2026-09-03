package main

import "fmt"

func main(){
	// 整数
	age := 25
	// 浮点数
	price := 19.99

	// 类型转换
	total := float64(age) + price
	fmt.Println(total)
	// 数字转字符串
	msg := fmt.Sprintf("年龄是%d岁", age)
	fmt.Println(msg)

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
	// 中文字符在UTF-8中占三个字节,直接切可能会乱码
	// 转成 []rune后,每个元素就是一个字符
	s := "你好李四!"
	sRune := []rune(s)
	fmt.Println("再见"+string(sRune[2:]))
}