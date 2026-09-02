package main

import "fmt"

func main() {
	age := 34
	price := 19.99
	// %T显示类型 %v显示值 Printf不带换行所以得手动添加\n
	fmt.Printf("age:%T %v\n",age,age)
	fmt.Printf("price:%T %v\n",price,price)

	fmt.Printf("Price with 2 decimal places:%2f\n",price)
}