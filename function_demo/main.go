package main

import (
	"errors"
	"fmt"
)

// 普通函数
func add(a int,b int) int {
	return a + b
}

// 调用时可以传入0或者多个int,返回值也为int
func sum(nums ...int) int {
	total := 0
	// _表示忽略的变量,这里表示忽略下标
	for _,n := range nums{
		total += n
	}
	return total
}
// 多返回值函数
func divide(a int,b int) (int,error){
	if b == 0{
		return 0, errors.New("除数不能为0")
	}
	return a/b,nil
}
func main(){
	result := add(1,2)
	fmt.Println(result)
	// 匿名函数
	addFunc := func(a int,b int) int {
		return a + b
	}
	fmt.Println(addFunc(3,4))

	// 调用多返回值函数
	result,err := divide(10,2)
	if err != nil{
		fmt.Println("计算失败:",err)
		return
	} 
		fmt.Println(result)
	sums := []int{20,30,40}
	fmt.Println(sum(sums...))
}
