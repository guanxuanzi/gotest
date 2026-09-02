package main

// 导包用来输出输入
import "fmt"

func main() {
	// 显式声明类型
	var x int =10
	var name string = "Golang"

	// 类型推导,编译器根据初始值判断类型
	var y = 3.14  // float64

	// 短变量说明(只有函数内能使用,让Go自己进行类型判断)
	z := true

	// 多变量也可以
	a,b := "hello",100

	// 重新赋值用 =
	x = 6
	y = 5

	fmt.Println(x,name,y,z,a,b)
}